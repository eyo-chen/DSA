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
	// Handle edge cases
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// maxMoneyTwoHousesBack: maximum money we could rob from houses ending 2 positions ago
	// This represents dp[i-2] in traditional DP notation
	maxMoneyTwoHousesBack := 0

	// maxMoneyOneHouseBack: maximum money we could rob from houses ending 1 position ago
	// This represents dp[i-1] in traditional DP notation
	maxMoneyOneHouseBack := nums[0]

	// We start the loop at index 1 (the second house) because we've already solved
	// the base case for index 0. At index 1, we encounter our first real decision:
	// should we rob this house or stick with our previous best? Our recurrence
	// relation dp[i] = max(dp[i-1], nums[i] + dp[i-2]) requires both dp[i-1]
	// and dp[i-2] to be available. Starting at index 1 ensures we have:
	// - dp[i-1] which is maxMoneyOneHouseBack (from house 0)
	// - dp[i-2] which is maxMoneyTwoHousesBack (from empty set, value 0)
	for currentHouseIndex := 1; currentHouseIndex < len(nums); currentHouseIndex++ {
		// Current choice: rob this house + best from two houses back,
		// OR skip this house and keep the best from one house back
		maxMoneyIncludingCurrentHouse := max(
			maxMoneyOneHouseBack,                          // Skip current house
			nums[currentHouseIndex]+maxMoneyTwoHousesBack, // Rob current house
		)

		// Slide our window forward: what was "one back" becomes "two back"
		maxMoneyTwoHousesBack = maxMoneyOneHouseBack
		// Current result becomes our new "one back" for next iteration
		maxMoneyOneHouseBack = maxMoneyIncludingCurrentHouse
	}

	// Return the maximum money we can rob from this linear range
	return maxMoneyOneHouseBack
}
