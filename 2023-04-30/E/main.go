package main

import (
	"bufio"
	"fmt"
	"os"
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
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var line string

	fmt.Fscanln(in, &line)

	lines := convert(line)

	for _, l := range lines {
		fmt.Fprintln(out, l)
	}

	fmt.Fprintln(out, "-")
}

func convert(line string) []string {
	syms := make([][]string, 1, 50)

	runes := []rune(line)

	posX := 0
	posROW := 0
	syms[0] = make([]string, 0)
	for _, r := range runes {
		s := string(r)
		switch s {
		// влево
		case "L":
			if posX > 0 {
				posX--
			}
			// вправо
		case "R":
			if posX < len(syms[posROW]) {
				posX++
			}
		//	вверх. Они перемещают курсор на одну позицию вверх.
		//	Если в соответствующем направлении нет строки, то операция игнорируется.
		//	Если строка есть, но в ней нужная позиция не существует, то курсор встаёт в конец строки.
		case "U":
			if posROW > 0 {
				posROW--
				if len(syms[posROW]) < posX {
					posX = len(syms[posROW])
				}
			}
		//	вниз. Они перемещают курсор на одну позицию вниз.
		//	Если в соответствующем направлении нет строки, то операция игнорируется.
		//	Если строка есть, но в ней нужная позиция не существует, то курсор встаёт в конец строки
		case "D":
			if posROW+1 < len(syms) {
				posROW++
				if len(syms[posROW]) < posX {
					posX = len(syms[posROW])
				}
			}
		// Home
		case "B":
			posX = 0
		// End
		case "E":
			posX = len(syms[posROW])
			//  Если курсор находился не в конце текущей строки, то она разрывается, и часть после курсора переносится в новую строку.
			//  Курсор после этой операции стоит в начале новой строки.
		case "N":
			if posX == len(syms[posROW]) {
				row := make([]string, 0)
				end := posROW + 1
				tail := append(syms[:end:end], row)
				syms = append(tail, syms[end:]...)
				posROW++
				posX = 0
			} else {
				beforeLines := syms[:posROW:posROW]
				afterLines := syms[posROW+1:]

				row := syms[posROW]

				rowBefore := row[:posX:posX]
				rowAfter := row[posX:]

				if len(beforeLines) > 0 {
					beforeLines = append(beforeLines, rowBefore)
				} else {
					beforeLines = [][]string{rowBefore}
				}

				beforeLines = append(beforeLines, rowAfter)

				syms = append(beforeLines, afterLines...)
				posROW++
				posX = 0
			}
		default:
			if posX == 0 {
				tail := []string{string(r)}

				syms[posROW] = append(tail, syms[posROW]...)
				posX++
				continue
			}
			// последний
			if posX == len(syms[posROW]) {
				syms[posROW] = append(syms[posROW], string(r))
				posX++
				continue
			}

			row := syms[posROW]
			end := posX
			tl := row[:end:end]
			tl = append(tl, s)
			tl = append(tl, row[end:]...)
			syms[posROW] = tl
			posX++
		}
	}

	lines := make([]string, len(syms))
	for i, l := range syms {
		lines[i] = strings.Join(l, "")
	}
	return lines
}
