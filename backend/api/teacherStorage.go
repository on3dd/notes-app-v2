package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (api *API) GetTeachers(c *gin.Context) {
	var teachers []Teacher
	api.db.Order("id asc").Find(&teachers)

	if len(teachers) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := json.NewEncoder(c.Writer).Encode(teachers); err != nil {
		log.Printf("Cannot encode teachers to JSON, error: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}