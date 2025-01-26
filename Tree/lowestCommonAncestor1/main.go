package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil || root == q || root == p {
		return root
	}

	// Go to my left and right subtree, and I'll do the further operation once I get the result from my children
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// If my left and right child both found `p` or `q`, return myself
	if left != nil && right != nil {
		return root
	}

	// If only my left child found `p` or `q`, return my left child
	if left != nil {
		return left
	}

	// If only my right child found `p` or `q`, return my right child
	if right != nil {
		return right
	}

	return nil
}
