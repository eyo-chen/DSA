package main

import "slices"

// Brute Force
// (1) Loop through nums1 and check if the number is in nums2 or ans[0]
// (2) Loop through nums2 and check if the number is in nums1 or ans[1]
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func FindDifference(nums1 []int, nums2 []int) [][]int {
	ans := make([][]int, 2)

	for _, n := range nums1 {
		if slices.Contains(nums2, n) || slices.Contains(ans[0], n) {
			continue
		}
		ans[0] = append(ans[0], n)
	}

	for _, n := range nums2 {
		if slices.Contains(nums1, n) || slices.Contains(ans[1], n) {
			continue
		}
		ans[1] = append(ans[1], n)
	}

	return ans
}

// Hash Table
// (1) Create a hash table for nums1 and nums2
// (2) Loop through hash table1 and check if the number is in hash table2
// (3) Loop through hash table2 and check if the number is in hash table1
// Note that hash table also helps us to avoid duplicates
// Time Complexity: O(n)
// Space Complexity: O(n)
func FindDifference2(nums1 []int, nums2 []int) [][]int {
	ans := make([][]int, 2)
	hashTable1 := map[int]bool{}
	hashTable2 := map[int]bool{}

	for _, n := range nums1 {
		hashTable1[n] = true
	}
	for _, n := range nums2 {
		hashTable2[n] = true
	}

	for key := range hashTable1 {
		if !hashTable2[key] {
			ans[0] = append(ans[0], key)
		}
	}

	for key := range hashTable2 {
		if !hashTable1[key] {
			ans[1] = append(ans[1], key)
		}
	}

	return ans
}
