package main

import (
	"asciart/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/ascii-art", handlers.Greethandler)
	log.Println("Server running on: http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
