package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var strCount int
	fmt.Fscanln(in, &strCount)

	var line string
	fmt.Fscanln(in, &line)

	rss := taskToPrint(line, strCount)
	fmt.Fprintln(out, rss)
}

func taskToPrint(line string, strCount int) string {
	parts := strings.Split(line, ",")

	pages := make([]bool, strCount)

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			pages[num-1] = true
			continue
		}

		pp := strings.Split(part, "-")

		leftS := pp[0]
		left, _ := strconv.Atoi(leftS)

		rS := pp[1]
		right, _ := strconv.Atoi(rS)

		for left <= right {
			pages[left-1] = true
			left++
		}
	}

	res := make([]string, 0)

	left := -1
	r := -1
	for i, printed := range pages {
		if printed {
			continue
		}
		if i == 0 {
			left = 0
		}

		// начало блока
		if i > 0 && pages[i-1] {
			left = i
		}

		// конец блока
		if i+1 != strCount && pages[i+1] {
			r = i
			if r == left {
				res = append(res, fmt.Sprintf("%d", left+1))
			} else {
				res = append(res, fmt.Sprintf("%d-%d", left+1, r+1))
			}
		}

		// последняя страница
		if i+1 == strCount {
			r = i
			if r == left {
				res = append(res, fmt.Sprintf("%d", left+1))
			} else {
				res = append(res, fmt.Sprintf("%d-%d", left+1, r+1))
			}
		}
	}

	return strings.Join(res, ",")
}
