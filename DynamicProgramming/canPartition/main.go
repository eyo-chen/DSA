package main

import "fmt"

// CanPartition - Brute force recursive approach
// Time: O(2^n), Space: O(n) for recursion stack
func CanPartition(nums []int) bool {
	// Calculate total sum of all numbers
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// Use recursive helper to try all possible partitions
	return partitionHelper(nums, 0, totalSum, 0)
}

// partitionHelper recursively tries to partition the array
// currentIndex: current position in nums array
// totalSum: sum of all numbers in nums
// currentSubsetSum: sum of current subset being built
func partitionHelper(nums []int, currentIndex int, totalSum int, currentSubsetSum int) bool {
	// Base case: reached end of array without finding valid partition
	if currentIndex >= len(nums) {
		return false
	}

	// Check if current subset sum equals remaining sum (equal partition found)
	remainingSum := totalSum - currentSubsetSum
	if remainingSum == currentSubsetSum {
		return true
	}

	// Try two choices: exclude current number OR include current number
	excludeCurrent := partitionHelper(nums, currentIndex+1, totalSum, currentSubsetSum)
	includeCurrent := partitionHelper(nums, currentIndex+1, totalSum, currentSubsetSum+nums[currentIndex])

	return excludeCurrent || includeCurrent
}

// CanPartition1 - Recursive approach with memoization
// Time: O(n * sum), Space: O(n * sum)
func CanPartition1(nums []int) bool {
	// Calculate total sum of all numbers
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// Create memoization map to cache results
	memoCache := map[string]bool{}

	return partitionHelperWithMemo(nums, 0, totalSum, 0, memoCache)
}

// partitionHelperWithMemo recursively tries to partition with memoization
func partitionHelperWithMemo(nums []int, currentIndex int, totalSum int, currentSubsetSum int, memoCache map[string]bool) bool {
	// Base case: reached end of array without finding valid partition
	if currentIndex >= len(nums) {
		return false
	}

	// Check if current subset sum equals remaining sum (equal partition found)
	remainingSum := totalSum - currentSubsetSum
	if remainingSum == currentSubsetSum {
		return true
	}

	// Create unique key for memoization (index + current sum)
	cacheKey := fmt.Sprintf("%d-%d", currentIndex, currentSubsetSum)
	if cachedResult, exists := memoCache[cacheKey]; exists {
		return cachedResult
	}

	// Try two choices: exclude current number OR include current number
	excludeCurrent := partitionHelperWithMemo(nums, currentIndex+1, totalSum, currentSubsetSum, memoCache)
	includeCurrent := partitionHelperWithMemo(nums, currentIndex+1, totalSum, currentSubsetSum+nums[currentIndex], memoCache)

	result := excludeCurrent || includeCurrent
	memoCache[cacheKey] = result

	return result
}

// CanPartition2 - Bottom-up approach using set of possible sums
// Time: O(n * sum), Space: O(sum)
func CanPartition2(nums []int) bool {
	// Calculate total sum of all numbers
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is odd, cannot partition into two equal subsets
	if totalSum%2 != 0 {
		return false
	}

	targetSum := totalSum / 2

	// Use map as set to track all possible subset sums
	possibleSums := make(map[int]bool)
	possibleSums[0] = true // Base case: empty subset has sum 0

	// For each number, update all possible sums
	for _, currentNum := range nums {
		// Create slice of current possible sums to avoid modifying map while iterating
		existingSums := make([]int, 0, len(possibleSums))
		for sum := range possibleSums {
			existingSums = append(existingSums, sum)
		}

		// Add new possible sums by including current number
		for _, existingSum := range existingSums {
			newSum := existingSum + currentNum
			possibleSums[newSum] = true
		}
	}

	// Check if target sum is achievable
	return possibleSums[targetSum]
}

// CanPartition3 - 2D Dynamic Programming approach
// Time: O(n * sum), Space: O(n * sum)
func CanPartition3(nums []int) bool {
	// Calculate total sum of all numbers
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is odd, cannot partition into two equal subsets
	if totalSum%2 != 0 {
		return false
	}

	targetSum := totalSum / 2

	// Create 2D DP table: dp[i][j] = can we achieve sum j using first i numbers
	numCount := len(nums)
	dp := make([][]bool, numCount+1)
	for i := range dp {
		dp[i] = make([]bool, targetSum+1)
	}

	// Base case: we can always achieve sum 0 with empty subset
	for numIndex := range dp {
		dp[numIndex][0] = true
	}

	// Fill the DP table
	for numIndex := 1; numIndex < len(dp); numIndex++ {
		currentNumber := nums[numIndex-1] // nums is 0-indexed, dp is 1-indexed

		for sum := 1; sum < len(dp[0]); sum++ {
			// Option 1: don't include current number
			dp[numIndex][sum] = dp[numIndex-1][sum]

			// Option 2: include current number (if possible)
			if sum >= currentNumber {
				canIncludeCurrent := dp[numIndex-1][sum-currentNumber]
				dp[numIndex][sum] = dp[numIndex][sum] || canIncludeCurrent
			}
		}
	}

	// Return result: can we achieve target sum using all numbers
	finalNumIndex := len(dp) - 1
	finalSum := len(dp[0]) - 1
	return dp[finalNumIndex][finalSum]
}
