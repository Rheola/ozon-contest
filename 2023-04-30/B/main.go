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

	for c := 0; c < n; c++ {
		runCase(in, out)
	}
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscanln(in, &s)

	fmt.Fprintln(out, pl(s))
}

func pl(word string) string {
	case1 := []string{"s", "sh", "ch", "x", "z"}

	for _, suf := range case1 {
		has := strings.HasSuffix(word, suf)
		if has {
			return word + "es"
		}
	}

	has := strings.HasSuffix(word, "y")
	if has {
		sogl := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z"}
		for _, suf := range sogl {
			suf = suf + "y"
			has := strings.HasSuffix(word, suf)
			if has {
				l := utf8.RuneCount([]byte(word))
				return string([]rune(word)[:l-1]) + "ies"
			}
		}
	}

	return word + "s"
}
