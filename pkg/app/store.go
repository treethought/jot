package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type NoteReadWriter interface {
	Read(id string) (*Note, error)
	Write(*Note) error
	List() ([]*Note, error)
}

type FileStore struct {
	Dir string
}

func NewFileStore(dir string) *FileStore {
	var fs *FileStore = new(FileStore)
	curdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir = fmt.Sprintf("%s/%s", curdir, dir)
	fs.Dir = dir
	return fs

}

func (s *FileStore) Read(id string) (note *Note, err error) {
	path := fmt.Sprintf("%s/%s", s.Dir, id)
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	// info, _ := file.Stat()

	scanner := bufio.NewScanner(file)

	// name is the first line of the file
	var name string

	var content string

	i := 0
	for scanner.Scan() {
		if i == 0 {
			name = scanner.Text()
			i += 1
			continue
		}

		content += scanner.Text()
	}

	note = &Note{
		ID:      id,
		Name:    name,
		Content: content,
	}
	return

}

func (s *FileStore) Write(note *Note) (err error) {
	path := fmt.Sprintf("%s/%s", s.Dir, note.ID)

	fileContent := fmt.Sprintf("%s\n\n%s", note.Name, note.Content)
	data := []byte(fileContent)

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return
	}

	// fmt.Println("Wrote note %s" + note.ID)

	return
}

func (s *FileStore) List() (notes []*Note, err error) {
	files, err := ioutil.ReadDir(s.Dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		n, rerr := s.Read(f.Name())
		if rerr != nil {
			return notes, rerr
		}
		notes = append(notes, n)
	}
	return

}
