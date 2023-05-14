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
		isValid := runCase(in)
		if isValid {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func runCase(in *bufio.Reader) bool {
	var m, n int
	fmt.Fscan(in, &m, &n)

	color := make([][]string, m)
	for i := 0; i < m; i++ {
		var s string
		fmt.Fscan(in, &s)

		arr := []rune(s)
		color[i] = make([]string, len(arr))
		for k, r := range arr {
			el := string([]rune{r})
			color[i][k] = el
		}
	}

	return validate(color)
}

func validate(color [][]string) bool {
	usedColors := make(map[string]struct{})

	for i := 0; i < len(color); i++ {
		for j := 0; j < len(color[i]); j++ {
			if color[i][j] == "." || color[i][j] == "-" {
				continue
			}
			_, ok := usedColors[color[i][j]]
			if ok {
				return false
			}

			usedColors[color[i][j]] = struct{}{}

			fill(color, i, j)
		}
	}

	return true
}

func fill(color [][]string, i int, j int) {
	s := stack{}
	var p point
	s = s.Push(newPoint(i, j))
	currentColor := color[i][j]

	for !s.IsEmpty() {
		s, p = s.Pop()
		directions := getDirections(p.x, p.y)
		for _, d := range directions {
			if d.x < 0 || d.y < 0 {
				continue
			}
			if d.x >= len(color) || d.y >= len(color[d.x]) {
				continue
			}
			if color[d.x][d.y] != currentColor {
				continue
			}

			s = s.Push(newPoint(d.x, d.y))
			color[d.x][d.y] = "-"

		}

	}

}

func getDirections(i int, j int) []point {
	res := make([]point, 6)

	res[0] = newPoint(i-1, j-1)
	res[1] = newPoint(i-1, j+1)
	res[2] = newPoint(i, j-2)
	res[3] = newPoint(i, j+2)
	res[4] = newPoint(i+1, j-1)
	res[5] = newPoint(i+1, j+1)

	return res
}

type point struct {
	x int
	y int
}

func newPoint(x int, y int) point {
	return point{x: x, y: y}
}

type stack []point

func (s stack) Push(v point) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, point) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}
