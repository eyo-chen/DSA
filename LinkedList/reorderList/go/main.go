package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use Array to store the nodes
func ReorderList(head *ListNode) {
	ptr := head
	s := []*ListNode{}
	for ptr != nil {
		// add the node to the array
		s = append(s, ptr)

		// hold the reference of next node
		next := ptr.Next

		// break the link
		ptr.Next = nil

		// move to the next node
		ptr = next
	}

	// use two pointers to re-wire the nodes
	left, right := 0, len(s)-1
	for left < right {
		// make the left pointer point to the right pointer
		// e.g. 1 -> 2 -> 3 -> 4 -> 5
		// s[left] = 1, s[right] = 5
		// 1 -> 5
		s[left].Next = s[right]

		// update the left pointer
		// left update to 2
		left++

		// check if the left pointer equals to the right pointer
		if left == right {
			break
		}

		// make the right pointer point to the left pointer
		// make 5 -> 2
		s[right].Next = s[left]

		// update the right pointer
		// right update to 4
		right--
	}
}

// Use Two Pointers to find the middle of the list and reverse the second half
func ReorderList2(head *ListNode) {
	// find the middle of the list
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse the second half of the list
	var prev *ListNode
	// start at middle's next node
	cur := slow.Next
	for cur != nil {
		next := cur.Next

		cur.Next = prev

		prev = cur
		cur = next
	}

	// break the link of the first half
	// slow.Next is the head of the second half
	slow.Next = nil

	// merge the two halves
	list1, list2 := head, prev

	// list2 is guaranteed to be shorter than or equal to list1
	for list2 != nil {
		next1, next2 := list1.Next, list2.Next

		list1.Next = list2
		list2.Next = next1

		list1 = next1
		list2 = next2
	}
}
