package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedListToBST converts a sorted linked list into a height-balanced binary search tree.
//
// Approach:
// - Use a recursive divide-and-conquer strategy
// - Find the middle node of the current list segment using the slow-fast pointer technique
// - Make the middle node the root to ensure balance
// - Recursively build left subtree from the left half and right subtree from the right half
//
// Time Complexity: O(n log n)
// - Finding the middle takes O(n) time for each level of recursion
// - There are O(log n) levels in the recursion tree
//
// Space Complexity: O(log n)
// - Recursion stack depth is O(log n) for a balanced tree
func SortedListToBST(head *ListNode) *TreeNode {
	return buildBST(head, nil)
}

// buildBST recursively constructs a BST from the linked list segment [start, end).
// The segment is treated as half-open interval: includes start but excludes end.
func buildBST(start *ListNode, end *ListNode) *TreeNode {
	// Base case: empty segment
	if start == end {
		return nil
	}

	// Find the middle node using slow-fast pointer technique
	// slow moves 1 step at a time, fast moves 2 steps
	// When fast reaches the end, slow will be at the middle
	slow, fast := start, start
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// Create root node with the middle element
	root := &TreeNode{Val: slow.Val}

	// Recursively build left subtree from [start, middle)
	root.Left = buildBST(start, slow)

	// Recursively build right subtree from [middle.Next, end)
	root.Right = buildBST(slow.Next, end)

	return root
}
