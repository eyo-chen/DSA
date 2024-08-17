package main

// Brute force solution
func MaxProfit(prices []int) int {
	maxProfit := 0

	// Try each day as a buying day
	for i := 0; i < len(prices); i++ {
		buyPrice := prices[i]

		// Try each day after the buying day as a selling day
		for j := i + 1; j < len(prices); j++ {
			sellPrice := prices[j]
			maxProfit = max(maxProfit, sellPrice-buyPrice)
		}
	}

	return maxProfit
}

// Optimized solution
func MaxProfit1(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices {
		// What is the minimum price so far?
		minPrice = min(minPrice, p)

		// What is the maximum profit so far?
		maxProfit = max(maxProfit, p-minPrice)
	}

	return maxProfit
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
