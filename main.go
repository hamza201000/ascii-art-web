package main

import (
	"fmt"
	"log"
	"net/http"

	"asciart/handlers"
)

func main() {
	// Handle the root URL "/"
	http.HandleFunc("/", handlers.Handler)
	// Handle the "/ascii-art" endpoint
	http.HandleFunc("/ascii-art", handlers.Greethandler)
	// Serve static files from "/static/"
	http.HandleFunc("/static/", handlers.HandlerStatic)
	// Log server start message
	log.Println("Server running on: http://localhost:8080")
	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
