package domain

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	createdAt time.Time
	updatedAt time.Time
	text      string
	id        uuid.UUID
}

type NoteOption func(n Note) Note

func NewNote(opts ...NoteOption) Note {
	n := Note{}

	for _, opt := range opts {
		n = opt(n)
	}

	if n.id == uuid.Nil {
		n.id = uuid.New()
	}

	if n.createdAt.IsZero() {
		n.createdAt = time.Now()
	}

	if n.updatedAt.IsZero() {
		n.updatedAt = time.Now()
	}

	return n
}

func WithID(id uuid.UUID) NoteOption {
	return func(n Note) Note {
		n.id = id
		return n
	}
}

func WithText(t string) NoteOption {
	return func(n Note) Note {
		n.SetText(t)
		return n
	}
}

func WithCreatedAt(t time.Time) NoteOption {
	return func(n Note) Note {
		n.createdAt = t
		return n
	}
}

func WithUpdatedAt(t time.Time) NoteOption {
	return func(n Note) Note {
		n.updatedAt = t
		return n
	}
}

func (n Note) ID() uuid.UUID {
	return n.id
}

func (n Note) Text() string {
	return n.text
}

func (n *Note) SetText(t string) {
	n.text = t
}

func (n Note) CreatedAt() time.Time {
	return n.createdAt
}

func (n Note) UpdatedAt() time.Time {
	return n.updatedAt
}
