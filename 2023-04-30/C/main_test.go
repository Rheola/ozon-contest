package main

import (
	"testing"
)

func Test_runCase(t *testing.T) {

	tests := []struct {
		name string
		arr  []int
		root int
	}{
		{
			name: "",
			arr:  []int{3, 0, 1, 0, 5, 2, 2, 6, 4, 3, 5, 1, 3, 2, 0, 6, 0},
			root: 4,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := buildTree(tt.arr)
			if got := findRoot(tr); got != tt.root {
				t.Errorf("solve() = %v, want %v", got, tt.root)
			}
		})
	}
}
