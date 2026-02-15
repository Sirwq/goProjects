package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	pages := getJson("gopher.json")
	log.Fatal(http.ListenAndServe(":8182", StoryHandler(pages, http.DefaultServeMux)))
}
