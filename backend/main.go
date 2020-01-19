package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(static.Serve("/", static.LocalFile("client/dist", true)))
	router.LoadHTMLGlob("client/dist/*.html")

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

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
