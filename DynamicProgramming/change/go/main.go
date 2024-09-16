package main

import "fmt"

func Change(amount int, coins []int) int {
	return helper(amount, coins, 0)
}

func helper(amount int, coins []int, index int) int {
	if amount == 0 {
		return 1
	}

	if amount < 0 {
		return 0
	}

	total := 0
	for i := index; i < len(coins); i++ {
		total += helper(amount-coins[i], coins, i)
	}

	return total
}

func Change1(amount int, coins []int) int {
	memo := map[string]int{}
	return helper1(amount, coins, 0, memo)
}

func helper1(amount int, coins []int, index int, memo map[string]int) int {
	if amount == 0 {
		return 1
	}

	if amount < 0 {
		return 0
	}

	key := fmt.Sprintf("%d-%d", amount, index)
	if v, ok := memo[key]; ok {
		return v
	}

	total := 0
	for i := index; i < len(coins); i++ {
		total += helper1(amount-coins[i], coins, i, memo)
	}

	memo[key] = total
	return total
}

func Change2(amount int, coins []int) int {
	// Create a 2D table to store the number of ways to make change
	// rows: number of coins + 1, columns: amount + 1
	table := make([][]int, len(coins)+1)
	for i := range table {
		table[i] = make([]int, amount+1)
		table[i][0] = 1 // There's one way to make 0 amount
	}

	// Iterate through each coin
	for i, coin := range coins {
		// For each amount from 1 to the target amount
		for currentAmount := 1; currentAmount <= amount; currentAmount++ {
			// Calculate ways without using the current coin
			waysWithoutCoin := table[i][currentAmount]

			// Calculate ways using the current coin
			waysWithCoin := 0
			if currentAmount >= coin {
				waysWithCoin = table[i+1][currentAmount-coin]
			}

			// Total ways is the sum of both
			table[i+1][currentAmount] = waysWithoutCoin + waysWithCoin
		}
	}

	// Return the total number of ways to make the target amount
	return table[len(coins)][amount]
}

// Space Optimized O(m), where m is the amount
// Same idea as Change2, but use two 1D array to store the current and next table
// No need to store the entire table, only the current and next table
func Change3(amount int, coins []int) int {
	// Initialize two tables for dynamic programming
	currentTable := make([]int, amount+1)
	nextTable := make([]int, amount+1)
	currentTable[0] = 1
	nextTable[0] = 1

	for _, coin := range coins {
		for currentAmount := 1; currentAmount <= amount; currentAmount++ {
			waysWithoutCoin := currentTable[currentAmount]

			waysWithCoin := 0
			if currentAmount >= coin {
				waysWithCoin = nextTable[currentAmount-coin]
			}

			nextTable[currentAmount] = waysWithCoin + waysWithoutCoin
		}

		// Swap tables for the next iteration
		currentTable, nextTable = nextTable, currentTable
	}

	return currentTable[amount]
}
