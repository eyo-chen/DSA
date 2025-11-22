package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev, cur := dummyHead, head

	for cur != nil && cur.Next != nil {
		prev.Next = cur.Next
		prev = prev.Next

		cur.Next = prev.Next
		prev.Next = cur

		prev = prev.Next
		cur = cur.Next
	}

	return dummyHead.Next
}
