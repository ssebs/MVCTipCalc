package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ssebs/MVCTipCalc/internal"
)

const TITLE = "MVC Tip Calc"

func main() {
	// fmt.Println(TITLE)
	myApp := app.New()
	win := myApp.NewWindow(TITLE)

	tipModel := internal.NewTipModel()
	tipView := internal.NewTipView()
	tipController := internal.NewTipController(tipModel, tipView)

	win.SetContent(tipController.TipView)

	win.Resize(fyne.NewSquareSize(400))
	win.CenterOnScreen()
	win.Show()
	myApp.Run()
}
