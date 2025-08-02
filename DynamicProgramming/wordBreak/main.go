package main

import "strings"

// WordBreak solves the word break problem using recursive brute force approach.
// It tries each word in the dictionary as a potential prefix and recursively
// checks if the remaining string can be broken into valid dictionary words.
//
// Time Complexity: O(2^n * m) where n is length of string, m is average word length
// Space Complexity: O(n) for recursion stack depth
func WordBreak(s string, wordDict []string) bool {
	// Base case: empty string can always be segmented
	if len(s) == 0 {
		return true
	}

	// Try each word in dictionary as potential starting word
	for _, currentWord := range wordDict {
		// Skip if current word doesn't match the beginning of string
		if !strings.HasPrefix(s, currentWord) {
			continue
		}

		// Extract remaining string after removing the matched word
		remainingString := s[len(currentWord):]

		// Recursively check if remaining string can be segmented
		if WordBreak(remainingString, wordDict) {
			return true
		}
	}

	// No valid segmentation found
	return false
}

// WordBreak2 solves the word break problem using memoization to optimize
// the recursive approach by caching results for previously computed substrings.
//
// Time Complexity: O(n^2 * m) where n is length of string, m is average word length
// Space Complexity: O(n^2) for memoization map + O(n) for recursion stack
func WordBreak2(s string, wordDict []string) bool {
	// Memoization map to store results for previously computed strings
	memoCache := make(map[string]bool)
	return wordBreakHelper(s, wordDict, memoCache)
}

// wordBreakHelper is the recursive helper function with memoization support
func wordBreakHelper(currentString string, wordDict []string, memoCache map[string]bool) bool {
	// Base case: empty string can always be segmented
	if len(currentString) == 0 {
		return true
	}

	// Check if result for current string is already computed
	if cachedResult, exists := memoCache[currentString]; exists {
		return cachedResult
	}

	// Try each word in dictionary as potential starting word
	for _, candidateWord := range wordDict {
		// Skip if candidate word doesn't match the beginning
		if !strings.HasPrefix(currentString, candidateWord) {
			continue
		}

		// Extract remaining string after removing the matched word
		remainingString := currentString[len(candidateWord):]

		// Recursively check if remaining string can be segmented
		if wordBreakHelper(remainingString, wordDict, memoCache) {
			// Cache positive result and return
			memoCache[currentString] = true
			return true
		}
	}

	// Cache negative result and return
	memoCache[currentString] = false
	return false
}

// WordBreak3 solves the word break problem using dynamic programming approach.
// It builds up the solution iteratively by checking each position in the string
// and determining if a valid word can be placed starting from that position.
//
// Time Complexity: O(n * m * k) where n is string length, m is dict size, k is avg word length
// Space Complexity: O(n) for the DP table
func WordBreak3(s string, wordDict []string) bool {
	stringLength := len(s)

	// DP table where dpTable[i] indicates if s[0:i] can be segmented
	dpTable := make([]bool, stringLength+1)
	dpTable[0] = true // Empty string base case

	// Iterate through each position in the string
	for currentPosition := 0; currentPosition <= stringLength; currentPosition++ {
		// Skip positions that cannot be reached through valid segmentation
		if !dpTable[currentPosition] {
			continue
		}

		// Try placing each dictionary word starting from current position
		for _, dictionaryWord := range wordDict {
			wordLength := len(dictionaryWord)
			endPosition := currentPosition + wordLength

			// Check if word fits within string bounds and matches the substring
			if endPosition <= stringLength && s[currentPosition:endPosition] == dictionaryWord {
				dpTable[endPosition] = true
			}
		}
	}

	// Return whether the entire string can be segmented
	return dpTable[stringLength]
}
