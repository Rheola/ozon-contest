package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var caseCount int
	fmt.Fscanln(in, &caseCount)

	for i := 0; i < caseCount; i++ {
		var n, r int
		fmt.Fscan(in, &n)
		fmt.Fscanln(in, &r)

		k := moveCount(n, r)

		fmt.Fprintln(out, k)
	}
}

func moveCount(n, r int) int {
	if r >= n {
		return 1
	}

	k := n / 2

	re := int(math.Ceil(float64(k) / float64(r)))

	if re < 1 {
		return 1
	}
	return re
}
