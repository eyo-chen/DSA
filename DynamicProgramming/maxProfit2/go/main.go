package main

func MaxProfit(prices []int) int {
	maxProfitTable := make([]int, len(prices))

	// start from the second day
	// because we know the first day's profit is always 0
	for i := 1; i < len(prices); i++ {
		sellToday := prices[i] - prices[i-1] + maxProfitTable[i-1]
		doNotSellToday := maxProfitTable[i-1]
		maxProfitTable[i] = max(sellToday, doNotSellToday)
	}

	return maxProfitTable[len(prices)-1]
}

func MaxProfit1(prices []int) int {
	maxProfit := 0

	// start from the second day
	// because we know the first day's profit is always 0
	for i := 1; i < len(prices); i++ {
		sellToday := prices[i] - prices[i-1] + maxProfit
		doNotSellToday := maxProfit
		maxProfit = max(sellToday, doNotSellToday)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
