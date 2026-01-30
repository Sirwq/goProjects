package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string

	fmt.Printf("uninitialized: %v, s == nil - %v, len == 0 - %v\n",
		s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	for i := range s {
		s[i] = string('A' + rune(i))
	}

	s = append(s, "text val")
	fmt.Println(s)
	fmt.Println()

	for i := range len(s) {
		fmt.Printf("%d - %s\n", i, s[i])
	}

	fmt.Println()

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)
	c = append(c, "str", "123", "abc")
	fmt.Println("copy:", c)
	fmt.Println(c[4:])

	fmt.Println()

	t1 := []string{"1", "2", "3", "4", "5"}
	t2 := []string{"1", "2", "3", "4", "5"}

	fmt.Println("Slice can be compared with nil -", t1 == nil)
	fmt.Println("But not with each other DIRECTLY")
	fmt.Println("That means u can actually compare them")

	if slices.Equal(t1, t2) {
		fmt.Println("t1 == t2")
	} else {
		fmt.Println("t1 != t2")
	}

	// One liners are no Go and there's no ternary operator :C
	fmt.Println(map[bool]string{true: "Yes", false: "No"}[slices.Equal(t1, t1)])

	slice2D := make([][]int, 3)

	for i := range slice2D {
		innerLen := 3
		slice2D[i] = make([]int, innerLen)

		for j := range innerLen {
			slice2D[i][j] = i + j
		}
	}

	fmt.Println("2d:", slice2D)
}
