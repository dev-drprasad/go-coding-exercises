package main

import (
	"fmt"
	"math"
)

type node struct {
	data  float64
	left  *node
	right *node
}

func dfs(node *node, min float64, max float64, diff float64) float64 {
	if node == nil {
		return diff
	}

	diff1 := math.Abs(float64(node.data - min))
	diff2 := math.Abs(float64(node.data - max))

	diff = math.Max(diff1, diff2)
	leftdiff := dfs(node.left, math.Min(node.data, min), math.Max(node.data, max), diff)
	rightdiff := dfs(node.right, math.Min(node.data, min), math.Max(node.data, max), diff)
	return math.Max(leftdiff, rightdiff)
}

func main() {
	n := node{8,
		&node{3,
			&node{1,
				nil, nil},
			&node{6,
				&node{4, nil, nil},
				&node{7, nil, nil}}},
		&node{10,
			nil, &node{14,
				&node{16, nil, nil},
				nil}}}

	diff := dfs(&n, n.data, n.data, 0)
	fmt.Println("diff", diff)
}
