package main

import "fmt"

func main() {

	var a = "string"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Printf("%d + %d = %d\n", b, c, b+c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Print(e, " - declaration without initialization ")
	fmt.Println("are zero-valued")

	f := "something"
	fmt.Println(f)

}
