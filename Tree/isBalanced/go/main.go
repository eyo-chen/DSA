package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := helper(node.Left)
	if left == -1 {
		return -1
	}

	right := helper(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) > 1 {
		return -1
	}

	return max(left, right) + 1
}
