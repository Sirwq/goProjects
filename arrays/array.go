package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var arr [5]int32
	fmt.Println(arr)

	arr[2] = 901
	fmt.Println(arr)
	fmt.Println(arr[2])
	fmt.Println("len:", len(arr))
	fmt.Println()
	//println(arr) Doenst work
	//println outputs to stderr so basicly never use it

	empty_arr := []int{}
	fmt.Println(empty_arr)
	println(empty_arr, '\n')

	arr2 := []int{10, -2, 412, 0, 5 / 2, 515}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	arr3 := [...]int{1, 2, 3, 4, 5, 10: 11}
	fmt.Println("idx", arr3)

	var arr2D [3][4]int
	fmt.Println(arr2D)
	fmt.Println(len(arr2D)) // rows

	for i := range len(arr2D) {
		for j := range len(arr2D[0]) {
			if i == 0 || j == 0 {
				arr2D[i][j] = i + j
			} else {
				arr2D[i][j] = i * j
			}
		}
	}
	fmt.Println(arr2D)

	var arr3D [3][3][3]int
	for i := range 3 {
		for j := range 3 {
			for k := range 3 {
				arr3D[i][j][k] = (i + j) * k
			}
		}
	}

	printAny(arr3D)
}

func printAny(arr any) {
	fmt.Print("\n--- ARRAY ---\n")
	b, _ := json.Marshal(arr)
	fmt.Println(string(b))
}
