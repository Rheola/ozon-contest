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
		runCase(in, out)
	}
	out.Flush()
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var k int
	fmt.Fscanln(in, &k)

	mmap := make(map[string]int)

	for i := 0; i < k; i++ {
		var s string
		fmt.Fscanln(in, &s)

		ns := simplify(s)
		mmap[ns]++
	}

	l := len(mmap)
	fmt.Fprintln(out, l)
}

func simplify(s string) string {
	bb := []rune(s)
	res := make([]rune, 0)

	for i, v := range bb {
		if i+2 < len(bb) {
			if v == bb[i+1] && v == bb[i+2] {
				continue
			}
		}
		res = append(res, v)
	}

	//товарные знаки Booble и Boooble похожи — берём в первом товарном знаке две подряд идущие буквы o и добавляем к ним ещё одну букву o, так получим второй товарный знак Boooble;
	//товарные знаки yyyess и yyessss похожи — сначала изменим второй товарный знак: yyessss →
	//yyyessss, потом два раза изменим первый товарный знак: yyyess → yyyesss →
	//yyyessss. Так с помощью последовательности операций получилось сделать оба знака равными yyyessss;
	//товарные знаки oooops и oooops похожи — операций производить не надо, знаки уже равны;
	//товарные знаки oooooopppss и ooooppssss похожи — например, сначала изменим второй товарный знак ooooppssss →
	//oooooppssss → ooooooppssss → oooooopppssss, затем изменим первый: oooooopppss → oooooopppsss → oooooopppssss.

	return string(res)
}
