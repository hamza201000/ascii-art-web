package handlers

import (
	"html/template"
	"net/http"
)

type Message_Error struct {
	Status  int
	Message string
}

func render(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("template/error.html")
	w.WriteHeader(status)
	if err != nil {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
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
	mes := Message_Error{
		Status:  status,
		Message: message,
	}

	tmp.Execute(w, mes)
}
