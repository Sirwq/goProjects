package main

import "fmt"

func main() {

	fmt.Println(plus(10, minus(32, 12)))
	fmt.Println(divmod(32, 3))
	_, v := divmod(10, 20)
	fmt.Println(v)

	fmt.Println(strcon("blabla", " ", "bla"))

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	wrapped := logger(func() {
		fmt.Println(plus(4, 6))
	})

	wrapped()

	loggerNoWrap("minus", minus(3, 2))

	fmt.Println("fact:", fact(10))
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

func strcon(s ...string) string {
	var ret []rune

	for _, str := range s {
		for _, ch := range str {
			ret = append(ret, ch)
		}
	}
	return string(ret)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func logger(f func()) func() {
	return func() {
		fmt.Println("start")
		f()
		fmt.Println("end")
	}
}

func loggerNoWrap(name string, result any) {
	fmt.Println("start")
	fmt.Printf("%s: %v\n", name, result)
	fmt.Println("end")
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
