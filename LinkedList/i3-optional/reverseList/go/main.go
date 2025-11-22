package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		// hold the reference of next node
		next := cur.Next

		// reverse
		cur.Next = prev

		// update prev and cur
		prev = cur
		cur = next
	}

	return prev
}
