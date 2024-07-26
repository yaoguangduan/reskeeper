package internal

import (
	"reskeeper/internal/excelx"
	"reskeeper/internal/protox"
	"reskeeper/internal/writex"
	"reskeeper/res_toml"
)

func Gen(tml string) {
	config := res_toml.ParseConfig(tml)
	protoFiles := protox.ParseProtoSet(config)
	excelx.GenExcelFiles(config, protoFiles)
	writex.GenerateAll(config, protoFiles)
}
