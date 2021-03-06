package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetUser(c *gin.Context) {
	// Check if the article ID is valid
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var user User
		api.db.Where("id = ?", id).First(&user)

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if err := json.NewEncoder(c.Writer).Encode(user); err != nil {
			log.Printf("Cannot encode user to JSON, error: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (api *API) GetUsers(c *gin.Context) {
	var users []User
	api.db.Order("id asc").Find(&users)

	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := json.NewEncoder(c.Writer).Encode(users); err != nil {
		log.Printf("Cannot encode users to JSON, error: %v\n", err)
	}
}
