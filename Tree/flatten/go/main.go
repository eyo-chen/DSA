package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// using slice to store the nodes in pre-order
func Flatten(root *TreeNode) {
	nodes := []*TreeNode{}
	genPreOrderNodes(root, &nodes)

	// reconstruct the tree
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]

		// always set left node to nil
		node.Left = nil

		// if the current node is the last node, set right node to nil
		// otherwise, set right node to the next node
		var nextNode *TreeNode
		if i != len(nodes)-1 {
			nextNode = nodes[i+1]
		}
		node.Right = nextNode
	}
}

func genPreOrderNodes(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root)
	genPreOrderNodes(root.Left, nodes)
	genPreOrderNodes(root.Right, nodes)
}

// using helper function to return the tail node of the current subtree
func Flatten1(root *TreeNode) {
	helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftTail := helper(root.Left)
	rightTail := helper(root.Right)

	// when the leftTail is not nil, we need to do the rewiring
	if leftTail != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightTail != nil {
		return rightTail
	}

	if leftTail != nil {
		return leftTail
	}

	return root
}
