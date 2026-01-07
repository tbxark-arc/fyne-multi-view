package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("5. Split")

	left := widget.NewMultiLineEntry()
	left.SetText("Left panel content...")

	right := widget.NewLabel("Right")
	split := container.NewHSplit(left, right)
	split.Offset = 0.3

	w.SetContent(split)
	w.ShowAndRun()
}
