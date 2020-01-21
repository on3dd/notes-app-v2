package api

import "time"

type Note struct {
	Id          int       `json:"id,omitempty"`
	AuthorId    int       `json:"author_id,omitempty"`
	CategoryId  int       `json:"category_id,omitempty"`
	SubjectId   int       `json:"subject_id,omitempty"`
	TeacherId   int       `json:"teacher_id,omitempty"`
	PostedAt    time.Time `json:"posted_at,omitempty"`
	Title       string    `json:"title,omitempty"`
	Descirption string    `json:"description,omitempty"` // TODO: change it back to description
	Link        string    `json:"link,omitempty"`
}
