package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "problems.csv" // add change path later
	//readCsv(path)

	for line := range readCsvChanneling(path) {
		fmt.Println(line[0], "|||", line[1])
	}

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
			fmt.Printf("rrror of type: %s\n", err)
			break
		}

		fmt.Printf("val1: %s, val2: %s\n", line[0], line[1])

		q, _ := strconv.Atoi(strings.Split(line[0], "+")[0])
		q1, _ := strconv.Atoi(strings.Split(line[0], "+")[1])
		q = q + q1
		a, _ := strconv.Atoi(line[1])

		fmt.Printf("VAL1: %d with type %T\n", q, q)
		fmt.Printf("VAL2: %d with type %T\n", a, a)
	}
}

func readCsvChanneling(path string) <-chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("can't open file")
			return
		}
		reader := csv.NewReader(file)

		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("rrror of type: %s\n", err)
			}
			ch <- line
		}
	}()
	return ch
}
