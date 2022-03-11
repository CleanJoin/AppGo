package page

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Form1(myWindow fyne.Window) *fyne.Container {
	login := widget.NewEntry()
	login.SetPlaceHolder("Enter Login...")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Password...")
	input.Password = true
	//fmt.Println(i)

	content := container.NewVBox(login, input, widget.NewButton("Login", func() {
		log.Println("Content was Pass:", input.Text)
		log.Println("Content was Login:", login.Text)
	}), widget.NewButton("Register", func() {
		log.Println("Content was Pass:", input.Text)
		log.Println("Content was Login:", login.Text)
	}))

	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 170)),
		content)
	myWindow.SetContent(grid)
	myWindow.Show()
	return grid
}

// Test оповещения
// otherGroup := widget.NewCard("Other", "",
// 		widget.NewButton("Notification", func() {
// 			fyne.CurrentApp().SendNotification(&fyne.Notification{
// 				Title:   "Fyne Demo",
// 				Content: "Testing notifications...",
// 			})
