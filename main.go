package main

import (
	"appGo/page"

	"fyne.io/fyne/v2/app"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Chat")
	page.Form1(myWindow, myApp)
	// centered, text := page.Form2(myWindow)
	// d1 := makeSplitTab(myWindow)
	// myWindow.SetContent(container.New(layout.NewVBoxLayout(), centered, text))

	myApp.Run()

}
