package main

// Brute Force
// Time Complexity O(n^2)
// Space Complexity O(1)
func LongestOnes(nums []int, k int) int {
	ans := 0

	// Iterate through the array with the starting pointer
	for i := 0; i < len(nums); i++ {
		// For each starting pointer, we initialize the temporary answer and flip count
		tmp, flip := 0, 0

		// Iterate through the array with the ending pointer
		for j := i; j < len(nums); j++ {
			// If we encounter a 1, we increment the temporary answer
			if nums[j] == 1 {
				tmp++
				continue
			}

			// If we encounter a 0 and we have flipped k 0s already, we break
			if flip == k {
				break
			}

			// Otherwise, we flip the 0 and increment the temporary answer
			flip++
			tmp++
		}

		ans = max(ans, tmp)
	}

	return ans
}

// Sliding Window
// Time Complexity O(n)
// Space Complexity O(1)
func LongestOnes2(nums []int, k int) int {
	ans, flip := 0, 0
	left, right := 0, 0

	for right < len(nums) {
		// First checks if we encounter a 0 and update the flip count
		if nums[right] == 0 {
			flip++
		}

		// Then ensures that the flip count does not exceed k by shrinking the window from the left
		for flip > k {
			// Only decrement the flip count if we encounter a 0
			if nums[left] == 0 {
				flip--
			}
			left++
		}

		// Only then updates the answer with the current valid window size
		ans = max(ans, right-left+1)

		// Always moves the right pointer forward
		right++
	}

	return ans
}

// Wrong Answer
// This is my initial attempt at the problem
// This is wrong and way too complicated
func LongestOnes3(nums []int, k int) int {
	ans, flip := 0, 0
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 1 {
			ans = max(ans, right-left+1)
			right++
			continue
		}

		if flip <= k {
			ans = max(ans, right-left+1)
			flip++
			right++
			continue
		}

		for left < right && flip > k {
			if nums[left] == 0 {
				flip--
			}
			left++
		}

		right++
	}

	return ans
}
