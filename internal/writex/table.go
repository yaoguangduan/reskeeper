package writex

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
)

type ColHead struct {
	Name       string
	Col        int
	Additional string
}

type SheetTable struct {
	Heads map[string]ColHead
	Data  [][]string
}

func ParseToSheetTable(data [][]string) SheetTable {
	st := SheetTable{Heads: make(map[string]ColHead), Data: [][]string{}}

	useful := lo.Filter(data, func(item []string, index int) bool {
		return len(item) > 0 && !strings.HasPrefix(item[0], "#")
	})
	for col, cell := range useful[0] {
		if cell != "" {
			var idx = strings.Index(cell, "{")
			idx = lo.If(idx == -1, len(cell)).Else(idx)
			st.Heads[cell[0:idx]] = ColHead{Name: cell[0:idx], Col: col, Additional: cell[idx:]}
		}
	}
	useful = useful[1:]
	for i := range useful {
		if len(useful[i]) < len(st.Heads) {
			useful[i] = append(useful[i], make([]string, len(st.Heads)-len(useful[i]))...)
		}
	}
	st.Data = useful
	fmt.Println(st.Data)
	return st
}

func removeIgnoreColAndGetHeadLines(data [][]string) ([][]string, map[int]string) {
	if len(data) == 0 {
		return data, map[int]string{}
	}

	var colsToRemove = make([]int, 1)
	var headLines = make(map[int]string)
	for colIndex, value := range data[0] {
		if strings.HasPrefix(value, "#") {
			colsToRemove = append(colsToRemove, colIndex)
		}
	}
	var newData [][]string
	for idx, row := range data {
		if strings.HasPrefix(row[0], "$head:") {
			headLines[idx] = row[0][6:]
		}
		var newRow []string
		for colIndex, value := range row {
			if !slices.Contains(colsToRemove, colIndex) {
				newRow = append(newRow, value)
			}
		}
		newData = append(newData, newRow)
	}
	return newData, headLines
}
