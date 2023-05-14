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
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscanln(in, &n)

	input := make([]map[int]struct{}, n)

	for i := 0; i < n; i++ {
		s, _ := Readln(in)

		s = strings.Trim(s, "\n")
		syms := strings.Split(s, " ")
		nums := make(map[int]struct{})
		for _, d := range syms {
			dd, _ := strconv.Atoi(d)
			nums[dd] = struct{}{}
		}
		input[i] = nums
	}

	//fmt.Fprintln(out, len(input))
	//
	//fmt.Fprintln(out, "=====")

	/*	for i, m := range input {
		fmt.Fprintln(out, "id", i)

		for j := range m {
			fmt.Fprint(out, j, " ")
		}
		fmt.Fprintln(out)
	}*/

	solve(input, out)
}

func solve(input []map[int]struct{}, out *bufio.Writer) {
	n := len(input)

	subordinates := make(map[int]Node)
	leaders := make(map[int]Node)

	for i, line := range input {
		if len(line) > 1 {
			for staffID := range line {
				node := Node{val: staffID}
				leaders[staffID] = node
			}
			continue
		}

		for staffID := range line {
			node := Node{val: staffID}
			subordinates[staffID] = node
		}
		// обнулили
		input[i] = map[int]struct{}{}
	}

	if len(subordinates) == 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	originalLen := make([]int, len(input))
	for i, m := range input {
		originalLen[i] = len(m)
	}

	for generateSubordinates(input, subordinates, leaders, originalLen) {

	}

	counthead := 0
	head := &Node{}
	for _, v := range subordinates {
		v := v
		if v.parent == nil {
			counthead++
			head = &v
		}
	}

	if counthead > 1 {
		fmt.Fprintln(out, "NO")
		return
	}

	if len(subordinates) != n {
		fmt.Fprintln(out, "NO")
		return
	}

	for _, v := range input {
		if len(v) > 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
	head.Print(out)
}

func generateSubordinates(input []map[int]struct{},
	subordinates map[int]Node,
	leaders map[int]Node,
	originalLen []int,
) bool {
	for i, line := range input {
		if len(line) == 0 {
			continue
		}

		line2 := make(map[int]struct{})
		for key, value := range line {
			line2[key] = value
		}

		// удаляем подчиненных из копии
		for staffID := range line {
			for subID := range subordinates {
				if subID == staffID {
					delete(line2, staffID)
				}
			}
		}

		if len(line2) > 1 {
			continue
		}

		if len(line2) == 0 {
			return false
		}

		// в списке остался только руководитель
		leadID := 0
		for id := range line2 {
			leadID = id
		}

		leadNode := leaders[leadID]

		for staffID := range line {
			if staffID == leadID {
				continue
			}

			// подчиненный
			nd := subordinates[staffID]
			leadNode.child = append(leadNode.child, &nd)

			nd.parent = &leadNode

			subordinates[staffID] = nd
			subordinates[leadID] = leadNode

			deleteSubFromList(input, staffID)
		}

		if leadNode.LenChain()+1 != originalLen[i] {
			return false
		}

		delete(leaders, leadID)

		input[i] = map[int]struct{}{}

		return true
	}

	return false
}

func deleteSubFromList(input []map[int]struct{}, staffID int) {
	for _, lines := range input {
		delete(lines, staffID)
	}
}

type Node struct {
	val    int
	parent *Node
	child  []*Node
}

func (n *Node) Print(out *bufio.Writer) {
	if n == nil {
		return
	}

	for _, c := range n.child {
		fmt.Fprintln(out, n.val, c.val)
		c.Print(out)
	}
}

func (n *Node) LenChain() int {
	if len(n.child) == 0 {
		return 0
	}

	t := 0
	for _, c := range n.child {
		t += c.LenChain()
	}

	t += len(n.child)

	return t
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
