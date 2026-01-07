package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("7. Main Window")

	mainView := widget.NewLabel("This is the Main Window")
	btn := widget.NewButton("Open Settings Window", func() {
		sw := a.NewWindow("Settings")
		sw.SetContent(widget.NewLabel("Settings Window Content"))
		sw.Resize(fyne.NewSize(300, 200))
		sw.Show()
	})

	w.SetContent(container.NewVBox(mainView, btn))
	w.ShowAndRun()
}
