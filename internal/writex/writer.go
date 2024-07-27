package writex

import (
	"encoding/base64"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func GenerateAll(config configs.ResProtoFiles, files protox.ProtoFiles) {
	excelTableMap := make(map[string][]configs.ResTableConfig)
	for _, resProto := range config {
		for _, table := range resProto.Tables {
			excelFullName := filepath.Join(resProto.Opt.GetExcelPath(), table.GetExcelName())
			_, exist := excelTableMap[excelFullName]

			if !exist {
				excelTableMap[excelFullName] = []configs.ResTableConfig{}
			}
			excelTableMap[excelFullName] = append(excelTableMap[excelFullName], table)
		}
	}
	wait := sync.WaitGroup{}
	wait.Add(len(excelTableMap))
	for excel, table := range excelTableMap {
		go GenerateOneExcel(excel, table, &wait, files)
	}
	wait.Wait()
}

func GenerateOneExcel(excel string, tables []configs.ResTableConfig, s *sync.WaitGroup, files protox.ProtoFiles) {
	file := lo.Must(excelize.OpenFile(excel))
	defer func() {
		lo.Must0(file.Close())
		s.Done()
	}()
	for _, table := range tables {
		var idx = lo.Must(file.GetSheetIndex(table.GetSheetName()))
		if idx == -1 {
			idx = lo.Must(file.GetSheetIndex("#" + table.GetSheetName()))
		}
		if idx == -1 {
			panic(fmt.Sprintf("can not find table %s in excel", table.GetSheetName()))
		}
		GenerateOneSheet(file, table, files)
	}
}

func GenerateOneSheet(file *excelize.File, table configs.ResTableConfig, protos protox.ProtoFiles) {
	tableMessageDesc := protos.GetMessage(table.TableName)
	messageDesc := protos.GetMessage(table.MessageName)
	st := ParseToSheetTable(lo.Must(file.GetRows(table.GetSheetName())))
	for _, tag := range table.Opt.GetMarshalTags() {
		generateOneTable(tag, st, tableMessageDesc, messageDesc, table)
	}
}

func generateOneTable(tag string, data SheetTable, tableMsgDesc protoreflect.MessageDescriptor, msgDesc protoreflect.MessageDescriptor, table configs.ResTableConfig) {
	tableMsg := dynamicpb.NewMessage(tableMsgDesc)
	msgField := protox.GetFieldByMsgType(msgDesc, tableMsgDesc)
	var msg protoreflect.Message = nil
	msgList := tableMsg.Mutable(msgField).List()
	for idx, line := range data.Data {
		if needCreateNewMsg(msgDesc, data.Heads, line) { //一条新消息
			msg = msgList.AppendMutable().Message()
		}
		pressOneLineIntoMsg(msg, idx, data, msgDesc)
	}
	fmt.Println(protojson.Format(tableMsg))
	var formats = table.Belong.Opt.GetMarshalFormats()
	filename := filepath.Join(table.Belong.Opt.GetMarshalPath(), table.MessageName+"_"+tag+".")
	if formats != nil && len(formats) > 0 {
		for _, format := range formats {
			switch format {
			case resproto.ResMarshalFormat_Bin:
				marshal, err := proto.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+"bin", marshal, os.ModePerm))
			case resproto.ResMarshalFormat_Json:
				marshal, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+"json", marshal, os.ModePerm))
			case resproto.ResMarshalFormat_Text:
				marshal, err := prototext.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+"txt", marshal, os.ModePerm))
			}
		}
	}
}

// desc 里number=1得field的值不是空
func needCreateNewMsg(desc protoreflect.MessageDescriptor, nameColMap map[string]ColHead, line []string) bool {
	keyFieldDesc := protox.GetFieldByNumber(1, desc)
	keyName := string(keyFieldDesc.Name())
	keyValue := line[nameColMap[keyName].Col]
	return keyValue != ""
}

func pressOneLineIntoMsg(msg protoreflect.Message, lineIndex int, data SheetTable, desc protoreflect.MessageDescriptor) {
	noRecurCols := lo.PickBy(data.Heads, func(key string, value ColHead) bool {
		return !strings.Contains(key, ".")
	})
	line := data.Data[lineIndex]
	for _, colHead := range noRecurCols {
		value := line[colHead.Col]
		if value == "" {
			continue
		}
		field := desc.Fields().ByName(protoreflect.Name(colHead.Name))
		if field.IsList() {
			list := msg.Mutable(field).List()
			for _, val := range strings.Split(value, "|") {
				list.Append(getFieldValueFromStr(field, val, lineIndex, colHead))
			}
		} else {
			msg.Set(field, getFieldValueFromStr(field, value, lineIndex, colHead))
		}
	}
	commonPrefixMap := make(map[string]map[string]ColHead)
	for _, colHead := range data.Heads {
		if !strings.Contains(colHead.Name, ".") {
			continue
		}
		prefix := colHead.Name[0:strings.Index(colHead.Name, ".")]
		suffix := colHead.Name[strings.Index(colHead.Name, ".")+1:]
		_, ok := commonPrefixMap[prefix]
		if !ok {
			commonPrefixMap[prefix] = make(map[string]ColHead)
		}
		commonPrefixMap[prefix][suffix] = ColHead{Name: suffix, Col: colHead.Col, Additional: colHead.Additional}
	}
	for prefix, ncm := range commonPrefixMap {
		field := desc.Fields().ByName(protoreflect.Name(prefix))
		var hasValue = false
		for _, col := range ncm {
			if line[col.Col] != "" {
				hasValue = true
				break
			}
		}
		if !hasValue {
			continue
		}
		if field.IsMap() {
			mapField := msg.Mutable(field).Map()
			keyStr := getPrevNoEmptyVal(data.Data, lineIndex, ncm["mapKey"].Col)
			if keyStr == "" {
				panic(fmt.Sprintf("missing map key for line :%d,col:%d", lineIndex, ncm["mapKey"]))
			}
			delete(ncm, "mapKey")
			mapValDesc := field.MapValue()
			keyVal := protoreflect.MapKey(getFieldValueFromStr(field.MapKey(), keyStr, lineIndex, ncm["mapKey"]))
			if mapValDesc.Kind() == protoreflect.MessageKind {
				var valVal = mapField.Mutable(keyVal)
				pressOneLineIntoMsg(valVal.Message(), lineIndex, SheetTable{Heads: ncm, Data: data.Data}, mapValDesc.Message())
			} else {
				valVal := getFieldValueFromStr(mapValDesc, line[ncm["mapVal"].Col], lineIndex, ncm["mapVal"])
				mapField.Set(keyVal, valVal)
			}
		} else if field.IsList() {
			list := msg.Mutable(field).List()
			if needCreateNewMsg(field.Message(), ncm, line) {
				var fieldMsg = list.AppendMutable().Message()
				pressOneLineIntoMsg(fieldMsg, lineIndex, SheetTable{Heads: ncm, Data: data.Data}, field.Message())
			} else {
				var fieldMsg = list.Get(list.Len() - 1).Message()
				pressOneLineIntoMsg(fieldMsg, lineIndex, SheetTable{Heads: ncm, Data: data.Data}, field.Message())
			}
		} else {
			fieldMsg := msg.Mutable(field).Message()
			pressOneLineIntoMsg(fieldMsg, lineIndex, SheetTable{Heads: ncm, Data: data.Data}, field.Message())
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

func getFieldValueFromStr(field protoreflect.FieldDescriptor, cell string, lineIndex int, head ColHead) protoreflect.Value {
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
		fieldHead := strings.Split(strings.TrimSuffix(strings.TrimPrefix(head.Additional, "{"), "}"), ";")
		line := strings.Split(cell, ";")
		if len(fieldHead) != len(line) {
			panic(fmt.Sprintf("cell value error msg %v field not match value %s", field.Name(), cell))
		}
		tmp := dynamicpb.NewMessage(fmd)
		nameColMap := make(map[string]ColHead)
		for i, f := range fieldHead {
			nameColMap[f] = ColHead{Name: f, Col: i, Additional: ""}
		}
		pressOneLineIntoMsg(tmp, 0, SheetTable{Heads: nameColMap, Data: [][]string{line}}, fmd)
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
