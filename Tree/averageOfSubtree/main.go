package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// AverageOfSubtree counts the number of nodes where the node's value equals
// the average of all values in its subtree (rounded down).
//
// Approach: Use post-order DFS traversal to calculate subtree sum and count
// from bottom to top. For each node, check if its value matches the average
// of its subtree.
//
// Time Complexity: O(n) - visit each node exactly once
// Space Complexity: O(h) - recursion stack depth, where h is tree height
//
//	O(n) in worst case for skewed tree, O(log n) for balanced tree
func AverageOfSubtree(root *TreeNode) int {
	matchCount := 0
	calculateSubtreeStats(root, &matchCount)
	return matchCount
}

// calculateSubtreeStats performs post-order DFS to compute subtree statistics.
// Returns (sum, nodeCount) where:
//   - sum: total sum of all node values in the subtree rooted at current node
//   - nodeCount: total number of nodes in the subtree
//
// Side effect: increments matchCount when node value equals subtree average
func calculateSubtreeStats(node *TreeNode, matchCount *int) (sum int, nodeCount int) {
	// Base case: empty subtree has sum 0 and count 0
	if node == nil {
		return 0, 0
	}

	// Recursively calculate stats for left subtree
	leftSum, leftNodeCount := calculateSubtreeStats(node.Left, matchCount)

	// Recursively calculate stats for right subtree
	rightSum, rightNodeCount := calculateSubtreeStats(node.Right, matchCount)

	// Calculate total sum: left subtree + right subtree + current node
	totalSum := leftSum + rightSum + node.Val

	// Calculate total count: left subtree + right subtree + current node (1)
	totalNodeCount := leftNodeCount + rightNodeCount + 1

	// Calculate average (integer division automatically rounds down)
	average := totalSum / totalNodeCount

	// Check if current node's value matches its subtree average
	if node.Val == average {
		*matchCount++
	}

	// Return subtree statistics to parent node
	return totalSum, totalNodeCount
}
