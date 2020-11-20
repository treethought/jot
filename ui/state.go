package ui

import (
	"github.com/treethought/jot/app"
)

type State struct {
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

func (s *State) SetCurrentNote(note *app.Note) {
	s.currentNote = note
}

func (s *State) CurrentNote() (note *app.Note) {
	return s.currentNote
}
