package main

// LengthOfLIS finds the length of the longest increasing subsequence using recursive approach with backtracking.
// This approach explores all possible subsequences by making two choices at each element: include it or skip it.
// Time Complexity: O(2^n) - exponential, as we explore all possible combinations
// Space Complexity: O(n) - recursion stack depth and temporary subsequence storage
func LengthOfLIS(nums []int) int {
	return helper(nums, []int{}, 0)
}

// helper is a recursive function that builds subsequences and finds the maximum length
func helper(nums []int, currentSubsequence []int, currentIndex int) int {
	// Base case: if we've processed all elements, return the length of current subsequence
	if currentIndex >= len(nums) {
		return len(currentSubsequence)
	}

	// Option 1: Skip the current element and move to next
	skipCurrentElement := helper(nums, currentSubsequence, currentIndex+1)

	// Option 2: Include the current element (if valid)
	includeCurrentElement := -1

	// Check if we can include current element in the increasing subsequence
	// We can include if: subsequence is empty OR current element is greater than last element
	if len(currentSubsequence) == 0 || nums[currentIndex] > currentSubsequence[len(currentSubsequence)-1] {
		// Create new subsequence with current element appended
		newSubsequence := append(currentSubsequence, nums[currentIndex])
		includeCurrentElement = helper(nums, newSubsequence, currentIndex+1)
	}

	// Return the maximum length from both choices
	return max(skipCurrentElement, includeCurrentElement)
}

// LengthOfLIS2 finds the length of the longest increasing subsequence using dynamic programming.
// This approach builds up the solution by storing the length of LIS ending at each position.
// Time Complexity: O(n^2) - nested loops through the array
// Space Complexity: O(n) - DP array to store intermediate results
func LengthOfLIS2(nums []int) int {
	// DP array where dp[i] represents the length of LIS ending at index i
	dp := make([]int, len(nums))
	maxLISLength := 1 // At minimum, each element forms a subsequence of length 1

	// Initialize: each element by itself forms a subsequence of length 1
	for i := range dp {
		dp[i] = 1
	}

	// Fill the DP array by checking all previous elements
	for currentPos := 1; currentPos < len(dp); currentPos++ {
		// Check all elements before current position
		for prevPos := 0; prevPos < currentPos; prevPos++ {
			// If current element is greater than previous element,
			// we can extend the LIS ending at prevPos
			if nums[currentPos] > nums[prevPos] {
				// Update LIS length ending at current position
				dp[currentPos] = max(dp[currentPos], dp[prevPos]+1)
				// Keep track of the overall maximum LIS length found so far
				maxLISLength = max(maxLISLength, dp[currentPos])
			}
		}
	}

	return maxLISLength
}
