package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Standard Reversal Approach
// Approach: Find the nodes before and after the reversal range, reverse the sublist
// in-place using standard reversal technique, then reconnect the segments.
// Time Complexity: O(n) where n is the position of right node
// Space Complexity: O(1) - only uses pointers
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	// Create dummy head to handle edge cases where left = 1
	dummy := &ListNode{Next: head}

	// Initialize pointers to traverse and locate key positions
	leftPtr, rightPtr := dummy, dummy
	nodeBeforeLeft := dummy
	leftIndex, rightIndex, beforeLeftIndex := left, right, left-1

	// Traverse to find three key positions:
	// 1. nodeBeforeLeft: node right before the reversal start (to reconnect later)
	// 2. leftPtr: first node to be reversed
	// 3. rightPtr: last node to be reversed
	for {
		if leftIndex > 0 {
			leftPtr = leftPtr.Next
			leftIndex--
		}
		if rightIndex > 0 {
			rightPtr = rightPtr.Next
			rightIndex--
		}
		if beforeLeftIndex > 0 {
			nodeBeforeLeft = nodeBeforeLeft.Next
			beforeLeftIndex--
		}
		// Stop when all three positions are found
		if leftIndex == 0 && rightIndex == 0 && beforeLeftIndex == 0 {
			break
		}
	}

	// Set prev to the node after rightPtr (where reversed list will connect)
	prev := rightPtr.Next
	current := leftPtr
	reversalSteps := right - left + 1

	// Reverse the sublist from left to right position
	// Standard reversal: redirect each node's Next pointer to previous node
	for range reversalSteps {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}

	// Connect the node before left to the new head of reversed sublist
	nodeBeforeLeft.Next = prev

	return dummy.Next
}

// Solution 2: Incremental Move Approach
// Approach: Instead of reversing in bulk, move nodes one by one from their current
// position to the front of the sublist. This builds the reversed list incrementally.
// Time Complexity: O(n) where n is the position of right node
// Space Complexity: O(1) - only uses pointers
func ReverseBetweenV2(head *ListNode, left int, right int) *ListNode {
	// Create dummy head to simplify edge case handling
	dummy := &ListNode{Next: head}

	// Find the node right before the reversal range starts
	nodeBeforeReversalRange := dummy
	for range left - 1 {
		nodeBeforeReversalRange = nodeBeforeReversalRange.Next
	}

	// current points to the first node in the reversal range
	// This node will end up at the end of the reversed sublist
	current := nodeBeforeReversalRange.Next

	// Move each subsequent node to the front of the sublist
	// Repeat (right - left) times to process all nodes in range except the first
	for range right - left {
		// Extract the next node to move
		nodeToMove := current.Next

		// Remove nodeToMove from its current position
		current.Next = nodeToMove.Next

		// Insert nodeToMove at the front of the sublist
		// (right after nodeBeforeReversalRange)
		nodeToMove.Next = nodeBeforeReversalRange.Next
		nodeBeforeReversalRange.Next = nodeToMove
	}

	return dummy.Next
}
