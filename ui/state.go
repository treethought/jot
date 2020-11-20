package ui

import (
	"github.com/treethought/jot/app"
)

type State struct {
	currentWidget WidgetRenderer
	currentNote   *app.Note
	currentIdx    int
}

func NewState() *State {
	state := &State{
		currentWidget: nil,
		currentNote:   &app.Note{},
	}
	return state
}

func (s *State) SetCurrentNote(note *app.Note) {
	s.currentNote = note
}

func (s *State) CurrentNote() (note *app.Note) {
	if s.currentNote == nil {
		return &app.Note{}

	}
	return s.currentNote
}
