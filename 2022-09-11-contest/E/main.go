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

	var cases int
	fmt.Fscan(in, &cases)

	for caseI := 0; caseI < cases; caseI++ {
		var x, y int
		fmt.Fscan(in, &x, &y)

		myMap := make([][]string, x)

		for i := 0; i < x; i++ {
			var s string
			fmt.Fscan(in, &s)

			arr := []rune(s)
			myMap[i] = make([]string, len(arr))
			for k, r := range arr {
				el := string([]rune{r})
				myMap[i][k] = el
			}
		}

		/*		for z := 0; z < x; z++ {
				for r := 0; r < y; r++ {
					fmt.Fprint(out, myMap[z][r])
				}
				fmt.Fprintln(out)
			}*/

		start := NewPoint(0, 0)

		flag := false
		for i := 0; i < x && flag == false; i++ {
			for j := 0; j < y && flag == false; j++ {
				if myMap[i][j] == "*" {
					flag = isStart(NewPoint(i, j), myMap)
					if flag {
						start = NewPoint(i, j)
						break
					}
				}
			}
		}

		way := findWay(start, myMap)

		for _, w := range way {
			fmt.Fprint(out, w)
		}

		fmt.Fprintln(out)

	}
}

func findWay(p Point, myMap [][]string) []string {
	res := make([]string, 0)
	end := false

	way := Stack{}

	way.Push(p)

	for !way.IsEmpty() && end == false {
		next, _ := way.Pop()

		myMap[next.x][next.y] = "+"

		dirs := getDirections(next.x, next.y)
		for _, dir := range dirs {
			if dir.x < 0 || dir.y < 0 {
				continue
			}

			if dir.x >= len(myMap) || dir.y >= len(myMap[0]) {
				continue
			}

			if myMap[dir.x][dir.y] == "*" {
				res = append(res, dir.letter)
				way.Push(NewPoint(dir.x, dir.y))
				if isStart(NewPoint(dir.x, dir.y), myMap) {
					end = true
				}
				continue
			}

		}
	}

	return res
}

func isStart(p Point, myMap [][]string) bool {
	dirs := getDirections(p.x, p.y)

	way := 0
	for _, dir := range dirs {
		if dir.x < 0 || dir.y < 0 {
			continue
		}

		if dir.x >= len(myMap) || dir.y >= len(myMap[0]) {
			continue
		}

		if myMap[dir.x][dir.y] == "*" || myMap[dir.x][dir.y] == "+" {
			way++
		}
	}

	return way == 1
}

func getDirections(x, y int) []Point {
	res := make([]Point, 4)
	res[0] = NewPoint(x, y+1)
	res[0].letter = "R"

	res[1] = NewPoint(x, y-1)
	res[1].letter = "L"

	res[2] = NewPoint(x+1, y)
	res[2].letter = "D"

	res[3] = NewPoint(x-1, y)
	res[3].letter = "U"

	return res
}

type Point struct {
	x      int
	y      int
	letter string
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

type Stack []Point

// Push a new value onto the stack
func (s *Stack) Push(str Point) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (Point, bool) {
	if s.IsEmpty() {
		return Point{}, false
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true
}
