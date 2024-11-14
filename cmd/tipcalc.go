package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ssebs/MVCTipCalc/internal"
)

const TITLE = "MVC Tip Calc"

func main() {
	fmt.Println(TITLE)
	myApp := app.New()
	win := myApp.NewWindow(TITLE)

	// tipModel := internal.NewTipModel()
	tipView := internal.NewTipView()

	win.SetContent(
		tipView,
	)
	win.Resize(fyne.NewSquareSize(200))
	win.CenterOnScreen()
	win.Show()
	myApp.Run()
}
