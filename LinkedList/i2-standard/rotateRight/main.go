package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// RotateRight rotates the linked list to the right by k positions.
//
// Approach:
// 1. Find the length of the list to handle cases where k > length
// 2. Calculate effective rotations: k % length (to avoid unnecessary full rotations)
// 3. Use two pointers with a gap of k nodes to find the new tail
// 4. Break the list at the new tail and reconnect to form the rotated list
//
// Time Complexity: O(n) where n is the number of nodes in the list
//   - First pass to calculate length: O(n)
//   - Second pass to find the breaking point: O(n)
//
// Space Complexity: O(1) - only using pointers, no extra data structures
func RotateRight(head *ListNode, k int) *ListNode {
	// Handle empty list edge case
	if head == nil {
		return head
	}

	// Calculate the length of the linked list
	length := 0
	for current := head; current != nil; current = current.Next {
		length++
	}

	// Calculate effective rotations (handle k > length)
	effectiveRotations := k % length

	// If no rotation needed, return original list
	if effectiveRotations == 0 {
		return head
	}

	// Find the new tail (node before the new head)
	// Use two pointers: fast pointer moves k steps ahead first,
	// then both move together until fast reaches the end
	slow := head
	fast := head

	// Move fast pointer k steps ahead
	for range effectiveRotations {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the last node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Perform the rotation:
	// 1. The new head is the node after slow
	newHead := slow.Next
	// 2. Connect the old tail (fast) to the old head
	fast.Next = head
	// 3. Break the connection at the new tail
	slow.Next = nil

	return newHead
}
