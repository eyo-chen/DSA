package main

import "math"

type node struct {
	val     int
	listIdx int
	idx     int
}

func SmallestRange(nums [][]int) []int {
	globalMax := math.MinInt

	heap := NewMinHeapNode(len(nums))
	for i, n := range nums {
		heap.Insert(&node{val: n[0], listIdx: i, idx: 0})
		globalMax = max(globalMax, n[0])
	}

	start, end := 0, math.MaxInt32
	for heap.Len() == len(nums) {
		minVal := heap.Pull()
		if globalMax-minVal.val < end-start {
			start, end = minVal.val, globalMax
		}

		if minVal.idx < len(nums[minVal.listIdx])-1 {
			nextVal := nums[minVal.listIdx][minVal.idx+1]
			heap.Insert(&node{val: nextVal, listIdx: minVal.listIdx, idx: minVal.idx + 1})
			globalMax = max(globalMax, nextVal)
		}
	}

	return []int{start, end}
}
