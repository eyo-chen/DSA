package main

import "slices"

// Brute force
func ContainsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// Sorting
func ContainsDuplicate2(nums []int) bool {
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

// Hash table
func ContainsDuplicate3(nums []int) bool {
	hashTable := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := hashTable[nums[i]]; ok {
			return true
		}

		hashTable[nums[i]] = true
	}

	return false
}
