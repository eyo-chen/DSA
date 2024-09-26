package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Naive Approach
func MergeKLists(lists []*ListNode) *ListNode {
	minHeap := NewMinHeap(0)

	// Insert all values into a min heap
	for _, l := range lists {
		for l != nil {
			minHeap.Insert(l.Val)
			l = l.Next
		}
	}

	dummyHead := &ListNode{}
	ptrNode := dummyHead

	// Pop all values from the min heap and create a new linked list
	for minHeap.Len() > 0 {
		n := &ListNode{Val: minHeap.Pull()}
		ptrNode.Next = n
		ptrNode = ptrNode.Next
	}

	return dummyHead.Next
}

// Optimized Approach
func MergeKLists1(lists []*ListNode) *ListNode {
	minHeap := NewMinHeapNode(len(lists))
	dummyHead := &ListNode{}
	nodePtr := dummyHead

	// Insert all heads of the linked lists into a min heap
	for _, l := range lists {
		// Avoid nil pointer
		// Because we're going to reference the Next pointer in the next step
		if l != nil {
			minHeap.Insert(l)
		}
	}

	// Keep pulling the smallest node from the min heap
	for minHeap.Len() > 0 {
		// Pull the smallest node from the min heap
		minNode := minHeap.Pull()

		// Add the smallest node to the new linked list
		nodePtr.Next = minNode

		// Move the pointer to the next node
		nodePtr = nodePtr.Next

		// If the smallest node has a next node, insert the next node into the min heap
		if minNode.Next != nil {
			minHeap.Insert(minNode.Next)
		}
	}

	return dummyHead.Next
}

// Divide and Conquer Approach
func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return helper(lists, 0, len(lists)-1)
}

func helper(lists []*ListNode, start, end int) *ListNode {
	// when start and end are the same, return the list at that index
	if start == end {
		return lists[start]
	}

	// when start and end are consecutive, merge the two lists
	if start+1 == end {
		return mergeList(lists[start], lists[end])
	}

	// find the middle index
	mid := start + (end-start)/2
	left := helper(lists, start, mid)
	right := helper(lists, mid+1, end)
	return mergeList(left, right)
}

// Iterative Approach
func MergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		// append the merged list to the end of the list
		lists = append(lists, mergeList(lists[0], lists[1]))

		// remove the first two lists
		lists = lists[2:]
	}

	return lists[0]
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil {
		return l2
	}

	dummyHead := &ListNode{}
	curNode := dummyHead

	// While both lists are not nil
	for l1 != nil && l2 != nil {
		// Compare the values of the current nodes of both lists
		if l1.Val > l2.Val {
			// If the value of l1 is greater than l2, append l2 to the new list
			curNode.Next = l2
			l2 = l2.Next
		} else {
			// If the value of l1 is less than l2, append l1 to the new list
			curNode.Next = l1
			l1 = l1.Next
		}

		curNode = curNode.Next
	}

	// If one of the lists is not nil, append the remaining nodes to the new list
	// We can guarantee only one of the list is not nil because the for-loop condition is `&&`
	if l1 != nil {
		curNode.Next = l1
	}
	if l2 != nil {
		curNode.Next = l2
	}

	return dummyHead.Next
}
