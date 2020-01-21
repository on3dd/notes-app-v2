package main

import (
	"net/http"

	"notes-app-v2/backend/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))
	router.LoadHTMLGlob("../frontend/dist/*.html")

	pageRoutes := []string{"/", "/about", "/upload", "/notes", "/notes/:id"}

	pageRouter := router.Group("/")
	for _, route := range pageRoutes {
		pageRouter.GET(route, func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}

	api := api.New(db)

	apiRouter := router.Group("/api/v1/")
	{
		apiRouter.Use(HeadersMiddleware())

		apiRouter.GET("notes", api.GetNotes)
		apiRouter.GET("notes/:id", api.GetNote)
		apiRouter.POST("notes", api.AddNote)

		apiRouter.GET("users", api.GetUsers)
		apiRouter.GET("users/:id", api.GetUser)

		apiRouter.GET("categories", api.GetCategories)
		apiRouter.GET("categories/:id", api.GetCategory)

		apiRouter.GET("subjects", api.GetSubjects)
		apiRouter.GET("subjects/:id", api.GetSubject)

		apiRouter.GET("teachers", api.GetTeachers)
		apiRouter.GET("teachers/:id", api.GetTeacher)
	}

	return router
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
