package convert

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"github.com/yaoguangduan/reskeeper/internal/convert/pson"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/internal/validate"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func GenerateAll(config configs.ResProtoFiles, files protox.ProtoFiles, list []string) {
	excelTableMap := make(map[string][]configs.ResTableConfig)
	for _, resProto := range config {
		for _, table := range resProto.Tables {
			excelFullName := table.GetExcelName()
			excelNameWithSheet := table.GetExcelName() + "#" + table.GetSheetName()
			if lo.NoneBy(list, func(item string) bool {
				if item == table.GetExcelName() || item == excelNameWithSheet || item == table.GetSheetName() {
					return true
				}
				exp := regexp.MustCompile(item)
				return exp.MatchString(table.GetExcelName()) || exp.MatchString(excelNameWithSheet)
			}) {
				continue
			}
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
	validate.Validator.PrintValidateResult()
}

func GenerateOneExcel(excel string, tables []configs.ResTableConfig, s *sync.WaitGroup, files protox.ProtoFiles) {
	log.Printf("start convert excel : [%s]", excel)
	file := lo.Must(excelize.OpenFile(excel))
	defer func() {
		lo.Must0(file.Close())
		s.Done()
		log.Printf("end convert excel : [%s]", excel)
	}()
	for _, table := range tables {
		var idx = lo.Must(file.GetSheetIndex(table.GetSheetName()))
		if idx == -1 {
			idx = lo.Must(file.GetSheetIndex("#" + table.GetSheetName()))
		}
		if idx == -1 {
			panic(fmt.Sprintf("can not find table %s in excel", table.GetSheetName()))
		}
		convertOneSheet(file, table, files)
	}
}

func convertOneSheet(file *excelize.File, table configs.ResTableConfig, protos protox.ProtoFiles) {
	log.Printf("start convert excel sheet : %s#%s", table.GetExcelName(), table.GetSheetName())
	tableMessageDesc := protos.GetMessage(table.TableName)
	messageDesc := protos.GetMessage(table.MessageName)
	st := configs.ParseToSheetTable(lo.Must(file.GetRows(table.GetSheetName())))
	ctx := configs.CvtContext{Table: table, Protos: protos, Sheet: st, TableDesc: tableMessageDesc, DataDesc: messageDesc}
	for _, tag := range table.GetGenerateTags() {
		ctx.Tag = tag
		convertOneTable(ctx)
	}
	log.Printf("end convert excel sheet : %s#%s", table.GetExcelName(), table.GetSheetName())
}

func convertOneTable(ctx configs.CvtContext) {
	data := ctx.Sheet
	tag := ctx.Tag
	table := ctx.Table
	tableMsg := dynamicpb.NewMessage(ctx.TableDesc)
	msgField := protox.GetFieldByMsgType(ctx.DataDesc, ctx.TableDesc)
	var msg protoreflect.Message = nil
	msgList := tableMsg.Mutable(msgField).List()
	for idx, line := range data.Lines {
		if needCreateNewMsg(data.Heads, line, ctx.DataDesc) { //一条新消息
			msg = msgList.AppendMutable().Message()
		}
		parseOneLineIntoMsg(tag, msg, idx, data)
	}
	log.Printf("start to do resource validate %s#%s %s", table.GetExcelName(), table.GetSheetName(), tag)
	validate.Validator.Validate(tableMsg, ctx)

	var outPrefix = table.MessageName
	if table.GenerateName != "" {
		outPrefix = table.GenerateName
	}
	var formats = table.Belong.GetMarshalFormats()
	filename := filepath.Join(table.Belong.GetGeneratePath(), outPrefix+"."+tag+".")
	if formats != nil && len(formats) > 0 {
		for _, format := range formats {
			switch format {
			case "bin":
				marshaller := proto.MarshalOptions{Deterministic: true}
				marshal, err := marshaller.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+"bin", marshal, os.ModePerm))
			case "json":
				marshal, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+"json", marshal, os.ModePerm))
			case "txt":
				marshal, err := prototext.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(tableMsg)
				if err != nil {
					panic(err)
				}
				lo.Must0(os.WriteFile(filename+"txt", marshal, os.ModePerm))
			}
		}
	}
}

// 任何一个非repeated和map的字段有值，则需要创建新message
func needCreateNewMsg(nameColMap map[string]configs.ColHead, line []string, descriptor protoreflect.MessageDescriptor) bool {
	selfFields := lo.PickBy(nameColMap, func(key string, value configs.ColHead) bool {
		return !strings.Contains(key, ".") && !descriptor.Fields().ByName(protoreflect.Name(key)).IsMap() && !descriptor.Fields().ByName(protoreflect.Name(key)).IsList()
	})
	var newMsg = false
	for _, colHead := range selfFields {
		var value = line[colHead.Col]
		if value != "" {
			newMsg = true
			break
		}
	}
	return newMsg
}

func parseOneLineIntoMsg(tag string, msg protoreflect.Message, lineIndex int, data configs.SheetData) {
	desc := msg.Descriptor()
	line := data.Lines[lineIndex]
	selfFields := lo.PickBy(data.Heads, func(key string, value configs.ColHead) bool {
		return !strings.Contains(key, ".")
	})
	for _, colHead := range selfFields {
		var value = line[colHead.Col]
		field := desc.Fields().ByName(protoreflect.Name(colHead.Name))
		if value == "" {
			if proto.HasExtension(field.Options(), resproto.E_ResDefault) {
				defVal := proto.GetExtension(field.Options(), resproto.E_ResDefault).(string)
				val := GetFieldValueFromStr("", field, defVal, configs.ColHead{
					Name:       string(field.Name()),
					Col:        0,
					NestFields: "",
				})
				msg.Set(field, val)
			}
			continue
		}
		if field == nil && desc.Oneofs().ByName(protoreflect.Name(colHead.Name)) != nil {
			fn := value[0:strings.Index(value, "{")]
			field = desc.Fields().ByName(protoreflect.Name(fn))
			value = strings.TrimSuffix(value[strings.Index(value, "{")+1:], "}")
		}
		if field == nil {
			log.Panicf("can not find message field def:%s %s", msg.Descriptor().Name(), colHead.Name)
		}
		if protox.IgnoreCurField(tag, protox.GetMsgTagIgnoreInfo(desc), field) {
			continue
		}
		if field.IsList() {
			list := msg.Mutable(field).List()
			for _, val := range strings.Split(value, "|") {
				list.Append(GetFieldValueFromStr(tag, field, val, colHead))
			}
		} else if field.IsMap() {
			mapTmp := GetFieldValueFromStr(tag, field, value, colHead).Map()
			mapCur := msg.Mutable(field).Map()
			mapTmp.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
				mapCur.Set(key, value)
				return true
			})
		} else {
			msg.Set(field, GetFieldValueFromStr(tag, field, value, colHead))
		}
	}
	processNestedFields(tag, msg, lineIndex, data, desc)
}

func processNestedFields(tag string, msg protoreflect.Message, lineIndex int, data configs.SheetData, desc protoreflect.MessageDescriptor) {
	line := data.Lines[lineIndex]
	commonPrefixMap := make(map[string]map[string]configs.ColHead)
	for _, colHead := range data.Heads {
		if !strings.Contains(colHead.Name, ".") {
			continue
		}
		prefix := colHead.Name[0:strings.Index(colHead.Name, ".")]
		suffix := colHead.Name[strings.Index(colHead.Name, ".")+1:]
		if _, ok := commonPrefixMap[prefix]; !ok {
			commonPrefixMap[prefix] = make(map[string]configs.ColHead)
		}
		commonPrefixMap[prefix][suffix] = configs.ColHead{Name: suffix, Col: colHead.Col, NestFields: colHead.NestFields}
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
		if !hasValue || protox.IgnoreCurField(tag, protox.GetMsgTagIgnoreInfo(desc), field) {
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
					parseOneLineIntoMsg(tag, valVal.Message(), 0, parseHeadAndCellToSheetData(valHead.NestFields, line[valHead.Col]))
				} else {
					parseOneLineIntoMsg(tag, valVal.Message(), lineIndex, configs.SheetData{Heads: ncm, Lines: data.Lines})
				}
			} else {
				valHead, exist := ncm["map-val"]
				if !exist {
					mapField.Set(keyVal, field.Default())
				} else {
					valVal := GetFieldValueFromStr(tag, mapValDesc, line[valHead.Col], valHead)
					mapField.Set(keyVal, valVal)
				}
			}
		} else if field.IsList() {
			list := msg.Mutable(field).List()
			if needCreateNewMsg(ncm, line, field.Message()) {
				parseOneLineIntoMsg(tag, list.AppendMutable().Message(), lineIndex, configs.SheetData{Heads: ncm, Lines: data.Lines})
			} else {
				parseOneLineIntoMsg(tag, list.Get(list.Len()-1).Message(), lineIndex, configs.SheetData{Heads: ncm, Lines: data.Lines})
			}
		} else {
			parseOneLineIntoMsg(tag, msg.Mutable(field).Message(), lineIndex, configs.SheetData{Heads: ncm, Lines: data.Lines})
		}
	}
}

func parseHeadAndCellToSheetData(fields, values string) configs.SheetData {
	fieldHead := strings.Split(strings.TrimSuffix(strings.TrimPrefix(fields, "{"), "}"), ";")
	currentLine := strings.Split(values, ";")
	if len(fieldHead) != len(currentLine) {
		panic(fmt.Sprintf("cell value error msg %v field not match value %s", fields, values))
	}
	nameColMap := make(map[string]configs.ColHead)
	for i, f := range fieldHead {
		nameColMap[f] = configs.ColHead{Name: f, Col: i, NestFields: ""}
	}
	return configs.SheetData{Heads: nameColMap, Lines: [][]string{currentLine}}
}

func getMapKeyFromLine(tag string, ncm map[string]configs.ColHead, data configs.SheetData, lineIndex int, field protoreflect.FieldDescriptor) protoreflect.MapKey {
	var headCol = configs.ColHead{Col: -1}
	if _, exist := ncm["key"]; exist {
		headCol = ncm["key"]
	} else {
		if field.MapValue().Kind() != protoreflect.MessageKind {
		}
		mapValDesc := field.MapValue().Message()
		keyField := protox.GetMsgKeyField(mapValDesc)
		if keyField == nil {
			panic(fmt.Sprintf("missing map key for line :%d,msg:%v", lineIndex, field.Name()))
		}
		keyColHead, ok := ncm[string((*keyField).Name())]
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
	delete(ncm, "key")
	return protoreflect.MapKey(GetFieldValueFromStr(tag, field.MapKey(), keyStr, headCol))
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

func GetFieldValueFromStr(tag string, field protoreflect.FieldDescriptor, cell string, head configs.ColHead) protoreflect.Value {
	if field.Kind() == protoreflect.MessageKind {
		var value protoreflect.Value
		var fmd = field.Message()
		if field.IsMap() {
			fmd = field.MapValue().Message()
		}
		line := strings.Split(cell, ":")
		if field.IsMap() && field.MapValue().Kind() != protoreflect.MessageKind {
			parentMsg := field.Parent().(protoreflect.MessageDescriptor)
			mapField := dynamicpb.NewMessage(parentMsg).Mutable(field).Map()
			key := protoreflect.MapKey(GetFieldValueFromStr(tag, field.MapKey(), line[0], head))
			val := GetFieldValueFromStr(tag, field.MapValue(), line[1], head)
			mapField.Set(key, val)
			return protoreflect.ValueOfMap(mapField)
		}
		tmp := dynamicpb.NewMessage(fmd)
		if head.NestFields != "" {
			sheetData := parseHeadAndCellToSheetData(head.NestFields, cell)
			parseOneLineIntoMsg(tag, tmp, 0, sheetData)
			if field.IsMap() {
				parentMsg := field.Parent().(protoreflect.MessageDescriptor)
				mapField := dynamicpb.NewMessage(parentMsg).Mutable(field).Map()
				keyVal := tmp.Get(*protox.GetMsgKeyField(fmd)) //todo
				mapField.Set(protoreflect.MapKey(keyVal), protoreflect.ValueOfMessage(tmp))
				value = protoreflect.ValueOfMap(mapField)
			} else {
				value = protoreflect.ValueOfMessage(tmp)
			}
		} else {
			pson.Decode(tag, tmp, cell)
			value = protoreflect.ValueOfMessage(tmp)
		}
		return value
	} else {
		return pson.ValueOfField(tag, field, cell)
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
