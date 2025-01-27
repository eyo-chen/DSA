package main

import (
	"math"
	"sort"
)

// Time Complexity O(n^3)
// Space Complexity O(1)
func ThreeSumClosest(nums []int, target int) int {
	ans := math.MaxInt
	closest := math.MaxInt

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				// calculate the sum of the triplet
				sum := nums[i] + nums[j] + nums[k]
				// calculate the difference between the sum and the target
				diff := int(math.Abs(float64(sum - target)))
				// if the difference is smaller than the smallest difference we have found so far
				// update the smallest difference and the closest sum
				if diff < closest {
					closest = diff
					ans = sum
				}
			}
		}
	}

	return ans
}

// Time Complexity O(n^2)
// Space Complexity O(1)
func ThreeSumClosest1(nums []int, target int) int {
	// sum is the closest sum to target, which is the answer
	sum := 0
	// closest is the smallest difference between sum and target
	// technically, we don't need this one, but it can make the code more readable
	closest := math.MaxInt

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// i only need to iterate through the array until the third last element
	// because we need at least 3 numbers to form a triplet
	// e.g. [1, 2, 3, 4, 5], we only need to iterate through [1, 2, 3]
	// don't need to iterate through 4 and 5 because we need to form a triplet
	for i := 0; i < len(nums)-2; i++ {
		// keep increment i until the value is different
		// avoid duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			// calculate the current sum of the triplet
			curSum := nums[left] + nums[right] + nums[i]
			// calculate the difference between the current sum and the target
			diff := int(math.Abs(float64(target - curSum)))

			// if the difference is smaller than the smallest difference we have found so far
			// update the smallest difference and the closest sum
			if diff < closest {
				closest = diff
				sum = curSum
			}

			// if the target is less than the current sum
			// move the right pointer to the left
			if target < curSum {
				// this is to avoid duplicates
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if target > curSum {
				// if the target is greater than the current sum
				// move the left pointer to the right

				// this is to avoid duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				// if the target is equal to the current sum
				// we have found the closest sum, return it
				return sum
			}
		}
	}

	return sum
}
