package api

import "time"

type NoteCard struct {
	Id           int       `json:"id"`
	AuthorId     int       `json:"author_id"`
	AuthorName   string    `json:"author_name"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	SubjectId    int       `json:"subject_id"`
	SubjectName  string    `json:"subject_name"`
	TeacherId    int       `json:"teacher_id"`
	TeacherName  string    `json:"teacher_name"`
	PostedAt     time.Time `json:"posted_at"`
	Title        string    `json:"title"`
}
