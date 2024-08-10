package main

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := 0
	prevEnd := intervals[0][1]

	for _, v := range intervals[1:] {
		start, end := v[0], v[1]
		if start >= prevEnd {
			prevEnd = end
			continue
		}

		res++
		if end < prevEnd {
			prevEnd = end
		}
	}

	return res
}
