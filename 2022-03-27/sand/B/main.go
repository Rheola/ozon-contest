package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	num := make([]int, 0, 3)
	num = append(num, n-10*(n/10))
	n = n / 10
	num = append(num, n-10*(n/10))
	n = n / 10
	num = append(num, n)
	sort.Ints(num)
	for _, n1 := range num {
		for _, n2 := range num {
			if n1 == n2 {
				continue
			}
			for _, n3 := range num {
				if n1 == n3 || n2 == n3 {
					continue
				}

				fmt.Fprintln(out, 100*n1+10*n2+n3)
			}
		}
	}
}
