package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", viewHandle("abc"))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
