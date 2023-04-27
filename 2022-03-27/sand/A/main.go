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

	var i int

	fmt.Fscan(in, &i)

	i = i / 10000000

	f := i / 1000

	rest := i - f*1000
	fmt.Fprintln(out, rest)
}
