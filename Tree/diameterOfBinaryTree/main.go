package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, ans)
	right := helper(root.Right, ans)

	*ans = max(*ans, left+right)

	return max(left, right) + 1
}
