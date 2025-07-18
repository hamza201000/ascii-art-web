package main

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/basic.html"))
	tmpl.Execute(w, nil)
}

func greethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	name := r.FormValue("name")
	data := PageData{Name: name}
	tmpl := template.Must(template.ParseFiles("template/greet.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/greet", greethandler)
	http.ListenAndServe(":8080", nil)
}
