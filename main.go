package main

import (
	"net/http"
	"text/template"

	"asciart/asciart"
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
	// fmt.Println(name)
	data := PageData{Name: asciart.Fmain(name)}
	tmpl := template.Must(template.ParseFiles("template/basic.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii", greethandler)
	http.ListenAndServe(":8080", nil)
}
