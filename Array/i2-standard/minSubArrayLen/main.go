package main

// Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func MinSubArrayLen(target int, nums []int) int {
	ans := 0

	// Loop through the array
	for i := range len(nums) {
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

// MinSubArrayLen3 finds the minimal length of a contiguous subarray whose sum is >= target.
// Uses a sliding window approach: expand the window by moving right pointer, then shrink
// from the left while the sum remains >= target.
// Time Complexity: O(n) - each element is visited at most twice (once by right, once by left)
// Space Complexity: O(1) - only uses constant extra space
func MinSubArrayLen3(target int, nums []int) int {
	minLength := 0 // Track the minimum subarray length found (0 means no valid subarray yet)
	windowSum := 0 // Current sum of elements in the sliding window
	left := 0      // Left boundary of the sliding window

	// Expand the window by moving the right boundary
	for right := range len(nums) {
		// Add the current element to the window sum
		windowSum += nums[right]

		// Shrink the window from the left while sum is still >= target
		// This ensures we find the smallest possible window ending at right
		for windowSum >= target {
			// Calculate current window size
			currentLength := right - left + 1

			// Update minLength if this is the first valid window or smaller than previous
			if minLength == 0 || currentLength < minLength {
				minLength = currentLength
			}

			// Shrink window: remove the leftmost element and move start pointer right
			windowSum -= nums[left]
			left++
		}

		right++
	}

	return minLength
}
