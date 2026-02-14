package main

import (
	"log"
	"net/http"
)

func main() {
	//pages := getJson("gopher.json")
	log.Fatal(http.ListenAndServe(":8181", viewHandle("/abc")))
}
