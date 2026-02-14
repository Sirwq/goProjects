package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m)
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

func main() {
	pages := getJson("gopher.json")

	for i, p := range pages {
		fmt.Println(i, p.Title)
	}

	http.ListenAndServe(":8181", StoryHandler(pages, http.DefaultServeMux))
}
