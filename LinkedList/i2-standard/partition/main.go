package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Partition reorders the linked list so that all nodes with values less than x
// come before nodes with values greater than or equal to x, while preserving
// the relative order of nodes in each partition.
//
// Approach: Use two separate linked lists to collect nodes less than x and
// nodes greater than or equal to x, then connect them together.
//
// Time Complexity: O(n) - single pass through the list
// Space Complexity: O(1) - only uses pointers, no additional data structures
func Partition(head *ListNode, x int) *ListNode {
	// Create dummy heads for two separate lists
	// lessDummy: collects nodes with values < x
	// greaterOrEqualDummy: collects nodes with values >= x
	lessDummy := &ListNode{}
	greaterOrEqualDummy := &ListNode{}

	// Pointers to track the current tail of each list
	lessTail := lessDummy
	greaterOrEqualTail := greaterOrEqualDummy

	// Traverse the original list and distribute nodes into two lists
	current := head
	for current != nil {
		if current.Val >= x {
			// Append to the greater-or-equal list
			greaterOrEqualTail.Next = current
			greaterOrEqualTail = greaterOrEqualTail.Next
		} else {
			// Append to the less-than list
			lessTail.Next = current
			lessTail = lessTail.Next
		}
		current = current.Next
	}

	// Connect the less-than list to the greater-or-equal list
	lessTail.Next = greaterOrEqualDummy.Next

	// Terminate the greater-or-equal list to avoid cycles
	greaterOrEqualTail.Next = nil

	// Return the head of the partitioned list (skip the dummy node)
	return lessDummy.Next
}
