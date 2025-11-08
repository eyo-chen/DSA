package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MinDepth returns the minimum depth of a binary tree using BFS (level-order traversal).
// The minimum depth is the number of nodes along the shortest path from the root to the nearest leaf node.
//
// Approach: Use BFS to traverse the tree level by level. Return immediately when the first leaf node is found,
// as BFS guarantees this will be at the minimum depth.
//
// Time Complexity: O(n) in the worst case, where n is the number of nodes. In the best case (balanced tree),
// we only visit nodes up to the first leaf, which could be O(log n) nodes.
// Space Complexity: O(w) where w is the maximum width of the tree (for the queue).
func MinDepth(root *TreeNode) int {
	// Base case: empty tree has depth 0
	if root == nil {
		return 0
	}

	// Initialize queue with root node for BFS traversal
	queue := []*TreeNode{root}
	depth := 0

	// Process nodes level by level
	for len(queue) > 0 {
		// Get the number of nodes at current level
		levelSize := len(queue)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			// Dequeue the front node
			currentNode := queue[0]
			queue = queue[1:]

			// Check if this is a leaf node (no children)
			if currentNode.Left == nil && currentNode.Right == nil {
				// First leaf found - this is the minimum depth
				return depth + 1
			}

			// Add left child to queue if it exists
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			// Add right child to queue if it exists
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		// Move to next level
		depth++
	}

	// This line should never be reached for a valid tree
	return depth
}

// MinDepthRecursive returns the minimum depth of a binary tree using DFS (recursive approach).
// The minimum depth is the number of nodes along the shortest path from the root to the nearest leaf node.
//
// Approach: Use DFS to recursively calculate the minimum depth. Handle edge cases where a node has only
// one child - in such cases, we must go down the existing child's path.
//
// Time Complexity: O(n) where n is the number of nodes (must visit all nodes in worst case).
// Space Complexity: O(h) where h is the height of the tree (recursion call stack).
func MinDepthRecursive(root *TreeNode) int {
	// Base case: empty tree has depth 0
	if root == nil {
		return 0
	}

	// Recursively calculate minimum depth of left and right subtrees
	leftDepth := MinDepthRecursive(root.Left)
	rightDepth := MinDepthRecursive(root.Right)

	// Case 1: Leaf node (no children) - depth is 1
	if leftDepth == 0 && rightDepth == 0 {
		return 1
	}

	// Case 2: Only right child exists - must go down right path
	if leftDepth == 0 {
		return rightDepth + 1
	}

	// Case 3: Only left child exists - must go down left path
	if rightDepth == 0 {
		return leftDepth + 1
	}

	// Case 4: Both children exist - take minimum depth and add current node
	return min(leftDepth, rightDepth) + 1
}
