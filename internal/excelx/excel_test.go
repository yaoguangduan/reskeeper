package excelx

import (
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestExcel(t *testing.T) {
	f := excelize.NewFile()
	f.NewSheet("sheet222")
	f.NewSheet("sheet33")
	f.SaveAs("test.xlsx")
}
