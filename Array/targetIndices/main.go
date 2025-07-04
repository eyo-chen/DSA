package main

import "slices"

// Time: O(n*log(n))
func TargetIndices(nums []int, target int) []int {
	ans := []int{}

	slices.Sort(nums)
	for i, n := range nums {
		if n == target {
			ans = append(ans, i)
		}
	}

	return ans
}

// Time: O(n)
func TargetIndices1(nums []int, target int) []int {
	sameCount := 0
	lessCount := 0

	for _, n := range nums {
		if n == target {
			sameCount++
		}

		if n < target {
			lessCount++
		}
	}

	// The indices will be from smaller to smaller+count-1
	result := make([]int, sameCount)
	for i := 0; i < sameCount; i++ {
		result[i] = lessCount + i
	}

	return result
}
