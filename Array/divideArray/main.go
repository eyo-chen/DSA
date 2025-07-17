package main

import "slices"

func DivideArray(nums []int, k int) [][]int {
	ans := [][]int{}
	slices.Sort(nums)

	for i := 0; i <= len(nums)-3; i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}

		ans = append(ans, nums[i:i+3])
	}

	return ans
}
