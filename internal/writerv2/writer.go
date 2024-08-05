package writerv2

import (
	"github.com/samber/lo"
	"github.com/yaoguangduan/reskeeper/internal/writex"
	"strings"
)

func parse(data writex.SheetData) {
	for _, line := range data.Lines {
		lineMap := make(map[string]interface{})
		convertLineToNestMap(line, lineMap, data.Heads)
	}
}

func convertLineToNestMap(line []string, lineMap map[string]interface{}, heads map[string]writex.ColHead) {
	selfFields := lo.PickBy(heads, func(key string, value writex.ColHead) bool {
		return !strings.Contains(key, ".")
	})
	for key, colH := range selfFields {

		if line[colH.Col] != "" {
			lineMap[key] = line[colH.Col]
		}
	}
	commonPrefixMap := make(map[string]map[string]writex.ColHead)
	for _, colHead := range heads {
		if !strings.Contains(colHead.Name, ".") {
			continue
		}
		prefix := colHead.Name[0:strings.Index(colHead.Name, ".")]
		suffix := colHead.Name[strings.Index(colHead.Name, ".")+1:]
		if _, ok := commonPrefixMap[prefix]; !ok {
			commonPrefixMap[prefix] = make(map[string]writex.ColHead)
		}
		commonPrefixMap[prefix][suffix] = writex.ColHead{Name: suffix, Col: colHead.Col, NestFields: colHead.NestFields}
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
