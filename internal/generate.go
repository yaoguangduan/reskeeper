package internal

import (
	"fmt"
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"github.com/yaoguangduan/reskeeper/internal/excelx"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/internal/writex"
)

func Gen(pbDirList []string) {
	protoFiles := protox.ParseProtoSet(pbDirList)
	config := configs.ResolveCfgFromFiles(pbDirList, protoFiles)
	fmt.Println(config)
	excelx.GenExcelFiles(config, protoFiles)
	writex.GenerateAll(config, protoFiles)
}
