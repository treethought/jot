package ui

import (
	"fmt"
	"os"
	"os/exec"

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
	w.view.SetInputCapture(w.HandleInput)
	w.view.SetDynamicColors(true)
	w.view.SetChangedFunc(func() {
		ui.app.Draw()
	})
	w.view.Highlight("0").ScrollToHighlight()

	// regions := w.view.GetRegionText(regionID string)
	w.view.SetBackgroundColor(tcell.ColorDarkCyan)

	w.note = ui.state.CurrentNote()

	return w
}

func (w *NoteView) Render(grid *tview.Grid) (err error) {
	w.view.Clear().SetText(w.note.Name)
	grid.AddItem(w.view, 1, 1, 2, 2, 0, 0, false)
	// w.view.ScrollToBeginning()
	return nil
}

func (w *NoteView) HandleInput(event *tcell.EventKey) *tcell.EventKey {
	// if event.Key() == tcell.
	key := event.Key()
	switch key {

	case tcell.KeyRune:
		switch event.Rune() {
		case 'e': // Home.
			w.ui.app.Suspend(func() {
				path := fmt.Sprintf("notetest/%s", w.ui.state.CurrentNote())
				cmd := exec.Command("vim", path)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					return
				}
			})
		}

	}

	return nil
}

func (w *NoteView) View() tview.Primitive {
	return w.view

}
