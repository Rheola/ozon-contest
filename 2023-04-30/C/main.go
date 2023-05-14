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
}

func runCase(in *bufio.Reader, out *bufio.Writer) {
	var k int
	fmt.Fscanln(in, &k)

	nodes := make([]int, k)
	for i := 0; i < k; i++ {
		var node int

		fmt.Fscan(in, &node)
		nodes[i] = node
	}
	fmt.Fscanln(in)

	t := buildTree(nodes)
	r := findRoot(t)
	fmt.Fprintln(out, r)
}

func buildTree(a []int) map[int]Node {
	tree := make(map[int]Node)

	node := Node{}

	for i := 0; i < len(a); i++ {
		val := a[i]
		node.num = val
		node.chCount = a[i+1]
		node.children = make([]int, 0, node.chCount)
		for k := i + 2; k < i+2+node.chCount; k++ {
			node.children = append(node.children, a[k])
		}
		tree[val] = node
		i = i + 1 + a[i+1]
	}

	return tree
}

type Node struct {
	num      int
	chCount  int
	children []int
}

func findRoot(tree map[int]Node) int {
	for r := range tree {
		isRoot := true
		for rr, val := range tree {
			if r == rr {
				continue
			}
			for _, cd := range val.children {
				if r == cd {
					isRoot = false
				}
			}
		}

		if isRoot {
			return r
		}
	}

	return 0
}
