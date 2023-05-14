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
		res := runCase(in)
		if res {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func runCase(in *bufio.Reader) bool {
	var k int
	fmt.Fscanln(in, &k)

	tasks := make([]int, k)
	for i := 0; i < k; i++ {
		var task int

		fmt.Fscan(in, &task)
		tasks[i] = task
	}
	fmt.Fscanln(in)

	return validateTasks(tasks)
}

func validateTasks(tasks []int) bool {
	if len(tasks) == 1 {
		return true
	}

	tasksMap := make(map[int]int)
	for i, task := range tasks {
		_, ok := tasksMap[task]
		if !ok {
			tasksMap[task]++
			continue
		}

		if i > 0 {
			if tasks[i-1] != task {
				return false
			}
		}
	}

	return true
}
