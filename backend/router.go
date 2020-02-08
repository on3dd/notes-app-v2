package main

import (
	"net/http"

	apipkg "notes-app-v2/backend/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB, sk []byte) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))
	router.LoadHTMLGlob("../frontend/dist/*.html")

	pageRoutes := []string{"/", "about", "upload", "notes", "notes/:id", "login", "join"}

	pageRouter := router.Group("/")
	for _, route := range pageRoutes {
		pageRouter.GET(route, func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}

	api := apipkg.New(db, sk)

	apiRouter := router.Group("/api/v1/")
	{
		apiRouter.Use(HeadersMiddleware())

		apiRouter.OPTIONS("notes/:id", OptionsHandler)
		//apiRouter.OPTIONS("notes", OptionsHandler)

		//apiRouter.GET("notes", api.GetNotes)
		apiRouter.GET("notes", api.GetNoteCards)
		apiRouter.GET("notes/:id", api.GetNote)
		apiRouter.POST("notes", api.AddNote)
		apiRouter.PUT("notes/:id", api.UpdateNote)
		apiRouter.DELETE("notes/:id", api.DeleteNote)

		apiRouter.GET("users", api.GetUsers)
		apiRouter.GET("users/:id", api.GetUser)
		apiRouter.GET("userLastNotes/:id", api.GetUserLastNotes)

		apiRouter.GET("categories", api.GetCategories)
		apiRouter.GET("categories/:id", api.GetCategory)

		apiRouter.GET("subjects", api.GetSubjects)
		apiRouter.GET("subjects/:id", api.GetSubject)

		apiRouter.GET("teachers", api.GetTeachers)
		apiRouter.GET("teachers/:id", api.GetTeacher)

		apiRouter.POST("register", api.SignUp)
		apiRouter.POST("login", api.SignIn)
	}

	return router
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

func OptionsHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
}
