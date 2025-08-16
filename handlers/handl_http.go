package handlers

import (
	"html/template"
	"net/http"
	"os"

	"asciart/asciart"
)

// Handler serves the main index page for GET requests to "/"
func Handler(w http.ResponseWriter, r *http.Request) {
	// Check if the URL path is exactly "/"
	if r.URL.Path != "/" {
		render(w, http.StatusNotFound)
		return
	}
	// Only allow GET requests
	if r.Method != http.MethodGet {
		render(w, http.StatusMethodNotAllowed)
		return
	}
	// Parse and execute the index template
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		render(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		render(w, http.StatusInternalServerError)
		return
	}
}

// Greethandler processes POST requests to "/ascii-art" and generates ASCII art
func Greethandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		render(w, http.StatusMethodNotAllowed)
		return
	}
	// Validate input length
	if len(r.FormValue("name")) > 10000 {
		render(w, http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	Banner := r.FormValue("mySelect")
	// Generate ASCII art using Fmain
	data, status_code := asciart.Fmain(name, Banner)
	if status_code != 200 {
		render(w, status_code)
		return
	}
	// Parse and execute the index template with generated data
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		render(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		render(w, http.StatusInternalServerError)
		return
	}
}

// HandlerStatic serves static files for GET requests to "/static/"
func HandlerStatic(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		render(w, http.StatusMethodNotAllowed)
		return
	} else {
		// Check if the requested file exists and is not a directory
		info, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			render(w, http.StatusNotFound)
			return
		} else if info.IsDir() {
			render(w, http.StatusNotFound)
			return
		} else {
			// Serve the static file
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}
