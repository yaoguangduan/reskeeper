package excelx

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
	"google.golang.org/protobuf/reflect/protoreflect"
	"os"
	"path/filepath"
	"reskeeper/internal/excelx/styles"
	"reskeeper/internal/protox"
	"reskeeper/internal/tools"
	"reskeeper/res_toml"
)

func GenExcelFiles(config res_toml.Config, files protox.ProtoFiles) {
	excelPath := config.ExcelPath
	_, err := os.Stat(excelPath)
	if err != nil && os.IsNotExist(err) {
		panic(fmt.Sprintf("table path %s does not exist", excelPath))
	}
	if err != nil {
		panic(err)
	}
	excelTableMap := make(map[string][]res_toml.Table)
	for _, table := range config.Tables {
		_, exist := excelTableMap[table.Excel]
		if !exist {
			excelTableMap[table.Excel] = []res_toml.Table{}
		}
		excelTableMap[table.Excel] = append(excelTableMap[table.Excel], table)
	}
	for excel := range excelTableMap {
		fp := filepath.Join(excelPath, excel)
		_, err := os.Stat(fp)
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("table file %s state err", fp))
		}
		if err != nil {
			//文件不存在
			newExcelWithConfig(excel, excelTableMap[excel], config, files)
		} else {
			//文件存在
			adjustExcelWithConfig(excel, excelTableMap[excel], config, files)
		}
	}
}

func adjustExcelWithConfig(excel string, tables []res_toml.Table, config res_toml.Config, files protox.ProtoFiles) {
	file, err := excelize.OpenFile(filepath.Join(config.ExcelPath, excel))
	if err != nil {
		panic(err)
	}
	defer func() {
		lo.Must0(file.Close())
	}()
	for _, table := range tables {
		var index = lo.Must(file.GetSheetIndex(table.SheetName))
		if index == -1 {
			index = lo.Must(file.GetSheetIndex("#" + table.SheetName))
		}
		if index == -1 {
			createTableSheetOnExcel(file, table, files, config)
		} else {
			refreshTableSheetComment(file, table)
		}
	}
}

func newExcelWithConfig(excel string, tables []res_toml.Table, config res_toml.Config, files protox.ProtoFiles) {
	fp := filepath.Join(config.ExcelPath, excel)
	excelFile := excelize.NewFile()

	defer func() {
		_ = excelFile.DeleteSheet("Sheet1")
		if err := excelFile.SaveAs(fp); err != nil {
			panic(err)
		}
	}()
	for _, table := range tables {
		createTableSheetOnExcel(excelFile, table, files, config)
	}
}

func createTableSheetOnExcel(excelFile *excelize.File, table res_toml.Table, files protox.ProtoFiles, config res_toml.Config) {
	_, err := excelFile.NewSheet(table.SheetName)
	if err != nil {
		panic(err)
	}
	msgD := files.GetMessage(table.MessageName)

	refreshTableSheetComment(excelFile, table)
	fullFieldFlat := make([]interface{}, 0)
	fullFieldFlat = flatFieldName(fullFieldFlat, msgD, "", config)
	lo.Must0(excelFile.SetSheetRow(table.SheetName, "B2", &fullFieldFlat))
	lo.Must0(excelFile.SetCellStr(table.SheetName, "A2", "$head:full"))

	lo.Must0(excelFile.SetRowStyle(table.SheetName, 2, 65535, styles.FontAlignCenter(excelFile)))
	lo.Must0(excelFile.SetCellStyle(table.SheetName, "A2", "A3", styles.FontKeywords(excelFile)))

	styles.AdjustColumnWidth(excelFile, table.SheetName, len(fullFieldFlat)+1)
}

func refreshTableSheetComment(excelFile *excelize.File, table res_toml.Table) {
	fw := tools.NewFileWriter()
	fw.PL("the first line and first column is for control usage")
	fw.PL("the line or column will be ignore if it's first cell value start with #")
	fw.PL("cell value usage for generate if the cell in first line or column and start with $:")
	fw.PL("$head is the message's fields def and generate name suffix")
	fw.PLF("protoName:%s", table.Proto)
	fw.PLF("tableName:%s", table.TableName)
	fw.PLF("messageMame:%s", table.MessageName)
	lo.Must0(excelFile.SetCellStr(table.SheetName, "A1", fw.String()))
	// 方法1：自动换行
	style, err := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{WrapText: true},
		Font:      &excelize.Font{Italic: true},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	lo.Must(excelFile.GetStyle(lo.Must(excelFile.GetCellStyle(table.SheetName, "A1"))))
	rh := lo.Must(excelFile.GetRowHeight(table.SheetName, 1))
	lo.Must0(excelFile.SetCellStyle(table.SheetName, "A1", "A1", style))
	lo.Must0(excelFile.SetRowHeight(table.SheetName, 1, rh))
}

func flatFieldName(fullFieldFlat []interface{}, msgD protoreflect.MessageDescriptor, prefix string, config res_toml.Config) []interface{} {
	for i := 0; i < msgD.Fields().Len(); i++ {
		f := msgD.Fields().Get(i)
		if f.Kind() == protoreflect.MessageKind {
			if f.IsMap() {
				fullFieldFlat = append(fullFieldFlat, string(f.Name())+".mapKey")
				mapValDesc := f.MapValue()
				if mapValDesc.Kind() == protoreflect.MessageKind {
					fullFieldFlat = flatFieldName(fullFieldFlat, mapValDesc.Message(), lo.If(prefix == "", string(f.Name())).Else(prefix+"."+string(f.Name())), config)
				} else {
					fullFieldFlat = append(fullFieldFlat, string(f.Name())+".mapVal")
				}
			} else {
				fm := f.Message()
				var noMsgMapList = true
				for j := 0; j < fm.Fields().Len(); j++ {
					fmf := fm.Fields().Get(j)
					if fmf.IsMap() || fmf.IsList() || fmf.Kind() == protoreflect.MessageKind {
						noMsgMapList = false
						break
					}
				}
				if noMsgMapList && fm.Fields().Len() <= config.CompatMsgWithMaxFields {
					fullFieldFlat = append(fullFieldFlat, lo.If(prefix == "", string(f.Name())).Else(prefix+"."+string(f.Name())))
				} else {
					fullFieldFlat = flatFieldName(fullFieldFlat, f.Message(), lo.If(prefix == "", string(f.Name())).Else(prefix+"."+string(f.Name())), config)
				}
			}
		} else {
			fullFieldFlat = append(fullFieldFlat, lo.If(prefix == "", string(f.Name())).Else(prefix+"."+string(f.Name())))
		}
	}
	return fullFieldFlat
}
