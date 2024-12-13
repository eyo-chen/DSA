package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Store leaf sequence in strings and compare them
// Time Complexity: O(N)
// Space Complexity: O(N)
func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	str1, str2 := strings.Builder{}, strings.Builder{}
	genLeafSequence(root1, &str1)
	genLeafSequence(root2, &str2)

	return str1.String() == str2.String()
}

func genLeafSequence(root *TreeNode, strB *strings.Builder) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		strB.WriteString(strconv.Itoa(root.Val))
		strB.WriteString("-")
		return
	}

	genLeafSequence(root.Left, strB)
	genLeafSequence(root.Right, strB)
}

// Stack-based iteration
// Time Complexity: O(N)
// Space Complexity: O(h) where h is the height of the tree
// (Best Case: O(logN) for balanced tree, Worst Case: O(N) for skewed tree)
// This solution is more space-efficient than the string-based approach because:
// - It only stores nodes in the current path (O(h)) vs storing all leaf values
// - No string conversion or concatenation overhead
// - No need to store complete leaf sequences before comparison
func LeafSimilar1(root1 *TreeNode, root2 *TreeNode) bool {
	// Create stack for each tree
	stack1 := []*TreeNode{root1}
	stack2 := []*TreeNode{root2}

	// Compare leaves one by one
	for len(stack1) > 0 && len(stack2) > 0 {
		// Get next leaf from first tree
		leaf1 := getNextLeaf(&stack1)
		// Get next leaf from second tree
		leaf2 := getNextLeaf(&stack2)

		// If leaves don't match, trees aren't similar
		if leaf1 != leaf2 {
			return false
		}
	}

	// Both trees should run out of leaves at the same time
	return len(stack1) == 0 && len(stack2) == 0
}

// Helper function to get next leaf value using stack-based iteration
func getNextLeaf(stack *[]*TreeNode) int {
	for len(*stack) > 0 {
		// Pop the top node
		node := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]

		// If it's a leaf, return its value
		if node.Left == nil && node.Right == nil {
			return node.Val
		}

		// Add children to stack (right first so left is processed first)
		if node.Right != nil {
			*stack = append(*stack, node.Right)
		}
		if node.Left != nil {
			*stack = append(*stack, node.Left)
		}
	}
	return -1 // Should never reach here if trees are valid
}
