package main

import (
	"html/template"
	"net/http"
)

func viewHandle(page Page) http.HandlerFunc {
	tmpl, err := template.ParseFiles("templates/index.html")
	check(err, "parsing template")

	return func(w http.ResponseWriter, r *http.Request) {
		data := page
		err = tmpl.Execute(w, data)
		check(err, "executing template")
	}
}

func StoryHandler(pages map[string]Page, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arc := r.URL.Path[1:]

		if arc == "" {
			arc = "intro"
		}

		if page, ok := pages[arc]; ok {
			viewHandle(page)(w, r)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
