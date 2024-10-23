package main

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	// Sort the array in ascending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := [][]int{}

	for i := 0; i < len(nums); i++ {
		// Keep updating i until the value is different
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// If the current value is greater than 0, then we can break
		// Because the array is sorted in ascending order, and we are using 3 positive numbers to sum up to 0
		// If the current value is greater than 0, then the rest of the values must be greater than 0
		// So, we can break
		if nums[i] > 0 {
			break
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// If the sum is 0, then we find one answer
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				// Move left and right pointer until the value is different
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
			}

			// If the sum is less than 0, then we need to increase the sum
			// So, we move the left pointer to the right
			if sum < 0 {
				left++
			} else {
				// If the sum is greater than 0, then we need to decrease the sum
				// So, we move the right pointer to the left
				right--
			}
		}
	}

	return ans
}
