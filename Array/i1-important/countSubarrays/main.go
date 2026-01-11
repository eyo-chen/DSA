package main

// CountSubarrays counts the number of subarrays containing at least k occurrences
// of the maximum element in the array.
//
// Approach: Brute force - check all possible subarrays
// Time Complexity: O(n²) where n is the length of nums
// Space Complexity: O(1) - only using constant extra space
func CountSubarrays(nums []int, k int) int64 {
	// Find the maximum element in the array
	maxElement := nums[0]
	for _, num := range nums {
		maxElement = max(num, maxElement)
	}

	var result int64 = 0

	// Try all possible starting positions
	for start := 0; start < len(nums); start++ {
		maxCount := 0

		// Extend the subarray from current starting position
		for end := start; end < len(nums); end++ {
			// Count occurrences of max element in current subarray
			if nums[end] == maxElement {
				maxCount++
			}

			// If we have at least k occurrences, this subarray is valid
			if maxCount >= k {
				result++
			}
		}
	}

	return result
}

// CountSubarraysOptimized counts the number of subarrays containing at least k occurrences
// of the maximum element in the array.
//
// Approach: Sliding window - maintain a window with at least k occurrences of max element.
// For each valid window starting at 'left', all subarrays extending from 'right' to the end
// are also valid, giving us (len(nums) - right) valid subarrays.
//
// Time Complexity: O(n) where n is the length of nums
// Space Complexity: O(1) - only using constant extra space
func CountSubarraysOptimized(nums []int, k int) int64 {
	// Find the maximum element in the array
	maxElement := nums[0]
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}

	var result int64 = 0
	maxCount := 0 // Count of maxElement in the current window
	left := 0     // Left boundary of the sliding window

	// Expand the window by moving the right boundary
	for right := range len(nums) {
		// Increment count when we encounter the max element
		if nums[right] == maxElement {
			maxCount++
		}

		// Shrink the window while we have at least k occurrences of max element
		// For each valid window position, count all valid subarrays
		for maxCount >= k {
			// Key insight: If window [left...right] has k occurrences of maxElement,
			// then all subarrays [left...j] where j >= right are also valid.
			// Example: nums = [1,3,2,3,3], k=2, maxElement=3
			// When window is [3,2,3] (indices 1-3), valid subarrays are:
			//   - [3,2,3] (extends to index 3)
			//   - [3,2,3,3] (extends to index 4)
			// Therefore, we add (len(nums) - right) = 5 - 3 = 2 subarrays
			result += int64(len(nums) - right)

			// Move left boundary to find the next valid window
			if nums[left] == maxElement {
				maxCount--
			}
			left++
		}
	}

	return result
}
