package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// IsPalindrome checks if a linked list is a palindrome using O(n) space approach
// Time Complexity: O(n), Space Complexity: O(n)
func IsPalindrome(head *ListNode) bool {
	// Convert linked list to array for easy comparison
	values := []int{}

	// Traverse the linked list and collect all values
	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	// Use two pointers to check if the array is a palindrome
	leftPointer, rightPointer := 0, len(values)-1
	for leftPointer < rightPointer {
		if values[leftPointer] != values[rightPointer] {
			return false
		}
		leftPointer++
		rightPointer--
	}

	return true
}

// IsPalindrome2 checks if a linked list is a palindrome using O(1) space approach
// Time Complexity: O(n), Space Complexity: O(1)
func IsPalindrome2(head *ListNode) bool {
	// Step 1: Find the middle of the linked list using two pointers
	fastPointer, slowPointer := head, head

	// Create a dummy node to help track the node before the middle
	dummyHead := &ListNode{Next: head}
	nodeBeforeMiddle := dummyHead

	// Fast pointer moves 2 steps, slow pointer moves 1 step
	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next
		nodeBeforeMiddle = nodeBeforeMiddle.Next
	}

	// Step 2: Split the list into two halves
	// Cut the connection between first half and second half
	nodeBeforeMiddle.Next = nil

	// Step 3: Reverse the second half of the linked list
	var previousNode *ListNode
	currentNode := slowPointer
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	// Step 4: Compare the first half with the reversed second half
	firstHalfHead := head
	secondHalfHead := previousNode
	for firstHalfHead != nil {
		if firstHalfHead.Val != secondHalfHead.Val {
			return false
		}
		firstHalfHead = firstHalfHead.Next
		secondHalfHead = secondHalfHead.Next
	}

	return true
}
