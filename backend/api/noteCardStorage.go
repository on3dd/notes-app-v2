package api

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func (api *API) GetNoteCards(c *gin.Context) {
	var noteCards []NoteCard

	searchItems := "notes.id, notes.author_id, users.name, notes.category_id, categories.name, notes.subject_id, subjects.name, " +
		"notes.teacher_id, teachers.name, notes.posted_at, notes.title"

	err := api.db.Table("notes").Select(searchItems).
		Joins("INNER JOIN users ON notes.author_id = users.id").
		Joins("INNER JOIN categories ON notes.category_id = categories.id").
		Joins("INNER JOIN subjects ON notes.subject_id = subjects.id").
		Joins("INNER JOIN teachers ON notes.teacher_id = teachers.id").
		Order("notes.posted_at DESC").
		Scan(&noteCards).Error

	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(c.Writer).Encode(noteCards); err != nil {
		log.Printf("Cannot encode note cards to JSON, error: %v\n", err)
	}
}