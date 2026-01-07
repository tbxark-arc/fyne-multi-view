package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("8. Dialog")

	btn := widget.NewButton("Open Dialog", func() {
		d := dialog.NewCustom("Info", "Close", widget.NewLabel("Hello from Dialog!"), w)
		d.Show()
	})
	w.SetContent(container.NewCenter(btn))
	w.ShowAndRun()
}
