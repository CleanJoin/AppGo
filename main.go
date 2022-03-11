package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Chat")

	centered, text := page.Form2(myWindow)
	// d1 := makeSplitTab(myWindow)
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), centered, text))
	// myWindow.SetContent(d1)
	myWindow.ShowAndRun()

}
