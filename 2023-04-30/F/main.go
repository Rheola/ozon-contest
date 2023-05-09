package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		runCase(in, out)
	}
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var k int
	fmt.Fscanln(in, &k)

	arr := make([]int, k)
	for i := 0; i < k; i++ {
		var t int
		fmt.Fscan(in, &t)
		arr[i] = t

	}
	fmt.Fscanln(in)

	res := solve(arr)

	fmt.Fprintln(out, res)
}

func solve(a []int) string {
	if len(a) == 1 {
		return "!"
	}

	minpos := minPos(a)
	if minpos == 0 {
		return "!" + solve(a[1:])
	}

	if minpos > len(a)/2 {
		rotateL(a, minpos)
		tail := a[1:]
		return strings.Repeat("R", len(a)-minpos) + "!" + solve(tail)
	}

	rotateL(a, minpos)
	return strings.Repeat("L", minpos) + "!" + solve(a[1:])
}

func minPos(a []int) int {
	index := 0
	min := a[index]
	for i, v := range a {
		if v < min {
			min = v
			index = i
		}
	}

	return index
}

func rotateL(nums []int, k int) {
	if k > len(nums) {
		r := k % len(nums)
		k = r
	}

	for i := 0; i < k/2; i++ {
		a := nums[i]
		nums[i] = nums[k-1-i]
		nums[k-1-i] = a
	}
	r := len(nums)

	mid := (len(nums) - k) / 2
	for i := 0; i < mid; i++ {
		a := nums[k+i]
		nums[k+i] = nums[r-1-i]
		nums[r-1-i] = a
	}

	for i := 0; i < r/2; i++ {
		a := nums[i]
		nums[i] = nums[r-1-i]
		nums[r-1-i] = a
	}
}
