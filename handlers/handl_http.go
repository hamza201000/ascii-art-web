package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"asciart/asciart"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		render(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		render(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func Greethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		render(w, 400)
		return
	}
	start := time.Now() // Start timer here
	if len(r.FormValue("name")) > 10000 {
		render(w, 400)
		return
	}
	name := r.FormValue("name")
	Banner := r.FormValue("mySelect")
	data, status_code := asciart.Fmain(name, Banner)
	if status_code != 200 {
		render(w, status_code)
		return
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
	elapsed := time.Since(start).Seconds()
	log.Printf("Handled /ascii in %.3f seconds\n", elapsed)
}
