package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "problems.csv" // add change path later
	readCsv(path)

}

func readCsv(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error of type: %s\n", err)
			break
		}

		fmt.Printf("val1: %s, val2: %s\n", line[0], line[1])
	}
}
