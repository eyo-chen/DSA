package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ConnectLevelOrderBFS connects nodes at the same level using BFS (Breadth-First Search).
// This approach uses a queue to process nodes level by level from left to right.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(w) where w is the maximum width of the tree (worst case O(n) for a complete binary tree)
func ConnectLevelOrderBFS(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Initialize queue with root node for BFS traversal
	queue := []*Node{root}

	// Process each level of the tree
	for len(queue) > 0 {
		currentLevelSize := len(queue)

		var previousNodeInLevel *Node

		// Process all nodes in the current level
		for range currentLevelSize {
			// Dequeue the first node
			currentNode := queue[0]
			queue = queue[1:]

			// Add children to queue for next level processing
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

			// Connect the previous node to current node (skip first node in level)
			if previousNodeInLevel != nil {
				previousNodeInLevel.Next = currentNode
			}
			previousNodeInLevel = currentNode
		}
	}

	return root
}

// ConnectRecursiveDFS connects nodes using DFS (Depth-First Search) with recursion.
// This approach leverages the existing Next pointers to find connections for child nodes.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree due to recursion stack
func ConnectRecursiveDFS(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Connect left child to right child if both exist
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			// If no right child, find next available child in the same level
			root.Left.Next = findNextAvailableChild(root.Next)
		}
	}

	// Connect right child to the next available child in the same level
	if root.Right != nil {
		root.Right.Next = findNextAvailableChild(root.Next)
	}

	// Process right subtree first to ensure Next pointers are available for left subtree
	ConnectRecursiveDFS(root.Right)
	ConnectRecursiveDFS(root.Left)

	return root
}

// ConnectIterativeDFS connects nodes using iterative DFS with constant extra space.
// This approach processes the tree level by level using existing Next pointers for navigation.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(1) constant extra space (not counting the tree itself)
func ConnectIterativeDFS(root *Node) *Node {
	// Start from the leftmost node of each level
	leftmostNodeInLevel := root

	// Process each level until no more levels exist
	for leftmostNodeInLevel != nil {
		currentNodeInLevel := leftmostNodeInLevel

		// Traverse all nodes in the current level using Next pointers
		for currentNodeInLevel != nil {
			// Connect left child to right child or next available child
			if currentNodeInLevel.Left != nil {
				if currentNodeInLevel.Right != nil {
					currentNodeInLevel.Left.Next = currentNodeInLevel.Right
				} else {
					currentNodeInLevel.Left.Next = findNextAvailableChild(currentNodeInLevel.Next)
				}
			}

			// Connect right child to next available child in the same level
			if currentNodeInLevel.Right != nil {
				currentNodeInLevel.Right.Next = findNextAvailableChild(currentNodeInLevel.Next)
			}

			// Move to the next node in the current level
			currentNodeInLevel = currentNodeInLevel.Next
		}

		// Move to the leftmost node of the next level
		leftmostNodeInLevel = findNextAvailableChild(leftmostNodeInLevel)
	}

	return root
}

// findNextAvailableChild finds the next available child node by traversing
// through nodes at the same level using Next pointers.
// This helper function is used to connect nodes across different parent subtrees.
func findNextAvailableChild(startNode *Node) *Node {
	if startNode == nil {
		return nil
	}

	// Traverse nodes in the same level to find the first available child
	currentNode := startNode
	for currentNode != nil {
		// Return the first child found (left has priority)
		if currentNode.Left != nil {
			return currentNode.Left
		}
		if currentNode.Right != nil {
			return currentNode.Right
		}

		// Move to the next node in the same level
		currentNode = currentNode.Next
	}

	// No child found in the remaining nodes of this level
	return nil
}
