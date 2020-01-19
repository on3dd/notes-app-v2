package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))
	router.LoadHTMLGlob("../frontend/dist/*.html")

	// Pages router
	pr := router.Group("/")
	{
		pr.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		pr.GET("/about", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		pr.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}

	return router
}
