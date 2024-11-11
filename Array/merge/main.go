package main

// Use Temporary Array
func Merge(nums1 []int, m int, nums2 []int, n int) {
	// sorted is a temporary array to store the sorted elements
	sorted := make([]int, len(nums1))

	idx := 0
	ptr1, ptr2 := 0, 0

	// keep looping until one of the pointers reach the end of the array
	for ptr1 < m && ptr2 < n {
		// if the elements in nums1 is less than the elements in nums2,
		// we can put the elements in nums1 to the sorted array.
		if nums1[ptr1] < nums2[ptr2] {
			sorted[idx] = nums1[ptr1]
			ptr1++
		} else {
			// otherwise, we can put the elements in nums2 to the sorted array.
			sorted[idx] = nums2[ptr2]
			ptr2++
		}

		idx++
	}

	// Note that technically, only one of the following loops will be executed.
	// Because we are looping until one of the pointers reach the end of the array.

	// if the elements in nums1 is not used up, we can put the remaining elements to the sorted array.
	for ptr1 < m {
		sorted[idx] = nums1[ptr1]
		ptr1++
		idx++
	}

	// if the elements in nums2 is not used up, we can put the remaining elements to the sorted array.
	for ptr2 < n {
		sorted[idx] = nums2[ptr2]
		ptr2++
		idx++
	}

	// copy the sorted elements back to the original array(nums1)
	copy(nums1, sorted)
}

// Use Two Pointers
func Merge2(nums1 []int, m int, nums2 []int, n int) {
	// idx is the pointer to the end of the sorted array(m + n - 1)
	// ptr1 is the pointer to the end of the first array(m - 1)
	// ptr2 is the pointer to the end of the second array(n - 1)
	idx, ptr1, ptr2 := len(nums1)-1, m-1, n-1

	// keep looping until the sorted pointer reach the beginning of the array
	for idx >= 0 {
		// if the elements in nums2 is used up, or the elements in nums1 is greater than the elements in nums2,
		// we can put the elements in nums1 to the sorted array.
		// ptr1 >= 0 is for safety check, because ptr1 may be negative.
		if ptr2 < 0 || (ptr1 >= 0 && nums1[ptr1] > nums2[ptr2]) {
			nums1[idx] = nums1[ptr1]
			ptr1--
		} else {
			// otherwise, we can put the elements in nums2 to the sorted array.
			nums1[idx] = nums2[ptr2]
			ptr2--
		}

		// move the sorted pointer to the beginning of the array
		idx--
	}
}

// Note that this solution is try to fill up the sorted elements from the beginning of the array.
// But it will lead to the problem
/*
The main logical problem in this solution is in the swapping logic. Here's why it's wrong:
The current implementation tries to swap elements between nums1 and nums2 whenever it finds a smaller element in nums2, but this approach is problematic because:
- It modifies nums2 during the process, which shouldn't happen
- The swapping logic can lead to incorrect ordering
- It might miss some elements or swap them multiple times

Let's see an example where it fails:
nums1 = [1,3,5,0,0,0], m = 3
nums2 = [2,4,6], n = 3

When running your current code:
- First iteration (i=0): 1 is compared with [2,4,6], no swap happens
- Second iteration (i=1): 3 is compared with [2,4,6], swaps with 2
	- nums1 becomes [1,2,5,0,0,0]
	- nums2 becomes [3,4,6]
- Third iteration (i=2): 5 is compared with [3,4,6], swaps with 3
	- nums1 becomes [1,2,3,0,0,0]
	- nums2 becomes [5,4,6]

The final result will be incorrect because nums2 got modified and is no longer sorted.
*/
func Merge3(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	for i := 0; i < m; i++ {
		n1 := nums1[i]

		k := 0
		for ; k < n; k++ {
			if nums2[k] < n1 {
				break
			}
		}

		nums1[i], nums2[k] = nums2[k], nums1[i]
	}

	for i, n := range nums2 {
		nums1[m+i] = n
	}
}
