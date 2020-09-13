package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/treethought/jot/pkg/app"
)

type NoteView struct {
	Widget
	view *tview.TextView
	note *app.Note
}

func NewNoteView(ui *UI) *NoteView {
	w := &NoteView{
		Widget: Widget{ui: ui},
	}

	w.view = tview.NewTextView().SetScrollable(true)
	w.view.SetBackgroundColor(tcell.ColorDarkCyan)

	w.note = ui.state.CurrentNote()

	return w
}

func (w *NoteView) Render(grid *tview.Grid) (err error) {
	w.view.Clear().SetText(w.note.Name)
	grid.AddItem(w.view, 1, 1, 2, 2, 0, 0, false)
	return nil
}

func (w *NoteView) View() tview.Primitive {
	return w.view

}
