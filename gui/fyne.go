package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"strconv"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Layout Example")

	// 示例数据绑定
	protoDirStr := binding.NewString()
	excelSheetList := binding.NewStringList()
	outputContent := binding.NewString()
	// 创建组件
	label := widget.NewLabel("input proto dirs")
	entry := widget.NewEntryWithData(protoDirStr)
	loadButton := &widget.Button{
		Text:       "Load",
		Icon:       theme.ViewRefreshIcon(),
		Importance: widget.HighImportance,
		OnTapped:   func() { fmt.Println("high importance button") },
	}
	topRow := container.NewBorder(nil, nil, label, loadButton, entry)

	list := widget.NewListWithData(excelSheetList,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	entry1 := widget.NewMultiLineEntry()
	entry1.Disable()
	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * 1500)
			outputContent.Set(lo.Must(outputContent.Get()) + "\r\n" + strconv.Itoa(i))
		}
	}()
	entry1.Wrapping = fyne.TextWrapWord
	entry1.Bind(outputContent)
	rightContent := container.NewBorder(container.NewVBox(container.NewHBox(
		&widget.Label{
			Text:       "Setss",
			Importance: widget.HighImportance,
		},
		widget.NewButton("gen excel", func() {

		}),
		widget.NewButton("gen sheet", func() {

		}),
		widget.NewButton("marshal", func() {

		}),
	), widget.NewSeparator()), nil, nil, nil, entry1)
	split := container.NewHSplit(
		container.NewScroll(list),
		rightContent,
	)
	split.Offset = 0.35 // 设置分割比例

	// 使用 Border 布局让组件充满窗口
	mainContainer := container.NewBorder(container.NewVBox(topRow, widget.NewSeparator()), nil, nil, nil, split)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(900, 600))
	myWindow.ShowAndRun()
}
