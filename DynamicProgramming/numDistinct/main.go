package main

import "fmt"

// NumDistinct - Brute Force with Subsequence Generation
// Approach: Generates all possible subsequences of s and counts how many equal t
// Time Complexity: O(2^n) where n = len(s) - generates all 2^n subsequences
// Space Complexity: O(n * 2^n) - stores all subsequences, each of max length n
func NumDistinct(s string, t string) int {
	currentSubsequence := []byte{}
	return generateSubsequencesHelper(s, t, 0, currentSubsequence)
}

func generateSubsequencesHelper(sourceStr string, targetStr string, currentIndex int, currentSubsequence []byte) int {
	// Base case: processed all characters in source string
	if currentIndex >= len(sourceStr) {
		// Check if current subsequence matches target
		if string(currentSubsequence) == targetStr {
			return 1
		}
		return 0
	}

	// Two choices at each position:
	// 1. Include current character in subsequence
	includeChar := generateSubsequencesHelper(sourceStr, targetStr, currentIndex+1, append(currentSubsequence, sourceStr[currentIndex]))

	// 2. Exclude current character from subsequence
	excludeChar := generateSubsequencesHelper(sourceStr, targetStr, currentIndex+1, currentSubsequence)

	return includeChar + excludeChar
}

// NumDistinct1 - Optimized Brute Force with Two Pointers
// Approach: Uses two pointers to match characters without generating actual subsequences
// Time Complexity: O(2^n) where n = len(s) - in worst case, explores all combinations
// Space Complexity: O(n + m) - recursion stack depth, where n = len(s), m = len(t)
func NumDistinct1(s string, t string) int {
	return twoPointersHelper(s, t, 0, 0)
}

func twoPointersHelper(sourceStr string, targetStr string, sourceIndex int, targetIndex int) int {
	// Base case: successfully matched all characters in target
	if targetIndex == len(targetStr) {
		return 1
	}

	// Base case: exhausted source string but haven't matched all of target
	if sourceIndex >= len(sourceStr) {
		return 0
	}

	// Decision 1: Skip this character
	// We're saying "I don't need this character to build my target word.
	// Let me see how many ways I can build the target using everything that comes after this character."
	skipCurrentChar := twoPointersHelper(sourceStr, targetStr, sourceIndex+1, targetIndex)

	matchCount := 0
	// Decision 2: Use this character (if it matches what I need)
	if sourceStr[sourceIndex] == targetStr[targetIndex] {
		// We're saying "This character is exactly what I need for the next letter in my target word.
		// If I use it, then I need to figure out how many ways I could build the remaining target
		// using the rest of the source string."
		matchCount = twoPointersHelper(sourceStr, targetStr, sourceIndex+1, targetIndex+1)
	}

	return skipCurrentChar + matchCount
}

// NumDistinct2 - Memoized Recursion (Top-Down Dynamic Programming)
// Approach: Same as NumDistinct1 but caches results to avoid redundant calculations
// Time Complexity: O(n * m) where n = len(s), m = len(t) - each state computed once
// Space Complexity: O(n * m) - memoization table + O(n + m) recursion stack
func NumDistinct2(s string, t string) int {
	memoizationCache := make(map[string]int)
	return memoizedHelper(s, t, 0, 0, memoizationCache)
}

func memoizedHelper(sourceStr string, targetStr string, sourceIndex int, targetIndex int, cache map[string]int) int {
	// Base case: successfully matched all characters in target
	if targetIndex == len(targetStr) {
		return 1
	}

	// Base case: exhausted source string but haven't matched all of target
	if sourceIndex >= len(sourceStr) {
		return 0
	}

	// Create unique key for current state
	stateKey := fmt.Sprintf("%d-%d", sourceIndex, targetIndex)
	if cachedResult, exists := cache[stateKey]; exists {
		return cachedResult
	}

	// Decision 1: Skip this character
	// We're saying "I don't need this character to build my target word.
	// Let me see how many ways I can build the target using everything that comes after this character."
	skipCurrentChar := memoizedHelper(sourceStr, targetStr, sourceIndex+1, targetIndex, cache)

	matchCount := 0
	// Decision 2: Use this character (if it matches what I need)
	if sourceStr[sourceIndex] == targetStr[targetIndex] {
		// We're saying "This character is exactly what I need for the next letter in my target word.
		// If I use it, then I need to figure out how many ways I could build the remaining target
		// using the rest of the source string."
		matchCount = memoizedHelper(sourceStr, targetStr, sourceIndex+1, targetIndex+1, cache)
	}

	result := skipCurrentChar + matchCount
	cache[stateKey] = result // Cache the result for future use
	return result
}

// NumDistinct3 - Bottom-Up Dynamic Programming (Tabulation)
// Approach: Builds solution iteratively using 2D DP table
// Time Complexity: O(n * m) where n = len(s), m = len(t)
// Space Complexity: O(n * m) - 2D DP table
func NumDistinct3(s string, t string) int {
	sourceLen := len(s)
	targetLen := len(t)

	// Create DP table: dp[i][j] = number of ways to form t[0:j] using s[0:i]
	dp := make([][]int, sourceLen+1)
	for i := range dp {
		dp[i] = make([]int, targetLen+1)
		// Empty target can be formed in exactly 1 way (by choosing nothing)
		dp[i][0] = 1
	}

	// Fill DP table bottom-up
	for sourcePos := 1; sourcePos <= sourceLen; sourcePos++ {
		for targetPos := 1; targetPos <= targetLen; targetPos++ {
			// Decision 1: Skip this character
			// We're saying "I don't need this character to build my target word.
			// Let me see how many ways I can build the target using everything that comes after this character."
			dp[sourcePos][targetPos] = dp[sourcePos-1][targetPos]

			// Decision 2: Use this character (if it matches what I need)
			// We're saying "This character is exactly what I need for the next letter in my target word.
			// If I use it, then I need to figure out how many ways I could build the remaining target
			// using the rest of the source string."
			if s[sourcePos-1] == t[targetPos-1] {
				dp[sourcePos][targetPos] += dp[sourcePos-1][targetPos-1]
			}
		}
	}

	// Return the number of ways to form entire target using entire source
	return dp[sourceLen][targetLen]
}
