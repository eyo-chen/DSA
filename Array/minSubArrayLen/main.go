package main

// Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func MinSubArrayLen(target int, nums []int) int {
	ans := 0

	// Loop through the array
	for i := 0; i < len(nums); i++ {
		acc := 0

		// For each element, we loop through the rest of the array
		// to find the smallest subarray that sums to the target
		for k := i; k < len(nums); k++ {
			// update accumulated sum
			acc += nums[k]

			// if the accumulated sum is greater than or equal to the target
			// we update the answer
			if ans == 0 || acc >= target {
				ans = k - i + 1
			}

			// break if the accumulated sum is greater than or equal to the target
			if acc >= target {
				break
			}
		}
	}

	return ans
}

// Sliding Window
// Time Complexity: O(n)
// Space Complexity: O(1)
func MinSubArrayLen2(target int, nums []int) int {
	ans := 0
	acc := 0
	right, left := 0, 0

	for right < len(nums) {
		// update accumulated sum
		acc += nums[right]

		// if the accumulated sum is greater than or equal to the target
		// we update the answer
		// Also, we keep shrinking the window from the left until the accumulated sum is less than the target
		for acc >= target {
			curSize := right - left + 1
			if ans == 0 || curSize < ans {
				ans = curSize
			}

			// move the left pointer to the right
			// to shrink the window
			acc -= nums[left]
			left++
		}

		// move the right pointer to the right
		// to expand the window
		right++
	}

	return ans
}
