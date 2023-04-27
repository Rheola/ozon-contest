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

	var m int // кол-во друзей
	fmt.Fscanln(in, &m)

	for i := 0; i <= m; i++ {
		var n int // кол-во районов
		var t int // модуль

		fmt.Fscan(in, &n)
		fmt.Fscanln(in, &t)

		fmt.Fprintln(out, fact(n, t))
	}
}

func fact(n, t int) int {
	if n >= t {
		return 0
	}

	res := 1
	for i := 2; i <= n; i++ {
		res = (res * i) % t
	}

	return res
}
