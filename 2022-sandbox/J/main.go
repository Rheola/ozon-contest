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

	var n int
	fmt.Fscanln(in, &n)

	dict := make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanln(in, &s)
		dict[i] = s
	}

	var k int
	fmt.Fscanln(in, &k)

	queries := make([]string, k)

	for i := 0; i < k; i++ {
		var s string
		fmt.Fscanln(in, &s)
		queries[i] = s
	}

	for _, s := range queries {
		max := 0
		res := ""
		for _, d := range dict {
			if d != s {
				r := rifm(s, d)
				if r > max {
					max = r
					res = d
				}
			}
		}

		if max == 0 {
			for _, d := range dict {
				if d != s {
					res = d
					break
				}
			}
		}

		fmt.Fprintln(out, res)

	}

	// Для каждого запроса выведите одну строку — слово из словаря, которое не совпадает с заданным в
	// запросе и имеет с ним максимальную зарифмованность (если таких несколько — выведите любое).
}

func rifm(a, b string) (res int) {
	ab := []rune(a)
	bb := []rune(b)

	maxA := len(ab)
	maxB := len(bb)

	res = 0
	for i := 0; i < maxA; i++ {
		if i >= maxB {
			return
		}
		if ab[maxA-1-i] == bb[maxB-1-i] {
			res++
		} else {
			return
		}
	}

	return
}
