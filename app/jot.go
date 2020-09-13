package app

type App struct {
	Store NoteReadWriter
}

func NewApp() *App {
	store := NewFileStore("notetest")
	app := &App{}
	app.Store = store
	return app

}

func (a *App) AddNote(name string, content string) (id string, err error) {
	note := NewNote(name)
	note.Content = content
	err = a.Store.Write(note)
	if err != nil {
		return
	}

	return note.ID, nil
}

func (a *App) GetNote(id string) (note *Note, err error) {
	return a.Store.Read(id)
}

func (a *App) ListNotes() (notes []*Note, err error) {
	return a.Store.List()
}
