package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	out.Flush()
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var k int
	fmt.Fscanln(in, &k)

	res := make([]people, k)
	for i := 0; i < k; i++ {
		var p int
		fmt.Fscan(in, &p)

		ppl := people{
			num:    i + 1,
			score:  p,
			score2: p,
		}

		res[i] = ppl
	}

	fmt.Fscanln(in)

	sort.Slice(res, func(i, j int) bool {
		return res[i].score < res[j].score
	})
	changed := normalise(res)
	for changed {
		changed = normalise(res)
	}

	place := 1
	current := res[0].score2
	nextPlace := place
	for i, ppl := range res {
		if ppl.score2 == current {
			res[i].place = place
			nextPlace++
			continue
		}
		place = nextPlace
		current = res[i].score2
		res[i].place = place
		nextPlace++
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].num < res[j].num
	})

	for _, re := range res {
		fmt.Fprint(out, re.place, " ")
	}

	fmt.Fprintln(out)
}

func normalise(res []people) bool {
	changed := false
	k := len(res)
	for i := k - 1; i > 0; i-- {
		if i == 0 {
			break
		}

		r := res[i].score2
		l := res[i-1].score2

		if r-1 == l {
			res[i].score2--
			changed = true
		}
	}

	return changed
}

type people struct {
	num    int
	score  int
	score2 int
	place  int
}
