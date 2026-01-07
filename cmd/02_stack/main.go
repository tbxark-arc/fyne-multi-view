package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("2. Stack")

	page1 := container.NewVBox(widget.NewLabel("Page 1"))
	page2 := container.NewVBox(widget.NewLabel("Page 2"))

	stack := container.NewStack(page1, page2)
	page2.Hide() // Initially hide page2

	btn1 := widget.NewButton("Show Page 1", func() {
		page2.Hide()
		page1.Show()
	})

	btn2 := widget.NewButton("Show Page 2", func() {
		page1.Hide()
		page2.Show()
	})

	w.SetContent(container.NewBorder(
		container.NewHBox(btn1, btn2), nil, nil, nil,
		stack,
	))
	w.ShowAndRun()
}
