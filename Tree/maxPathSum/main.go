package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxPathSum finds the maximum sum of any path in a binary tree.
// A path is defined as a sequence of nodes where each pair of adjacent nodes
// has an edge connecting them, and each node appears at most once.
//
// Approach: Uses post-order DFS traversal where each node considers two things:
// 1. Update global maximum with the path that passes through current node
// 2. Return the maximum single-direction path sum to parent for extension
//
// Time Complexity: O(n) where n is the number of nodes - visit each node once
// Space Complexity: O(h) where h is the height of tree - recursion stack space
func MaxPathSum(root *TreeNode) int {
	globalMaxSum := math.MinInt32
	calculateMaxPath(root, &globalMaxSum)
	return globalMaxSum
}

// calculateMaxPath performs DFS traversal and returns the maximum path sum
// that starts from the current node and goes down in one direction.
// It also updates the global maximum with paths that pass through current node.
//
// Parameters:
//   - node: current node being processed
//   - globalMaxSum: pointer to track the overall maximum path sum
//
// Returns: maximum path sum starting from current node going down (single direction)
func calculateMaxPath(node *TreeNode, globalMaxSum *int) int {
	// Base case: null nodes contribute 0 to any path
	if node == nil {
		return 0
	}

	// Recursively get maximum path sums from left and right subtrees
	// Use max(0, ...) to ignore negative contributions - better to skip than include negative paths
	leftMaxPath := max(0, calculateMaxPath(node.Left, globalMaxSum))
	rightMaxPath := max(0, calculateMaxPath(node.Right, globalMaxSum))

	// Update global maximum: consider the path that goes through current node
	// This path connects left subtree -> current node -> right subtree
	pathThroughCurrent := leftMaxPath + rightMaxPath + node.Val
	*globalMaxSum = max(*globalMaxSum, pathThroughCurrent)

	// Return to parent: the maximum single-direction path from current node
	// Parent can extend this path further up the tree
	// We can only go either left OR right, not both (to maintain valid tree path)
	maxSinglePath := max(leftMaxPath, rightMaxPath) + node.Val
	return maxSinglePath
}
