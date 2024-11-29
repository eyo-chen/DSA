package main

import (
	"fmt"
)

func EqualPairs(grid [][]int) int {
	// Use string representation to avoid collisions
	hashTable := make(map[string]int)
	ans := 0

	// Count occurrences of each row pattern
	for _, row := range grid {
		// Convert row to string
		// e.g. slice = [1 2 3] -> string = "[1 2 3]"
		key := fmt.Sprint(row)

		// Increment count of the row pattern
		hashTable[key]++
	}

	// Check each column against stored rows
	for c := 0; c < len(grid[0]); c++ {
		// Get column as a slice
		col := make([]int, len(grid))
		for r := 0; r < len(grid); r++ {
			col[r] = grid[r][c]
		}

		// Convert column to string
		// e.g. slice = [1 2 3] -> string = "[1 2 3]"
		key := fmt.Sprint(col)

		// Add count of matching rows
		ans += hashTable[key]
	}

	return ans
}
