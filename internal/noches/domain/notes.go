package domain

import "github.com/google/uuid"

type Notes map[uuid.UUID]Note

func NewNotes() Notes {
	return make(Notes)
}
