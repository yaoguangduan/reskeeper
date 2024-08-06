package writerv2

import (
	"github.com/samber/lo"
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"strings"
)

func parse(data configs.SheetData) {
	for _, line := range data.Lines {
		lineMap := make(map[string]interface{})
		convertLineToNestMap(line, lineMap, data.Heads)
	}
}

func convertLineToNestMap(line []string, lineMap map[string]interface{}, heads map[string]configs.ColHead) {
	selfFields := lo.PickBy(heads, func(key string, value configs.ColHead) bool {
		return !strings.Contains(key, ".")
	})
	for key, colH := range selfFields {

		if line[colH.Col] != "" {
			lineMap[key] = line[colH.Col]
		}
	}
	commonPrefixMap := make(map[string]map[string]configs.ColHead)
	for _, colHead := range heads {
		if !strings.Contains(colHead.Name, ".") {
			continue
		}
		prefix := colHead.Name[0:strings.Index(colHead.Name, ".")]
		suffix := colHead.Name[strings.Index(colHead.Name, ".")+1:]
		if _, ok := commonPrefixMap[prefix]; !ok {
			commonPrefixMap[prefix] = make(map[string]configs.ColHead)
		}
		commonPrefixMap[prefix][suffix] = configs.ColHead{Name: suffix, Col: colHead.Col, NestFields: colHead.NestFields}
	}
	for prefix, ncm := range commonPrefixMap {
		var hasValue = false
		for _, col := range ncm {
			if line[col.Col] != "" {
				hasValue = true
				break
			}
		}
		if !hasValue {
			continue
		}
		subMap := make(map[string]interface{})
		convertLineToNestMap(line, subMap, ncm)
		lineMap[prefix] = subMap
	}
}
