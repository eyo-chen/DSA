package main

import "sort"

func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	count := len(points)
	prev := points[0]

	// start at 2nd index because we use the first index as the initial prev
	for i := 1; i < len(points); i++ {
		cur := points[i]

		// not overlapping
		// prev[1] -> end of the previous interval
		// cur[0]  -> start of the current interval
		if prev[1] < cur[0] {
			prev = cur
			continue
		}

		// overlapping
		count--

		// merge the intervals
		newPrevStart := cur[0]
		newPrevEnd := cur[1]
		if prev[1] < newPrevEnd {
			newPrevEnd = prev[1]
		}
		prev = []int{newPrevStart, newPrevEnd}
	}

	return count
}
