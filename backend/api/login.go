package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (api *API) SignIn(c *gin.Context) {
	var user User
	user.Name = c.Request.FormValue("name")
	user.Password = c.Request.FormValue("password")

	var tempUser User
	if err := api.db.Where("name = ?", user.Name).First(&tempUser).Error; err != nil {
		log.Printf("Cannot find user by name, error: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	if ok := CheckPasswordHash(user.Password, tempUser.Password); !ok {
		log.Println("Passwords do not match")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session := sessions.Default(c)
	session.Set("name", user.Name)
	if err := session.Save(); err != nil {
		log.Printf("Error saving cookie: %v\n", err)
	}
}
