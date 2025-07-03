package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use Array
func PairSum(head *ListNode) int {
	hashTable := []int{}
	ans := 0

	// Store the values of the linked list in an array
	for node := head; node != nil; node = node.Next {
		hashTable = append(hashTable, node.Val)
	}

	// Find the maximum sum of pairs
	// Sum = hashTable[i] + hashTable[len(hashTable)-1-i]
	for i := 0; i < len(hashTable)/2; i++ {
		ans = max(ans, hashTable[i]+hashTable[len(hashTable)-1-i])
	}

	return ans
}

// Use Stack
func PairSum2(head *ListNode) int {
	ans := 0
	stack := []*ListNode{}
	fast, slow := head, head

	// Store the first half of the linked list in a stack
	// And also move the slow pointer to the middle of the linked list
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow)
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Find the maximum sum of pairs
	// Sum = slow.Val + stack[len(stack)-1].Val
	for slow != nil {
		ans = max(ans, slow.Val+stack[len(stack)-1].Val)

		slow = slow.Next
		stack = stack[:len(stack)-1] // Pop the stack
	}

	return ans
}

// Reverse the Second Half
func PairSum3(head *ListNode) int {
	ans := 0
	slow, fast := head, head

	// Find the middle of the linked list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the linked list
	var prev *ListNode
	cur := slow
	for cur != nil {
		next := cur.Next

		cur.Next = prev

		prev = cur
		cur = next
	}

	// Find the maximum sum of pairs
	// Sum = head.Val + prev.Val
	// It's important to use prev as the condition, instead of head
	// because head still references the original linked list
	// e.g. 1 -> 2 -> 3 -> 4 -> 5 -> 6
	// After reversing the second half, the linked list becomes 1 -> 2 -> 3 -> 4 -> x
	// If we use head as the condition, I will iterate three times, which is not correct
	// Instead, we use prev as the condition, which will iterate twice
	// because prev is the first node of the reversed second half: 6 -> 5 -> 4 -> x
	// In this case, we only iterate twice, which is correct
	for prev != nil {
		ans = max(ans, head.Val+prev.Val)
		head = head.Next
		prev = prev.Next
	}

	return ans
}
