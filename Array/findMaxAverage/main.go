package main

import (
	"math"
)

// Brute Force
// Time Complexity: O(n * k)
// Space Complexity: O(1)
func FindMaxAverage(nums []int, k int) float64 {
	ans := -math.MaxFloat64

	// Iterate through the array to find each subarray of length k
	for i := 0; i <= len(nums)-k; i++ {
		sum := 0

		// For each element, sum up the elements in the subarray(loop through k elements)
		for j := 0; j < k && i+j < len(nums); j++ {
			sum += nums[i+j]
		}

		ans = max(ans, float64(sum)/float64(k))
	}

	return ans
}

// Sliding Window
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMaxAverage1(nums []int, k int) float64 {
	ans := -math.MaxFloat64
	sum := 0

	// Calculate initial window sum
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans = float64(sum) / float64(k)

	// Slide the window
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		ans = max(ans, float64(sum)/float64(k))
	}

	return ans
}
