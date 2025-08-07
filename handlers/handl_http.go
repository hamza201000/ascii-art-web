package handlers

import (
	"html/template"
	"net/http"
	"os"

	"asciart/asciart"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		render(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		render(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		render(w, 500)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		render(w, http.StatusNotFound)
		return

	}
}

func Greethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		render(w, 400)
		return
	}
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
	if err := tmpl.Execute(w, data); err != nil {
		render(w, http.StatusNotFound)
		return
	}
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

