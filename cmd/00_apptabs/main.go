package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("0. AppTabs")

	tabs := container.NewAppTabs(
		container.NewTabItem("Home", widget.NewLabel("Home View")),
		container.NewTabItem("Settings", widget.NewLabel("Settings View")),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}
