package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	w, h float64
}

type circle struct {
	radius float64
}

func main() {
	r := rect{10, 3}
	c := circle{4}

	detectCircle(r)
	detectCircle(c)

	measure(r)
	measure(c)
}

func (r rect) area() float64 {
	return r.w * r.h
}
func (r rect) perim() float64 {
	return 2*r.w + 2*r.h
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("---   MEASURE  ---")
	fmt.Printf("Type: %T\n", g)
	fmt.Println("vals:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
	fmt.Println("------------------")
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius: ", c.radius)
	}
}
