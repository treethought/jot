package ui

import (
	"github.com/rivo/tview"
)

type Widget struct {
	ui   *UI
	name string
}

func (w *Widget) Name() string {
	return w.name
}

type WidgetRenderer interface {
	Render(grid *tview.Grid) error
	View() tview.Primitive
	Name() string
}
