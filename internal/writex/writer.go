package writex

import (
	"cmp"
	"encoding/base64"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"os"
	"path/filepath"
	"reskeeper/internal/protox"
	"reskeeper/res_toml"
	"slices"
	"strings"
	"sync"
)

func GenerateAll(config res_toml.Config, files protox.ProtoFiles) {
	excelTableMap := make(map[string][]res_toml.Table)
	for _, table := range config.Tables {
		_, exist := excelTableMap[table.Excel]
		if !exist {
			excelTableMap[table.Excel] = []res_toml.Table{}
		}
		excelTableMap[table.Excel] = append(excelTableMap[table.Excel], table)
	}
	wait := sync.WaitGroup{}
	wait.Add(len(excelTableMap))
	for excel, table := range excelTableMap {
		go GenerateOneExcel(excel, table, &wait, config, files)
	}
	wait.Wait()
}

func GenerateOneExcel(excel string, tables []res_toml.Table, s *sync.WaitGroup, config res_toml.Config, files protox.ProtoFiles) {
	file := lo.Must(excelize.OpenFile(filepath.Join(config.ExcelPath, excel)))
	defer func() {
		lo.Must0(file.Close())
		s.Done()
	}()
	for _, table := range tables {
		var idx = lo.Must(file.GetSheetIndex(table.SheetName))
		if idx == -1 {
			idx = lo.Must(file.GetSheetIndex("#" + table.SheetName))
		}
		if idx == -1 {
			panic(fmt.Sprintf("can not find table %s in excel", table.SheetName))
		}
		GenerateOneSheet(file, excel, table, config, files)
	}
}

func GenerateOneSheet(file *excelize.File, excel string, table res_toml.Table, config res_toml.Config, protos protox.ProtoFiles) {
	tableMessageDesc := protos.GetMessage(table.TableName)
	messageDesc := protos.GetMessage(table.MessageName)
	st := ParseToSheetTable(lo.Must(file.GetRows(table.SheetName)))
	for suffix, header := range st.Heads {
		generateOneTable(suffix, header, st.Data, tableMessageDesc, messageDesc, table, config)
	}
}

func generateOneTable(suffix string, nameColMap map[string]int, lines [][]string, tableMsgDesc protoreflect.MessageDescriptor, msgDesc protoreflect.MessageDescriptor, table res_toml.Table, config res_toml.Config) {
	tableMsg := dynamicpb.NewMessage(tableMsgDesc)
	msgField := protox.GetFieldByMsgType(msgDesc, tableMsgDesc)
	var msg protoreflect.Message = nil
	msgList := tableMsg.Mutable(msgField).List()
	for idx, line := range lines {
		if needCreateNewMsg(msgDesc, nameColMap, line) { //一条新消息
			msg = msgList.AppendMutable().Message()
		}
		pressOneLineIntoMsg(msg, idx, line, nameColMap, msgDesc, lines)
	}
	var formats = table.OutFormats
	if formats == nil {
		formats = config.OutFormats
	}
	filename := filepath.Join(config.OutPath, table.MessageName+"_"+suffix+".")
	if formats != nil && len(*formats) > 0 {
		for _, format := range *formats {
			switch format {
			case "bin":
				marshal, err := proto.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+format, marshal, os.ModePerm))
			case "json":
				marshal, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+format, marshal, os.ModePerm))
			case "text":
				marshal, err := prototext.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+format, marshal, os.ModePerm))
			}
		}
	}
}

// desc 里number=1得field的值不是空
func needCreateNewMsg(desc protoreflect.MessageDescriptor, nameColMap map[string]int, line []string) bool {
	keyFieldDesc := protox.GetFieldByNumber(1, desc)
	keyName := string(keyFieldDesc.Name())
	keyValue := line[nameColMap[keyName]]
	return keyValue != ""
}

func pressOneLineIntoMsg(msg protoreflect.Message, lineIndex int, line []string, nameColMap map[string]int, desc protoreflect.MessageDescriptor, lines [][]string) {
	noRecurCols := lo.PickBy(nameColMap, func(key string, value int) bool {
		return !strings.Contains(key, ".")
	})
	for name, col := range noRecurCols {
		value := line[col]
		if value == "" {
			continue
		}
		field := desc.Fields().ByName(protoreflect.Name(name))
		if field.IsList() {
			list := msg.Mutable(field).List()
			for _, val := range strings.Split(value, "|") {
				list.Append(getFieldValueFromStr(field, val, lines, lineIndex))
			}
		} else {
			msg.Set(field, getFieldValueFromStr(field, value, lines, lineIndex))
		}
	}
	commonPrefixMap := make(map[string]map[string]int)
	for name, col := range nameColMap {
		if !strings.Contains(name, ".") {
			continue
		}
		prefix := name[0:strings.Index(name, ".")]
		suffix := name[strings.Index(name, ".")+1:]
		_, ok := commonPrefixMap[prefix]
		if !ok {
			commonPrefixMap[prefix] = make(map[string]int)
		}
		commonPrefixMap[prefix][suffix] = col
	}
	for prefix, ncm := range commonPrefixMap {
		field := desc.Fields().ByName(protoreflect.Name(prefix))
		var hasValue = false
		for _, col := range ncm {
			if line[col] != "" {
				hasValue = true
				break
			}
		}
		if !hasValue {
			continue
		}
		if field.IsMap() {
			mapField := msg.Mutable(field).Map()
			keyStr := getPrevNoEmptyVal(lines, lineIndex, ncm["mapKey"])
			if keyStr == "" {
				panic(fmt.Sprintf("missing map key for line :%d,col:%d", lineIndex, ncm["mapKey"]))
			}
			delete(ncm, "mapKey")
			mapValDesc := field.MapValue()
			keyVal := protoreflect.MapKey(getFieldValueFromStr(field.MapKey(), keyStr, lines, lineIndex))
			if mapValDesc.Kind() == protoreflect.MessageKind {
				var valVal = mapField.Mutable(keyVal)
				pressOneLineIntoMsg(valVal.Message(), lineIndex, line, ncm, mapValDesc.Message(), lines)
			} else {
				valVal := getFieldValueFromStr(mapValDesc, line[ncm["mapVal"]], lines, lineIndex)
				mapField.Set(keyVal, valVal)
			}
		} else if field.IsList() {
			list := msg.Mutable(field).List()
			if needCreateNewMsg(field.Message(), ncm, line) {
				var fieldMsg = list.AppendMutable().Message()
				pressOneLineIntoMsg(fieldMsg, lineIndex, line, ncm, field.Message(), lines)
			} else {
				var fieldMsg = list.Get(list.Len() - 1).Message()
				pressOneLineIntoMsg(fieldMsg, lineIndex, line, ncm, field.Message(), lines)
			}
		} else {
			fieldMsg := msg.Mutable(field).Message()
			pressOneLineIntoMsg(fieldMsg, lineIndex, line, ncm, field.Message(), lines)
		}
	}
}

// 获取从line开始倒数，col列的第一个非零string
func getPrevNoEmptyVal(lines [][]string, line int, col int) string {
	for i := line; i >= 0; i-- {
		if lines[i][col] != "" {
			return lines[i][col]
		}
	}
	return ""
}

func getFieldValueFromStr(field protoreflect.FieldDescriptor, cell string, lines [][]string, lineIndex int) protoreflect.Value {
	var value protoreflect.Value
	switch field.Kind() {
	case protoreflect.StringKind:
		value = protoreflect.ValueOfString(cell)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		value = protoreflect.ValueOfUint32(cast.ToUint32(cell))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		value = protoreflect.ValueOfInt32(cast.ToInt32(cell))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		value = protoreflect.ValueOfUint64(cast.ToUint64(cell))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		value = protoreflect.ValueOfInt64(cast.ToInt64(cell))
	case protoreflect.BoolKind:
		value = protoreflect.ValueOfBool(cast.ToBool(cell))
	case protoreflect.DoubleKind:
		value = protoreflect.ValueOfFloat64(cast.ToFloat64(cell))
	case protoreflect.FloatKind:
		value = protoreflect.ValueOfFloat32(cast.ToFloat32(cell))
	case protoreflect.EnumKind:
		value = protoreflect.ValueOfEnum(toEnumNumber(field, cell))
	case protoreflect.BytesKind:
		value = protoreflect.ValueOfBytes(lo.Must(base64.StdEncoding.DecodeString(cell)))
	case protoreflect.MessageKind:
		fmd := field.Message()
		fields := fmd.Fields()
		line := strings.Split(cell, ";")
		if fields.Len() != len(line) {
			panic(fmt.Sprintf("cell value error msg %v field not match value %s", field.Name(), cell))
		}
		tmp := dynamicpb.NewMessage(fmd)
		fieldSlice := make([]protoreflect.FieldDescriptor, 0)
		for i := 0; i < fields.Len(); i++ {
			fieldSlice = append(fieldSlice, fields.Get(i))
		}
		slices.SortFunc(fieldSlice, func(a, b protoreflect.FieldDescriptor) int {
			return cmp.Compare(a.Number(), b.Number())
		})
		nameColMap := make(map[string]int)
		for i, f := range fieldSlice {
			nameColMap[string(f.Name())] = i
		}
		pressOneLineIntoMsg(tmp, lineIndex, line, nameColMap, fmd, lines)
		value = protoreflect.ValueOfMessage(tmp)
	}
	return value
}

func toEnumNumber(field protoreflect.FieldDescriptor, cell string) protoreflect.EnumNumber {
	evs := field.Enum().Values()
	i32, err := cast.ToInt32E(cell)
	if err == nil {
		return evs.ByNumber(protoreflect.EnumNumber(i32)).Number()
	} else {
		return evs.ByName(protoreflect.Name(cell)).Number()
	}
}

// splitIgnoringBraces 按照 `;` 分割字符串，但忽略在花括号 `{}` 内的 `;`
func splitIgnoringBraces(input string) []string {
	var result []string
	var currentSegment strings.Builder
	bracesCount := 0

	for _, char := range input {
		switch char {
		case '{':
			bracesCount++
			currentSegment.WriteRune(char)
		case '}':
			bracesCount--
			currentSegment.WriteRune(char)
		case ';':
			if bracesCount > 0 {
				currentSegment.WriteRune(char)
			} else {
				result = append(result, currentSegment.String())
				currentSegment.Reset()
			}
		default:
			currentSegment.WriteRune(char)
		}
	}

	// Append the last segment if it's not empty
	if currentSegment.Len() > 0 {
		result = append(result, currentSegment.String())
	}

	return result
}
