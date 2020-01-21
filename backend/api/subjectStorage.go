package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetSubject(c *gin.Context) {
	// Check if the article ID is valid
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var subject Subject
		api.db.Where("id = ?", id).First(&subject)

		if subject.Id == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if err := json.NewEncoder(c.Writer).Encode(subject); err != nil {
			log.Printf("Cannot encode subject to JSON, error: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (api *API) GetSubjects(c *gin.Context) {
	var subjects []Subject

	if keys := c.Request.URL.Query()["category_id"]; len(keys[0]) != 0 {
		api.db.Joins("inner join categories on categories.id = subjects.category_id").
			Where("categories.id = ?", keys[0]).Find(&subjects)
	} else {
		api.db.Order("id asc").Find(&subjects)
	}

	if len(subjects) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := json.NewEncoder(c.Writer).Encode(subjects); err != nil {
		log.Printf("Cannot encode subjects to JSON, error: %v\n", err)
	}
}
