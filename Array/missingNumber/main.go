package main

// Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func MissingNumber(nums []int) int {
	// Iterate through the range of numbers from 0 to len(nums)
	// Note that we are iterating from 0 to len(nums) inclusive
	// because we are looking for the missing number in the range of numbers from 0 to len(nums)
	for i := 0; i <= len(nums); i++ {
		found := false

		// For each number, we try to check if the number is in the array
		// If it is, we set found to true and break the loop
		// If it is not, we continue the loop
		for k := 0; k < len(nums); k++ {
			if nums[k] == i {
				found = true
				break
			}
		}

		// If the number is not found in the array, we return the number
		if !found {
			return i
		}
	}

	return -1
}

// Hash Table
// Time Complexity: O(n)
// Space Complexity: O(n)
func MissingNumber2(nums []int) int {
	// Create a hash table to store the numbers in the array
	// The hash table is initialized with a size of len(nums)+1
	// because we are looking for the missing number in the range of numbers from 0 to len(nums)
	hashTable := make([]bool, len(nums)+1)

	// Iterate through the array and mark the numbers in the hash table
	for _, n := range nums {
		hashTable[n] = true
	}

	// Iterate through the range of numbers from 0 to len(nums)
	// If the number is not found in the hash table, we return the number
	for i := 0; i <= len(nums); i++ {
		if !hashTable[i] {
			return i
		}
	}

	return -1
}

// Cyclic Sort
// Time Complexity: O(n)
// Space Complexity: O(1)
func MissingNumber3(nums []int) int {
	nums = append(nums, len(nums))

	// Iterate through the array and swap the numbers to their correct positions
	// For each iteration, all we wanna do is to make sure each index has the correct number
	// If the current index NOT have the correct number
	// AND
	// If the the index gonna swap(nums[nums[i]]) NOT have the correct number
	// We swap the numbers
	for i := 0; i < len(nums)-1; i++ {
		for i != nums[i] && nums[nums[i]] != nums[i] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	// Iterate through the array and check if the number is in the correct position
	// If it is not, we return the number
	for i := 0; i < len(nums)-1; i++ {
		if i != nums[i] {
			return i
		}
	}

	return len(nums) - 1
}
