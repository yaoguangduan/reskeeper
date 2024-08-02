package internal

import (
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"github.com/yaoguangduan/reskeeper/internal/excelx"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/internal/writex"
	"log"
)

func Gen(pbDirList []string, excelSheetNameList []string, excelSheetGen bool) {
	protoFiles := protox.ParseProtoSet(pbDirList)
	log.Println("parse protos finish")
	config := configs.ResolveCfgFromFiles(pbDirList, protoFiles)
	if excelSheetGen {
		excelx.GenExcelFiles(config, protoFiles)
	} else {
		log.Printf("skip excel and sheet fix gen_excel:%v", excelSheetGen)
	}
	if len(excelSheetNameList) <= 0 {
		log.Printf("skip marshal data,because of marshal args empty")
		return
	}
	writex.GenerateAll(config, protoFiles, excelSheetNameList)
}
