package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return CountNodes(root.Right) + CountNodes(root.Left) + 1
}
