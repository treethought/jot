package ui

import (
	"github.com/rivo/tview"
)

type Widget struct {
	ui *UI
}

type WidgetRenderer interface {
	Render(grid *tview.Grid) error
	View() tview.Primitive
}
