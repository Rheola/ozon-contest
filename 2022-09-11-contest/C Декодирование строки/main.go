package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)

		res, _ := decode(s)
		for _, r := range res {
			fmt.Fprint(out, r)
		}
		fmt.Fprintln(out)

	}
}

func decode(s string) ([]string, error) {
	res := make([]string, 0)

	if utf8.RuneCountInString(s) == 1 {
		return []string{}, fmt.Errorf("wrong")
	}

	if utf8.RuneCountInString(s) == 2 {
		if s == "00" {
			return []string{"a"}, nil
		}

		if s == "11" {
			return []string{"d"}, nil
		}

		return []string{}, fmt.Errorf("wrong")
	}

	if utf8.RuneCountInString(s) == 3 {
		if s == "100" {
			return []string{"b"}, nil
		}

		if s == "101" {
			return []string{"c"}, nil
		}

		return []string{}, fmt.Errorf("wrong")
	}

	result := []rune(s)       // convert string to runes
	s2 := string(result[0:2]) // convert runes back to string

	if s2 == "00" {
		r2, err := decode(string(result[2:]))
		if err == nil {
			res = append(res, "a")
			res = append(res, r2...)
			return res, nil
		}
	}

	if s2 == "11" {
		r2, err := decode(string(result[2:]))
		if err == nil {
			res = append(res, "d")
			res = append(res, r2...)
			return res, nil
		}
	}

	s3 := string(result[0:3]) // convert runes back to string
	if s3 == "100" {
		r3, err := decode(string(result[3:]))
		if err == nil {
			res = append(res, "b")
			res = append(res, r3...)
			return res, nil
		}
	}

	if s3 == "101" {
		r3, err := decode(string(result[3:]))
		if err == nil {
			res = append(res, "c")
			res = append(res, r3...)
			return res, nil
		}
	}

	return []string{}, fmt.Errorf("wrong")
}

//каждая буква 'c' была закодирована как 101;
//каждая буква 'd' была закодирована как 11.
