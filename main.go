package main

import (
	"fmt"
	"log"
	"net/http"

	"asciart/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/ascii-art", handlers.Greethandler)
	http.HandleFunc("/static/", handlers.HandlerStatic)
	log.Println("Server running on: http://localhost:8080")
if err:=http.ListenAndServe(":8080", nil);err!=nil{
	fmt.Print("internal server.")
	return 
}
	
}
