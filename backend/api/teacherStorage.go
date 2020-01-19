package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (api *API) GetTeacher(c *gin.Context) {
	// Check if the article ID is valid
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var teacher Teacher
		api.db.Where("id = ?", id).First(&teacher)

		if teacher.Id == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if err := json.NewEncoder(c.Writer).Encode(teacher); err != nil {
			log.Printf("Cannot encode teacher to JSON, error: %v\n", err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

