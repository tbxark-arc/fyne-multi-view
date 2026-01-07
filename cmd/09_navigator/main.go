package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ViewManager manages switching between different views
type ViewManager struct {
	host  *fyne.Container
	views map[string]fyne.CanvasObject
}

// NewViewManager creates a new ViewManager
func NewViewManager() *ViewManager {
	return &ViewManager{
		host:  container.NewStack(),
		views: map[string]fyne.CanvasObject{},
	}
}

// Register adds a new view with a name
func (vm *ViewManager) Register(name string, v fyne.CanvasObject) {
	vm.views[name] = v
}

// Show switches the display to the named view
func (vm *ViewManager) Show(name string) {
	if v, ok := vm.views[name]; ok {
		vm.host.Objects = []fyne.CanvasObject{v}
		vm.host.Refresh()
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("9. Navigator / ViewManager")

	vm := NewViewManager()
	vm.Register("home", widget.NewLabel("Home View"))
	vm.Register("settings", widget.NewLabel("Settings View"))
	vm.Show("home")

	btnHome := widget.NewButton("Home", func() { vm.Show("home") })
	btnSet := widget.NewButton("Settings", func() { vm.Show("settings") })

	w.SetContent(container.NewBorder(nil, nil, container.NewVBox(btnHome, btnSet), nil, vm.host))
	w.ShowAndRun()
}
