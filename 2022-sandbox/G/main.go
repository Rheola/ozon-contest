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

	var peopleCount int
	fmt.Fscan(in, &peopleCount)

	var n int
	fmt.Fscan(in, &n)

	type friendList []int
	data := make([]friendList, peopleCount+1)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		data[a-1] = append(data[a-1], b)
		data[b-1] = append(data[b-1], a)
	}

	for i := 1; i <= peopleCount; i++ {
		frListI := data[i-1]
		friendMap := make(map[int]struct{})

		for _, friendNum := range frListI {
			friendMap[friendNum] = struct{}{}
		}

		fr := make(map[int]int)
		for _, friendNum := range frListI {
			// перебор друзей
			for _, pf := range data[friendNum-1] {
				if pf == i {
					continue
				}
				_, ok := friendMap[pf]
				if ok {
					continue
				}

				fr[pf]++
			}
		}

		for _, s := range probableFriends(fr) {
			fmt.Fprint(out, s, " ")
		}

		fmt.Fprintln(out)
	}
}

func probableFriends(friends map[int]int) []int {
	if len(friends) == 0 {
		return []int{0}
	}

	max := 0
	for _, count := range friends {
		if max <= count {
			max = count
		}
	}

	res := make([]int, 0)
	for fr, count := range friends {
		if max == count {
			res = append(res, fr)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
