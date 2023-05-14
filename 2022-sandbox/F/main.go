package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countCases int
	fmt.Fscan(in, &countCases)

	for i := 0; i < countCases; i++ {
		val := runCase(in)
		if val {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}

type tRange struct {
	from time.Time
	to   time.Time
}

func runCase(in *bufio.Reader) bool {
	var n int
	fmt.Fscan(in, &n)

	cases := make([]tRange, n)
	layout := "2006-01-02 15:04:05"

	hardError := false
	for i := 0; i < n; i++ {
		var l string
		fmt.Fscan(in, &l)

		times := strings.Split(l, "-")

		from, err := time.Parse(layout, "2006-01-02 "+times[0])
		if err != nil {
			hardError = true
			continue
		}

		to, err := time.Parse(layout, "2006-01-02 "+times[1])
		if err != nil {
			hardError = true
			continue
		}

		if to.Before(from) {
			hardError = true
			continue
		}

		cases[i] = tRange{
			from: from,
			to:   to,
		}
	}
	if hardError {
		return false
	}

	sort.Slice(cases, func(i, j int) bool {
		return cases[i].from.Before(cases[j].from)
	})

	for j, t := range cases {
		if j+1 < len(cases) {
			next := cases[j+1].from

			if t.to.Equal(next) {
				return false
			}

			if t.to.After(next) {
				return false
			}
		}
	}

	return true
}
