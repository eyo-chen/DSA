package main

// Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func ContainsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && j-i <= k {
				return true
			}
		}
	}

	return false
}

// Hash Table
// Time Complexity: O(n)
// Space Complexity: O(n)
func ContainsNearbyDuplicate2(nums []int, k int) bool {
	hashTable := map[int]int{}

	for i, n := range nums {
		if idx, ok := hashTable[n]; ok && i-idx <= k {
			return true
		}

		hashTable[n] = i
	}

	return false
}
