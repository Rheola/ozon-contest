package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	fmt.Fscanln(in, &n)

	for c := 0; c < n; c++ {
		runCase(in, out)
	}
}

type dataStruct struct {
	n          int
	m          int
	clickCount int
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	fmt.Fscanln(in)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var ds dataStruct

	ds.n = n
	ds.m = m
	arr := make([][]int, ds.n)
	for i := 0; i < ds.n; i++ {
		arr[i] = make([]int, ds.m)
		for j := 0; j < ds.m; j++ {
			var t int
			fmt.Fscan(in, &t)
			arr[i][j] = t
		}
		fmt.Fscanln(in)
	}

	fmt.Fscan(in, &ds.clickCount)

	clicks := make([]int, ds.clickCount)

	for i := 0; i < ds.clickCount; i++ {
		var t int
		fmt.Fscan(in, &t)

		clicks[i] = t
	}

	for _, row := range clicks {
		sort.SliceStable(arr, func(i, j int) bool {
			ai := arr[i][row-1]
			aj := arr[j][row-1]
			return ai < aj
		})
	}

	for _, line := range arr {
		chars := make([]string, len(line))
		for i, el := range line {
			chars[i] = strconv.Itoa(el)
			fmt.Fprint(out, el, " ")
		}
		fmt.Fprintln(out)
	}

	fmt.Fprintln(out)
	fmt.Fscanln(in)
}
