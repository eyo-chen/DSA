package main

import (
	"math"
	"sort"
)

// Brute Force O(n^2)
// The reason this works is because we are guaranteed that each integer only appears once or twice
// For each element, we are asking "From your position to the end of the array, is there any other duplicate?"
// If we add one duplicate to the answer, it means we've found the duplicate, and the later value won't have any duplicate because each integer only appears once or twice
// Also, we won't add the same value twice to the answer
// e.g. [4,3,2,7,8,2,3,1], when we hit first 3, we add it to the answer, and the second 3 will not be added to the answer
// because when we hit the second 3, there's no way there's another 3 in the rest of the array
func FindDuplicates(nums []int) []int {
	ans := []int{}

	for i := 0; i < len(nums); i++ {
		// Start from the next element of i
		for j := i + 1; j < len(nums); j++ {
			// Found the duplicate
			// Add it to the answer and break the loop early
			if nums[i] == nums[j] {
				ans = append(ans, nums[i])
				break
			}
		}
	}

	return ans
}

// Sorting O(n*logn)
// The reason this works is similar to the brute force solution
// e.g. [4,3,2,7,8,2,3,1] -> [1,2,2,3,3,4,7,8]
// Here, we start from the second element and compare it with the previous one
// If they are the same, it means we've found a duplicate, and we add it to the answer
// Otherwise, we move on to the next element
// In this manner, we can make sure we won't add the same value twice to the answer
// because each integer only appears once or twice
func FindDuplicates2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ans = append(ans, nums[i])
		}
	}

	return ans
}

// Hash Table O(n), but space complexity is O(n)
// The reason this works is also similar to the brute force solution
// We use a hash table to store the value we've seen
// For each value, we check if it's in the hash table
// If it is, we've found a duplicate, and we add it to the answer
// Otherwise, we add it to the hash table
// We won't add the same value twice to the answer
// because each integer only appears once or twice
func FindDuplicates3(nums []int) []int {
	hashTable := map[int]bool{}
	ans := []int{}

	for _, n := range nums {
		if _, ok := hashTable[n]; ok {
			ans = append(ans, n)
		}

		hashTable[n] = true
	}

	return ans
}

// Marking Visited Numbers O(n)
func FindDuplicates4(nums []int) []int {
	ans := []int{}

	for _, val := range nums {
		// get the absolute value of the number
		v := int(math.Abs(float64(val)))

		// get the index of the number
		// e.g. [4,3,2,7,8,2,3,1]
		// when we hit 4, idx = 4 - 1 = 3
		idx := v - 1

		// if the number is negative, it means we have seen it before
		// so we add it to the answer
		if nums[idx] < 0 {
			ans = append(ans, v)
		} else {
			// otherwise, we mark the number as visited by multiplying it by -1
			nums[idx] *= -1
		}
	}

	return ans
}

func FindDuplicates5(nums []int) []int {
	ans := []int{}

	// Sorting with swapping
	for i := 0; i < len(nums); i++ {
		// nums[i] != i+1 means the number is not in its correct position
		// e.g. [3,2,3,4,8,2,7,1], when i = 1, nums[i] = 2
		// value 2 is at the correct position(idx = 1)
		// nums[i] != nums[nums[i]-1] means the value we're gonna swap to is already have the correct number
		// e.g. [3,2,3,4,8,2,7,1], when i = 0, nums[i] = 3
		// value 3 is not in the correct position(idx = 2)
		// However, idx 2 already has the correct number(value = 3)
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// After sorting, the numbers that are not in their correct positions are the duplicates
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, nums[i])
		}
	}

	return ans
}
