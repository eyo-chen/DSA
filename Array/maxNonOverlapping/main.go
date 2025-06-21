package main

func MaxNonOverlapping(nums []int, target int) int {
	// Initialize count of valid non-overlapping subarrays
	subarrayCount := 0

	// Iterate through possible starting indices
	for i := 0; i < len(nums); i++ {
		// Initialize hash map to store prefix sums and their indices
		// Key: prefix sum, Value: index where it occurs
		// Start with 0 at index -1 to handle subarrays starting at index 0
		prefixSums := map[int]int{0: -1}
		// Track current prefix sum for the segment
		currentSum := 0

		// Scan from i to find a subarray summing to target
		for k := i; k < len(nums); k++ {
			// Update prefix sum by adding current element
			currentSum += nums[k]

			// Check if there’s a prefix sum that makes a subarray sum to target
			if _, exists := prefixSums[currentSum-target]; exists {
				// Found a subarray from prevIdx+1 to k that sums to target
				subarrayCount++
				// Move i to k to start a new non-overlapping segment
				i = k
				break
			}

			// Store the current prefix sum with its index
			prefixSums[currentSum] = k
		}
	}

	// Return the total number of non-overlapping subarrays found
	return subarrayCount
}

func MaxNonOverlapping1(nums []int, target int) int {
	prefixMap := map[int]int{0: -1} // Maps prefix sum to earliest index
	currSum := 0
	count := 0
	lastEnd := -1 // Tracks the end index of the last selected subarray

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]

		// Check if there’s a prefix sum such that currSum - target exists
		if prevIndex, exists := prefixMap[currSum-target]; exists && prevIndex >= lastEnd {
			// Found a non-overlapping subarray ending at i
			count++
			lastEnd = i // Update the last used end index
		}

		// Store the current prefix sum with its index
		prefixMap[currSum] = i
	}

	return count
}
