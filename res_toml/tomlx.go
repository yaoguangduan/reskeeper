package res_toml

import "github.com/BurntSushi/toml"

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

func ParseConfig(file string) Config {
	var config Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		panic(err)
	}
	return config
}
