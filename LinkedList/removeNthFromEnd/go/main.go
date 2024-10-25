package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Calculate the length of list
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	ptr := head
	for ptr != nil {
		size++
		ptr = ptr.Next
	}

	dummyHead := &ListNode{Next: head}
	ptr = dummyHead
	for i := 0; i < size-n; i++ {
		ptr = ptr.Next
	}
	ptr.Next = ptr.Next.Next

	return dummyHead.Next
}

// Use two pointers
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dummyHead.Next
}
