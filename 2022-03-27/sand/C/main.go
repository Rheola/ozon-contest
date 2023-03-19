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

	fmt.Fprintln(out, stepCount(n))
}

func stepCount(n int) int {
	p := fmt.Sprintf("%d", n)
	d := 0

	for !isPalindrome(p) {
		n++
		d++
		p = fmt.Sprintf("%d", n)
	}

	return d
}

func isPalindrome(s string) bool {
	d := []rune(s)
	n := len(d)
	for i := 0; i < n/2; i++ {
		if d[i] != d[n-1-i] {
			return false
		}
	}
	return true
}
