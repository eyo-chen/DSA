package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseKGroup reverses nodes in groups of k using an array-based approach.
//
// Approach: Convert the linked list to an array, reverse elements in groups of k
// within the array, then reconstruct the linked list from the modified array.
//
// Time Complexity: O(n) where n is the number of nodes - we traverse the list once
// to build the array, once to reverse groups, and once to rebuild the list.
// Space Complexity: O(n) for storing all nodes in the array.
func ReverseKGroup(head *ListNode, k int) *ListNode {
	// Convert linked list to array for easier manipulation
	nodes := []*ListNode{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	// Reverse each complete group of k nodes in the array
	for groupStart := 0; groupStart < len(nodes); groupStart += k {
		left := groupStart
		right := groupStart + k - 1

		// Only reverse if we have a complete group of k nodes
		for left < right && right < len(nodes) {
			nodes[left], nodes[right] = nodes[right], nodes[left]
			left++
			right--
		}
	}

	// Rebuild the linked list by reconnecting nodes in array order
	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Next = nodes[i]
	}

	// Set the last node's Next pointer to nil to terminate the list
	nodes[len(nodes)-1].Next = nil

	return nodes[0]
}

// ReverseKGroup1 reverses nodes in groups of k using a recursive approach.
//
// Approach: Check if at least k nodes exist, reverse the first k nodes,
// then recursively reverse the remaining list and connect the segments.
//
// Time Complexity: O(n) where n is the number of nodes - each node is visited once.
// Space Complexity: O(n/k) for the recursion stack, where each call processes k nodes.
func ReverseKGroup1(head *ListNode, k int) *ListNode {
	// Check if we have at least k nodes remaining
	current := head
	for range k {
		if current == nil {
			// Less than k nodes left, return the list unchanged
			return head
		}
		current = current.Next
	}

	// Reverse the first k nodes
	var previous *ListNode
	current = head
	for range k {
		nextNode := current.Next
		current.Next = previous // Reverse the pointer
		previous = current
		current = nextNode
	}

	// After reversal:
	// - 'previous' points to the new head of the reversed segment
	// - 'head' points to the new tail of the reversed segment
	// - 'current' points to the first node of the next segment

	// Recursively reverse the rest and connect it to the tail of current segment
	head.Next = ReverseKGroup1(current, k)

	// Return the new head of the reversed segment
	return previous
}

// ReverseKGroup2 reverses nodes in groups of k using an iterative approach with dummy head.
//
// Approach: Use a dummy head to simplify edge cases. For each group of k nodes,
// find the kth node, reverse the group in place, and update pointers to connect
// the segments. Continue until fewer than k nodes remain.
//
// Time Complexity: O(n) where n is the number of nodes - each node is visited twice
// (once to find kth, once to reverse).
// Space Complexity: O(1) - only using pointers, no additional data structures.
func ReverseKGroup2(head *ListNode, k int) *ListNode {
	// Use a dummy head to simplify edge cases
	dummyHead := &ListNode{Next: head}

	// Points to the node before the current group to be reversed
	previousGroupTail := dummyHead

	for {
		// Find the kth node from previousGroupTail
		kthNode := findKthNode(previousGroupTail, k)
		if kthNode == nil {
			// Less than k nodes remaining, stop reversing
			break
		}

		// Reverse the current group of k nodes
		// 'previous' starts as the node after the group (to connect reversed nodes to)
		// 'current' starts as the first node of the group
		previous := kthNode.Next
		current := previousGroupTail.Next

		for range k {
			nextNode := current.Next
			current.Next = previous // Reverse the pointer
			previous = current
			current = nextNode
		}

		// After reversal, update pointers to connect the segments
		// The original first node of the group is now the tail
		newGroupTail := previousGroupTail.Next

		// Connect previous segment to the new head of reversed group (kthNode)
		previousGroupTail.Next = kthNode

		// Move previousGroupTail to the tail of the reversed group for next iteration
		previousGroupTail = newGroupTail
	}

	return dummyHead.Next
}

// findKthNode returns the kth node after the given start node.
// Returns nil if there are fewer than k nodes after start.
func findKthNode(start *ListNode, k int) *ListNode {
	current := start
	for range k {
		if current == nil {
			return nil
		}
		current = current.Next
	}
	return current
}
