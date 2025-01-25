package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return dfs(root, math.MaxInt, math.MinInt)
}

func dfs(node *TreeNode, maxVal, minVal int) bool {
	if node == nil {
		return true
	}

	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}

	if !dfs(node.Left, node.Val, minVal) || !dfs(node.Right, maxVal, node.Val) {
		return false
	}

	return true

	// Or we can simply return
	// return dfs(node.Left, node.Val, minVal) && dfs(node.Right, maxVal, node.Val)
}

func IsValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
