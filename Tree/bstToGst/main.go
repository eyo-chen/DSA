package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BstToGst converts a Binary Search Tree to a Greater Sum Tree.
//
// Approach: Perform a reverse in-order traversal (right -> node -> left) of the BST.
// This visits nodes in descending order, allowing us to maintain a running sum of all
// visited nodes. For each node, we update its value to include all greater values.
//
// Time Complexity: O(n) where n is the number of nodes (visit each node once)
// Space Complexity: O(h) where h is the height of the tree (recursion stack depth)
func BstToGst(root *TreeNode) *TreeNode {
	runningSum := 0
	reverseInOrderTraversal(root, &runningSum)
	return root
}

// reverseInOrderTraversal performs a reverse in-order traversal (right -> node -> left)
// and updates each node's value to include the sum of all greater nodes.
func reverseInOrderTraversal(node *TreeNode, runningSum *int) {
	// Base case: reached a null node
	if node == nil {
		return
	}

	// Step 1: Traverse right subtree first (all greater values)
	reverseInOrderTraversal(node.Right, runningSum)

	// Step 2: Process current node
	// Save current running sum (sum of all nodes greater than current)
	sumOfGreaterNodes := *runningSum

	// Update running sum to include current node's value
	*runningSum += node.Val

	// Update current node's value: original value + sum of all greater nodes
	node.Val += sumOfGreaterNodes

	// Step 3: Traverse left subtree (all smaller values)
	reverseInOrderTraversal(node.Left, runningSum)
}
