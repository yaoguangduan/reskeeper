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
	"strings"
)

type ResProtoFiles []ResProtoFileConfig

type ResProtoFileConfig struct {
	FullName string
	Opt      *resproto.ResourceFileOpt
	Tables   []ResTableConfig
}
type ResTableConfig struct {
	TableName   string
	MessageName string
	Opt         *resproto.ResourceTableOpt
	Belong      ResProtoFileConfig
}

func (t ResTableConfig) GetExcelName() string {
	return t.Opt.GetExcelAndSheetName()[0:strings.Index(t.Opt.GetExcelAndSheetName(), "#")]
}

func (t ResTableConfig) GetSheetName() string {
	return t.Opt.GetExcelAndSheetName()[strings.Index(t.Opt.GetExcelAndSheetName(), "#")+1:]
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
		if proto.HasExtension(fo, resproto.E_ResFileOpt) {
			rfOpt := proto.GetExtension(fo, resproto.E_ResFileOpt).(*resproto.ResourceFileOpt)
			resProtoFile := ResProtoFileConfig{FullName: string(fd.FullName())}
			resProtoFile.Opt = rfOpt
			for i := 0; i < fd.Messages().Len(); i++ {
				msg := fd.Messages().Get(i)
				mo := msg.Options().(*descriptorpb.MessageOptions)
				if !proto.HasExtension(mo, resproto.E_ResTableOpt) {
					continue
				}
				resTable := ResTableConfig{Belong: resProtoFile}
				resTable.TableName = string(msg.Name())
				resTable.Opt = proto.GetExtension(mo, resproto.E_ResTableOpt).(*resproto.ResourceTableOpt)
				inner := msg.Fields().ByNumber(1)
				resTable.MessageName = string(inner.Message().Name())
				resProtoFile.Tables = append(resProtoFile.Tables, resTable)
			}
			fmt.Println(fd.Path())
			toPath := filepath.Join(findFileInDirectories(fd.Path(), list), *resProtoFile.Opt.ExcelPath)

			resProtoFile.Opt.ExcelPath = proto.String(toPath)
			resProtoFile.Opt.MarshalPath = proto.String(filepath.Join(findFileInDirectories(fd.Path(), list), *resProtoFile.Opt.MarshalPath))
			resProtoFiles = append(resProtoFiles, resProtoFile)
		}
		return true
	})
	return resProtoFiles
}

func GetMsgOpt(fm protoreflect.MessageDescriptor) *resproto.ResourceMsgOpt {
	if proto.HasExtension(fm.Options(), resproto.E_ResMsgOpt) {
		return proto.GetExtension(fm.Options(), resproto.E_ResMsgOpt).(*resproto.ResourceMsgOpt)
	}
	return nil
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
