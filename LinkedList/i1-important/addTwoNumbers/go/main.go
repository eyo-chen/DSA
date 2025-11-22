package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	node := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		newNode := &ListNode{}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			newNode.Val = sum % 10
			carry = sum / 10
		} else {
			newNode.Val = sum
			carry = 0
		}

		node.Next = newNode
		node = node.Next
	}

	if carry != 0 {
		newNode := &ListNode{Val: carry}
		node.Next = newNode
	}

	return dummyHead.Next
}
