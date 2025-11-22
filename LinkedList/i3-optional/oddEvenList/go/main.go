package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {
	oddDummyHead, evenDummyHead := &ListNode{}, &ListNode{}
	oddPtr, evenPtr := oddDummyHead, evenDummyHead

	idx := 1
	for head != nil {
		node := head
		if idx%2 == 0 {
			evenPtr.Next = node
			evenPtr = evenPtr.Next
		} else {
			oddPtr.Next = node
			oddPtr = oddPtr.Next
		}

		idx++
		head = head.Next

		// It's not necessary to set cut off the node's next pointer
		// but it's a good practice to do so in case of cycle
		node.Next = nil
	}

	oddPtr.Next = evenDummyHead.Next
	return oddDummyHead.Next
}

func OddEvenList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddPtr := head
	evenPtr := head.Next
	evenHead := head.Next
	for evenPtr != nil && evenPtr.Next != nil {
		oddPtr.Next = evenPtr.Next
		oddPtr = oddPtr.Next

		evenPtr.Next = oddPtr.Next
		evenPtr = evenPtr.Next
	}

	oddPtr.Next = evenHead
	return head
}
