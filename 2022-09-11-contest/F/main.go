package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cases int
	fmt.Fscan(in, &cases)

	for caseI := 0; caseI < cases; caseI++ {
		var lines int
		fmt.Fscan(in, &lines)

		myMap := make(map[int]InPut)

		f := 0
		for i := 0; i < lines; i++ {
			var a, b, c int

			fmt.Fscan(in, &a, &b, &c)
			if f == 0 {
				f = a
			}

			myMap[a] = InPut{
				a: b,
				b: c,
			}
		}

		first := Node{
			val: f,
		}

		current := &first
		finished := false
		for finished == false {
			key := current.val
			el, ok := myMap[current.val]
			if !ok {
				finished = true
			}

			val := el.a
			_, ok = myMap[el.a]
			if !ok {
				val = el.b
			}

			next := Node{
				val:  val,
				prev: current,
			}

			current.next = &next
			current = current.next
			delete(myMap, key)
		}

		start := first
		middle := first

		for i := 0; i < lines/2; i++ {
			middle = *middle.next
		}

		for i := 0; i < lines/2; i++ {
			fmt.Fprintln(out, start.val, middle.val)
			start = *start.next
			middle = *middle.next
		}

		fmt.Fprintln(out)
	}
}

type InPut struct {
	a int
	b int
}

type Node struct {
	val  int
	prev *Node
	next *Node
}
