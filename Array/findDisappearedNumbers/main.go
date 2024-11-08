package main

import "math"

// Brute Force
func FindDisappearedNumbers(nums []int) []int {
	ans := []int{}

	// Loop through the numbers from 1 to len(nums)
	for i := 0; i < len(nums); i++ {
		found := false
		// Check if the number is in nums
		for _, n := range nums {
			if n == i+1 {
				found = true
				break
			}
		}
		// If the number is not in nums, add it to the result
		if !found {
			ans = append(ans, i+1)
		}
	}

	return ans
}

// Hash Table
func FindDisappearedNumbers2(nums []int) []int {
	// Create a hash table to store the numbers that are in nums
	hashTable := make([]bool, len(nums))
	ans := []int{}

	// Iterate through nums and set the index of the number to true in the hash table
	for _, n := range nums {
		hashTable[n-1] = true
	}

	// Iterate through the hash table and add the index to the result if the value is false
	for i, b := range hashTable {
		if !b {
			ans = append(ans, i+1)
		}
	}

	return ans
}

// Sorting With Swapping
func FindDisappearedNumbers3(nums []int) []int {
	ans := []int{}
	// Sort the array with swapping
	for i := 0; i < len(nums); i++ {
		for i+1 != nums[i] && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// After sorting, the numbers that are not in their correct positions are the missing numbers
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

// Marking With Negatives
func FindDisappearedNumbers4(nums []int) []int {
	ans := []int{}
	// Mark the number as seen by setting the value at the index of the number to negative
	// Note that we need to use math.Abs to get the absolute value of the number
	// because the number at the index could be negative(already marked)
	// e.g. nums = [4,3,2,7,8,2,3,1], when i = 3, the nums becomes [4,3,2,-7,8,2,3,1]
	// can you see that value at index 3 is negative?
	// because we've already marked it as seen when i = 0
	// in order to get the correct index, we need to use math.Abs to get the absolute value of the number
	// then minus 1 because the index is 0-based to get the correct index
	for _, n := range nums {
		idx := int(math.Abs(float64(n))) - 1
		nums[idx] = int(math.Abs(float64(nums[idx]))) * -1
	}

	// Iterate through the array and add the index to the result if the value is positive
	for i, n := range nums {
		if n > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
