package ui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/treethought/jot/app"
)

type NoteList struct {
	Widget
	view  *tview.List
	notes []*app.Note
}

func NewNoteList(ui *UI, notes []*app.Note) *NoteList {
	w := &NoteList{
		Widget: Widget{ui: ui, name: "note_list"},
		notes:  notes,
	}

	w.view = tview.NewList()

	w.view.SetTitle("Notes").
		SetInputCapture(w.HandleInput)
	// .SetMouseCapture(w.H

	return w
}

func (l *NoteList) HandleInput(event *tcell.EventKey) *tcell.EventKey {
	// if event.Key() == tcell.
	key := event.Key()
	switch key {
	case tcell.KeyEnter:
		viewWidget, ok := l.ui.Widgets["view"]
		if !ok {
			panic(errors.New("View not found"))
		}
		view := viewWidget.View()
		// l.ui.app.SetRoot(view, true)
		l.ui.app.SetFocus(view)

	case tcell.KeyRune:
		switch event.Rune() {
		case 'g': // Home.
			l.view.SetCurrentItem(0)
		case 'G': // End.
			l.view.SetCurrentItem(-1)
		case 'j': // Down.
			cur := l.view.GetCurrentItem()
			l.view.SetCurrentItem(cur + 1)
		case 'k': // Up.
			cur := l.view.GetCurrentItem()
			l.view.SetCurrentItem(cur - 1)
			// case 'h': // Left.
			// 	t.columnOffset--
			// case 'l': // Right.
			// 	t.columnOffset++
		}

	}
	idx := l.view.GetCurrentItem()
	note := l.notes[idx]
	l.ui.state.SetCurrentNote(note)
	viewWidget, ok := l.ui.getWidget("view")
	if !ok {
		panic(errors.New("View not found"))
	}
	view := viewWidget.View()
	w, ok := view.(*tview.TextView)
	if !ok {
		return nil
	}
	text := fmt.Sprintf("%s\n\n%s", note.Name, note.Content)
	w.Clear().SetText(text)

	return nil
}

func (l *NoteList) SetNotes(notes []*app.Note) {
	l.notes = notes
}

func (l *NoteList) Render(grid *tview.Grid) (err error) {
	i := 0
	for _, n := range l.notes {
		firstLine := strings.Split(n.Content, "\n")[0]
		l.view.AddItem(n.Name, firstLine, rune(i), nil)
	}
	grid.AddItem(l.view, 1, 0, 1, 1, 0, 0, false)
	l.ui.app.SetFocus(l.view)

	return
}

func (w *NoteList) View() tview.Primitive {
	return w.view

}
