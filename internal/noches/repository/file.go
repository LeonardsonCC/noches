package repository

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/LeonardsonCC/noches/internal/noches/domain"
	"github.com/google/uuid"
)

type FileRepository struct {
	buffer   io.Reader
	filePath string
}

func NewFileRepository(filePath string) (Repository, error) {
	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		f, err = os.Create(filePath)
	}
	if err != nil {
		return nil, err
	}

	return &FileRepository{
		filePath: filePath,
		buffer:   f,
	}, nil
}

func (fr *FileRepository) Save(notes domain.Notes) error {
	n := mapNoteToNoteModel(notes)

	b, err := json.Marshal(n)
	if err != nil {
		return err
	}

	err = os.WriteFile(fr.filePath, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileRepository) List() (domain.Notes, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 100))
	tee := io.TeeReader(fr.buffer, buf)
	fr.buffer = buf

	content, err := io.ReadAll(tee)
	if err != nil {
		return nil, err
	}

	var notes map[string]NoteModel

	err = json.Unmarshal(content, &notes)
	if err != nil {
		return nil, err
	}

	n := mapNoteModelToNote(notes)

	return n, nil
}

func (fr *FileRepository) Delete(notes domain.Notes, id uuid.UUID) error {
	if _, found := notes[id]; !found {
		return ErrNoteNotFound
	}
	delete(notes, id)

	return nil
}
