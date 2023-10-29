package repository

import (
	"errors"

	"github.com/LeonardsonCC/noches/internal/noches/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Save(notes domain.Notes) error
	List() (domain.Notes, error)
	Delete(notes domain.Notes, id uuid.UUID) error
}

var ErrNoteNotFound = errors.New("note not found in notes list")
