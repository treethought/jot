package app

import (
	uuid "github.com/satori/go.uuid"
)

type Note struct {
	ID      string
	Name    string
	Content string
}

// Creates a new note by name
func NewNote(name string) *Note {
	id := uuid.NewV4().String()
	return &Note{ID: id, Name: name}
}
