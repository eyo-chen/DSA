package main

import (
	"slices"
)

// Brute Force
// Time Complexity: O(n^3)
// Space Complexity: O(1)
func ThreeSum(nums []int) [][]int {
	ans := [][]int{}

	slices.Sort(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for k := i + 1; k < len(nums); k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}

			for j := k + 1; j < len(nums); j++ {
				if j > k+1 && nums[j] == nums[j-1] {
					continue
				}

				if nums[i]+nums[k]+nums[j] == 0 {
					ans = append(ans, []int{nums[i], nums[k], nums[j]})
				}
			}
		}
	}

	return ans
}

func ThreeSum2(nums []int) [][]int {
	// Sort the array in ascending order
	slices.Sort(nums)

	ans := [][]int{}

	for i := range nums {
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

// ThreeSum3 finds all unique triplets in the array that sum to zero.
//
// Approach:
// 1. Sort the array to enable two-pointer technique and easy duplicate handling
// 2. For each element, use two pointers (left and right) to find pairs that sum to -element
// 3. Skip duplicate elements to ensure unique triplets
//
// Time Complexity: O(n²) - O(n log n) for sorting + O(n²) for nested loops
// Space Complexity: O(1) - excluding the output array, only constant extra space is used
func ThreeSum3(nums []int) [][]int {
	result := [][]int{}

	// Sort array to enable two-pointer approach and handle duplicates
	slices.Sort(nums)

	// Iterate through each element as the first element of the triplet
	for firstIdx := 0; firstIdx < len(nums); firstIdx++ {
		// Use two pointers to find pairs that sum with current element to zero
		leftPtr, rightPtr := firstIdx+1, len(nums)-1

		for leftPtr < rightPtr {
			tripletSum := nums[firstIdx] + nums[leftPtr] + nums[rightPtr]

			if tripletSum == 0 {
				// Found a valid triplet
				result = append(result, []int{nums[firstIdx], nums[leftPtr], nums[rightPtr]})

				// Skip duplicate values for left pointer
				for leftPtr < rightPtr && nums[leftPtr] == nums[leftPtr+1] {
					leftPtr++
				}
				leftPtr++

				// Skip duplicate values for right pointer
				for leftPtr < rightPtr && nums[rightPtr] == nums[rightPtr-1] {
					rightPtr--
				}
				rightPtr--
			} else if tripletSum < 0 {
				// Sum too small, move left pointer right to increase sum
				leftPtr++
			} else {
				// Sum too large, move right pointer left to decrease sum
				rightPtr--
			}
		}

		// Skip duplicate values for the first element to avoid duplicate triplets
		// This can be done at the beginning or end of iteration (logic is equivalent)
		for firstIdx+1 < len(nums) && nums[firstIdx] == nums[firstIdx+1] {
			firstIdx++
		}
	}

	return result
}
