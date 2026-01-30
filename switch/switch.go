package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)

	// fmt.Println(t.Hour())
	// fmt.Println(t.Day())
	// fmt.Println(t.Date())
	// fmt.Println(t.Clock())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("It's the weekend. YAY! ")
	default:
		fmt.Print("It's a weekday. ")
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}

	whatAmI(true != false)
	whatAmI(231)
	whatAmI("Hello")
	whatAmI(uint16(123))

}
