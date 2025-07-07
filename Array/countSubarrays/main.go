package main

func CountSubarrays(nums []int, k int) int64 {
	maxEle := nums[0]
	for _, n := range nums {
		maxEle = max(n, maxEle)
	}

	var ans int64 = 0
	for i := 0; i < len(nums); i++ {
		count := 0

		for j := i; j < len(nums); j++ {
			if nums[j] == maxEle {
				count++
			}

			if count >= k {
				ans++
			}
		}
	}

	return ans
}

func CountSubarrays1(nums []int, k int) int64 {
	// Step 1: Find the maximum element in the array
	maxEle := nums[0]
	for _, n := range nums {
		if n > maxEle {
			maxEle = n
		}
	}

	// Step 2: Sliding window to count subarrays
	var ans int64 = 0
	count := 0 // Count of maxEle in the current window
	left := 0  // Left pointer of the sliding window

	for right := 0; right < len(nums); right++ {
		// Increment count if we find maxEle
		if nums[right] == maxEle {
			count++
		}

		// Shrink window while count >= k
		for count >= k && left <= right {
			// Count all subarrays from current window to the end
			// e.g. nums = [1,3,2,3,4,5]
			// when window size is [1,2,3,3]
			// since we know the current window size is a valid subarray
			// so "ANY extensive subarray from right ptr is also valid"
			// both [1,2,3,3,4] and [1,2,3,3,4,5] are all valid
			// Therefore, the number of subarrays in the current valid window
			// Is the number of element from right to the end -> (len(nums) - right)
			ans += int64(len(nums) - right)

			if nums[left] == maxEle {
				count--
			}

			left++ // Shrink the window
		}
	}

	return ans
}
