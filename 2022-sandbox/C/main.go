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
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		runCase(in, out)
	}
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	staffs := make([]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		staffs[i] = t
	}

	res := funcName(staffs)

	for i := 0; i < n/2; i++ {
		fmt.Fprintln(out, res[i][0], res[i][1])
	}

	fmt.Fprintln(out)
}

func funcName(staffs []int) [][]int {
	n := len(staffs)
	res := make([][]int, n/2)
	for i := 0; i < n/2; i++ {
		res[i] = make([]int, 2)
		firstID := -1
		secondID := -1
		firstLevel := 0
		diff := 1001
		for j := 0; j < n; j++ {
			if firstID == -1 && staffs[j] > 0 {
				firstID = j
				firstLevel = staffs[j]

				staffs[j] = 0
			}

			if firstID >= 0 && firstID != j && staffs[j] > 0 {
				d := abc(firstLevel, staffs[j])
				if d < diff {
					diff = d
					secondID = j
				}
			}
		}
		staffs[secondID] = 0
		res[i][0] = firstID + 1
		res[i][1] = secondID + 1

	}
	return res
}

func abc(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
