package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open("ex5.html")

	if err != nil {
		log.Fatal("Error while opening file")
	}
	defer f.Close()

	var linkText strings.Builder
	tokenizer := html.NewTokenizer(f)
	var counter int = 0

	for {

		tt := tokenizer.Next()
		tagName := tokenizer.Token().Data

		switch tt {
		case html.StartTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				counter++
				for _, attr := range token.Attr {
					fmt.Println("href:", attr.Val)
				}
			}
		case html.TextToken:
			if (counter > 0 && counter < 2) && tagName == "a" {
				linkText.WriteString(tokenizer.Token().Data)
			}
		case html.EndTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				counter--
				fmt.Println("Text:", strings.TrimSpace(linkText.String()))
				linkText.Reset()
			}
		case html.ErrorToken:
			err := tokenizer.Err()

			if err == io.EOF {
				return
			}

		}

	}
}
