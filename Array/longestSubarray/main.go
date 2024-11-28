package main

/*
The whole idea is very similar to the longestOnes problem.
Both brute force and sliding window solutions are very similar to the longestOnes problem.
*/

// Brute Force
// Time Complexity O(n^2)
// Space Complexity O(1)
func LongestSubarray(nums []int) int {
	ans := 0

	// Iterate through the array with the starting pointer
	for i := 0; i < len(nums); i++ {
		// For each starting pointer, we initialize the temporary answer and flip count
		flip := false
		tmp := 0

		// Iterate through the array with the ending pointer
		for k := i; k < len(nums); k++ {
			// If we encounter a 1, we increment the temporary answer
			if nums[k] == 1 {
				tmp++
				continue
			}

			// If we encounter a 0 and we have not flipped yet, we flip it and increment the temporary answer
			if !flip {
				tmp++
				flip = true
				continue
			}

			// If we encounter a 0 and we have already flipped, we break
			break
		}

		// Update the answer with the maximum of the current valid subarray length
		ans = max(ans, tmp)
	}

	// We subtract 1 to represent the removal of at most one 0
	return ans - 1
}

// Sliding Window
// Time Complexity O(n)
// Space Complexity O(1)
func LongestSubarray2(nums []int) int {
	ans := 0
	left, right := 0, 0
	flip := false

	for right < len(nums) {
		// If we encounter a 0 and we have flipped, we shrink the window from the left
		if nums[right] == 0 && flip {
			for flip {
				// Only set flip to false if we encounter a 0
				if nums[left] == 0 {
					flip = false
				}
				left++
			}
		}

		// If we encounter a 0, we set flip to true
		if nums[right] == 0 {
			flip = true
		}

		// Update the answer with the current valid subarray length
		ans = max(ans, right-left+1)
		right++
	}

	// We subtract 1 to represent the removal of at most one 0
	return ans - 1
}

// Sliding Window 2 (Using Flip Count)
// Time Complexity O(n)
// Space Complexity O(1)
func LongestSubarray3(nums []int) int {
	ans, flip := 0, 0
	left, right := 0, 0

	for right < len(nums) {
		// If we encounter a 0, we increment the flip count
		if nums[right] == 0 {
			flip++
		}

		// If we have flipped more than once, we shrink the window from the left
		for flip == 2 {
			if nums[left] == 0 {
				flip--
			}
			left++
		}

		// Update the answer with the current valid subarray length
		ans = max(ans, right-left+1)
		right++
	}

	// We subtract 1 to represent the removal of at most one 0
	return ans - 1
}
