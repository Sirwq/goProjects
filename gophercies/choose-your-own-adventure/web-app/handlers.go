package main

import (
	"fmt"
	"net/http"
)

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
