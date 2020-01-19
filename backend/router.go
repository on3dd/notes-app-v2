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

	pageRouter := router.Group("/")
	{
		pageRouter.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		pageRouter.GET("/about", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		pageRouter.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}

	api := api.New(db)

	apiRouter := router.Group("/api/v1/")
	{
		apiRouter.Use(HeadersMiddleware())

		apiRouter.GET("notes/:id", api.GetNote)

		apiRouter.GET("users/:id", api.GetUser)

		apiRouter.GET("categories", api.GetCategories)
		apiRouter.GET("categories/:id", api.GetCategory)

		apiRouter.GET("subjects/:id", api.GetSubject)

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