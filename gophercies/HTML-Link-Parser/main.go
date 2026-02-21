package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open("ex1.html")

	if err != nil {
		log.Fatal("Error while opening file")
	}
	defer f.Close()

	tokenizer := html.NewTokenizer(f)
	insideLink := false

	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return // err
		}

		if tt == html.StartTagToken {
			fmt.Print("Start tag: ")
			fmt.Println(tokenizer.Token().Data)
			//fmt.Println(tokenizer.Token())
		}

		if tt == html.EndTagToken {
			fmt.Print("End tag")
		}

	}
}
