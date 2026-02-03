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
	qCounter := 0
	rAnswers := 0
	var uInput string

	for line := range readCsvChanneling(path) {
		qCounter++
		//fmt.Println(line[0], "|||", line[1])
		_, answer := parseLine(line[0], line[1])
		fmt.Print(line[0], "=")
		fmt.Scanln(&uInput)
		fmt.Println()
		num, _ := strconv.Atoi(uInput)
		if num == answer {
			rAnswers++
		}
	}
	fmt.Printf("You got right %d out of %d questions\n", rAnswers, qCounter)
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

// add other operations +,-,*,/
func parseLine(math, answer string) (val1, val2 int) {
	val1, _ = strconv.Atoi(strings.Split(math, "+")[0])
	temp, _ := strconv.Atoi(strings.Split(math, "+")[1])
	val1 += temp
	val2, _ = strconv.Atoi(answer)
	return val1, val2
}
