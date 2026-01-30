package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	for i := range 10 {
		s := fmt.Sprintf("key%d", i+1)
		m[s] = i + 1
	}
	fmt.Println(m)

	val1 := m["somethingNonExistant"]
	fmt.Println(val1)

	delete(m, "key10")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
}
