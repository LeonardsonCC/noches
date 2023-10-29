package domain_test

import (
	"testing"
	"time"

	"github.com/LeonardsonCC/noches/internal/noches/domain"
)

func TestNote(t *testing.T) {
	t.Run("note constructor", func(t *testing.T) {
		tt := []struct {
			createdAt time.Time
			updatedAt time.Time
			text      string
		}{
			{
				text:      "a lot of text",
				createdAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
				updatedAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
		}

		for _, tc := range tt {
			opts := make([]domain.NoteOption, 0, 3)

			if tc.text != "" {
				opts = append(opts, domain.WithText(tc.text))
			}
			opts = append(opts, domain.WithCreatedAt(tc.createdAt))
			opts = append(opts, domain.WithUpdatedAt(tc.updatedAt))

			n := domain.NewNote(opts...)

			if tc.text != n.Text() {
				t.Errorf("note text is different: expected [%s] but received [%s]", tc.text, n.Text())
			}

			if tc.createdAt != n.CreatedAt() {
				t.Errorf("note created-at is different: expected [%s] but received [%s]", tc.createdAt, n.CreatedAt())
			}

			if tc.updatedAt != n.UpdatedAt() {
				t.Errorf("note updated-at is different: expected [%s] but received [%s]", tc.updatedAt, n.UpdatedAt())
			}
		}
	})
}
