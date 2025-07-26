package main

func Rob(nums []int) int {
	// Edge case: if no houses, return 0
	if len(nums) == 0 {
		return 0
	}

	// Create DP array with 2 extra spaces at the beginning
	// This allows us to handle the base cases elegantly:
	// maxMoney[0] = 0 (no houses robbed yet)
	// maxMoney[1] = 0 (still no houses robbed)
	// maxMoney[2] corresponds to decision for first house (nums[0])
	maxMoney := make([]int, len(nums)+2)

	// Fill the DP array for each house
	for houseIndex := 0; houseIndex < len(nums); houseIndex++ {
		// Map house index to DP array index (offset by 2)
		dpIndex := houseIndex + 2

		// For each house, we have two choices:
		// 1. Don't rob this house: take maxMoney[dpIndex-1] (previous maximum)
		// 2. Rob this house: take maxMoney[dpIndex-2] + nums[houseIndex]
		//    (maximum from 2 houses ago + current house money)
		// We choose the maximum of these two options
		maxMoney[dpIndex] = max(
			maxMoney[dpIndex-1],                  // Don't rob current house
			maxMoney[dpIndex-2]+nums[houseIndex], // Rob current house
		)
	}

	// Return the maximum money we can rob from all houses
	return maxMoney[len(maxMoney)-1]
}

// Alternative cleaner approach with more intuitive variable names
func Rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// Use more intuitive variable names
	// prevPrev: max money if we end 2 houses ago
	// prev: max money if we end at previous house
	prevPrev := 0
	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		current := max(prev, prevPrev+nums[i])
		prevPrev = prev
		prev = current
	}

	return prev
}
