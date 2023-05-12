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
		s, _ := in.ReadString('\n')
		s = strings.Trim(s, "\n")
		syms := strings.Split(s, " ")
		nums := make(map[int]struct{})

		for _, d := range syms {
			dd, _ := strconv.Atoi(d)
			nums[dd] = struct{}{}
		}
		input[i] = nums
	}

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

	for funcName(input, subordinates, leaders) {

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
	//fmt.Fprintln(out)
}

func funcName(input []map[int]struct{}, subordinates map[int]Node, leaders map[int]Node) bool {
	for lnn, line := range input {
		if len(line) == 0 {
			continue
		}

		line2 := make(map[int]struct{})
		for key, value := range line {
			line2[key] = value
		}

		// удаляем подчиненных
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

		// в списке остался только руководитель
		leadID := 0
		for id := range line2 {
			leadID = id
		}

		for staffID := range line {
			if staffID == leadID {
				continue
			}

			// подчиненный
			nd := subordinates[staffID]
			leadNode := leaders[leadID]
			leadNode.child = append(leadNode.child, &nd)
			leaders[leadID] = leadNode

			nd.parent = &leadNode

			subordinates[staffID] = nd
			delete(line, staffID)

			deleteSubFromList(input, staffID)

			subordinates[leadID] = leadNode

			t := leadNode.LenChain()
			t++
			if t != len(line2) {
				return false
			}
		}

		delete(leaders, leadID)

		input[lnn] = map[int]struct{}{}

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
