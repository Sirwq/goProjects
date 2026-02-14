package main

import "fmt"

type rect struct {
	w int
	h int
}

func main() {
	r := rect{10, 20}

	fmt.Println(r)
	fmt.Println(r.area())
}

func (r *rect) area() int {
	return r.h * r.w
}

func (r rect) perim() int {
	return 2*r.h + 2*r.w
}
