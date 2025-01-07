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

// Use Two Pointers (Updated 2025/01/07)
/*
This is the same logic as Merge2, but it's more readable.
We use ptr1 and ptr2 as our core condition
After the first for-loop, we only need to handle the remaining elements in nums2.
Why we don't need to handle the remaining elements in nums1?
Because the elements in nums1 are already sorted, and we already put the elements in nums1 to the sorted array.
Suppose the
nums1 = [1,3,5,0,0,0], m = 3
nums2 = [3,4,6], n = 3

- First iteration, ptr1 = 2, ptr2 = 2, idx = 5
  - compare 5 and 6, 6 is larger, so we put 6 to the sorted array.
	- nums1 = [1,3,5,0,0,6]
	- ptr2--, idx--
- Second iteration, ptr1 = 2, ptr2 = 1, idx = 4
  - compare 5 and 4, 5 is larger, so we put 5 to the sorted array.
	- nums1 = [1,3,5,0,5,6]
	- ptr1--, idx--
- Third iteration, ptr1 = 1, ptr2 = 1, idx = 3
  - compare 3 and 4, 4 is larger, so we put 4 to the sorted array.
	- nums1 = [1,3,5,4,5,6]
	- ptr2--, idx--
- Fourth iteration, ptr1 = 1, ptr2 = 0, idx = 2
  - compare 3 and 3, it's equal, so we put 3 of nums2 to the sorted array.
	- nums1 = [1,3,3,4,5,6]
	- ptr2--, idx--

Now, ptr2 is less than 0, so we can break the loop.
At this point, ptr1 is at 1, ptr2 is at -1, idx is at 0.
We can see that even thought we do not finish iterating the nums1, it's still sorted.
So, we don't need to handle the remaining elements in nums1.
*/

func Merge3(nums1 []int, m int, nums2 []int, n int) {
	ptr1, ptr2 := m-1, n-1
	idx := m + n - 1

	// Keep looping until one of the pointers reach the beginning of the array
	for ptr1 >= 0 && ptr2 >= 0 {
		if nums1[ptr1] > nums2[ptr2] {
			nums1[idx] = nums1[ptr1]
			ptr1--
		} else {
			nums1[idx] = nums2[ptr2]
			ptr2--
		}
		idx--
	}

	// If there are remaining elements in nums2, we can put them to the sorted array.
	for ptr2 >= 0 {
		nums1[idx] = nums2[ptr2]
		ptr2--
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
func Merge4(nums1 []int, m int, nums2 []int, n int) {
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

// Updated 12/15/2024
// This is the swapping solution from the beginning of the array when I try to solve this problem second time.
// The logic is just like the Merge3, but it solves the problem.
// After swapping, we make sure the nums2 is sorted.
// For example, nums1 = [1,8,9,0,0,0], m = 3, nums2 = [2,4,6], n = 3
// We first compare 1 with 2, 1 is less than 2, so we don't need to swap.
// Then we compare 8 with 2, 8 is greater than 2, so we swap 8 and 2.
// After swapping, nums1 becomes [1,2,9,0,0,0], nums2 becomes [8,4,6].
// In this case, nums2 is not sorted, so we need to sort it.
// nums2 = [8,4,6] -> [4,6,8]
// Then we continue the same logic
// The most important thing is that we need to make sure the nums2 is sorted after swapping.
func Merge5(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	for i := 0; i < m; i++ {
		// When nums1 element is less than nums2 element, we don't need to swap.
		if nums1[i] <= nums2[0] {
			continue
		}

		// Swap the nums1 element with the nums2 element.
		nums1[i], nums2[0] = nums2[0], nums1[i]

		// Sort the nums2 array.
		for k := 0; k < len(nums2)-1; k++ {
			if nums2[k] > nums2[k+1] {
				nums2[k], nums2[k+1] = nums2[k+1], nums2[k]
			}
		}
	}

	// Copy the remaining elements in nums2 to the end of nums1.
	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}
}
