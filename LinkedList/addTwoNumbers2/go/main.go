package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers - Stack-based approach
// Time: O(n + m), Space: O(n + m) where n, m are lengths of input lists
// This approach uses stacks to process digits from right to left (least to most significant)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Convert linked lists to stacks (arrays) for easier right-to-left processing
	stack1, stack2 := genStack(l1), genStack(l2)

	// Start from the rightmost digits (least significant)
	idx1, idx2 := len(stack1)-1, len(stack2)-1
	carry := 0

	// Build result list from right to left by prepending nodes
	var curNode *ListNode

	// Process digits from both stacks until exhausted
	for idx1 >= 0 || idx2 >= 0 {
		// Start with carry from previous calculation
		sum := carry

		// Add digit from first number if available
		if idx1 >= 0 {
			sum += stack1[idx1]
			idx1--
		}

		// Add digit from second number if available
		if idx2 >= 0 {
			sum += stack2[idx2]
			idx2--
		}

		// Handle carry: if sum >= 10, carry to next position
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}

		// Create new node and prepend to result (builds list in correct order)
		node := &ListNode{Val: sum, Next: curNode}
		curNode = node
	}

	// If there's remaining carry, add it as most significant digit
	if carry != 0 {
		node := &ListNode{Val: carry, Next: curNode}
		curNode = node
	}

	return curNode
}

// genStack converts a linked list to a slice for stack-like access
func genStack(head *ListNode) []int {
	res := []int{}

	// Traverse list and collect all values
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// AddTwoNumbers2 - Direct manipulation approach
// Time: O(n + m), Space: O(1) extra space (not counting output)
// This approach builds the result in reverse order, then handles carry during reversal
func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// Find lengths to determine which list is longer
	len1, len2 := countLength(l1), countLength(l2)

	// Identify longer and shorter lists for alignment
	maxList, minList := l1, l2
	maxLen, minLen := max(len1, len2), min(len1, len2)
	if len2 > len1 {
		maxList, minList = l2, l1
	}

	// Phase 1: Build initial sum list in reverse order (most significant first)
	idx := 0
	var curNode *ListNode
	for idx < maxLen {
		// Always add digit from longer list
		sum := maxList.Val

		// Add digit from shorter list only when aligned
		// (idx >= maxLen-minLen) ensures we start adding shorter list digits
		// at the correct position (right-aligned)
		if idx >= maxLen-minLen {
			sum += minList.Val
			minList = minList.Next
		}
		maxList = maxList.Next

		// Build list in reverse order (prepend each node)
		node := &ListNode{Val: sum, Next: curNode}
		curNode = node
		idx++
	}

	// Phase 2: Handle carry while reversing the list
	// This processes from most significant to least significant in the reverse list,
	// which correctly handles carry propagation
	node := curNode
	var prev *ListNode
	carry := 0

	for node != nil {
		// Add carry from previous digit
		node.Val += carry

		// Calculate new carry and adjust current digit
		if node.Val >= 10 {
			carry = node.Val / 10
			node.Val = node.Val % 10
		} else {
			carry = 0
		}

		// Reverse the list: point current node to previous
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	// If final carry exists, add it as most significant digit
	if carry >= 1 {
		node := &ListNode{Val: carry, Next: prev}
		prev = node
	}

	return prev
}

// countLength returns the number of nodes in a linked list
func countLength(head *ListNode) int {
	res := 0
	for head != nil {
		head = head.Next
		res++
	}
	return res
}
