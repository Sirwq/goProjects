package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Page struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func getJson(filePath string) map[string]Page {
	file, err := os.Open(filePath)

	check(err, "reading file")
	defer file.Close()

	var p map[string]Page
	byteValue, err := io.ReadAll(file)

	check(err, "parsing json")
	json.Unmarshal(byteValue, &p)

	return p
}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
