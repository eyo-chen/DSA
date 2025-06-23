package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion Approach
func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = InsertIntoBST(root.Left, val)
	} else {
		root.Right = InsertIntoBST(root.Right, val)
	}

	return root
}

// Iterative Approach
func InsertIntoBST1(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}

	curNode := root
	for curNode != nil {
		if curNode.Val > val {
			if curNode.Left == nil {
				curNode.Left = node
				break
			}

			curNode = curNode.Left
			continue
		}

		if curNode.Right == nil {
			curNode.Right = node
			break
		}
		curNode = curNode.Right
	}

	return root
}
