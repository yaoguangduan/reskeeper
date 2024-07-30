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
	NestFields string
}

type SheetData struct {
	Heads map[string]ColHead
	Lines [][]string
}

func ParseToSheetTable(data [][]string) SheetData {
	st := SheetData{Heads: make(map[string]ColHead), Lines: [][]string{}}

	useful := lo.Filter(data, func(item []string, index int) bool {
		return len(item) > 0 && !strings.HasPrefix(item[0], "#")
	})
	for col, cell := range useful[0] {
		if cell != "" {
			ch := parseCellValueToColHead(cell)
			ch.Col = col
			st.Heads[ch.Name] = ch
		}
	}
	useful = useful[1:]
	for i := range useful {
		if len(useful[i]) < len(st.Heads) {
			useful[i] = append(useful[i], make([]string, len(st.Heads)-len(useful[i]))...)
		}
	}
	st.Lines = useful
	fmt.Println(st.Lines)
	return st
}

func parseCellValueToColHead(cell string) ColHead {
	var idx = strings.Index(cell, "{")
	if idx == -1 {
		idx = strings.Index(cell, "(")
		if idx == -1 {
			return ColHead{Name: cell}
		} else {
			return ColHead{Name: cell[0:idx]}
		}
	} else {
		name := cell[0:idx]
		var nf = cell[idx+1:]
		if strings.Contains(nf, "(") {
			for strings.Contains(nf, "(") {
				b := strings.Index(nf, "(")
				e := strings.Index(nf, ")")
				nf = nf[0:b] + nf[e+1:]
			}
			return ColHead{Name: name, Col: idx, NestFields: nf}
		} else {
			return ColHead{Name: name, Col: idx, NestFields: nf}
		}
	}
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
