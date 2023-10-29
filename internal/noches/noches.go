package noches

import (
	"github.com/LeonardsonCC/noches/internal/noches/repository"
	"github.com/LeonardsonCC/noches/internal/noches/service"
)

type Noches struct {
	service.NoteService
}

func NewNoches(filePath string) (Noches, error) {
	repo, err := repository.NewFileRepository(filePath)
	if err != nil {
		return Noches{}, err
	}

	svc, err := service.NewNoteService(repo)
	if err != nil {
		return Noches{}, err
	}

	return Noches{
		svc,
	}, nil
}
