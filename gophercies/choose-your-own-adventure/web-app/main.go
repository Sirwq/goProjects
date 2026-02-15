package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	pages := getJson("gopher.json")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8182", StoryHandler(pages, mux)))
}
