package handlers

import (
	"html/template"
	"net/http"
)

// Message_Error is used to pass error status and message to the template
type Message_Error struct {
	Status  int    // HTTP status code (e.g., 404, 400)
	Message string // Error message to display
}

// render displays an error page with the given status code
func render(w http.ResponseWriter, status int) {
	// Parse the error template file
	tmp, err := template.ParseFiles("template/error.html")
	// Set the HTTP status code in the response
	w.WriteHeader(status)
	// If there is an error loading the template, show a simple error message
	if err != nil {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	// Prepare the error message based on the status code
	message := ""
	switch status {
	case 400:
		message = "Bad Request"
	case 404:
		message = "Not Found"
	case 405:
		message = "Status Method Not Allowed"
	default:
		message = "Status Internal Server Error"
	}
	// Create a struct with status and message to pass to the template
	mes := Message_Error{
		Status:  status,
		Message: message,
	}
	// Execute the template and display the error page
	err = tmp.Execute(w, mes)
	if err != nil {
		render(w, http.StatusInternalServerError)
		return
	}
}
