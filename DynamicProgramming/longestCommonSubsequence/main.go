package main

import "fmt"

// longestCommonSubsequence implements a naive recursive approach to find LCS length.
// This approach explores all possible combinations by checking each character match.
// Time Complexity: O(2^(m+n)) - exponential due to overlapping subproblems
// Space Complexity: O(m+n) - recursion stack depth
func LongestCommonSubsequence(text1 string, text2 string) int {
	return lcsRecursive(text1, text2, len(text1)-1, len(text2)-1)
}

func lcsRecursive(text1 string, text2 string, index1 int, index2 int) int {
	// Base case: if either string is exhausted, no common subsequence possible
	if index1 < 0 || index2 < 0 {
		return 0
	}

	// If characters match, include this character in LCS and move both pointers back
	if text1[index1] == text2[index2] {
		return lcsRecursive(text1, text2, index1-1, index2-1) + 1
	}

	// If characters don't match, try excluding either character and take maximum
	excludeFromText1 := lcsRecursive(text1, text2, index1-1, index2)
	excludeFromText2 := lcsRecursive(text1, text2, index1, index2-1)
	return max(excludeFromText1, excludeFromText2)
}

// longestCommonSubsequence1 implements top-down dynamic programming (memoization) approach.
// It caches previously computed results to avoid redundant calculations.
// Time Complexity: O(m*n) - each subproblem is solved only once
// Space Complexity: O(m*n) - memoization table + O(m+n) recursion stack
func LongestCommonSubsequence1(text1 string, text2 string) int {
	memoCache := map[string]int{}
	return lcsWithMemoization(text1, text2, len(text1)-1, len(text2)-1, memoCache)
}

func lcsWithMemoization(text1 string, text2 string, index1 int, index2 int, memoCache map[string]int) int {
	// Base case: if either string is exhausted, no common subsequence possible
	if index1 < 0 || index2 < 0 {
		return 0
	}

	// Check if result is already computed and cached
	cacheKey := fmt.Sprintf("%d-%d", index1, index2)
	if cachedResult, exists := memoCache[cacheKey]; exists {
		return cachedResult
	}

	var result int
	// If characters match, include this character in LCS and move both pointers back
	if text1[index1] == text2[index2] {
		result = lcsWithMemoization(text1, text2, index1-1, index2-1, memoCache) + 1
	} else {
		// If characters don't match, try excluding either character and take maximum
		excludeFromText1 := lcsWithMemoization(text1, text2, index1-1, index2, memoCache)
		excludeFromText2 := lcsWithMemoization(text1, text2, index1, index2-1, memoCache)
		result = max(excludeFromText1, excludeFromText2)
	}

	// Cache the result before returning
	memoCache[cacheKey] = result
	return result
}

// longestCommonSubsequence3 implements bottom-up dynamic programming approach.
// It builds the solution iteratively from smaller subproblems to larger ones.
// Time Complexity: O(m*n) - nested loops through both string lengths
// Space Complexity: O(m*n) - 2D DP table
func LongestCommonSubsequence3(text1 string, text2 string) int {
	text1Len, text2Len := len(text1), len(text2)

	// Create DP table with extra row/column for empty string cases
	// dp[i][j] represents LCS length of text1[0...i-1] and text2[0...j-1]
	dp := make([][]int, text1Len+1)
	for i := range dp {
		dp[i] = make([]int, text2Len+1)
	}
	// Base case: dp[0][j] = 0 and dp[i][0] = 0 (already initialized to 0)

	// Fill the DP table bottom-up
	for row := 1; row <= text1Len; row++ {
		for col := 1; col <= text2Len; col++ {
			// If characters match, extend the LCS from diagonal (both strings reduced by 1)
			if text1[row-1] == text2[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				// If characters don't match, take maximum from excluding either character
				excludeFromText1 := dp[row-1][col] // Exclude current char from text1
				excludeFromText2 := dp[row][col-1] // Exclude current char from text2
				dp[row][col] = max(excludeFromText1, excludeFromText2)
			}
		}
	}

	// Return LCS length for complete strings
	return dp[text1Len][text2Len]
}
