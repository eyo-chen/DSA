package main

// The key insight: In a circular arrangement, the first and last houses are neighbors,
// creating a constraint that we cannot rob both. We solve this by breaking the problem
// into two linear subproblems:
//  1. Rob houses 1 to n-1 (exclude first house, can include last)
//  2. Rob houses 0 to n-2 (include first house, exclude last)
//
// By taking the maximum of these two scenarios, we ensure we never violate the
// circular adjacency constraint while still finding the optimal solution.
//
// Time complexity: O(n) - we traverse the array twice
// Space complexity: O(1) - only using constant extra space
func Rob(nums []int) int {
	// Base case: only one house, no adjacency constraint
	if len(nums) == 1 {
		return nums[0]
	}

	// Scenario 1: Exclude first house (index 0), consider houses 1 to n-1
	// This allows us to potentially rob the last house without violating constraints
	maxWithoutFirst := robLinearRange(nums[1:])

	// Scenario 2: Exclude last house (index n-1), consider houses 0 to n-2
	// This allows us to potentially rob the first house without violating constraints
	maxWithoutLast := robLinearRange(nums[:len(nums)-1])

	// Return the maximum money from either scenario
	return max(maxWithoutFirst, maxWithoutLast)
}

// robLinearRange solves the classic linear house robber problem using space-optimized DP.
//
// For any house i, we face a choice:
//   - Rob house i: get nums[i] + max money from houses up to i-2
//   - Skip house i: keep max money from houses up to i-1
//
// Instead of storing all previous results, we only track the last two values
// since that's all we need for the recurrence relation.
func robLinearRange(nums []int) int {
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

	// Process each house starting from index 1
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
