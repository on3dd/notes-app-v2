package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
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

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(api.signingKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	session := sessions.Default(c)
	session.Set("name", user.Name)
	session.Set("value", tokenString)
	session.Set("expires", 86400)

	if err := session.Save(); err != nil {
		log.Printf("Error saving cookie: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	res := &Response{
		Auth:  true,
		Token: tokenString,
		User:  user,
	}

	if err := json.NewEncoder(c.Writer).Encode(res); err != nil {
		log.Printf("Cannot encode regirstred user to JSON, error: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
