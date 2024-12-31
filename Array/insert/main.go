package main

import "sort"

// Add new interval and sort by start time
func Insert(intervals [][]int, newInterval []int) [][]int {
	// Add new interval
	intervals = append(intervals, newInterval)

	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}
	// Initialize with first interval
	current := intervals[0]

	// Iterate through all intervals
	for i := 1; i < len(intervals); i++ {
		// If current interval overlaps with next interval
		if current[1] >= intervals[i][0] {
			// Merge intervals by taking min of starts and max of ends
			current[1] = max(current[1], intervals[i][1])
		} else {
			// No overlap, add current to result and move to next
			ans = append(ans, current)
			current = intervals[i]
		}
	}

	// Don't forget to add the last merged interval
	ans = append(ans, current)

	return ans
}

// Iterate through the intervals and merge if overlap
func Insert2(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}

	for i := 0; i < len(intervals); i++ {
		// newInterval is before the current interval
		if newInterval[1] < intervals[i][0] {
			ans = append(ans, newInterval)

			// Because we have already added the new interval to the result array,
			// we can simply append the rest of the intervals to the result array.
			ans = append(ans, intervals[i:]...)
			return ans
		}

		// newInterval is after the current interval
		if newInterval[0] > intervals[i][1] {
			ans = append(ans, intervals[i])
			continue
		}

		// newInterval is overlapping with the current interval
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}

	ans = append(ans, newInterval)

	return ans
}
