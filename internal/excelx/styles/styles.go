package styles

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
	"math"
)

func FontKeywords(f *excelize.File) int {
	style, err := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Color: "#cb6028"},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return style
}

func FontIdentifier(f *excelize.File) int {
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Font:      &excelize.Font{Color: "#6fafbd"},
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return style
}

func FontAlignCenter(f *excelize.File) int {
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return style
}

func FontBold(f *excelize.File) int {
	style, err := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return style
}

func AdjustColumnWidth(f *excelize.File, sheetName string, colCnt int) {
	maxWidth := 8.5
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i < colCnt+1; i++ {
		cellValue := rows[0][i-1]
		cellWidth := calculateWidth(cellValue)
		if cellWidth > maxWidth {
			maxWidth = cellWidth
		}
		f.SetColWidth(sheetName, lo.Must(excelize.ColumnNumberToName(i)), lo.Must(excelize.ColumnNumberToName(i)), maxWidth)
	}
}

// calculateWidth 根据字符数和默认字体大小计算列宽
func calculateWidth(text string) float64 {
	// 假设默认字体大小为 11 磅，使用 Calibri 字体
	defaultFontSize := 11.0
	// 估算每个字符的宽度（可以根据实际情况调整）
	charWidth := defaultFontSize * 0.8
	// 计算文本的总宽度
	textWidth := float64(len(text)) * charWidth
	// Excel 列宽单位转换
	excelColWidth := textWidth / 7.0
	// 返回列宽，保留一位小数
	return math.Ceil(excelColWidth*10) / 10
}
