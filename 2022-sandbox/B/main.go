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

	for i := 0; i < cases; i++ {
		runCase(in, out)
	}
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	prices := make(map[int]int)
	for j := 0; j < n; j++ {
		var price int

		fmt.Fscan(in, &price)
		prices[price]++
	}
	total := calcAll(prices)

	fmt.Fprintln(out, total)
}

func calcAll(prices map[int]int) int {
	total := 0
	for price, count := range prices {
		total += calcPosition(price, count)
	}

	return total
}

func calcPosition(price int, count int) int {
	return price * (count - count/3)
}
