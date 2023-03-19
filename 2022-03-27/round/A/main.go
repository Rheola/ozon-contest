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

	var count int

	fmt.Fscan(in, &count)
	n1, n24, n2416 := calcSizes(count)
	fmt.Fprintln(out, n1, n24, n2416)
}

func calcSizes(count int) (n1, n24, n2416 int) {
	n2416 = count / (24 * 16)

	n24 = (count - n2416*24*16) / 24

	n1 = (count - n2416*24*16) % 24

	total := calcTotal(n1, n24, n2416)

	total1 := calcTotal(0, n24+1, n2416)

	total2 := calcTotal(0, 0, n2416+1)

	if total2 < total1 && total2 < total {
		return 0, 0, n2416 + 1
	}

	if total1 < total {
		return 0, n24 + 1, n2416
	}

	return n1, n24, n2416
}

func calcTotal(n1, n24, n2416 int) int {
	return n1*55 + n24*1100 + n2416*16000
}
