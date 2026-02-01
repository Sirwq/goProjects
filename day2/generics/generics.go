package main

import (
	"fmt"
	"strconv"
)

type Stack[T any] struct {
	items []T
}

func (st *Stack[T]) Push(item T) {
	st.items = append(st.items, item)
}

func (st *Stack[T]) Pop() (val T, ok bool) {
	if !st.IsEmpty() {
		val := st.items[len(st.items)-1]
		st.items = st.items[:len(st.items)-1]
		return val, true
	}
	return val, false
}

func (st *Stack[T]) IsEmpty() bool {
	return len(st.items) == 0
}

func (st *Stack[T]) Peek() (val T, ok bool) {
	if !st.IsEmpty() {
		return st.items[len(st.items)-1], true
	}
	return val, false
}

func main() {
	v := make([]int, 5)

	for i := range v {
		if i%2 == 0 {
			v[i] = -i
		} else {
			v[i] = i
		}
	}

	fmt.Println(v)

	vFiltered := Filter(v, func(x int) bool { return x > 0 })
	fmt.Println(vFiltered)

	vStr := Map(v, strconv.Itoa)

	fmt.Printf("Slice: %s With type: %T\t", vStr, vStr)

}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	ret := make([]T, 0) // var ret []T   OR   []T{}
	for _, v := range slice {
		if predicate(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map[T, R any](slice []T, transformer func(T) R) []R {
	ret := make([]R, 0, len(slice))
	for _, v := range slice {
		ret = append(ret, transformer(v))
	}
	return ret
}
