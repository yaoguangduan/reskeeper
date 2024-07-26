package internal

import (
	"github.com/yaoguangduan/reskeeper/internal/excelx"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/internal/writex"
	"github.com/yaoguangduan/reskeeper/res_toml"
)

func Gen(tml string) {
	config := res_toml.ParseConfig(tml)
	protoFiles := protox.ParseProtoSet(config)
	excelx.GenExcelFiles(config, protoFiles)
	writex.GenerateAll(config, protoFiles)
}
