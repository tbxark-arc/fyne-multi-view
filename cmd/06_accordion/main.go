package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("6. Accordion")

	acc := widget.NewAccordion(
		widget.NewAccordionItem("General", widget.NewLabel("General settings...")),
		widget.NewAccordionItem("Advanced", widget.NewLabel("Advanced settings...")),
		widget.NewAccordionItem("About", widget.NewLabel("Version 1.0")),
	)

	w.SetContent(acc)
	w.ShowAndRun()
}
