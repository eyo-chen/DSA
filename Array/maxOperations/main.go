package main

import (
	"sort"
)

// Brute Force
func MaxOperations(nums []int, k int) int {
	hashTable := make([]bool, len(nums))
	ans := 0

	for i := 0; i < len(nums); i++ {
		if hashTable[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if hashTable[j] {
				continue
			}

			if nums[i]+nums[j] == k {
				ans++
				hashTable[i] = true
				hashTable[j] = true
				break
			}
		}
	}

	return ans
}

// Sorting and Two Pointers
func MaxOperations2(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	left, right := 0, len(nums)-1
	ans := 0
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			ans++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return ans
}

// Hash Table
func MaxOperations3(nums []int, k int) int {
	hashTable := map[int]int{}
	ans := 0

	for _, n := range nums {
		remaining := k - n
		if hashTable[remaining] > 0 {
			ans++
			hashTable[remaining]--
		} else {
			hashTable[n]++
		}
	}

	return ans
}
