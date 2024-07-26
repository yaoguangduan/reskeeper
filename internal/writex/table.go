package writex

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
)

type SheetTable struct {
	Heads map[string]map[string]int
	Data  [][]string
}

func ParseToSheetTable(data [][]string) SheetTable {
	st := SheetTable{Heads: map[string]map[string]int{}, Data: [][]string{}}
	var maxHeadSize = 0
	rows, headLines := removeIgnoreColAndGetHeadLines(data)
	ignoreRows := make([]int, 0)
	for idx, row := range rows {
		if len(row) <= 0 || idx == 0 {
			ignoreRows = append(ignoreRows, idx)
			continue
		}
		if slices.Contains(lo.Keys(headLines), idx) {
			head := make(map[string]int)
			for col, cell := range row {
				if cell != "" {
					head[cell] = col
				}
			}
			maxHeadSize = max(maxHeadSize, len(head))
			st.Heads[headLines[idx]] = head
			ignoreRows = append(ignoreRows, idx)
		}
		if strings.HasPrefix(row[0], "#") {
			ignoreRows = append(ignoreRows, idx)
		}
	}
	useful := lo.Filter(rows, func(item []string, index int) bool {
		return !slices.Contains(ignoreRows, index)
	})
	for i := range useful {
		if len(useful[i]) < maxHeadSize {
			useful[i] = append(useful[i], make([]string, maxHeadSize-len(useful[i]))...)
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
