package main

import (
	"log"
	"net/http"
	"text/template"

	"asciart/asciart"
)

type PageData struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("template/basic.html"))
	tmpl.Execute(w, nil)
}

func greethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	name := r.FormValue("name")
	data := PageData{Name: asciart.Fmain(name)}
	tmpl := template.Must(template.ParseFiles("template/basic.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii", greethandler)
	log.Println("Server running on: http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
