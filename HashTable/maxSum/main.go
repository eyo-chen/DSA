package main

func MaxSum(nums []int) int {
	// Use a set to track unique positive numbers
	// We only care about positive numbers since we want to maximize the sum
	positiveUniques := map[int]bool{}

	// Track the maximum number in case all numbers are non-positive
	maxElement := nums[0]

	// First pass: collect all unique positive numbers and find the maximum element
	for _, num := range nums {
		// Only add positive numbers to our set since they contribute positively to the sum
		if num > 0 {
			positiveUniques[num] = true
		}

		// Keep track of the maximum element for the edge case where all numbers are <= 0
		maxElement = max(maxElement, num)
	}

	// Edge case: if there are no positive numbers, we must select the largest element
	// (since we cannot make the array empty after deletions)
	if len(positiveUniques) == 0 {
		return maxElement
	}

	// Calculate the sum of all unique positive numbers
	// This is optimal because:
	// 1. We can delete all negative and zero elements
	// 2. We can delete duplicate positive elements (keeping one of each)
	// 3. The remaining unique positive elements form a valid subarray after deletions
	totalSum := 0
	for uniquePositive := range positiveUniques {
		totalSum += uniquePositive
	}

	return totalSum
}
