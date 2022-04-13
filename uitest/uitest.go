package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	text1 := canvas.NewText("Hello", color.Black)
	text2 := canvas.NewText("World", color.Black)
	text3 := canvas.NewText("!", color.Black)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	w.SetContent(content)
	w.ShowAndRun()
}
