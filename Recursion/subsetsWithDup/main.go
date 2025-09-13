package main

import "slices"

// SubsetsWithDupBacktrack generates all unique subsets from an array with duplicates
// using backtracking with a visited array to track element usage.
//
// Approach: Sort the array first, then use backtracking. For each element, we have
// two choices: include it or skip it. To avoid duplicates, we enforce the rule that
// if we have consecutive identical elements, we can only choose a later one if we
// have chosen all previous ones.
//
// Time Complexity: O(2^n * n) where n is the length of nums
// - 2^n possible subsets, each subset takes O(n) time to copy
// Space Complexity: O(n) for recursion stack and temporary arrays
func SubsetsWithDupBacktrack(nums []int) [][]int {
	result := [][]int{}
	visited := make([]bool, len(nums))
	slices.Sort(nums) // Sort to group duplicates together

	generateSubsetsBacktrack(nums, 0, visited, &result, []int{})
	return result
}

func generateSubsetsBacktrack(nums []int, currentIndex int, visited []bool, result *[][]int, currentSubset []int) {
	// Base case: processed all elements, add current subset to result
	if currentIndex >= len(nums) {
		subset := make([]int, len(currentSubset))
		copy(subset, currentSubset)
		*result = append(*result, subset)
		return
	}

	// Choice 1: Skip the current element
	generateSubsetsBacktrack(nums, currentIndex+1, visited, result, currentSubset)

	// Choice 2: Include the current element (with duplicate handling)
	// If current element equals previous element and previous element is not visited,
	// we cannot choose current element to avoid duplicates
	if currentIndex > 0 && nums[currentIndex-1] == nums[currentIndex] && !visited[currentIndex-1] {
		return
	}

	// Mark current element as visited and include it
	visited[currentIndex] = true
	generateSubsetsBacktrack(nums, currentIndex+1, visited, result, append(currentSubset, nums[currentIndex]))
	visited[currentIndex] = false // Backtrack
}

// SubsetsWithDupSkipping generates all unique subsets from an array with duplicates
// using index skipping to handle duplicate elements.
//
// Approach: Sort the array first, then use recursion. For each element, we have
// two choices: include it or skip it. When we skip an element, we skip all
// consecutive duplicates of that element to avoid generating duplicate subsets.
//
// Time Complexity: O(2^n * n) where n is the length of nums
// - 2^n possible subsets, each subset takes O(n) time to copy
// Space Complexity: O(n) for recursion stack and temporary arrays
func SubsetsWithDupSkipping(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums) // Sort to group duplicates together

	generateSubsetsSkipping(nums, 0, &result, []int{})
	return result
}

func generateSubsetsSkipping(nums []int, currentIndex int, result *[][]int, currentSubset []int) {
	// Base case: processed all elements, add current subset to result
	if currentIndex >= len(nums) {
		subset := make([]int, len(currentSubset))
		copy(subset, currentSubset)
		*result = append(*result, subset)
		return
	}

	// Choice 1: Include the current element
	generateSubsetsSkipping(nums, currentIndex+1, result, append(currentSubset, nums[currentIndex]))

	// Choice 2: Skip the current element and all its consecutive duplicates
	// Find the next index that has a different value
	nextDifferentIndex := currentIndex + 1
	for nextDifferentIndex < len(nums) && nums[nextDifferentIndex] == nums[nextDifferentIndex-1] {
		nextDifferentIndex++
	}

	generateSubsetsSkipping(nums, nextDifferentIndex, result, currentSubset)
}
