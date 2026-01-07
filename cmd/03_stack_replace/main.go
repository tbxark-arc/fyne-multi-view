package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("3. Stack (Object Replacement)")

	viewA := widget.NewLabel("View A")
	viewB := widget.NewLabel("View B")

	content := container.NewStack(viewA)

	btn := widget.NewButton("Switch", func() {
		if content.Objects[0] == viewA {
			content.Objects = []fyne.CanvasObject{viewB}
		} else {
			content.Objects = []fyne.CanvasObject{viewA}
		}
		content.Refresh()
	})

	w.SetContent(container.NewBorder(btn, nil, nil, nil, content))
	w.ShowAndRun()
}
