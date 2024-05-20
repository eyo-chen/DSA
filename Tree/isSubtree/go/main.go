package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	// at the current state of subtree, ask
	// if the current subtree is same as subRoot
	// if it is, return true
	// if not, go to left and right subtree and keep asking the same question
	if isSame(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
}
