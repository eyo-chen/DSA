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
