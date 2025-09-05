package main

// subsets generates all possible subsets of the given integer array using backtracking.
// The approach uses recursion where at each element, we make two choices:
// 1. Include the current element in the subset
// 2. Exclude the current element from the subset
// This creates a binary decision tree that explores all 2^n possibilities.
func Subsets(nums []int) [][]int {
	result := [][]int{}
	generateSubsets(nums, 0, &result, []int{})
	return result
}

// generateSubsets is a recursive helper function that builds all subsets using backtracking
func generateSubsets(nums []int, currentIndex int, result *[][]int, currentSubset []int) {
	// Base case: when we've processed all elements, add the current subset to results
	if currentIndex == len(nums) {
		// Create a deep copy of currentSubset to avoid reference issues
		subsetCopy := make([]int, len(currentSubset))
		copy(subsetCopy, currentSubset)
		*result = append(*result, subsetCopy)
		return
	}

	// Choice 1: Exclude the current element and move to the next
	generateSubsets(nums, currentIndex+1, result, currentSubset)

	// Choice 2: Include the current element and move to the next
	generateSubsets(nums, currentIndex+1, result, append(currentSubset, nums[currentIndex]))
}
