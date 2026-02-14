package main

import "fmt"

type node struct {
	next *node
	val  int
}

func nodeAppend(head *node, val int) {
	cur := head
	for {
		if cur.next != nil {
			cur = cur.next
		} else {
			break
		}
	}
	cur.next = new(node)
	cur.next.val = val
}

func nodePrint(head *node) {
	cur := head
	for {
		fmt.Println(cur.val)
		if cur.next != nil {
			cur = cur.next
		} else {
			break
		}
	}
}

func (head *node) printNode() {
	cur := head
	for {
		fmt.Println(cur.val)
		if cur.next != nil {
			cur = cur.next
		} else {
			break
		}
	}
}

func nodePrepend(head **node, val int) {
	n := new(node)
	n.val = val
	n.next = *head
	*head = n
}

func nodeReverse(head **node) {
	var prev *node = nil
	var cur *node = *head

	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	*head = prev
}

func main() {
	// n := 10
	// passByVal(n)
	// fmt.Println(n)
	// passByRef(&n)
	// fmt.Println(n)

	// z, x := 10, 20

	// fmt.Println(z, x)
	// swapVals(&z, &x)
	// fmt.Println(z, x)

	// sl := make([]int, 10)

	// for i := range sl {
	// 	sl[i] = i * 2
	// }

	// fmt.Println(sl)

	// var k int = 0

	// findMax(sl, &k)
	// fmt.Println(k)

	head := new(node)
	head.val = 2

	nodeAppend(head, 7)
	nodeAppend(head, 6)
	nodeAppend(head, 4)
	nodeAppend(head, 2)
	nodeAppend(head, 1)

	nodePrepend(&head, 10)
	nodeReverse(&head)

	passByRef(&head.val)
	head.printNode()

}

func passByVal(n int) {
	n++
}

func passByRef(n *int) {
	*n++
}

func swapVals(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func findMax(s []int, saveMax *int) {
	*saveMax = 0

	for i := range len(s) {
		if s[i] > *saveMax {
			*saveMax = s[i]
		}
	}

}
