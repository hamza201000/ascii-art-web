package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
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
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		render(w, 500)
		return
	}
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
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		render(w, 500)
		return
	}
	tmpl.Execute(w, data)
	elapsed := time.Since(start).Seconds()
	log.Printf("Handled /ascii in %.3f seconds\n", elapsed)
}

func HandlerStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		render(w, 400)
		return
	} else {
		info, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			render(w, http.StatusNotFound)
			return
		} else if info.IsDir() {
			render(w, http.StatusForbidden)
			return
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}
