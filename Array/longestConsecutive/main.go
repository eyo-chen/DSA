package main

import "slices"

// Bruth Force Solution
// Time Complexity O(n^3)
// The reason is because we have a for loop to iterate through each number, and for each number, we have a while loop to count the consecutive sequence.
// And the checking of the existence of the number in the array is O(n)
// e.g. [1,2,3,4]
// When we are at 1, we need to check if 2,3,4 exists.
// For each checking, we do O(n) work.
// Space Complexity O(1)
func LongestConsecutive(nums []int) int {
	maxLength := 0

	for _, num := range nums {
		currentLength := 0
		currentNum := num

		for slices.Contains(nums, currentNum) {
			currentNum++
			currentLength++
		}

		maxLength = max(maxLength, currentLength)
	}

	return maxLength
}

// LongestConsecutive2 finds the length of the longest consecutive sequence in an unsorted array
// by checking each number and its consecutive numbers in a forward direction
// Time Complexity O(n^2)
// Space Complexity O(n)
func LongestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Store all numbers in a hash set for O(1) lookup
	numSet := map[int]bool{}
	maxLength := 1

	for _, num := range nums {
		numSet[num] = true
	}

	for _, num := range nums {
		// Only process sequences that continue forward
		if numSet[num+1] {
			currentLength := 0
			currentNum := num

			// Count consecutive numbers
			for numSet[currentNum] {
				currentNum++
				currentLength++
			}

			maxLength = max(maxLength, currentLength)
		}
	}

	return maxLength
}

// Update at 2024-12-26
// It's the same as LongestConsecutive2, but it's more readable and easier to understand.
// We remove the if-statement in the for loop.
// Because if there's no num + i in the hash table, the for loop will not enter.
func LongestConsecutive2_1(nums []int) int {
	ans := 0
	hashTable := make(map[int]bool, len(nums))

	// Create hash table
	for _, n := range nums {
		hashTable[n] = true
	}

	// Find the longest consecutive sequence
	for i := 0; i < len(nums); i++ {
		curLen := 1
		curVal := nums[i] + 1

		for hashTable[curVal] {
			curLen++
			curVal++
		}

		ans = max(ans, curLen)
	}

	return ans
}

// LongestConsecutive3 finds the length of the longest consecutive sequence in an unsorted array
// by only starting the count from the smallest number in each sequence
// Time Complexity O(n)
// Space Complexity O(n)
func LongestConsecutive3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Store all numbers in a hash set for O(1) lookup
	numSet := map[int]bool{}
	maxLength := 0

	for _, num := range nums {
		numSet[num] = true
	}

	for _, num := range nums {
		// Skip if this number is not the start of a sequence
		// (i.e., if there exists a smaller number that precedes it)
		if numSet[num-1] {
			continue
		}

		sequenceLength := 0
		currentNum := num

		// Count consecutive numbers starting from the sequence beginning
		for numSet[currentNum] {
			sequenceLength++
			currentNum++
		}

		maxLength = max(maxLength, sequenceLength)
	}

	return maxLength
}
