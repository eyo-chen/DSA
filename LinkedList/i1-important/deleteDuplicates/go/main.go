package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	// If the list is empty or has only one node, no duplicates are possible, so return head
	if head == nil || head.Next == nil {
		return head
	}

	// Initialize two pointers: slow and fast, both starting at the head
	fast, slow := head, head

	// Continue until slow reaches the end or the second-to-last node
	for slow != nil && slow.Next != nil {
		// Move fast forward while nodes have the same value as slow
		for fast != nil && fast.Val == slow.Val {
			fast = fast.Next
		}

		// Link slow to the next unique value (fast), effectively skipping duplicates
		slow.Next = fast
		// Move slow to the next unique node
		slow = fast

		// Move fast to the next node if it exists
		if fast != nil {
			fast = fast.Next
		}
	}

	// Return the head of the modified list
	return head
}

func DeleteDuplicates1(head *ListNode) *ListNode {
	// If the list is empty or has only one node, no duplicates are possible, so return head
	if head == nil || head.Next == nil {
		return head
	}

	// Initialize two pointers: slow points to the current unique node, fast checks the next node
	slow, fast := head, head.Next

	// Iterate until fast reaches the end of the list
	for fast != nil {
		// If slow and fast point to nodes with the same value, skip the duplicate
		if slow.Val == fast.Val {
			fast = fast.Next // Move fast to the next node
			slow.Next = fast // Update slow's next to skip the duplicate
			continue         // Continue to check the next node
		}

		// If values differ, move both pointers forward
		fast = fast.Next // Move fast to the next node
		slow = slow.Next // Move slow to the next unique node
	}

	// Return the head of the modified list
	return head
}
