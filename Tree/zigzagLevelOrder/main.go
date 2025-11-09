package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ZigzagLevelOrderWithStack performs zigzag level order traversal using a two-stack approach.
// The algorithm uses a queue for BFS and a stack to control child insertion order.
// For even levels (0,2,4...), children are added right-first to achieve left-to-right reading.
// For odd levels (1,3,5...), children are added left-first to achieve right-to-left reading.
//
// Time Complexity: O(n) where n is the number of nodes - each node is visited exactly once
// Space Complexity: O(w) where w is the maximum width of the tree (queue + stack storage)
func ZigzagLevelOrderWithStack(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	nodeQueue := []*TreeNode{root}
	tempStack := []*TreeNode{}
	currentLevel := 0

	for len(nodeQueue) > 0 {
		levelSize := len(nodeQueue)
		levelValues := []int{}

		// Process all nodes at current level
		for range levelSize {
			currentNode := nodeQueue[0]
			nodeQueue = nodeQueue[1:] // Dequeue from front

			// Store node in stack and collect its value
			tempStack = append(tempStack, currentNode)
			levelValues = append(levelValues, currentNode.Val)
		}

		// Pop nodes from stack and add their children to queue
		// Stack ensures we process nodes in reverse order for child insertion
		for len(tempStack) > 0 {
			// Pop from stack (LIFO order)
			nodeToProcess := tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]

			// Add children based on current level parity
			if currentLevel%2 == 0 {
				// Even level: add right child first, then left child
				// This will result in left-to-right reading in next level
				if nodeToProcess.Right != nil {
					nodeQueue = append(nodeQueue, nodeToProcess.Right)
				}
				if nodeToProcess.Left != nil {
					nodeQueue = append(nodeQueue, nodeToProcess.Left)
				}
			} else {
				// Odd level: add left child first, then right child
				// This will result in right-to-left reading in next level
				if nodeToProcess.Left != nil {
					nodeQueue = append(nodeQueue, nodeToProcess.Left)
				}
				if nodeToProcess.Right != nil {
					nodeQueue = append(nodeQueue, nodeToProcess.Right)
				}
			}
		}

		result = append(result, levelValues)
		currentLevel++
	}

	return result
}

// ZigzagLevelOrderWithPositioning performs zigzag level order traversal using position calculation.
// The algorithm uses standard BFS but calculates the correct position for each node's value
// in the level array based on the traversal direction (left-to-right vs right-to-left).
//
// Time Complexity: O(n) where n is the number of nodes - each node is visited exactly once
// Space Complexity: O(w) where w is the maximum width of the tree (queue storage only)
func ZigzagLevelOrderWithPositioning(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	nodeQueue := []*TreeNode{root}
	isRightToLeft := false

	for len(nodeQueue) > 0 {
		levelSize := len(nodeQueue)
		// Pre-allocate array with exact size needed for current level
		levelValues := make([]int, levelSize)

		// Process all nodes at current level
		for nodeIndex := range levelSize {
			currentNode := nodeQueue[0]
			nodeQueue = nodeQueue[1:] // Dequeue from front

			// Always add children in same order (right first, then left)
			// This maintains consistent queue ordering
			if currentNode.Right != nil {
				nodeQueue = append(nodeQueue, currentNode.Right)
			}
			if currentNode.Left != nil {
				nodeQueue = append(nodeQueue, currentNode.Left)
			}

			// Calculate position based on traversal direction
			var targetPosition int
			if isRightToLeft {
				// Right-to-left: use natural index order
				targetPosition = nodeIndex
			} else {
				// Left-to-right: reverse the index order
				targetPosition = levelSize - nodeIndex - 1
			}

			// Place node value at calculated position
			levelValues[targetPosition] = currentNode.Val
		}

		result = append(result, levelValues)
		// Toggle direction for next level
		isRightToLeft = !isRightToLeft
	}

	return result
}
