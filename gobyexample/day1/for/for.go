package main

import "fmt"

func main() {
	i := 1

	fmt.Println("i loop")
	for i <= 3 {
		fmt.Print(i, " ")
		i++ // i = i + 1
	}

	fmt.Println("\nj loop")
	for j := 0; j < 3; j++ {
		fmt.Print(j, " ")
	}
	fmt.Println()

	for i := range 3 {
		str := fmt.Sprintf("range %d ", i)
		fmt.Print(str)
	}
	fmt.Println()

	for {
		fmt.Print("loolooloop")
		break
	}
	fmt.Println()

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
