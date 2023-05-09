package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
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
	var n int // n различных слов
	var b int // b синих слов
	var r int // r красных слов
	var f int // f — номер чёрного слова

	fmt.Fscan(in, &n, &b, &r, &f)
	fmt.Fscanln(in)

	blue := make([]string, 0, b)
	red := make([]string, 0, r)
	black := ""
	other := make([]string, 0, n-b-r-1)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanln(in, &s)

		if i+1 == f {
			black = s
		}
		if i < b {
			blue = append(blue, s)
			continue
		}

		if i < b+r {
			red = append(red, s)
			continue
		}

		other = append(other, s)
	}

	//redMap := make(map[string][]string)
	blueSubstMap := substringMap(blue)
	redMap := substringMap(red)

	bMap := make(map[string]struct{})
	rMap := make(map[string]struct{})
	oMap := make(map[string]struct{})

	for _, c := range blue {
		bMap[c] = struct{}{}
	}

	for _, c := range red {
		rMap[c] = struct{}{}
	}

	for _, c := range other {
		oMap[c] = struct{}{}
	}

	max := 0
	answer := ""

	for bSubSt, val := range blueSubstMap {
		_, ok := rMap[bSubSt]
		if ok {
			continue
		}

		_, ok = bMap[bSubSt]
		if ok {
			continue
		}

		_, ok = oMap[bSubSt]
		if ok {
			continue
		}

		if strings.Contains(black, bSubSt) {
			continue
		}

		bc := len(val.list)
		rc := 0

		rs, ok := redMap[bSubSt]
		if ok {
			rc = len(rs.list)
		}

		d := bc - rc
		if d > max {
			max = d
			answer = bSubSt
		}

		if d > 1 && d == max {
			if utf8.RuneCountInString(bSubSt) > utf8.RuneCountInString(answer) {
				answer = bSubSt
			}
		}
	}

	if answer == "" {
		for sb := range blueSubstMap {
			alf := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"}
			for _, sss := range alf {
				if utf8.RuneCountInString(sss) <= 9 {
					nw := sb + sss

					_, ok := redMap[nw]
					if ok {
						continue
					}

					_, ok = blueSubstMap[nw]
					if ok {
						continue
					}

					_, ok = oMap[nw]
					if ok {
						continue
					}

					if !strings.Contains(black, nw) {
						fmt.Fprintln(out, nw, max)
						return
					}
				}
			}
		}
	}

	fmt.Fprintln(out, answer, max)
}

func substringMap(blue []string) map[string]info {
	blueMap := make(map[string]info)

	for _, word := range blue {
		w := []rune(word)
		n := len(w)
		// Pick starting point
		for l := 1; l <= len(w); l++ {
			// Pick ending point
			for i := 0; i <= n-l; i++ {
				// Print characters from current
				// starting point to current ending
				// point.
				j := i + l - 1
				for k := i; k <= j; k++ {
					wrd := string(w[i : j+1])
					_, ok := blueMap[wrd]
					if ok {
						blueMap[wrd].list[word] = struct{}{}
					} else {
						elem := info{}
						elem.list = make(map[string]struct{})
						elem.list[word] = struct{}{}

						blueMap[wrd] = elem
					}
				}
			}
		}
	}

	return blueMap
}

type info struct {
	list map[string]struct{}
}
