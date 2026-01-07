package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("4. Border + Sidebar")

	view1 := widget.NewLabel("Dashboard Content")
	view2 := widget.NewLabel("Users Content")
	content := container.NewStack(view1)

	menuData := []string{"Dashboard", "Users"}
	menu := widget.NewList(
		func() int { return len(menuData) },
		func() fyne.CanvasObject { return widget.NewLabel("Item") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(menuData[i])
		},
	)

	menu.OnSelected = func(i int) {
		if i == 0 {
			content.Objects = []fyne.CanvasObject{view1}
		} else {
			content.Objects = []fyne.CanvasObject{view2}
		}
		content.Refresh()
	}

	w.SetContent(container.NewBorder(nil, nil, menu, nil, content))
	w.ShowAndRun()
}
