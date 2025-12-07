package main

// Brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}

// Two pointers
// Time complexity: O(n)
// Space complexity: O(1)
func TwoSum1(numbers []int, target int) []int {
	// Initialize two pointers
	left, right := 0, len(numbers)-1

	// Iterate until the two pointers meet or cross each other
	for left < right {
		sum := numbers[left] + numbers[right]

		// If the sum of the two numbers is equal to the target, return the indices
		if sum == target {
			return []int{left + 1, right + 1}
		}

		// If the sum is greater than the target
		// It means we need to move the right pointer to the left to decrease the sum
		// Because the array is sorted in non-decreasing order
		// If we move the left pointer to the right, the sum will be greater
		if sum > target {
			right--
		} else {
			// If the sum is less than the target
			// It means we need to move the left pointer to the right to increase the sum
			left++
		}
	}

	return []int{}
}

// TwoSum2 returns the 1-indexed positions of two numbers in the array
// that add up to the target value using a hash map for O(1) lookups.
//
// Approach: Single-pass hash map - store each number with its index as we iterate,
// checking if the complement (target - current number) already exists in the map.
//
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(n) - hash map stores at most n elements
func TwoSum2(numbers []int, target int) []int {
	// Map to store number -> original index mapping for quick complement lookups
	numToIndex := map[int]int{}

	for currentIndex, currentNum := range numbers {
		// Calculate the complement needed to reach target
		complement := target - currentNum

		// Check if complement exists in our map (found a pair)
		if complementIndex, exists := numToIndex[complement]; exists {
			// Return 1-indexed positions (problem requirement)
			return []int{complementIndex + 1, currentIndex + 1}
		}

		// Store current number and its index for future complement checks
		numToIndex[currentNum] = currentIndex
	}

	// No valid pair found (should not happen per problem constraints)
	return []int{}
}
