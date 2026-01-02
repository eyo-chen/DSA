package main

import "slices"

/*
Given an array of n integers nums and a target, find the number of index triplets i, j, k with 0 <= i < j < k < n that satisfy the condition nums[i] + nums[j] + nums[k] < target.

Example:

Input: nums = [-2,0,1,3], and target = 2
Output: 2
Explanation: Because there are two triplets which sums are less than 2:
             [-2,0,1]
             [-2,0,3]
*/

// ThreeSumSmaller counts the number of triplets (i, j, k) where i < j < k
// and nums[i] + nums[j] + nums[k] < target using a brute force approach.
//
// Approach: Check all possible triplet combinations using three nested loops.
// Time Complexity: O(n³) - three nested loops iterate through all combinations
// Space Complexity: O(1) - only uses a constant amount of extra space
func ThreeSumSmaller(nums []int, target int) int {
	count := 0

	// Iterate through all possible triplets (i, j, k) where i < j < k
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				// Check if the sum of the current triplet is less than target
				if nums[i]+nums[j]+nums[k] < target {
					count++
				}
			}
		}
	}

	return count
}

// ThreeSumSmaller1 counts the number of triplets (i, j, k) where i < j < k
// and nums[i] + nums[j] + nums[k] < target using a two-pointer approach.
//
// Approach: Sort the array first, then for each element, use two pointers (left and right)
// to find valid pairs. When sum < target, all elements between left and right form valid
// triplets with the current element, so we add (right - left) to the count.
// Time Complexity: O(n²) - O(n log n) for sorting + O(n²) for two nested loops
// Space Complexity: O(1) - sorting is done in-place (or O(n) depending on sort implementation)
func ThreeSumSmaller1(nums []int, target int) int {
	count := 0

	// Sort the array to enable two-pointer technique
	slices.Sort(nums)

	// Fix the first element of the triplet
	for first := 0; first < len(nums)-2; first++ {
		// Use two pointers to find valid pairs for the remaining elements
		left, right := first+1, len(nums)-1

		for left < right {
			currentSum := nums[first] + nums[left] + nums[right]

			if currentSum < target {
				// If sum < target, all elements between left and right are valid
				// because the array is sorted. For example, if nums[first] + nums[left] + nums[right] < target,
				// then nums[first] + nums[left] + nums[right-1], nums[first] + nums[left] + nums[right-2], etc.
				// are also valid triplets.
				count += right - left
				left++
			} else {
				// If sum >= target, we need a smaller sum, so move right pointer left
				right--
			}
		}
	}

	return count
}
