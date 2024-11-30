package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// BFS (Using Queue)
func Connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		var prevNode *Node

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// Wire-up the next pointer
			if prevNode != nil {
				prevNode.Next = node
			}
			prevNode = node

			// Add the children of the current node to the queue
			if prevNode.Left != nil {
				queue = append(queue, prevNode.Left)
			}
			if prevNode.Right != nil {
				queue = append(queue, prevNode.Right)
			}
		}
	}

	return root
}

// Recursion
func Connect1(root *Node) *Node {
	// Base case
	// Note that the problem assume that the tree is perfect balanced tree.
	// So, if the root is nil or the root's left or right is nil, we just return the root.
	// In theory, we can only need to check one of the children is nil.
	// Because the problem assume that the tree is perfect balanced tree.
	if root == nil || root.Left == nil || root.Right == nil {
		return root
	}

	// Wire-up the next pointer
	// e.g.   1
	//      2 -> 3
	root.Left.Next = root.Right

	// Wire-up the next pointer for the right child
	// e.g.    1
	//      2      3
	//  4    5 -> 6   7
	// Note that the root.Next might be nil in some cases.
	// For example, when we're processing node 3, root.Next is nil.
	// In this case, we don't need to do anything.
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	// Recursively call the function for the left and right children
	root.Left = Connect1(root.Left)
	root.Right = Connect1(root.Right)
	return root
}

// Iteration
func Connect2(root *Node) *Node {
	leftMostNode := root

	// Iterate through each level of the tree
	// This loop will go layer by layer, and only process the leftmost node of each level.
	// Note that we don't want to process that last level's leftmost node.
	// e.g.   1
	//      2    3
	//    4  5  6  7
	// In this case, we don't want to process node 4.(third level)
	// Because curNode.Left will be nil in this case.
	// Also, our logic is wiring up the next pointer for the left and right children.
	// For example, when we're processing node 2, we're wiring up node 4 and 5.
	// In other words, when we're at the second level, we're wiring up the third level.
	for leftMostNode != nil && leftMostNode.Left != nil {
		curNode := leftMostNode

		// Wire-up the next pointer for the left and right children of the current node
		// This loop will go through all the nodes in the current level.
		for curNode != nil {
			// Wire-up the next pointer for the left child
			curNode.Left.Next = curNode.Right

			// Wire-up the next pointer for the right child
			// It's the same concept as recursion approach
			if curNode.Next != nil {
				curNode.Right.Next = curNode.Next.Left
			}

			// Move to the next node in the same level
			curNode = curNode.Next
		}

		// Move to the next level(leftmost node of the next level)
		leftMostNode = leftMostNode.Left
	}

	return root
}
