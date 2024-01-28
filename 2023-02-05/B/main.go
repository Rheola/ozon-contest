package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		runCase(in, out)
	}

	out.Flush()
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	m := make(map[string]int)

	for i := 0; i < 10; i++ {
		var s string

		fmt.Fscan(in, &s)

		m[s]++
	}

	if validate(m) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func validate(m map[string]int) bool {
	if m["1"] != 4 {
		return false
	}

	if m["2"] != 3 {
		return false
	}

	if m["3"] != 2 {
		return false
	}

	if m["4"] != 1 {
		return false
	}

	return true
}
