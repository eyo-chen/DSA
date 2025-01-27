package main

import "sort"

func PermuteUnique(nums []int) [][]int {
	ans := [][]int{}
	hashTable := make([]bool, len(nums))

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	helper(nums, hashTable, &ans, []int{})
	return ans
}

func helper(nums []int, hashTable []bool, ans *[][]int, cur []int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if hashTable[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !hashTable[i-1] {
			continue
		}

		hashTable[i] = true
		helper(nums, hashTable, ans, append(cur, nums[i]))
		hashTable[i] = false
	}
}

func PermuteUnique1(nums []int) [][]int {
	ans := [][]int{}
	hashTable := make([]bool, len(nums))

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	helper1(nums, hashTable, &ans, []int{})
	return ans
}

func helper1(nums []int, hashTable []bool, ans *[][]int, cur []int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if hashTable[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !hashTable[i-1] {
			continue
		}

		hashTable[i] = true
		helper1(nums, hashTable, ans, append(cur, nums[i]))
		hashTable[i] = false
	}
}
