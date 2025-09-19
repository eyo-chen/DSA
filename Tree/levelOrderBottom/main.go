package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrderBottom returns the bottom-up level order traversal of a binary tree.
// Approach: Uses BFS with a queue to traverse level by level, then prepends each level
// to the result slice to achieve bottom-up order.
// Time Complexity: O(n), Space Complexity: O(n)
func LevelOrderBottom(root *TreeNode) [][]int {
	// Handle empty tree case
	if root == nil {
		return [][]int{}
	}

	// Initialize queue with root node for BFS traversal
	nodeQueue := []*TreeNode{root}
	result := [][]int{}

	// Process nodes level by level using BFS
	for len(nodeQueue) > 0 {
		currentLevelSize := len(nodeQueue)
		currentLevelValues := []int{}

		// Process all nodes at the current level
		for range currentLevelSize {
			// Dequeue the first node
			currentNode := nodeQueue[0]
			nodeQueue = nodeQueue[1:]

			// Add current node's value to current level
			currentLevelValues = append(currentLevelValues, currentNode.Val)

			// Enqueue children for next level processing
			if currentNode.Left != nil {
				nodeQueue = append(nodeQueue, currentNode.Left)
			}
			if currentNode.Right != nil {
				nodeQueue = append(nodeQueue, currentNode.Right)
			}
		}

		// Prepend current level to result to achieve bottom-up order
		result = append([][]int{currentLevelValues}, result...)
	}

	return result
}

// LevelOrderBottom1 returns the bottom-up level order traversal using a stack approach.
// Approach: Uses BFS to collect levels in top-down order, then uses a stack to reverse
// the order to bottom-up.
// Time Complexity: O(n), Space Complexity: O(n)
func LevelOrderBottom1(root *TreeNode) [][]int {
	// Handle empty tree case
	if root == nil {
		return [][]int{}
	}

	// Initialize queue for BFS and stack to reverse order
	nodeQueue := []*TreeNode{root}
	levelStack := [][]int{}

	// Collect all levels using BFS (top-down order)
	for len(nodeQueue) > 0 {
		currentLevelSize := len(nodeQueue)
		currentLevelValues := []int{}

		// Process all nodes at current level
		for range currentLevelSize {
			// Dequeue the first node
			currentNode := nodeQueue[0]
			nodeQueue = nodeQueue[1:]

			// Add current node's value to current level
			currentLevelValues = append(currentLevelValues, currentNode.Val)

			// Enqueue children for next level
			if currentNode.Left != nil {
				nodeQueue = append(nodeQueue, currentNode.Left)
			}
			if currentNode.Right != nil {
				nodeQueue = append(nodeQueue, currentNode.Right)
			}
		}

		// Push current level onto stack
		levelStack = append(levelStack, currentLevelValues)
	}

	// Pop levels from stack to build bottom-up result
	result := [][]int{}
	for len(levelStack) > 0 {
		// Pop the top level from stack and add to result
		result = append(result, levelStack[len(levelStack)-1])
		levelStack = levelStack[:len(levelStack)-1]
	}

	return result
}

// LevelOrderBottom2 returns the bottom-up level order traversal using recursion.
// Approach: First calculates tree depth, then uses DFS to fill result array from
// bottom levels to top levels by calculating the correct index for each level.
// Time Complexity: O(n), Space Complexity: O(h) where h is height of tree
func LevelOrderBottom2(root *TreeNode) [][]int {
	// Calculate the maximum depth of the tree
	maxDepth := calculateTreeDepth(root)

	// Pre-allocate result slice with correct size
	resultLevels := make([][]int, maxDepth)

	// Fill the result using DFS, starting from the deepest level index
	fillLevelsBottomUp(root, &resultLevels, maxDepth-1)

	return resultLevels
}

// fillLevelsBottomUp recursively fills the result slice from bottom to top.
// It uses DFS to traverse the tree and places each node's value at the correct
// level index (calculated to achieve bottom-up ordering).
func fillLevelsBottomUp(currentNode *TreeNode, resultLevels *[][]int, currentLevelIndex int) {
	// Base case: if node is nil, return
	if currentNode == nil {
		return
	}

	// Add current node's value to the appropriate level
	(*resultLevels)[currentLevelIndex] = append((*resultLevels)[currentLevelIndex], currentNode.Val)

	// Recursively process left and right children at the next level up
	fillLevelsBottomUp(currentNode.Left, resultLevels, currentLevelIndex-1)
	fillLevelsBottomUp(currentNode.Right, resultLevels, currentLevelIndex-1)
}

// calculateTreeDepth returns the maximum depth (height) of the binary tree.
// Uses recursive approach to find the deeper path between left and right subtrees.
func calculateTreeDepth(node *TreeNode) int {
	// Base case: empty tree has depth 0
	if node == nil {
		return 0
	}

	// Calculate depth of left and right subtrees, return max + 1
	leftDepth := calculateTreeDepth(node.Left)
	rightDepth := calculateTreeDepth(node.Right)

	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
