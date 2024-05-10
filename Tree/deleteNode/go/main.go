package deletenode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	// If the root is nil, return nil
	if root == nil {
		return nil
	}

	// When the key is greater than the root value, we need to delete the key
	// from the right subtree.
	if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
		return root
	}

	// When the key is less than the root value, we need to delete the key from
	// the left subtree.
	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
		return root
	}

	// When the key is equal to the root value, we need to delete the root node.

	// If the root node is a leaf node, we can simply delete it.(return nil)
	if root.Left == nil && root.Right == nil {
		return nil
	}

	// If the root node has only one child, we can return the non-nil child
	// node.
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}

	// If the root node has two children, we need to find the leftmost node in
	// the right subtree, replace the root value with the leftmost node value,
	// and delete the leftmost node from the right subtree.
	leftMostNodeInRightSubTree := root.Right
	for leftMostNodeInRightSubTree.Left != nil {
		leftMostNodeInRightSubTree = leftMostNodeInRightSubTree.Left
	}
	root.Val = leftMostNodeInRightSubTree.Val
	root.Right = DeleteNode(root.Right, leftMostNodeInRightSubTree.Val)
	return root
}
