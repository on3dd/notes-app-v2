package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (api *API) GetNote(c *gin.Context) {
	// Check if the article ID is valid
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var note Note
		api.db.Where("id = ?", id).First(&note)

		if note.Id == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if err := json.NewEncoder(c.Writer).Encode(note); err != nil {
			log.Printf("Cannot encode note to JSON, error: %v\n", err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}