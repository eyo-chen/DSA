package main

// countHillValley counts the number of hills and valleys in the given array
// A hill is an index where the value is greater than both its closest non-equal neighbors
// A valley is an index where the value is smaller than both its closest non-equal neighbors
// Adjacent indices with the same value are considered part of the same hill/valley
func CountHillValley(nums []int) int {
	hillValleyCount := 0

	// Iterate through each possible position (excluding first and last elements)
	// First and last elements cannot be hills/valleys since they need neighbors on both sides
	for currentIdx := 1; currentIdx < len(nums)-1; currentIdx++ {

		// Step 1: Find the closest non-equal neighbor to the left
		leftNeighborIdx := currentIdx - 1
		// Keep moving left until we find a value different from current value
		for leftNeighborIdx >= 0 && nums[leftNeighborIdx] == nums[currentIdx] {
			leftNeighborIdx--
		}
		// If we reached the beginning without finding a different value, skip this position
		if leftNeighborIdx < 0 {
			continue
		}

		// Step 2: Find the closest non-equal neighbor to the right
		rightNeighborIdx := currentIdx + 1
		// Keep moving right until we find a value different from current value
		for rightNeighborIdx < len(nums) && nums[rightNeighborIdx] == nums[currentIdx] {
			rightNeighborIdx++
		}
		// If we reached the end without finding a different value, skip this position
		if rightNeighborIdx >= len(nums) {
			continue
		}

		// Step 3: Check if current position forms a hill or valley
		currentValue := nums[currentIdx]
		leftNeighborValue := nums[leftNeighborIdx]
		rightNeighborValue := nums[rightNeighborIdx]

		// Check if it's a hill: current value is greater than both neighbors
		if currentValue > leftNeighborValue && currentValue > rightNeighborValue {
			hillValleyCount++
		}
		// Check if it's a valley: current value is smaller than both neighbors
		if currentValue < leftNeighborValue && currentValue < rightNeighborValue {
			hillValleyCount++
		}

		// Step 4: Skip all consecutive equal elements to avoid counting the same hill/valley multiple times
		// This is the key optimization: since adjacent equal values are part of the same hill/valley,
		// we only need to count once per group
		for currentIdx < len(nums)-1 && nums[currentIdx] == nums[currentIdx+1] {
			currentIdx++
		}
		// Note: currentIdx will be incremented again by the for loop, so we'll start checking
		// the next different value in the next iteration
	}

	return hillValleyCount
}
