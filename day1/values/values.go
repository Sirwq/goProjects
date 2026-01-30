package main // To have entry point every programm should have 'main' package or it'll not run even if compiled

import "fmt"

func main() {
	fmt.Println("go" + "concat" + "123")
	fmt.Println("1+1 = ", 1+1)
	fmt.Printf("10*12 = %d\n", 10*12)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
