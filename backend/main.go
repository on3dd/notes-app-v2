package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
