package main

import (
	"math"
)

// Bruth Force Solution
func FindMaxLength(nums []int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		hashTable := make([]int, 2)

		for k := i; k < len(nums); k++ {
			hashTable[nums[k]]++

			if hashTable[0] == hashTable[1] {
				ans = max(ans, k-i+1)
			}
		}
	}

	return ans
}

func FindMaxLength1(nums []int) int {
	// Map to store the first occurrence of each prefix sum
	// Key: prefix sum, Value: index where this sum first occurred
	prefixSumByIndex := make(map[int]int)

	// Initialize with 0 sum at index -1 to handle cases where
	// the subarray starts from the beginning of the array
	prefixSumByIndex[0] = -1

	maxLength := 0
	runningSum := 0

	// Traverse the array and calculate prefix sums
	for i, num := range nums {
		// Transform 0 to -1, keep 1 as 1
		// This allows us to find subarrays that sum to 0
		if num == 0 {
			runningSum -= 1
		} else {
			runningSum += 1
		}

		// Check if we've seen this prefix sum before
		if firstIndex, found := prefixSumByIndex[runningSum]; found {
			// If we've seen this sum before, the subarray between
			// firstIndex+1 and current index i has sum = 0
			// This means equal number of 0s and 1s
			currentLength := i - firstIndex
			maxLength = int(math.Max(float64(maxLength), float64(currentLength)))
		} else {
			// First time seeing this prefix sum, record its index
			prefixSumByIndex[runningSum] = i
		}
	}

	return maxLength
}
