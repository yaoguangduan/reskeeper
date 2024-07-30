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

func generateOneTable(tag string, data SheetData, tableMsgDesc protoreflect.MessageDescriptor, msgDesc protoreflect.MessageDescriptor, table configs.ResTableConfig) {
	tableMsg := dynamicpb.NewMessage(tableMsgDesc)
	msgField := protox.GetFieldByMsgType(msgDesc, tableMsgDesc)
	var msg protoreflect.Message = nil
	msgList := tableMsg.Mutable(msgField).List()
	for idx, line := range data.Lines {
		if needCreateNewMsg(msgDesc, data.Heads, line) { //一条新消息
			msg = msgList.AppendMutable().Message()
		}
		parseOneLineIntoMsg(tag, msg, idx, data)
	}
	var outPrefix = table.MessageName
	if table.Opt.MarshalPrefix != nil {
		outPrefix = *table.Opt.MarshalPrefix
	}
	var formats = table.Belong.Opt.GetMarshalFormats()
	filename := filepath.Join(table.Belong.Opt.GetMarshalPath(), outPrefix+"_"+tag+".")
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
	keyFieldDesc := protox.GetMsgKeyField(desc)
	keyName := string(keyFieldDesc.Name())
	keyValue := line[nameColMap[keyName].Col]
	return keyValue != ""
}

func parseOneLineIntoMsg(tag string, msg protoreflect.Message, lineIndex int, data SheetData) {
	desc := msg.Descriptor()
	msgOpt := protox.GetMsgOptOrDefault(desc)
	line := data.Lines[lineIndex]
	selfFields := lo.PickBy(data.Heads, func(key string, value ColHead) bool {
		return !strings.Contains(key, ".")
	})
	for _, colHead := range selfFields {
		value := line[colHead.Col]
		field := desc.Fields().ByName(protoreflect.Name(colHead.Name))
		if value == "" || protox.IgnoreCurField(tag, msgOpt, field) {
			continue
		}
		if field.IsList() {
			list := msg.Mutable(field).List()
			for _, val := range strings.Split(value, "|") {
				list.Append(getFieldValueFromStr(tag, field, val, colHead))
			}
		} else if field.IsMap() {
			mapTmp := getFieldValueFromStr(tag, field, value, colHead).Map()
			mapCur := msg.Mutable(field).Map()
			mapTmp.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
				mapCur.Set(key, value)
				return true
			})
		} else {
			msg.Set(field, getFieldValueFromStr(tag, field, value, colHead))
		}
	}
	processNestedFields(tag, msg, lineIndex, data, desc, msgOpt)
}

func processNestedFields(tag string, msg protoreflect.Message, lineIndex int, data SheetData, desc protoreflect.MessageDescriptor, msgOpt *resproto.ResourceMsgOpt) {
	line := data.Lines[lineIndex]
	commonPrefixMap := make(map[string]map[string]ColHead)
	for _, colHead := range data.Heads {
		if !strings.Contains(colHead.Name, ".") {
			continue
		}
		prefix := colHead.Name[0:strings.Index(colHead.Name, ".")]
		suffix := colHead.Name[strings.Index(colHead.Name, ".")+1:]
		if _, ok := commonPrefixMap[prefix]; !ok {
			commonPrefixMap[prefix] = make(map[string]ColHead)
		}
		commonPrefixMap[prefix][suffix] = ColHead{Name: suffix, Col: colHead.Col, NestFields: colHead.NestFields}
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
		if !hasValue || protox.IgnoreCurField(tag, msgOpt, field) {
			continue
		}
		if field.IsMap() {
			mapField := msg.Mutable(field).Map()
			keyVal := getMapKeyFromLine(tag, ncm, data, lineIndex, field)
			mapValDesc := field.MapValue()
			if mapValDesc.Kind() == protoreflect.MessageKind {
				var valVal = mapField.Mutable(keyVal)
				valHead, exist := ncm["map-val"]
				if exist {
					fieldHead := strings.Split(strings.TrimSuffix(strings.TrimPrefix(valHead.NestFields, "{"), "}"), ";")
					currentLine := strings.Split(line[valHead.Col], ";")
					if len(fieldHead) != len(currentLine) {
						panic(fmt.Sprintf("cell value error msg %v field not match value %s", field.Name(), line[valHead.Col]))
					}
					nameColMap := make(map[string]ColHead)
					for i, f := range fieldHead {
						nameColMap[f] = ColHead{Name: f, Col: i, NestFields: ""}
					}
					parseOneLineIntoMsg(tag, valVal.Message(), 0, SheetData{Heads: nameColMap, Lines: [][]string{currentLine}})
				} else {
					parseOneLineIntoMsg(tag, valVal.Message(), lineIndex, SheetData{Heads: ncm, Lines: data.Lines})
				}
			} else {
				valHead, exist := ncm["map-val"]
				if !exist {
					mapField.Set(keyVal, field.Default())
				} else {
					valVal := getFieldValueFromStr(tag, mapValDesc, line[valHead.Col], valHead)
					mapField.Set(keyVal, valVal)
				}
			}
		} else if field.IsList() {
			list := msg.Mutable(field).List()
			if needCreateNewMsg(field.Message(), ncm, line) {
				parseOneLineIntoMsg(tag, list.AppendMutable().Message(), lineIndex, SheetData{Heads: ncm, Lines: data.Lines})
			} else {
				parseOneLineIntoMsg(tag, list.Get(list.Len()-1).Message(), lineIndex, SheetData{Heads: ncm, Lines: data.Lines})
			}
		} else {
			parseOneLineIntoMsg(tag, msg.Mutable(field).Message(), lineIndex, SheetData{Heads: ncm, Lines: data.Lines})
		}
	}
}

func getMapKeyFromLine(tag string, ncm map[string]ColHead, data SheetData, lineIndex int, field protoreflect.FieldDescriptor) protoreflect.MapKey {
	var headCol = ColHead{Col: -1}
	if _, exist := ncm["map-key"]; exist {
		headCol = ncm["map-key"]
	} else {
		if field.MapValue().Kind() != protoreflect.MessageKind {
			panic(fmt.Sprintf("missing map key for line :%d,msg:%v", lineIndex, field.Name()))
		}
		mapValDesc := field.MapValue().Message()
		keyField := protox.GetMsgKeyField(mapValDesc)
		keyColHead, ok := ncm[string(keyField.Name())]
		if !ok {
			panic(fmt.Sprintf("missing map key for line :%d,msg:%v", lineIndex, field.Name()))
		}
		headCol = keyColHead
	}
	if headCol.Col == -1 {
		panic(fmt.Sprintf("missing map key for line :%d,msg:%v", lineIndex, field.Name()))
	}
	keyStr := getPrevNoEmptyVal(data.Lines, lineIndex, headCol.Col)
	if keyStr == "" {
		panic(fmt.Sprintf("missing map key for line :%d,msg:%v", lineIndex, field.Name()))
	}
	delete(ncm, "map-key")
	return protoreflect.MapKey(getFieldValueFromStr(tag, field.MapKey(), keyStr, headCol))
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

func getFieldValueFromStr(tag string, field protoreflect.FieldDescriptor, cell string, head ColHead) protoreflect.Value {
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
		var fmd = field.Message()
		if field.IsMap() {
			fmd = field.MapValue().Message()
		}
		fieldHead := strings.Split(strings.TrimSuffix(strings.TrimPrefix(head.NestFields, "{"), "}"), ";")
		line := strings.Split(cell, ";")
		if len(fieldHead) != len(line) {
			panic(fmt.Sprintf("cell value error msg %v field not match value %s", field.Name(), cell))
		}
		if field.IsMap() && field.MapValue().Kind() != protoreflect.MessageKind {
			parentMsg := field.Parent().(protoreflect.MessageDescriptor)
			mapField := dynamicpb.NewMessage(parentMsg).Mutable(field).Map()
			key := protoreflect.MapKey(getFieldValueFromStr(tag, field.MapKey(), line[0], head))
			val := getFieldValueFromStr(tag, field.MapValue(), line[1], head)
			mapField.Set(key, val)
			return protoreflect.ValueOfMap(mapField)
		}
		tmp := dynamicpb.NewMessage(fmd)
		nameColMap := make(map[string]ColHead)
		for i, f := range fieldHead {
			nameColMap[f] = ColHead{Name: f, Col: i, NestFields: ""}
		}
		parseOneLineIntoMsg(tag, tmp, 0, SheetData{Heads: nameColMap, Lines: [][]string{line}})
		if field.IsMap() {
			parentMsg := field.Parent().(protoreflect.MessageDescriptor)
			mapField := dynamicpb.NewMessage(parentMsg).Mutable(field).Map()
			keyVal := tmp.Get(protox.GetMsgKeyField(fmd))
			mapField.Set(protoreflect.MapKey(keyVal), protoreflect.ValueOfMessage(tmp))
			value = protoreflect.ValueOfMap(mapField)
		} else {
			value = protoreflect.ValueOfMessage(tmp)
		}
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
