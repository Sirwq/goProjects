package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	qCounter := 0
	rAnswers := 0
	var uInput string
	var ops atomic.Uint64
	start := time.Now()

	go func() {
		for {
			time.Sleep(time.Second)
			ops.Add(1)
		}
	}()

	go func() {
		<-time.After(time.Duration(*timeLimit) * time.Second)
		fmt.Printf("\nYou got right %d out of %d questions\n", rAnswers, qCounter)
		exit("\nTime ended!")
	}()

	for line := range readCsvChanneling(*csvFilename) {
		qCounter++
		fmt.Print(line[0], "=")
		fmt.Scanln(&uInput)
		fmt.Println()
		if uInput == line[1] {
			rAnswers++
		}
	}
	fmt.Printf("You got right %d out of %d questions\n", rAnswers, qCounter)
	fmt.Println("Time passed easy method:", time.Since(start).Round(time.Millisecond))
	fmt.Println("Time passed:", ops.Load())
}

func readCsvChanneling(path string) <-chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)
		file, err := os.Open(path)
		if err != nil {
			exit(fmt.Sprintf("error of type: %s\n", err))
		}
		defer file.Close()
		reader := csv.NewReader(file)

		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				exit(fmt.Sprintf("error of type: %s\n", err))
			}
			ch <- line
		}
	}()
	return ch
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
