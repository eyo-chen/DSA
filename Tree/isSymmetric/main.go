package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSymmetric checks if a binary tree is symmetric around its center.
// Approach: Recursively compare the left and right subtrees by checking if they are mirror images.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (due to recursion stack)
func IsSymmetric(root *TreeNode) bool {
	// A tree is symmetric if its left and right subtrees are mirrors of each other
	return isMirror(root.Left, root.Right)
}

// isMirror checks if two trees are mirror reflections of each other
func isMirror(leftSubtree *TreeNode, rightSubtree *TreeNode) bool {
	// Base case: both nodes are nil, they are mirrors
	if leftSubtree == nil && rightSubtree == nil {
		return true
	}

	// If only one is nil, they are not mirrors
	if leftSubtree == nil || rightSubtree == nil {
		return false
	}

	// Values must match for nodes to be mirrors
	if leftSubtree.Val != rightSubtree.Val {
		return false
	}

	// Recursively check:
	// - left's left child mirrors right's right child
	// - left's right child mirrors right's left child
	return isMirror(leftSubtree.Left, rightSubtree.Right) &&
		isMirror(leftSubtree.Right, rightSubtree.Left)
}

// nodePair represents a pair of nodes that should be mirrors of each other
type nodePair struct {
	leftNode  *TreeNode
	rightNode *TreeNode
}

// IsSymmetric1 checks if a binary tree is symmetric around its center.
// Approach: Use BFS with a queue to iteratively compare pairs of nodes that should be mirrors.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(w) where w is the maximum width of the tree (queue size)
func IsSymmetric1(root *TreeNode) bool {
	// Initialize queue with the root's left and right children as the first pair to compare
	queue := []*nodePair{{leftNode: root.Left, rightNode: root.Right}}

	// Process each pair of nodes that should be mirrors
	for len(queue) > 0 {
		// Dequeue the first pair
		pair := queue[0]
		queue = queue[1:]

		left, right := pair.leftNode, pair.rightNode

		// Both nodes are nil - they are symmetric, continue to next pair
		if left == nil && right == nil {
			continue
		}

		// Only one is nil - not symmetric
		if left == nil || right == nil {
			return false
		}

		// Values don't match - not symmetric
		if left.Val != right.Val {
			return false
		}

		// Enqueue child pairs that should be mirrors:
		// - left's left child should mirror right's right child
		// - left's right child should mirror right's left child
		queue = append(queue,
			&nodePair{leftNode: left.Left, rightNode: right.Right},
			&nodePair{leftNode: left.Right, rightNode: right.Left})
	}

	// All pairs matched - tree is symmetric
	return true
}
