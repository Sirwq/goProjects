package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open("ex1.html")

	if err != nil {
		log.Fatal("Error while opening file")
	}

	z := html.NewTokenizer(f)

	for {
		tt := z.Next()

		if tt == html.ErrorToken {
			fmt.Printf("error: %s", err)
			break
		}

		if tt == html.StartTagToken {
			fmt.Println(z.Token())
		}

	}
}
