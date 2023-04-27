package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscanln(in, &n)

	res := make([]bool, n)

	for i := 0; i < n; i++ {
		res[i] = runCase(in)
	}

	for _, val := range res {
		if val {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func runCase(in *bufio.Reader) bool {
	var k int
	fmt.Fscanln(in, &k)

	var s, t string
	fmt.Fscanln(in, &s)
	fmt.Fscanln(in, &t)
	d := check(s, t)
	return d
}

func check(s, t string) bool {
	as := strings.Split(s, "")
	at := strings.Split(t, "")

	sort.Strings(as)
	sort.Strings(at)
	at = reverse(at)

	s = strings.Join(as[:], "")
	t = strings.Join(at[:], "")

	return s < t
}

func reverse(numbers []string) []string {
	newNumbers := make([]string, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}

	return newNumbers
}
