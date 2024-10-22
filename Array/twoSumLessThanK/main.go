package main

import (
	"fmt"
	"sort"
)

// Brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSum(nums []int, k int) int {
	maxVal := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// Only update the maxVal if the sum is less than k
			if nums[i]+nums[j] < k {
				maxVal = max(maxVal, nums[i]+nums[j])
			}
		}
	}

	if maxVal == 0 {
		return -1
	}

	return maxVal
}

// Two pointers
// Time complexity: O(n*log(n))
// Space complexity: O(1)
func TwoSum1(nums []int, k int) int {
	// Sort the array in ascending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	left, right := 0, len(nums)-1
	maxVal := 0

	for left < right {
		sum := nums[left] + nums[right]

		// Only update the maxVal if the sum is less than k
		if sum < k {
			maxVal = max(maxVal, sum)
			left++
		} else {
			// If the sum is greater than or equal to k, move the right pointer to the left
			// So that the sum is decreased
			right--
		}
	}

	if maxVal == 0 {
		return -1
	}

	return maxVal
}

func main() {
	fmt.Println(TwoSum([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60))
	fmt.Println(TwoSum([]int{10, 20, 30}, 15))
}
