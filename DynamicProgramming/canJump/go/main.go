package main

func CanJump(nums []int) bool {
	memo := make(map[int]bool)
	return helper(nums, 0, memo)
}

func helper(nums []int, index int, memo map[int]bool) bool {
	// we are at the last index
	if index == len(nums)-1 {
		return true
	}

	// there is 0 jump, so we cannot move
	if nums[index] == 0 {
		return false
	}

	// memoization
	if v, ok := memo[index]; ok {
		return v
	}

	steps := nums[index]
	// we can jump from 1 to steps AND we cannot jump out of the array
	for i := 1; i <= steps && index+i < len(nums); i++ {
		if helper(nums, index+i, memo) {
			memo[index] = true
			return true
		}
	}

	memo[index] = false
	return false
}

func CanJump1(nums []int) bool {
	table := make([]bool, len(nums))
	table[0] = true

	for index := range nums {
		steps := nums[index]

		if !table[index] || steps <= 0 {
			continue
		}

		for i := 1; i <= steps && index+i < len(nums); i++ {
			table[index+i] = true
		}
	}

	return table[len(nums)-1]
}

func CanJump2(nums []int) bool {
	end := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		// we can jump from i to end
		// so we update the end to i
		if nums[i]+i >= end {
			end = i
		}
	}

	return end == 0
}
