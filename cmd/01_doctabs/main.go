package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("1. DocTabs")

	tabs := container.NewDocTabs(
		container.NewTabItem("Doc1", widget.NewLabel("Doc1 Content")),
		container.NewTabItem("Doc2", widget.NewLabel("Doc2 Content")),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}
