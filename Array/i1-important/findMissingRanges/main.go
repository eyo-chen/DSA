package main

import (
	"math"
)

// FindMissingRanges finds all missing ranges between lower and upper that are not in nums.
// Approach: Use a hash table to track present numbers, then iterate through [lower, upper]
// to identify missing ranges by tracking the start of each gap.
// Time Complexity: O(n + (upper - lower)) where n is the length of nums
// Space Complexity: O(n) for the hash table
func FindMissingRanges(nums []int, lower int, upper int) [][]int {
	missingRanges := [][]int{}

	// Track the start of the current missing range
	// Use sentinel value to indicate no range is currently being tracked
	rangeStart := math.MinInt

	// Build a hash table for O(1) lookup of numbers present in the array
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	// Scan through the entire range [lower, upper] to find missing numbers
	for current := lower; current <= upper; current++ {
		// If current number exists in the array
		if numSet[current] {
			// If we were tracking a missing range, close it
			if rangeStart != math.MinInt {
				missingRanges = append(missingRanges, []int{rangeStart, current - 1})
				rangeStart = math.MinInt // Reset to indicate no active range
			}
			continue
		}

		// Current number is missing - start tracking a new range if not already tracking
		if rangeStart == math.MinInt {
			rangeStart = current
		}
	}

	// If we're still tracking a missing range at the end, close it with upper bound
	if rangeStart != math.MinInt {
		missingRanges = append(missingRanges, []int{rangeStart, upper})
	}

	return missingRanges
}

// FindMissingRanges1 finds all missing ranges between lower and upper that are not in nums.
// Approach: Leverage the sorted property of nums to check gaps between consecutive elements,
// as well as gaps before the first element and after the last element.
// Assumption: nums is sorted in ascending order
// Time Complexity: O(n) where n is the length of nums
// Space Complexity: O(1) excluding the output array
func FindMissingRanges1(nums []int, lower int, upper int) [][]int {
	missingRanges := [][]int{}

	// Helper function to add a missing range [start, end] to the result
	addMissingRange := func(start, end int) {
		if start <= end {
			missingRanges = append(missingRanges, []int{start, end})
		}
	}

	// Edge case: if array is empty, the entire range [lower, upper] is missing
	if len(nums) == 0 {
		addMissingRange(lower, upper)
		return missingRanges
	}

	// Check if there's a missing range before the first element
	if lower < nums[0] {
		addMissingRange(lower, nums[0]-1)
	}

	// Check for missing ranges between consecutive elements in the array
	for i := 0; i < len(nums)-1; i++ {
		// If there's a gap of more than 1 between consecutive numbers
		if nums[i+1]-nums[i] > 1 {
			addMissingRange(nums[i]+1, nums[i+1]-1)
		}
	}

	// Check if there's a missing range after the last element
	if nums[len(nums)-1] < upper {
		addMissingRange(nums[len(nums)-1]+1, upper)
	}

	return missingRanges
}
