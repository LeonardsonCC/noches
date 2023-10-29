package repository

import (
	"fmt"
	"time"

	"github.com/LeonardsonCC/noches/internal/noches/domain"
	"github.com/google/uuid"
)

func mapNoteModelToNote(notes map[string]NoteModel) domain.Notes {
	n := make(domain.Notes, len(notes))

	for _, note := range notes {
		id, err := uuid.Parse(note.NoteID)
		if err != nil {
			fmt.Printf("failed to get id from note: %v", err)
			continue
		}

		createdAt, err := time.Parse(time.UnixDate, note.CreatedAt)
		if err != nil {
			fmt.Printf("failed to get created at date from note: %v", err)
			continue
		}

		updatedAt, err := time.Parse(time.UnixDate, note.UpdatedAt)
		if err != nil {
			fmt.Printf("failed to get updated at date from note: %v", err)
			continue
		}

		nn := domain.NewNote(
			domain.WithID(id),
			domain.WithText(note.Text),
			domain.WithCreatedAt(createdAt),
			domain.WithUpdatedAt(updatedAt),
		)

		n[nn.ID()] = nn
	}

	return n
}

func mapNoteToNoteModel(notes domain.Notes) map[string]NoteModel {
	n := make(map[string]NoteModel, len(notes))

	for _, note := range notes {
		nn := NoteModel{
			NoteID:    note.ID().String(),
			Text:      note.Text(),
			CreatedAt: note.CreatedAt().Format(time.UnixDate),
			UpdatedAt: note.UpdatedAt().Format(time.UnixDate),
		}

		n[nn.NoteID] = nn
	}

	return n
}
