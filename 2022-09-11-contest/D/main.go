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
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		tasks := make([]int, k)
		for j := 0; j < k; j++ {
			var val int
			fmt.Fscan(in, &val)
			tasks[j] = val
		}
		mt := maxTask(tasks)
		fmt.Fprintln(out, mt)
	}
}

func maxTask(numbers []int) int {
	max := 0
	if has2(numbers) {
		return len(numbers)
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i < j {
				if has2(numbers[i : j+1]) {
					ll := len(numbers[i : j+1])
					if ll > max {
						max = ll
					}
				}
			}
		}
	}

	return max
}

func has2(numbers []int) bool {
	if len(numbers) <= 1 {
		return true
	}

	a := numbers[0]
	b := numbers[1]
	for i := 1; i < len(numbers); i++ {
		if a == b {
			b = numbers[i]
			continue
		}
		if numbers[i] != a && numbers[i] != b {
			return false
		}
	}

	return true
}
