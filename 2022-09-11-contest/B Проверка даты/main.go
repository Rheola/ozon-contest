package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
	var d, m, y int
	fmt.Fscan(in, &d, &m, &y)

	if d <= 0 || m <= 0 || y <= 0 || m > 12 {
		fmt.Fprintln(out, "NO")
		return
	}

	mo := time.Month(m)

	dd := time.Date(y, mo, d, 0, 0, 0, 0, time.UTC)

	if dd.Day() != d || dd.Month() != mo || dd.Year() != y {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
}
