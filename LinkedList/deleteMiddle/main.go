package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Find the length first, then find the middle node
func DeleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	// Find the length of the list
	length := 0
	ptr := head
	for ptr != nil {
		ptr = ptr.Next
		length++
	}

	// Find the middle node
	/*
		  Why stepToNodeBeforeMiddle := (length / 2) - 1?
			Let's use three examples:
			1. 1 -> 2
			   In this case, we want to remove 2, so we need to reference 1. (2 / 2) - 1 = 0
			2. 1 -> 2 -> 3
			   In this case, we want to remove 2, so we need to reference 1. (3 / 2) - 1 = 0
			3. 1 -> 2 -> 3 -> 4
			   In this case, we want to remove 3, so we need to reference 2. (4 / 2) - 1 = 1
	*/
	stepToNodeBeforeMiddle := (length / 2) - 1
	ptr = head
	for i := 0; i < stepToNodeBeforeMiddle; i++ {
		ptr = ptr.Next
	}

	// Remove the middle node
	ptr.Next = ptr.Next.Next

	return head
}

// Use two pointers, one is slow, one is fast
// The fast pointer will move two steps each time, while the slow pointer will move one step each time
// When the fast pointer reach the end of the list(ptr == nil or ptr.Next == nil), the slow pointer will be at the previous node of the middle node
/*
Why slow starts from the dummy head while fast starts from the head?
Let's use four examples:
1. 1
   In this case, we want to remove 1, so we need to reference the node before 1, which is the dummy head.
	 slow(dummyHead): x, fast(head): 1
	 because fast.Next is nil, so we don't move anything
	 Finally, we set slow.Next = slow.Next.Next to remove 1
2. 1 -> 2
   In this case, we want to remove 2, so we need to reference the node before 2, which is 1.
	 slow(dummyHead): x, fast(head): 1
	 After moving, slow: 1, fast: x(tail)
	 Finally, we set slow.Next = slow.Next.Next to remove 2
3. 1 -> 2 -> 3
   In this case, we want to remove 2, so we need to reference the node before 2, which is 1.
	 slow(dummyHead): x, fast(head): 1
	 After moving, slow: 1, fast: 3
	 Finally, we set slow.Next = slow.Next.Next to remove 2
4. 1 -> 2 -> 3 -> 4
   In this case, we want to remove 3, so we need to reference the node before 3, which is 2.
	 slow(dummyHead): x, fast(head): 1
	 After moving, slow: 2, fast: x(tail)
	 Finally, we set slow.Next = slow.Next.Next to remove 3
*/
func DeleteMiddle2(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	// slow starts from the dummy head, while fast starts from the head
	slow, fast := dummyHead, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	// return the head of the list
	/*
		Note that we can't directly return head because we need to consider the case when the input list is only one node
		1 -> x
		In this case, slow starts from the dummy head, and fast starts from the head
		x -> 1 -> x
		slow: x, fast: 1
		Because fast.Next is nil, so we don't move anything
		Finally, we set slow.Next = slow.Next.Next to remove 1
		In the end, we should return dummyHead.Next, which is the head of the list
	*/
	return dummyHead.Next
}
