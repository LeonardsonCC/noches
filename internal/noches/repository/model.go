package repository

type NoteModel struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Text      string `json:"text,omitempty"`
	NoteID    string `json:"note_id,omitempty"`
}
