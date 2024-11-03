package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	curSum := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return curSum == 0
	}

	return HasPathSum(root.Left, curSum) ||
		HasPathSum(root.Right, curSum)
}
