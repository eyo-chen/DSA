package main

import "math"

// IncreasingTriplet - Brute Force approach
// Time complexity: O(n³)
// Checks every possible combination of three indices (i,j,k)
func IncreasingTriplet(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				// Check if we found an increasing sequence
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}

// IncreasingTriplet2 - Two-pointer approach
// Time complexity: O(n²)
// For each middle element, check if there exists smaller element before it
// and larger element after it
func IncreasingTriplet2(nums []int) bool {
	for i := 1; i < len(nums)-1; i++ {
		// Check if there's any smaller number before current position
		hasLower := false
		for k := 0; k < i; k++ {
			if nums[k] < nums[i] {
				hasLower = true
			}
		}

		// Check if there's any larger number after current position
		hasHigher := false
		for k := i + 1; k < len(nums); k++ {
			if nums[i] < nums[k] {
				hasHigher = true
			}
		}

		// If we found both smaller and larger numbers, we have a valid triplet
		if hasLower && hasHigher {
			return true
		}
	}
	return false
}

// IncreasingTriplet3 - Optimal solution
// Time complexity: O(n)
// Maintains two values: smallest number (i) and second smallest number (j)
func IncreasingTriplet3(nums []int) bool {
	// Initialize i and j to maximum possible value
	// i represents the smallest number seen so far
	// j represents the second smallest number seen so far
	i, j := math.MaxInt, math.MaxInt

	for _, n := range nums {
		if n <= i {
			// Update the smallest number seen so far
			i = n
		} else if n <= j {
			// Update the second smallest number
			// Note: this only happens when n > i, maintaining i < j
			j = n
		} else {
			// Found a number larger than both i and j
			// This means we found a valid triplet because:
			// 1. j was only updated when we found a number > i
			// 2. n is greater than both i and j
			return true
		}
	}
	return false
}
