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

// Merge3 takes a slice of intervals and merges all overlapping intervals.
// It sorts the intervals by start time, then iterates through them, merging intervals
// that overlap (i.e., where the current interval's end >= next interval's start).
//
// Time Complexity: O(n log n) - dominated by the sorting operation
// Space Complexity: O(n) - for storing the result slice
//
// Example: [[1,3],[2,6],[8,10],[15,18]] -> [[1,6],[8,10],[15,18]]
func Merge3(intervals [][]int) [][]int {
	// Sort intervals by their start time in ascending order
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Initialize with the first interval as the current interval being built
	currentInterval := intervals[0]
	mergedIntervals := [][]int{}

	// Iterate through remaining intervals starting from index 1
	for i := 1; i < len(intervals); i++ {
		nextInterval := intervals[i]

		// If current interval doesn't overlap with next interval (gap between them)
		if currentInterval[1] < nextInterval[0] {
			// Add the completed current interval to results
			mergedIntervals = append(mergedIntervals, currentInterval)
			// Start a new current interval
			currentInterval = nextInterval
			continue
		}

		// Intervals overlap, so merge them by extending the end time
		// Keep the earlier start time and take the maximum end time
		currentInterval = []int{
			currentInterval[0],
			max(currentInterval[1], nextInterval[1]),
		}
	}

	// Don't forget to add the last interval being built
	mergedIntervals = append(mergedIntervals, currentInterval)

	return mergedIntervals
}
