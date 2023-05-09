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

	for i := 0; i < n; i++ {
		runCase(in, out)
	}
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var n int // количество позиций в заказе.
	fmt.Fscanln(in, &n)

	var maxSum int // максимальная сумма по чеку.
	fmt.Fscanln(in, &maxSum)

	posMap := make(map[int]position)

	pos := 0
	for i := 0; i < n; i++ {
		var art, sum int
		fmt.Fscan(in, &art, &sum)
		fmt.Fscanln(in)

		val, ok := posMap[art]
		if ok {
			val.sum += sum
			posMap[art] = val
			continue
		}

		val.art = art
		val.sum = sum
		val.pos = pos
		posMap[art] = val
		pos++
	}

	receipts := convert(posMap, maxSum)

	fmt.Fprintln(out, len(receipts))
	for _, r := range receipts {
		fmt.Fprintln(out, len(r.s))
		for _, b := range r.s {
			fmt.Fprintln(out, b.art, b.sum)
		}
	}
}

func convert(posMap map[int]position, maxSum int) []receipt {
	posSlice := make([]position, len(posMap))
	for _, v := range posMap {
		posSlice[v.pos] = v
	}

	total := 0
	receipts := make([]receipt, 1)
	for i := 0; i < len(posSlice); {
		cur := posSlice[i]
		last := receipts[len(receipts)-1]
		// надо добавить к последнему чеку
		if total < maxSum {
			diff := maxSum - total
			ps := position{
				art: cur.art,
			}

			if cur.sum <= diff {
				ps.sum = cur.sum
				cur.sum = 0
				posSlice[i] = cur
				i++
			} else {
				ps.sum = diff
				cur.sum = cur.sum - diff
				posSlice[i] = cur
			}
			total += ps.sum
			last.s = append(last.s, ps)
			receipts[len(receipts)-1] = last
			continue
		}

		ps := position{
			art: cur.art,
		}
		r := receipt{}
		if cur.sum > maxSum {
			ps.sum = maxSum
			cur.sum = cur.sum - maxSum
			posSlice[i] = cur
		} else {
			ps.sum = cur.sum
			cur.sum = 0
			posSlice[i] = cur
			i++
		}
		total = ps.sum
		r.s = append(r.s, ps)
		receipts = append(receipts, r)
	}
	return receipts
}

type position struct {
	art int
	pos int
	sum int
}

type receipt struct {
	s []position
}
