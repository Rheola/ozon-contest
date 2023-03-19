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

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		var price int

		fmt.Fscan(in, &price)
		prices[i] = price
	}

	fmt.Fprintln(out, calcMaxSum(prices))
}

func calcMaxSum(prices []int) int {
	n := len(prices)

	if n == 1 {
		return prices[0]
	}

	max := prices[n-1]

	days := make(map[int]struct{})
	days[n-1] = struct{}{}
	for i := n - 2; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
			days[i] = struct{}{}
		}
	}

	total := 0
	prev := 0
	for i, price := range prices {
		_, ok := days[i]
		if ok {
			l := i - prev + 1
			prev = i + 1
			total += price * l
		}
	}

	return total
}
