package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ReverseOddLevels reverses node values at odd levels of a perfect binary tree
// using a recursive DFS approach with two pointers from opposite ends.
//
// Approach: Recursively traverse the tree with two pointers (left and right)
// that mirror each other. At odd levels, swap the values of mirrored nodes.
// The recursive calls ensure we process corresponding nodes from opposite ends.
//
// Time Complexity: O(n) - visit each node once
// Space Complexity: O(log n) - recursion stack depth equals tree height
func ReverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Start recursion from level 1 (first odd level) with left and right children
	reverseOddLevelsDFS(root.Left, root.Right, 1)
	return root
}

// reverseOddLevelsDFS recursively swaps values of mirrored nodes at odd levels
// leftNode: node from the left side of the current level
// rightNode: corresponding mirrored node from the right side
// currentLevel: current level in the tree (root is level 0)
func reverseOddLevelsDFS(leftNode *TreeNode, rightNode *TreeNode, currentLevel int) {
	// Base case: reached leaf level
	if leftNode == nil || rightNode == nil {
		return
	}

	// If current level is odd, swap values between mirrored nodes
	if currentLevel%2 == 1 {
		leftNode.Val, rightNode.Val = rightNode.Val, leftNode.Val
	}

	// Recursively process next level with mirrored pairs:
	// - leftNode's left child pairs with rightNode's right child (outer pair)
	// - leftNode's right child pairs with rightNode's left child (inner pair)
	reverseOddLevelsDFS(leftNode.Left, rightNode.Right, currentLevel+1)
	reverseOddLevelsDFS(leftNode.Right, rightNode.Left, currentLevel+1)
}

// ReverseOddLevelsBFS reverses node values at odd levels of a perfect binary tree
// using an iterative BFS (level-order traversal) approach.
//
// Approach: Use a queue to traverse level by level. At each odd level, collect
// all nodes and swap values from both ends toward the center.
//
// Time Complexity: O(n) - visit each node once
// Space Complexity: O(n) - queue holds up to n/2 nodes (last level width)
func ReverseOddLevelsBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	currentLevel := 0

	// Process tree level by level
	for len(queue) > 0 {
		levelSize := len(queue)
		nodesAtCurrentLevel := make([]*TreeNode, levelSize)
		shouldSwap := currentLevel%2 == 1

		// Store current level nodes if we need to swap them
		if shouldSwap {
			copy(nodesAtCurrentLevel, queue)
		}

		// Process all nodes at current level and add their children to queue
		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Swap values from both ends toward center if at odd level
		if shouldSwap {
			left, right := 0, levelSize-1
			for left < right {
				nodesAtCurrentLevel[left].Val, nodesAtCurrentLevel[right].Val =
					nodesAtCurrentLevel[right].Val, nodesAtCurrentLevel[left].Val
				left++
				right--
			}
		}

		currentLevel++
	}

	return root
}
