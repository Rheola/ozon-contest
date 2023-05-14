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
	var k int
	fmt.Fscanln(in, &k)

	arr := make([]int, k)
	for i := 0; i < k; i++ {
		var t int
		fmt.Fscan(in, &t)
		arr[i] = t

	}
	fmt.Fscanln(in)

	res := compress(arr)
	fmt.Fprintln(out, len(res))

	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
func compress(arr []int) []int {
	if len(arr) == 1 {
		return []int{arr[0], 0}
	}
	res := make([]int, 0)

	var ci int
	temp := arr[0]
	toWrite := temp
	for i, v := range arr {
		if i == 0 {
			continue
		}

		if (ci < 0 || ci == 0) && v+1 == temp {
			ci--
			if i == len(arr)-1 {
				res = append(res, toWrite)
				res = append(res, ci)
			}

			temp--
			continue
		}

		if (ci > 0 || ci == 0) && v-1 == temp {
			ci++
			if i == len(arr)-1 {
				res = append(res, toWrite)
				res = append(res, ci)
			}

			temp++
			continue
		}

		if ci > 0 || ci < 0 {
			res = append(res, toWrite)
			res = append(res, ci)
			ci = 0

			temp = v
			toWrite = temp
			if i+1 == len(arr) {
				res = append(res, v)
				res = append(res, 0)
			}
			continue
		}

		res = append(res, toWrite)
		res = append(res, 0)

		if i+1 != len(arr) {
			toWrite = v
			temp = v
		} else {
			res = append(res, v)
			res = append(res, 0)
		}
	}

	return res
}
