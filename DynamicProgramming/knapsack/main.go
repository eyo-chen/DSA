package main

// Knapsack solves the 0-1 knapsack problem using dynamic programming
func Knapsack(values []int, weights []int, maxWeightConstraint int) int {
	// Create a 2D DP table where dp[itemIndex][currentWeight] represents the maximum value
	// that can be obtained using the first 'itemIndex' items with weight limit 'currentWeight'
	// We use (len(weights)+1) to include the base case of 0 items
	dp := make([][]int, len(weights)+1)
	for itemIndex := 0; itemIndex < len(dp); itemIndex++ {
		// Each row has (maxWeightConstraint+1) columns to include weight 0 to maxWeightConstraint
		dp[itemIndex] = make([]int, maxWeightConstraint+1)
	}

	// Iterate through each item (starting from 1 since index 0 represents "no items")
	for itemIndex := 1; itemIndex < len(dp); itemIndex++ {
		currentItemWeight := weights[itemIndex-1] // itemIndex-1 because dp is 1-indexed but weights array is 0-indexed
		currentItemValue := values[itemIndex-1]   // itemIndex-1 because dp is 1-indexed but values array is 0-indexed

		// Iterate through each possible weight capacity (starting from 1 since weight 0 gives value 0)
		for currentWeight := 1; currentWeight < len(dp[0]); currentWeight++ {
			// Option 1: Include the current item (if it fits)
			valueIfItemIncluded := -1

			// Check if current item can fit in the remaining weight capacity
			if currentWeight >= currentItemWeight {
				remainingWeight := currentWeight - currentItemWeight
				// Add current item's value to the best value achievable with remaining weight and previous items
				valueIfItemIncluded = dp[itemIndex-1][remainingWeight] + currentItemValue
			}

			// Option 2: Don't include the current item (take the best value from previous items with same weight limit)
			valueIfItemNotIncluded := dp[itemIndex-1][currentWeight]

			// Choose the option that gives maximum value
			dp[itemIndex][currentWeight] = max(valueIfItemIncluded, valueIfItemNotIncluded)
		}
	}

	// Return the maximum value achievable with all items and full weight constraint
	return dp[len(weights)][maxWeightConstraint]
}

// knapsackOptimized uses space-optimized version with 1D array
// TODO: Review next time(noted at 2025/08/09)
func KnapsackOptimized(values []int, weights []int, maxWeight int) int {
	n := len(values)
	if n == 0 || maxWeight <= 0 {
		return 0
	}

	// Use only 1D array since we only need previous row
	dp := make([]int, maxWeight+1)

	for i := 0; i < n; i++ {
		// Traverse backwards to avoid using updated values
		for w := maxWeight; w >= weights[i]; w-- {
			dp[w] = max(dp[w], dp[w-weights[i]]+values[i])
		}
	}

	return dp[maxWeight]
}
