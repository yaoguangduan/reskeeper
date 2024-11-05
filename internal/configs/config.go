package configs

import (
	"fmt"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"os"
	"path/filepath"
)

type ResProtoFiles []ResProtoFileConfig

type ResProtoFileConfig struct {
	FullName     string
	Tables       []ResTableConfig
	ExcelPath    string
	GeneratePath string
	GenerateTags []string
	GenerateJson bool
	GenerateTxt  bool
}

func (c ResProtoFileConfig) GetMarshalFormats() []string {
	formats := []string{"bin"}
	if c.GenerateJson {
		formats = append(formats, "json")
	}
	if c.GenerateTxt {
		formats = append(formats, "txt")
	}
	return formats
}

func (c ResProtoFileConfig) GetGeneratePath() string {
	return c.GeneratePath
}

type ResTableConfig struct {
	TableName    string
	MessageName  string
	Belong       ResProtoFileConfig
	SheetName    string
	GenerateName string
}

func (t ResTableConfig) ExcelWithFieldType() bool {
	return false
}
func (t ResTableConfig) GetGenerateTags() []string {
	return t.Belong.GenerateTags
}

func (t ResTableConfig) GetExcelName() string {
	return t.Belong.ExcelPath
}

func (t ResTableConfig) GetSheetName() string {
	return t.SheetName
}

// HeadLineRangeAndSuffix 表示表头行范围和后缀
type HeadLineRangeAndSuffix map[int]string

// Table 表示一个 Excel 表格的配置
type Table struct {
	Excel          string    `toml:"excel"`
	Proto          string    `toml:"proto"`
	TableName      string    `toml:"table_name"`
	MessageName    string    `toml:"message_name"`
	SheetName      string    `toml:"sheet_name"`
	OutputName     string    `toml:"output_name"`
	OutFormats     *[]string `toml:"out_formats"`
	OutSuffixNames []string  `toml:"out_suffix_names"`
	OutFileName    *string   `toml:"out_file_name"`
}

// Config 表示整个配置文件
type Config struct {
	ProtoPath              string    `toml:"proto_path"`
	ExcelPath              string    `toml:"excel_path"`
	OutPath                string    `toml:"out_path"`
	OutFormats             *[]string `toml:"out_formats"`
	CompatMsgWithMaxFields int       `toml:"compat_message_with_max_fields"`
	Tables                 []Table   `toml:"tables"`
}

func ResolveCfgFromFiles(list []string, files protox.ProtoFiles) ResProtoFiles {
	resProtoFiles := ResProtoFiles{}
	files.RegFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		fo := fd.Options().(*descriptorpb.FileOptions)
		if !proto.HasExtension(fo, resproto.E_ResExcelPath) || !proto.HasExtension(fo, resproto.E_ResGeneratePath) {
			return true
		}
		resProtoFile := ResProtoFileConfig{}
		dir := findFileInDirectories(fd.Path(), list)
		resProtoFile.FullName = filepath.Join(dir, fd.Path())
		if proto.HasExtension(fo, resproto.E_ResExcelPath) {
			resProtoFile.ExcelPath = filepath.Join(dir, proto.GetExtension(fo, resproto.E_ResExcelPath).(string))
		}
		if proto.HasExtension(fo, resproto.E_ResGeneratePath) {
			resProtoFile.GeneratePath = filepath.Join(dir, proto.GetExtension(fo, resproto.E_ResGeneratePath).(string))
		}
		if proto.HasExtension(fo, resproto.E_ResGenerateTags) {
			resProtoFile.GenerateTags = proto.GetExtension(fo, resproto.E_ResGenerateTags).([]string)
		}
		if proto.HasExtension(fo, resproto.E_ResGenerateJson) {
			resProtoFile.GenerateJson = proto.GetExtension(fo, resproto.E_ResGenerateJson).(bool)
		}
		if proto.HasExtension(fo, resproto.E_ResGenerateTxt) {
			resProtoFile.GenerateTxt = proto.GetExtension(fo, resproto.E_ResGenerateTxt).(bool)
		}
		if proto.HasExtension(fo, resproto.E_ResExcelPath) && proto.HasExtension(fo, resproto.E_ResGeneratePath) {
			for i := 0; i < fd.Messages().Len(); i++ {
				msg := fd.Messages().Get(i)
				mo := msg.Options().(*descriptorpb.MessageOptions)
				if proto.HasExtension(mo, resproto.E_ResSheetName) && proto.HasExtension(mo, resproto.E_ResGenerateName) {
					resTable := ResTableConfig{Belong: resProtoFile}
					resTable.TableName = string(msg.Name())
					inner := msg.Fields().Get(0)
					resTable.MessageName = string(inner.Message().Name())
					resTable.SheetName = proto.GetExtension(mo, resproto.E_ResSheetName).(string)
					resTable.GenerateName = proto.GetExtension(mo, resproto.E_ResGenerateName).(string)
					resProtoFile.Tables = append(resProtoFile.Tables, resTable)
				}
			}
			resProtoFiles = append(resProtoFiles, resProtoFile)
		}
		return true
	})
	return resProtoFiles
}

// findFileInDirectories 在多个目录中查找包含特定文件的目录
func findFileInDirectories(targetFile string, directories []string) string {
	for _, dir := range directories {
		targetPath := filepath.Join(dir, targetFile)

		if _, err := os.Stat(targetPath); err == nil {
			return dir
		}
	}
	panic(fmt.Errorf("file %s not found in any of the provided directories", targetFile))
}
