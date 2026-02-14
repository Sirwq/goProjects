package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/manifoldco/promptui"
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

func main() {
	p := getJson("gopher.json")
	page, ok := p["intro"]
	if !ok {
		fmt.Println("error")
		return
	}

	for {
		x := readStory(page)

		if x == "home" {
			return
		}

		page, ok = p[x]
		if !ok {
			fmt.Println("error")
			return
		}
	}

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

func readStruct(p map[string]Page) {
	for i := range p {
		fmt.Println(p[i].Title)
		for j, opt := range p[i].Options {
			fmt.Println(j+1, opt.Text)
		}
		fmt.Println()
	}
}

func readStory(p Page) string {

	for i := range p.Story {
		fmt.Println(p.Story[i])
	}

	prompt := promptui.Select{
		Label: "Choose your option",
		Items: p.Options,
		Templates: &promptui.SelectTemplates{
			Active:   "-> {{ .Text}}",
			Inactive: "   {{ .Text}}",
			Selected: "   {{ .Text}}",
		},
	}
	i, _, err := prompt.Run()

	check(err, "Prompt failed %v\n")
	return p.Options[i].Arc
}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
