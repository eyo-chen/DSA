package main

// CountZeroSubarraysBruteForce counts all subarrays that contain only zeros using a brute force approach.
// For each zero found, it extends right to count all consecutive zeros.
// Time Complexity: O(n²) - nested loops in worst case when array is all zeros
// Space Complexity: O(1) - only uses constant extra space
func CountZeroSubarraysBruteForce(nums []int) int64 {
	var totalCount int64

	// Iterate through each element
	for i, num := range nums {
		// Skip non-zero elements
		if num != 0 {
			continue
		}

		// Count all zero-filled subarrays starting at index i
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				totalCount++
			} else {
				// Stop when we hit a non-zero element
				break
			}
		}
	}

	return totalCount
}

// CountZeroSubarraysSlidingWindow counts all subarrays that contain only zeros using a sliding window approach.
// Maintains a window that contains only zeros by tracking non-zero elements.
// Time Complexity: O(n) - each element is visited at most twice (by left and right pointers)
// Space Complexity: O(1) - only uses constant extra space
func CountZeroSubarraysSlidingWindow(nums []int) int64 {
	var totalCount int64
	nonZeroCount := 0
	left, right := 0, 0

	for right < len(nums) {
		// If current element is non-zero, increment counter and move right
		if nums[right] != 0 {
			nonZeroCount++
			right++
			continue
		}

		// Shrink window from left until all non-zero elements are removed
		for nonZeroCount > 0 {
			if nums[left] != 0 {
				nonZeroCount--
			}
			left++
		}

		// All subarrays from left to right (inclusive) are valid zero-filled subarrays
		totalCount += int64(right - left + 1)
		right++
	}

	return totalCount
}

// CountZeroSubarraysTwoPointers counts all subarrays that contain only zeros using a two-pointer approach.
// When a zero is found, it finds the entire consecutive zero segment and counts all subarrays within it.
// Time Complexity: O(n) - each element is visited at most once
// Space Complexity: O(1) - only uses constant extra space
func CountZeroSubarraysTwoPointers(nums []int) int64 {
	var totalCount int64 = 0

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			// Mark the start of a zero segment
			left = right

			// Extend right to find all consecutive zeros
			for right < len(nums) && nums[right] == 0 {
				// For each position, count subarrays ending at current right index
				// (right - left + 1) represents subarrays: [left:right+1], [left+1:right+1], ..., [right:right+1]
				totalCount += int64(right - left + 1)
				right++
			}
		} else {
			right++
		}
	}

	return totalCount
}

// CountZeroSubarraysOptimal counts all subarrays that contain only zeros using an optimal approach.
// Tracks consecutive zeros and accumulates the count incrementally.
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(1) - only uses constant extra space
func CountZeroSubarraysOptimal(nums []int) int64 {
	var totalCount, consecutiveZeros int64 = 0, 0

	for _, num := range nums {
		if num == 0 {
			// Increment consecutive zero count
			consecutiveZeros++
			// Add all subarrays ending at current position
			// If we have k consecutive zeros, we can form k new subarrays ending here
			totalCount += consecutiveZeros
		} else {
			// Reset counter when we encounter a non-zero element
			consecutiveZeros = 0
		}
	}

	return totalCount
}
