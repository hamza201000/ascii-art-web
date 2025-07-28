package main

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"asciart/asciart"
)

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
	start := time.Now() // Start timer here
	name := r.FormValue("name")
	Banner := r.FormValue("mySelect")
	data := asciart.Fmain(name, Banner)
	tmpl := template.Must(template.ParseFiles("template/basic.html"))
	tmpl.Execute(w, data)
	elapsed := time.Since(start).Seconds()
	log.Printf("Handled /ascii in %.3f seconds\n", elapsed)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii", greethandler)
	log.Println("Server running on: http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
