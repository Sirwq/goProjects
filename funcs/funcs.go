package main

import "fmt"

func main() {

	fmt.Println(plus(10, minus(32, 12)))
	fmt.Println(divmod(32, 3))
	_, v := divmod(10, 20)
	fmt.Println(v)
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func divmod(a, b int) (quotient float32, remainder int) {
	quotient = float32(a) / float32(b)
	remainder = a % b
	return
}
