package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func viewHandle(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:   "Some basic title",
			Message: "There's a message from tempalte",
		}
		tmpl, err := template.ParseFiles("templates/index.html")
		check(err, "parsing template")
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
			fmt.Fprint(w, page.Title)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
