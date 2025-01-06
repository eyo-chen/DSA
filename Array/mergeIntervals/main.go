package main

import (
	"sort"
)

// First Solution
// I came up on my own
func Merge(intervals [][]int) [][]int {
	ans := [][]int{}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Initialize the first interval
	prevInterval := intervals[0]

	// Iterate through the intervals
	for i := 1; i < len(intervals); i++ {
		// Check if the current interval overlaps with the previous one
		if prevInterval[1] >= intervals[i][0] {
			// Merge the intervals
			prevInterval[1] = max(prevInterval[1], intervals[i][1])
			continue
		}

		// When two intervals are not overlapping, add the previous interval to the result

		// Add the previous interval to the result
		ans = append(ans, prevInterval)

		// Update the previous interval to the current one
		prevInterval = intervals[i]
	}

	ans = append(ans, prevInterval)
	return ans
}

// Second Solution
// AI solution
func Merge2(intervals [][]int) [][]int {
	// 1. Edge case check
	if len(intervals) <= 1 {
		return intervals
	}

	// 2. Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 3. Initialize result with the first interval
	result := [][]int{intervals[0]}

	// 4. Iterate through remaining intervals
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := result[len(result)-1]

		// If current interval overlaps with last merged interval
		if current[0] <= lastMerged[1] {
			// Update the end time if current interval extends further
			lastMerged[1] = max(lastMerged[1], current[1])
		} else {
			// No overlap, add current interval to result
			result = append(result, current)
		}
	}

	return result
}
