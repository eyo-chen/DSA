package main

import "fmt"

// State constants for better readability
const (
	NOT_HOLDING = 0 // State when we don't own any stock
	HOLDING     = 1 // State when we own stock
)

func maxProfit(prices []int) int {
	// Start from day 0 in NOT_HOLDING state (we don't own stock initially)
	return calculateMaxProfit(prices, 0, NOT_HOLDING)
}

func calculateMaxProfit(prices []int, currentDay int, state int) int {
	// Base case: if we've reached the end of trading days, no more profit possible
	if currentDay >= len(prices) {
		return 0
	}

	// Option 1: Do nothing today (stay in current state, move to next day)
	doNothingProfit := calculateMaxProfit(prices, currentDay+1, state)

	if state == NOT_HOLDING {
		// We don't own stock, so we can choose to buy
		// Option 2: Buy stock today
		// - Cost: prices[currentDay] (negative profit)
		// - New state: HOLDING
		// - Next day: currentDay + 1
		buyProfit := calculateMaxProfit(prices, currentDay+1, HOLDING) - prices[currentDay]

		// Return the maximum profit between doing nothing and buying
		return max(doNothingProfit, buyProfit)
	}

	// state == HOLDING: We own stock, so we can choose to sell
	// Option 2: Sell stock today
	// - Gain: prices[currentDay] (positive profit)
	// - New state: NOT_HOLDING
	// - Next day: currentDay + 2 (skip next day due to cooldown)
	sellProfit := calculateMaxProfit(prices, currentDay+2, NOT_HOLDING) + prices[currentDay]

	// Return the maximum profit between doing nothing and selling
	return max(doNothingProfit, sellProfit)
}

// Time Complexity: O(n) - We solve each unique subproblem (day, state) exactly once
// Space Complexity: O(n) - Memoization table stores at most 2n entries + recursion stack depth of n
func maxProfit2(prices []int) int {
	// Initialize memoization table to store computed results
	memo := make(map[string]int)

	// Start from day 0 in NOT_HOLDING state (we don't own stock initially)
	return calculateMaxProfitWithMemo(prices, 0, NOT_HOLDING, memo)
}

// Time Complexity: O(n) - Each unique (day, state) combination is computed only once
// Space Complexity: O(n) - Memoization table + recursion stack
func calculateMaxProfitWithMemo(prices []int, currentDay int, state int, memo map[string]int) int {
	// Base case: if we've reached the end of trading days, no more profit possible
	if currentDay >= len(prices) {
		return 0
	}

	// Create unique key for memoization: "day-state"
	memoKey := fmt.Sprintf("%d-%d", currentDay, state)

	// Check if we've already computed this subproblem
	if cachedResult, exists := memo[memoKey]; exists {
		return cachedResult
	}

	// Option 1: Do nothing today (stay in current state, move to next day)
	doNothingProfit := calculateMaxProfitWithMemo(prices, currentDay+1, state, memo)

	var maxProfitForToday int

	if state == NOT_HOLDING {
		// We don't own stock, so we can choose to buy
		// Option 2: Buy stock today
		// - Cost: prices[currentDay] (negative profit)
		// - New state: HOLDING
		// - Next day: currentDay + 1
		buyProfit := calculateMaxProfitWithMemo(prices, currentDay+1, HOLDING, memo) - prices[currentDay]

		// Choose the maximum profit between doing nothing and buying
		maxProfitForToday = max(doNothingProfit, buyProfit)
	} else {
		// state == HOLDING: We own stock, so we can choose to sell
		// Option 2: Sell stock today
		// - Gain: prices[currentDay] (positive profit)
		// - New state: NOT_HOLDING
		// - Next day: currentDay + 2 (skip next day due to cooldown)
		sellProfit := calculateMaxProfitWithMemo(prices, currentDay+2, NOT_HOLDING, memo) + prices[currentDay]

		// Choose the maximum profit between doing nothing and selling
		maxProfitForToday = max(doNothingProfit, sellProfit)
	}

	// Store the computed result in memo table for future use
	memo[memoKey] = maxProfitForToday
	return maxProfitForToday
}

// Time Complexity: O(n) - Single pass through the prices array
// Space Complexity: O(n) - Three DP arrays of size n each
func maxProfit3(prices []int) int {
	// Edge case: can't make profit with 0 or 1 day of prices
	if len(prices) <= 1 {
		return 0
	}

	numDays := len(prices)

	// DP arrays to track maximum profit in each state for each day
	holdingStock := make([]int, numDays)  // Maximum profit when we own stock on day i
	justSoldStock := make([]int, numDays) // Maximum profit when we sold stock on day i
	restingState := make([]int, numDays)  // Maximum profit when we're resting (no stock, can buy next day)

	// Base case: Day 0 initialization
	holdingStock[0] = -prices[0] // Buy stock on day 0 (negative because it's a cost)
	justSoldStock[0] = 0         // Impossible to sell on day 0 (no stock to sell)
	restingState[0] = 0          // Start with no profit, no stock

	// Fill DP arrays for each day
	for currentDay := 1; currentDay < numDays; currentDay++ {
		// To hold stock today, we either:
		// 1) Already held stock yesterday (holdingStock[currentDay-1])
		// 2) Buy today after resting yesterday (restingState[currentDay-1] - prices[currentDay])
		holdingStock[currentDay] = max(
			holdingStock[currentDay-1],                    // Keep holding from yesterday
			restingState[currentDay-1]-prices[currentDay], // Buy today after resting
		)

		// To have just sold stock today:
		// We must have held stock yesterday and sell it today
		justSoldStock[currentDay] = holdingStock[currentDay-1] + prices[currentDay]

		// To be resting today, we either:
		// 1) Were already resting yesterday (restingState[currentDay-1])
		// 2) Are in cooldown after selling yesterday (justSoldStock[currentDay-1])
		restingState[currentDay] = max(
			restingState[currentDay-1],  // Continue resting from yesterday
			justSoldStock[currentDay-1], // Rest after selling yesterday (cooldown)
		)
	}

	// At the end, we want maximum profit without holding stock
	// Choose between being in resting state or having just sold on the last day
	finalDay := numDays - 1
	return max(justSoldStock[finalDay], restingState[finalDay])
}

// TODO: Add Space Optimized Approach
