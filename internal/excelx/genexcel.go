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
	"os"
	"path/filepath"
	"strings"
)

func GenExcelFiles(configs []configs.ResProtoFileConfig, files protox.ProtoFiles) {
	for _, config := range configs {
		genByOneConfigProto(config, files)
	}
}

func genByOneConfigProto(config configs.ResProtoFileConfig, files protox.ProtoFiles) {
	excelPath := config.Opt.GetExcelPath()
	_, err := os.Stat(excelPath)
	if err != nil && os.IsNotExist(err) {
		panic(fmt.Sprintf("table path %s does not exist", excelPath))
	}
	if err != nil {
		panic(err)
	}
	excelTableMap := make(map[string][]configs.ResTableConfig)
	for _, table := range config.Tables {
		_, exist := excelTableMap[table.GetExcelName()]
		if !exist {
			excelTableMap[table.GetExcelName()] = []configs.ResTableConfig{}
		}
		excelTableMap[table.GetExcelName()] = append(excelTableMap[table.GetExcelName()], table)
	}
	for excelName, tableConfigs := range excelTableMap {
		fp := filepath.Join(excelPath, excelName)
		_, err := os.Stat(fp)
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("table file %s state err", fp))
		}
		if err != nil {
			//文件不存在
			newExcelWithConfig(excelName, tableConfigs, config, files)
		} else {
			//文件存在
			adjustExcelWithConfig(excelName, tableConfigs, config, files)
		}
	}
}

func adjustExcelWithConfig(excel string, tables []configs.ResTableConfig, config configs.ResProtoFileConfig, files protox.ProtoFiles) {
	file, err := excelize.OpenFile(filepath.Join(config.Opt.GetExcelPath(), excel))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			lo.Must0(file.Save())
		}
	}()
	for _, table := range tables {
		var index = lo.Must(file.GetSheetIndex(table.GetSheetName()))
		if index == -1 {
			index = lo.Must(file.GetSheetIndex("#" + table.GetSheetName()))
		}
		if index == -1 {
			createTableSheetOnExcel(file, table, files, config)
		}
	}
}

func newExcelWithConfig(excel string, tables []configs.ResTableConfig, config configs.ResProtoFileConfig, files protox.ProtoFiles) {
	fp := filepath.Join(config.Opt.GetExcelPath(), excel)
	excelFile := excelize.NewFile()

	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			_ = excelFile.DeleteSheet("Sheet1")
			if err = excelFile.SaveAs(fp); err != nil {
				panic(err)
			}
		}
	}()
	for _, table := range tables {
		createTableSheetOnExcel(excelFile, table, files, config)
	}
}

func createTableSheetOnExcel(excelFile *excelize.File, table configs.ResTableConfig, files protox.ProtoFiles, config configs.ResProtoFileConfig) {
	_, err := excelFile.NewSheet(table.GetSheetName())
	if err != nil {
		panic(err)
	}
	msgD := files.GetMessage(table.MessageName)

	fullFieldFlat := make([]interface{}, 0)
	fullFieldFlat = flatFieldName(fullFieldFlat, msgD, "", table)

	lo.Must0(excelFile.SetSheetRow(table.GetSheetName(), "A1", &fullFieldFlat))
	lo.Must0(excelFile.SetRowStyle(table.GetSheetName(), 1, 65535, styles.FontAlignCenter(excelFile)))
	lo.Must0(excelFile.SetCellStyle(table.GetSheetName(), "A1", fmt.Sprintf("%s1", lo.Must(excelize.ColumnNumberToName(len(fullFieldFlat)))), styles.FontBold(excelFile)))
	styles.AdjustColumnWidth(excelFile, table.GetSheetName(), len(fullFieldFlat))
}

func flatFieldName(fullFieldFlat []interface{}, msgD protoreflect.MessageDescriptor, prefix string, table configs.ResTableConfig) []interface{} {
	for i := 0; i < msgD.Fields().Len(); i++ {
		f := msgD.Fields().Get(i)
		if f.Kind() == protoreflect.MessageKind {
			if f.IsMap() {
				var oneColumn = false
				if proto.HasExtension(f.Options(), resproto.E_ResOneColumn) && proto.GetExtension(f.Options(), resproto.E_ResOneColumn).(bool) {
					oneColumn = true
				}
				if f.MapValue().Kind() != protoreflect.MessageKind {
					if oneColumn {
						var colName = fmt.Sprintf("%s{map-key;map-val}", string(f.Name()))
						if table.ExcelWithFieldType() {
							colName = fmt.Sprintf("%s{map-key(%s);map-val(%s)}", string(f.Name()), f.MapKey().Kind().String(), f.MapValue().Kind().String())
						}
						fullFieldFlat = append(fullFieldFlat, lo.If(prefix == "", colName).Else(prefix+"."+colName))
					} else {
						fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name())+".map-key", f.MapKey(), table))
						fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name())+".map-val", f.MapValue(), table))
					}
				} else {
					if !proto.HasExtension(f.Options(), resproto.E_ResUseMsgKey) || !proto.GetExtension(f.Options(), resproto.E_ResUseMsgKey).(bool) {
						fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name())+".map-key", f.MapKey(), table))
					}
					fullFieldFlat = flatOneMessage(fullFieldFlat, f, prefix, table)
				}
			} else {
				fullFieldFlat = flatOneMessage(fullFieldFlat, f, prefix, table)
			}
		} else {
			fullFieldFlat = append(fullFieldFlat, fieldExcelHead(prefix, string(f.Name()), f, table))
		}
	}
	return fullFieldFlat
}

func fieldExcelHead(prefix string, name string, fd protoreflect.FieldDescriptor, table configs.ResTableConfig) string {
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
	if table.ExcelWithFieldType() {
		return ns + "(" + fs + ")"
	} else {
		return ns
	}
}

func flatOneMessage(fullFieldFlat []interface{}, f protoreflect.FieldDescriptor, prefix string, table configs.ResTableConfig) []interface{} {
	var fm = f.Message()
	if f.IsMap() {
		fm = f.MapValue().Message()
	}
	var noMsgMapList = true
	fields := make([]string, 0)
	for j := 0; j < fm.Fields().Len(); j++ {
		fmf := fm.Fields().Get(j)
		if fmf.IsMap() || fmf.IsList() || fmf.Kind() == protoreflect.MessageKind {
			noMsgMapList = false
			break
		}
		if table.ExcelWithFieldType() {
			fields = append(fields, fieldExcelHead("", string(fmf.Name()), fmf, table))
		} else {
			fields = append(fields, string(fmf.Name()))
		}
	}
	if noMsgMapList && proto.HasExtension(f.Options(), resproto.E_ResOneColumn) && proto.GetExtension(f.Options(), resproto.E_ResOneColumn).(bool) {
		var quoteName = string(f.Name()) + "{" + strings.Join(fields, ";") + "}"
		if f.IsMap() && (!proto.HasExtension(f.Options(), resproto.E_ResUseMsgKey) || !proto.GetExtension(f.Options(), resproto.E_ResUseMsgKey).(bool)) {
			quoteName = string(f.Name()) + ".map-val{" + strings.Join(fields, ";") + "}"
		}
		fullFieldFlat = append(fullFieldFlat, lo.If(prefix == "", quoteName).Else(prefix+"."+quoteName))
	} else {
		fullFieldFlat = flatFieldName(fullFieldFlat, fm, lo.If(prefix == "", string(f.Name())).Else(prefix+"."+string(f.Name())), table)
	}
	return fullFieldFlat
}
