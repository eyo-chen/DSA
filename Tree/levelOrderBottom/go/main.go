package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// using recursion to count max depth first
// then using recursion to fill the result
func LevelOrderBottom(root *TreeNode) [][]int {
	// get the max depth of the tree
	depth := countDepth(root)
	res := make([][]int, depth)
	helper(root, &res, depth-1)

	return res
}

func helper(node *TreeNode, res *[][]int, level int) {
	if node == nil {
		return
	}

	// fill the result from the bottom to the top
	(*res)[level] = append((*res)[level], node.Val)
	helper(node.Left, res, level-1)
	helper(node.Right, res, level-1)
}

func countDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return int(math.Max(float64(countDepth(node.Left)), float64(countDepth(node.Right)))) + 1
}
