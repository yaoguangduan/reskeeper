package excelx

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"github.com/yaoguangduan/reskeeper/internal/excelx/styles"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func GenExcelFiles(configs []configs.ResProtoFileConfig, files protox.ProtoFiles) {
	for _, config := range configs {
		genByOneConfigProto(config, files)
	}
}

func genByOneConfigProto(config configs.ResProtoFileConfig, files protox.ProtoFiles) {
	excelTableMap := make(map[string][]configs.ResTableConfig)
	for _, table := range config.Tables {
		_, exist := excelTableMap[table.GetExcelName()]
		if !exist {
			excelTableMap[table.GetExcelName()] = []configs.ResTableConfig{}
		}
		excelTableMap[table.GetExcelName()] = append(excelTableMap[table.GetExcelName()], table)
	}
	for excelName, tableConfigs := range excelTableMap {
		fp := filepath.Join(excelName)
		_, err := os.Stat(fp)
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("table file %s state err", fp))
		}
		if err != nil {
			log.Printf("create excel %s", excelName)
			newExcelWithConfig(excelName, tableConfigs, config, files)
		} else {
			adjustExcelWithConfig(excelName, tableConfigs, config, files)
		}
	}
}

func adjustExcelWithConfig(excel string, tables []configs.ResTableConfig, config configs.ResProtoFileConfig, files protox.ProtoFiles) {
	file, err := excelize.OpenFile(excel)
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		var index = lo.Must(file.GetSheetIndex(table.GetSheetName()))
		if index == -1 {
			index = lo.Must(file.GetSheetIndex("#" + table.GetSheetName()))
			if index == -1 {
				createTableSheetOnExcel(file, table, files, config)
			}
		} else {
			adjustColumnOnSheet(file, table, files, config)
		}
	}
	err = file.Save()
	if err != nil {
		panic(err)
	}
	lo.Must0(file.Close())
}
func newExcelWithConfig(excel string, tables []configs.ResTableConfig, config configs.ResProtoFileConfig, files protox.ProtoFiles) {
	excelFile := excelize.NewFile()

	defer func() {
		if err := recover(); err != nil {
			log.Panicf("create excel/sheet %s err %v", excel, err)
		} else {
			_ = excelFile.DeleteSheet("Sheet1")
			if err = excelFile.SaveAs(excel); err != nil {
				panic(err)
			}
		}
	}()
	for _, table := range tables {
		createTableSheetOnExcel(excelFile, table, files, config)
	}
}

type fieldExcelCellDesc struct {
	name    string
	comment string
}

func adjustColumnOnSheet(file *excelize.File, table configs.ResTableConfig, files protox.ProtoFiles, config configs.ResProtoFileConfig) {
	log.Printf("adjust column for sheet %s of %s", table.GetSheetName(), table.GetExcelName())
	rows, err := file.Rows(table.GetSheetName())
	if err != nil {
		panic(err)
	}
	if !rows.Next() {
		lo.Must0(rows.Close())
		lo.Must0(file.DeleteSheet(table.GetSheetName()))
		createTableSheetOnExcel(file, table, files, config)
		return
	}
	defer func() {
		lo.Must0(rows.Close())
	}()
	firstRow, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	fullFieldFlat := flatFieldName(nil, files.GetMessage(table.MessageName), "", table, "")
	missing := make([]int, 0)
	for i, cell := range fullFieldFlat {
		if slices.Contains(firstRow, cell.name) || slices.Contains(firstRow, "#"+cell.name) {
			continue
		}
		missing = append(missing, i)
		colName := lo.Must(excelize.ColumnNumberToName(i + 1))
		if i < len(firstRow) {
			lo.Must0(file.InsertCols(table.GetSheetName(), colName, 1))
			lo.Must0(file.SetCellValue(table.GetSheetName(), colName+"1", cell.name))
			firstRow = append(firstRow, cell.name)
		} else {
			colName = lo.Must(excelize.ColumnNumberToName(len(firstRow) + 1))
			firstRow = append(firstRow, cell.name)
			lo.Must0(file.SetCellValue(table.GetSheetName(), colName+"1", cell.name))
		}
	}
	for i := range fullFieldFlat {
		cellVal := lo.Must(file.GetCellValue(table.GetSheetName(), fmt.Sprintf("%s1", lo.Must(excelize.ColumnNumberToName(i+1)))))
		cell := lo.Must(lo.Find[fieldExcelCellDesc](fullFieldFlat, func(item fieldExcelCellDesc) bool {
			return item.name == cellVal
		}))
		lo.Must0(file.DeleteComment(table.SheetName, lo.Must(excelize.ColumnNumberToName(i+1))+"1"))
		lo.Must0(file.AddComment(table.GetSheetName(), excelize.Comment{
			Cell: lo.Must(excelize.ColumnNumberToName(i+1)) + "1",
			Text: cell.comment,
		}))
	}

	lo.Must0(file.SetRowStyle(table.GetSheetName(), 1, 65535, styles.FontAlignCenter(file)))
	lo.Must0(file.SetCellStyle(table.GetSheetName(), "A1", fmt.Sprintf("%s1", lo.Must(excelize.ColumnNumberToName(len(fullFieldFlat)))), styles.FontBold(file)))
	styles.AdjustColumnWidth(file, table.GetSheetName(), len(fullFieldFlat))
}

func createTableSheetOnExcel(excelFile *excelize.File, table configs.ResTableConfig, files protox.ProtoFiles, config configs.ResProtoFileConfig) {
	log.Printf("generate new sheet %s for excel %s", table.GetSheetName(), table.GetExcelName())
	_, err := excelFile.NewSheet(table.GetSheetName())
	if err != nil {
		panic(err)
	}
	msgD := files.GetMessage(table.MessageName)

	fullFieldFlat := make([]fieldExcelCellDesc, 0)
	fullFieldFlat = flatFieldName(fullFieldFlat, msgD, "", table, "")
	nameList := lo.Map[fieldExcelCellDesc, string](fullFieldFlat, func(item fieldExcelCellDesc, index int) string {
		return item.name
	})
	lo.Must0(excelFile.SetSheetRow(table.GetSheetName(), "A1", &nameList))
	for i := range fullFieldFlat {
		colRow := lo.Must1[string](excelize.ColumnNumberToName(i+1)) + "1"
		lo.Must0(excelFile.AddComment(table.GetSheetName(), excelize.Comment{
			Cell: colRow,
			Text: fullFieldFlat[i].comment,
		}))
	}
	lo.Must0(excelFile.SetRowStyle(table.GetSheetName(), 1, 65535, styles.FontAlignCenter(excelFile)))
	lo.Must0(excelFile.SetCellStyle(table.GetSheetName(), "A1", fmt.Sprintf("%s1", lo.Must(excelize.ColumnNumberToName(len(fullFieldFlat)))), styles.FontBold(excelFile)))
	styles.AdjustColumnWidth(excelFile, table.GetSheetName(), len(fullFieldFlat))
}

func flatFieldName(fullFieldFlat []fieldExcelCellDesc, msgD protoreflect.MessageDescriptor, prefix string, table configs.ResTableConfig, parentComment string) []fieldExcelCellDesc {
	for i := 0; i < msgD.Fields().Len(); i++ {
		f := msgD.Fields().Get(i)
		var comment = ""
		if proto.HasExtension(f.Options(), resproto.E_ResComment) {
			comment = proto.GetExtension(f.Options(), resproto.E_ResComment).(string)
		}
		if parentComment != "" {
			if comment == "" {
				comment = parentComment
			} else {
				comment = parentComment + " / " + comment
			}
		}
		if f.Kind() == protoreflect.MessageKind {
			if f.IsMap() {
				if f.MapValue().Kind() != protoreflect.MessageKind {
					fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name()), f, table, comment))
				} else {
					keyField := protox.GetMsgKeyField(f.MapValue().Message())
					if keyField == nil || f.MapKey().Kind() != (*keyField).Kind() {
						fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name())+".key", f.MapKey(), table, comment))
					}
					fullFieldFlat = flatOneMessage(fullFieldFlat, f, prefix, table, comment)
				}
			} else {
				fullFieldFlat = flatOneMessage(fullFieldFlat, f, prefix, table, comment)
			}
		} else {
			fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name()), f, table, comment))
		}
	}
	return fullFieldFlat
}

func fieldExcelHead(prefix string, name string, fd protoreflect.FieldDescriptor, table configs.ResTableConfig, parentComment string) fieldExcelCellDesc {
	if fd.IsMap() {
		colName := fmt.Sprintf("%s{key:value}", string(fd.Name()))
		colComments := fmt.Sprintf("map{%s:%s}", fd.MapKey().Kind().String(), fd.MapValue().Kind().String())
		comment := parentComment
		if comment != "" {
			colComments = comment + "\n" + colComments
		}
		ret := fieldExcelCellDesc{
			name:    lo.If(prefix == "", colName).Else(prefix + "." + colName),
			comment: colComments,
		}
		return ret
	}
	var ns = name
	if prefix != "" {
		ns = prefix + "." + ns
	}
	var fs = fd.Kind().String()
	if fd.Kind() == protoreflect.MessageKind {
		fs = string(fd.Message().Name())
	} else if fd.Kind() == protoreflect.EnumKind {
		fs = string(fd.Enum().Name())
	}
	if fd.IsList() {
		fs = "repeated " + fs
	}
	if parentComment != "" {
		fs = parentComment + " / " + fs
	}
	return fieldExcelCellDesc{
		name:    ns,
		comment: fs,
	}
}

func flatOneMessage(fullFieldFlat []fieldExcelCellDesc, f protoreflect.FieldDescriptor, prefix string, table configs.ResTableConfig, comment string) []fieldExcelCellDesc {
	var fm = f.Message()
	if f.IsMap() {
		fm = f.MapValue().Message()
	}
	fullFieldFlat = flatFieldName(fullFieldFlat, fm, lo.If(prefix == "", string(f.Name())).Else(prefix+"."+string(f.Name())), table, comment)
	return fullFieldFlat
}
