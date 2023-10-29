package service

import (
	"github.com/LeonardsonCC/noches/internal/noches/domain"
	"github.com/LeonardsonCC/noches/internal/noches/repository"
	"github.com/google/uuid"
)

type NoteService struct {
	repo  repository.Repository
	notes domain.Notes
}

func NewNoteService(repo repository.Repository) (NoteService, error) {
	notes, err := repo.List()
	if err != nil {
		return NoteService{}, err
	}

	return NoteService{
		repo,
		notes,
	}, nil
}

func (n NoteService) Create(note domain.Note) error {
	n.notes[note.ID()] = note
	return n.repo.Save(n.notes)
}

func (n NoteService) List() (domain.Notes, error) {
	return n.repo.List()
}

func (n NoteService) Delete(id uuid.UUID) error {
	err := n.repo.Delete(n.notes, id)
	if err != nil {
		return err
	}

	return n.repo.Save(n.notes)
}
