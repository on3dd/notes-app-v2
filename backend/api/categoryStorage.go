package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (api *API) GetCategory(c *gin.Context) {
	// Check if the article ID is valid
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var category Category
		api.db.Where("id = ?", id).First(&category)

		if category.Id == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if err := json.NewEncoder(c.Writer).Encode(category); err != nil {
			log.Printf("Cannot encode category to JSON, error: %v\n", err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (api *API) GetCategories(c *gin.Context) {
	var categories []Category
	api.db.Order("id asc").Find(&categories)

	if len(categories) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := json.NewEncoder(c.Writer).Encode(categories); err != nil {
		log.Printf("Cannot encode category to JSON, error: %v\n", err)
	}
}

