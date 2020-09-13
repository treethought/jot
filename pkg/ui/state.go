package ui

import (
	"github.com/treethought/jot/pkg/app"
)

type State struct {
	// sync.Mutex
	currentWidget WidgetRenderer
	currentNote   *app.Note
}

func NewState() *State {
	state := &State{
		currentWidget: nil,
		currentNote: &app.Note{
			Content: "",
			Name:    "Select a note",
		},
	}
	return state
}

func (s State) SetCurrentNote(note *app.Note) {
	// s.Unlock()
	s.currentNote = note
	// s.Lock()
}

func (s State) CurrentNote() (note *app.Note) {
	return s.currentNote
}
