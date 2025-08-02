package main

import (
	"math"
)

// APPROACH 1: RECURSIVE SOLUTION (Brute Force)
// Time Complexity: O(amount^coins) - exponential time due to repeated subproblems
// Space Complexity: O(amount) - due to recursion stack depth
//
// This approach explores all possible combinations by trying each coin at every step.
// It suffers from overlapping subproblems, making it inefficient for large inputs.
func CoinChange(coins []int, amount int) int {
	// Base case: if target amount is 0, we need 0 coins
	if amount == 0 {
		return 0
	}

	// Base case: if amount becomes negative, this path is invalid
	if amount < 0 {
		return -1
	}

	// Track the minimum number of coins needed across all possible coin choices
	minimumCoinsNeeded := math.MaxInt

	// Try using each available coin denomination
	for _, currentCoin := range coins {
		// Recursively find the minimum coins needed for the remaining amount
		// after using the current coin
		coinsForRemainingAmount := CoinChange(coins, amount-currentCoin)

		// If a valid solution exists for the remaining amount
		if coinsForRemainingAmount != -1 {
			// Calculate total coins needed: coins for remaining amount + 1 (current coin)
			totalCoinsWithCurrentCoin := coinsForRemainingAmount + 1
			// Keep track of the minimum across all coin choices
			minimumCoinsNeeded = min(minimumCoinsNeeded, totalCoinsWithCurrentCoin)
		}
	}

	// If no valid combination found, return -1
	if minimumCoinsNeeded == math.MaxInt {
		return -1
	}

	return minimumCoinsNeeded
}

// APPROACH 2: MEMOIZED RECURSION (Top-Down Dynamic Programming)
// Time Complexity: O(amount * coins) - each subproblem solved once
// Space Complexity: O(amount) - for memoization map + recursion stack
//
// This approach optimizes the recursive solution by storing results of subproblems
// to avoid redundant calculations, dramatically improving performance.
func CoinChange1(coins []int, amount int) int {
	// Create memoization map to store results of subproblems
	// Key: amount, Value: minimum coins needed for that amount
	memoizationCache := map[int]int{}
	return coinChangeHelper(coins, amount, memoizationCache)
}

func coinChangeHelper(coins []int, targetAmount int, memoizationCache map[int]int) int {
	// Base case: no coins needed to make amount 0
	if targetAmount == 0 {
		return 0
	}

	// Base case: negative amount is impossible to achieve
	if targetAmount < 0 {
		return -1
	}

	// Check if we've already solved this subproblem
	if cachedResult, existsInCache := memoizationCache[targetAmount]; existsInCache {
		return cachedResult
	}

	// Initialize minimum coins tracker
	minimumCoinsRequired := math.MaxInt

	// Try each coin denomination
	for _, currentCoin := range coins {
		// Recursively solve for the remaining amount after using current coin
		coinsNeededForRemainder := coinChangeHelper(coins, targetAmount-currentCoin, memoizationCache)

		// If a valid solution exists for the remainder
		if coinsNeededForRemainder != -1 {
			// Calculate total coins: remainder solution + 1 (for current coin)
			totalCoinsWithThisCoin := coinsNeededForRemainder + 1
			minimumCoinsRequired = min(minimumCoinsRequired, totalCoinsWithThisCoin)
		}
	}

	// Determine the result for this amount
	var resultForThisAmount int
	if minimumCoinsRequired == math.MaxInt {
		// No valid solution found
		resultForThisAmount = -1
	} else {
		resultForThisAmount = minimumCoinsRequired
	}

	// Store result in cache before returning
	memoizationCache[targetAmount] = resultForThisAmount
	return resultForThisAmount
}

// APPROACH 3: BOTTOM-UP DYNAMIC PROGRAMMING (Iterative Tabulation)
// Time Complexity: O(amount * coins) - nested loops over amount and coins
// Space Complexity: O(amount) - for the DP table only
//
// This approach builds solutions from the ground up, starting with smaller amounts
// and using those results to solve larger amounts. Most space-efficient approach.
func CoinChange2(coins []int, amount int) int {
	// Base case: 0 amount requires 0 coins
	if amount == 0 {
		return 0
	}

	// Create DP table where dpTable[i] = minimum coins needed to make amount i
	dpTable := make([]int, amount+1)

	// Initialize all positions (except 0) with an "impossible" value
	// We use (amount + 1) because it's larger than any possible legitimate answer
	// Since the worst case would be using 'amount' coins of denomination 1
	for currentAmount := 1; currentAmount < amount+1; currentAmount++ {
		dpTable[currentAmount] = amount + 1 // Sentinel value meaning "impossible"
	}
	// dpTable[0] remains 0 (base case: 0 coins needed for amount 0)

	// Build up solutions for each amount from 1 to target amount
	for currentAmount := 1; currentAmount <= amount; currentAmount++ {
		// Try each available coin denomination
		for _, coinValue := range coins {
			// Can we use this coin for the current amount?
			if currentAmount >= coinValue {
				// If yes, update the minimum coins needed for current amount
				// Either keep the existing solution OR use this coin + optimal solution for remainder
				coinsWithThisCoin := dpTable[currentAmount-coinValue] + 1
				dpTable[currentAmount] = min(dpTable[currentAmount], coinsWithThisCoin)
			}
		}
	}

	// Check if we found a valid solution for the target amount
	if dpTable[amount] == amount+1 {
		return -1 // Still has sentinel value, meaning impossible
	}

	return dpTable[amount]
}
