package main

import (
	"net/http"
)

func main() {
	pages := getJson("gopher.json")
	http.ListenAndServe(":8181", StoryHandler(pages, http.DefaultServeMux))
}
