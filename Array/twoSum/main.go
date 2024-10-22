package main

// Brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Hash table
// Time complexity: O(n)
// Space complexity: O(n)
func TwoSum1(nums []int, target int) []int {
	hashTable := map[int]int{}

	for i, n := range nums {
		if v, ok := hashTable[n]; ok {
			return []int{i, v}
		}

		hashTable[target-n] = i
	}

	return []int{}
}
