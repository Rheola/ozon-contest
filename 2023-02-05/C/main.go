package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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

		res, err := decode(s)
		if err != nil {
			fmt.Fprintln(out, "-")
			continue
		}

		for _, r := range res {
			fmt.Fprint(out, r, " ")
		}
		fmt.Fprintln(out)

	}

	out.Flush()
}

func decode(s string) ([]string, error) {
	res := make([]string, 0)

	if utf8.RuneCountInString(s) <= 3 {
		return []string{}, fmt.Errorf("wrong")
	}

	if utf8.RuneCountInString(s) == 4 {
		if valid4(s) {
			return []string{s}, nil

		}

		return []string{}, fmt.Errorf("wrong")
	}

	if utf8.RuneCountInString(s) == 5 {
		if valid5(s) {
			return []string{s}, nil
		}

		return []string{}, fmt.Errorf("wrong")
	}

	result := []rune(s)       // convert string to runes
	s4 := string(result[0:4]) // convert runes back to string

	if valid4(s4) {
		r4, err := decode(string(result[4:]))
		if err == nil {
			res = append(res, s4)
			res = append(res, r4...)
			return res, nil
		}
	}

	s5 := string(result[0:5]) // convert runes back to string

	if valid5(s5) {
		r5, err := decode(string(result[5:]))
		if err == nil {
			res = append(res, s5)
			res = append(res, r5...)
			return res, nil
		}
	}

	return []string{}, fmt.Errorf("wrong")
}

func valid4(s string) bool {
	// либо автомобильный номер имеет вид буква-цифра-буква-буква
	// (примеры корректных номеров второго вида: T7RR, A9PQ, O0OO).
	bytes := []rune(s)

	if !unicode.IsLetter(bytes[0]) {
		return false
	}

	if !unicode.IsDigit(bytes[1]) {
		return false
	}

	if !unicode.IsLetter(bytes[2]) {
		return false
	}

	if !unicode.IsLetter(bytes[3]) {
		return false
	}

	return true

}

func valid5(s string) bool {
	// 	  автомобильный номер имеет вид буква-цифра-цифра-буква-буква
	//	(примеры корректных номеров первого вида: R48FA, O00OO, A99OK);
	bytes := []rune(s)

	if !unicode.IsLetter(bytes[0]) {
		return false
	}

	if !unicode.IsDigit(bytes[1]) {
		return false
	}

	if !unicode.IsDigit(bytes[2]) {
		return false
	}

	if !unicode.IsLetter(bytes[3]) {
		return false
	}

	if !unicode.IsLetter(bytes[4]) {
		return false
	}

	return true

}
