package main

import (
	"sort"
)

func CanAttendMeetings(intervals [][]int) bool {
	// Sort the intervals by the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Check if there's any overlap
	for i := 1; i < len(intervals); i++ {
		// If the end time of the previous meeting is greater than the start time of the current meeting, then there's an overlap
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}

	return true
}

/*
return false if (second value of first interval > first value of second interval)
*/

/*
0       30
   5 10

   2   4
          7    10
*/

// This is my initial solution
// However, it's not working as expected
// The logic of this solution is to find the minimum and maximum value of the intervals
// Then, we check if
// (1) the start time of the current meeting is greater than the maximum value
// (2) the end time of the current meeting is less than the minimum value
// If either of these conditions is true, then there's NO overlap, and we continue
// Otherwise, we return false

// This is might seem like a good solution, but it's not
// The main reason is not working is because we CAN'T assume that the meetings are sorted by the start time
// AND we CAN'T update the min and max value when we find an overlap
// e.g. [[17 19] [8 14] [15 17]]
//
//	                       17  19
//	8       14
//	             15        17
//
// |--------------------------------|
// As we can see, there should be NO overlap, but the solution will return false
// Let's see why
// First, min = 17, max = 19
// Then, we're at [8 14]. Because interval[1] <= min, so min = interval[0]
// min = 8, max = 19
// Now, the new interval becomes [8, 19]
// So, the next interval [15 17] will be compared with [8, 19], which will lead to false
func CanAttendMeetings2(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	minVal, maxVal := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] >= maxVal {
			maxVal = interval[1]
			continue
		}
		if interval[1] <= minVal {
			minVal = interval[0]
			continue
		}

		return false
	}

	return true
}
