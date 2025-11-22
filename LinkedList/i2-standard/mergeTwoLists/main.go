package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	ptr := dummyHead

	// Fix 1: loop until one of the list is nil
	for list1 != nil && list2 != nil {
		// safely reference list1 and list2
		// because we can make sure both of them are not nil
		if list1.Val < list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else {
			ptr.Next = list2
			list2 = list2.Next
		}

		ptr = ptr.Next
	}

	// if one of them is not nil
	// just make the ptr to connect it
	if list1 != nil {
		ptr.Next = list1
	} else {
		ptr.Next = list2
	}

	return dummyHead.Next
}
