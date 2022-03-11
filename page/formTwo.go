package page

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list"}

var data2 = []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}

func Form2(myWindow fyne.Window) (*fyne.Container, *fyne.Container) {

	list1 := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	list2 := widget.NewList(
		func() int {
			return len(data2)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data2[i])
		})

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Texy...")
	input.Password = true
	//fmt.Println(i)
	textForm := container.NewVBox(input, widget.NewButton("Login", func() {
		log.Println("Content was Pass:", input.Text)
		log.Println("Content was Login:", input.Text)
	}))

	content := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 800)), list1)
	centeredMessage := container.New(layout.NewGridWrapLayout(fyne.NewSize(1630, 800)), list2)
	content2 := container.New(layout.NewHBoxLayout(), content, layout.NewSpacer(), centeredMessage)

	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(1900, 100)),
		textForm)

	centeredFormMessage := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), grid, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content2, centeredFormMessage))
	myWindow.Show()
	return content2, centeredFormMessage
}

func makeSplitTab(_ fyne.Window) fyne.CanvasObject {
	left := widget.NewMultiLineEntry()
	left.Wrapping = fyne.TextWrapWord
	left.SetText("Long text is looooooooooooooong")
	right := container.NewVSplit(
		widget.NewLabel("Label"),
		widget.NewButton("Button", func() { fmt.Println("button tapped!") }),
	)
	return container.NewHSplit(container.NewVScroll(left), right)
}
