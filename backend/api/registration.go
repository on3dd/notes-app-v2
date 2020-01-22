package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) SignUp(c *gin.Context) {
	var user User
	user.Name = c.Request.FormValue("name")
	user.Password = c.Request.FormValue("password")
	user.Email = c.Request.FormValue("email")
	user.About = c.Request.FormValue("about")

	var tempUser User
	if err := api.db.Last(&tempUser).Error; err != nil {
		log.Printf("Error getting last user from db, error: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	user.Id = tempUser.Id + 1

	var err error
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Cannot hash user's password, error: %v", err)
	}

	if err := api.db.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Cannot create new user, error: %v", err)
	}
}
