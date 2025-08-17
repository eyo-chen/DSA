package main

import "fmt"

// MinDistance calculates the minimum edit distance using brute force recursion
// Approach: Recursive solution that explores all three operations (insert, delete, replace)
// Time Complexity: O(3^(m+n)) where m and n are lengths of word1 and word2
// Space Complexity: O(m+n) due to recursion stack depth
func MinDistance(word1 string, word2 string) int {
	return editDistanceRecursive(word1, word2, 0, 0)
}

// editDistanceRecursive is the helper function for brute force recursive approach
func editDistanceRecursive(word1 string, word2 string, word1Index int, word2Index int) int {
	// Base case: if we've processed all characters in word1,
	// we need to insert all remaining characters from word2
	if word1Index == len(word1) {
		return len(word2) - word2Index
	}

	// Base case: if we've processed all characters in word2,
	// we need to delete all remaining characters from word1
	if word2Index == len(word2) {
		return len(word1) - word1Index
	}

	// If characters match, move to next character in both strings (no operation needed)
	if word1[word1Index] == word2[word2Index] {
		return editDistanceRecursive(word1, word2, word1Index+1, word2Index+1)
	}

	// Try all three operations and find the minimum cost
	// Insert: add word2[word2Index] to word1, advance word2Index
	insertCost := editDistanceRecursive(word1, word2, word1Index, word2Index+1) + 1

	// Delete: remove word1[word1Index], advance word1Index
	deleteCost := editDistanceRecursive(word1, word2, word1Index+1, word2Index) + 1

	// Replace: replace word1[word1Index] with word2[word2Index], advance both indices
	replaceCost := editDistanceRecursive(word1, word2, word1Index+1, word2Index+1) + 1

	return min(insertCost, deleteCost, replaceCost)
}

// MinDistance1 calculates the minimum edit distance using memoized recursion
// Approach: Top-down dynamic programming with memoization to avoid recomputing subproblems
// Time Complexity: O(m*n) where m and n are lengths of word1 and word2
// Space Complexity: O(m*n) for memoization table + O(m+n) for recursion stack
func MinDistance1(word1 string, word2 string) int {
	memoTable := map[string]int{}
	return editDistanceMemoized(word1, word2, 0, 0, memoTable)
}

// editDistanceMemoized is the helper function for memoized recursive approach
func editDistanceMemoized(word1 string, word2 string, word1Index int, word2Index int, memoTable map[string]int) int {
	// Base case: if we've processed all characters in word1,
	// we need to insert all remaining characters from word2
	if word1Index == len(word1) {
		return len(word2) - word2Index
	}

	// Base case: if we've processed all characters in word2,
	// we need to delete all remaining characters from word1
	if word2Index == len(word2) {
		return len(word1) - word1Index
	}

	// If characters match, move to next character in both strings (no operation needed)
	if word1[word1Index] == word2[word2Index] {
		return editDistanceMemoized(word1, word2, word1Index+1, word2Index+1, memoTable)
	}

	// Create a unique key for the current state (word1Index, word2Index)
	stateKey := fmt.Sprintf("%d-%d", word1Index, word2Index)

	// Check if we've already computed this subproblem
	if cachedResult, exists := memoTable[stateKey]; exists {
		return cachedResult
	}

	// Try all three operations and find the minimum cost
	// Insert: add word2[word2Index] to word1, advance word2Index
	insertCost := editDistanceMemoized(word1, word2, word1Index, word2Index+1, memoTable) + 1

	// Delete: remove word1[word1Index], advance word1Index
	deleteCost := editDistanceMemoized(word1, word2, word1Index+1, word2Index, memoTable) + 1

	// Replace: replace word1[word1Index] with word2[word2Index], advance both indices
	replaceCost := editDistanceMemoized(word1, word2, word1Index+1, word2Index+1, memoTable) + 1

	// Store the minimum cost in memoization table
	result := min(insertCost, deleteCost, replaceCost)
	memoTable[stateKey] = result

	return result
}

// MinDistance2 calculates the minimum edit distance using bottom-up dynamic programming
// Approach: Build a 2D DP table where dp[i][j] represents min operations to convert word1[0:i] to word2[0:j]
// Time Complexity: O(m*n) where m and n are lengths of word1 and word2
// Space Complexity: O(m*n) for the 2D DP table
func MinDistance2(word1 string, word2 string) int {
	word1Length := len(word1)
	word2Length := len(word2)

	// Create DP table: dp[i][j] = min operations to convert word1[0:i] to word2[0:j]
	dp := make([][]int, word1Length+1)
	for i := range dp {
		dp[i] = make([]int, word2Length+1)
	}

	// Initialize first column: converting empty string to word1[0:i] requires i deletions
	for row := range dp {
		dp[row][0] = row
	}

	// Initialize first row: converting word1[0:0] to word2[0:j] requires j insertions
	for col := range dp[0] {
		dp[0][col] = col
	}

	// Fill the DP table using the recurrence relation
	for row := 1; row < len(dp); row++ {
		for col := 1; col < len(dp[0]); col++ {
			// If characters match, no operation needed, copy diagonal value
			if word1[row-1] == word2[col-1] {
				dp[row][col] = dp[row-1][col-1]
				continue
			}

			// Characters don't match, try all three operations and take minimum
			// Insert: dp[row-1][col] + 1 (insert word2[col-1])
			insertCost := dp[row-1][col] + 1

			// Delete: dp[row][col-1] + 1 (delete word1[row-1])
			deleteCost := dp[row][col-1] + 1

			// Replace: dp[row-1][col-1] + 1 (replace word1[row-1] with word2[col-1])
			replaceCost := dp[row-1][col-1] + 1

			dp[row][col] = min(insertCost, deleteCost, replaceCost)
		}
	}

	// Return the result: min operations to convert entire word1 to entire word2
	return dp[word1Length][word2Length]
}
