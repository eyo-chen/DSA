package main

import (
	"strconv"
	"strings"
)

// CountAndSay generates the nth term of the count-and-say sequence using recursion.
// The sequence starts with "1" and each subsequent term describes the previous term
// by counting consecutive identical digits.
// Example: "1" -> "11" (one 1) -> "21" (two 1s) -> "1211" (one 2, one 1)
//
// Approach: Recursively builds the sequence from the base case up. Each call generates
// the (n-1)th term first, then processes it to create the nth term. This approach is
// more elegant but uses additional stack space for the recursive calls.
//
// Time Complexity: O(2^n) - each term can be roughly twice the length of the previous term
// Space Complexity: O(2^n) - for the recursion stack and string storage
func CountAndSay(n int) string {
	// Base case: the first term is "1"
	if n == 1 {
		return "1"
	}

	var result strings.Builder
	// Recursively get the previous term in the sequence
	prevTerm := CountAndSay(n - 1)
	i := 0

	// Process each group of consecutive identical digits
	for i < len(prevTerm) {
		// Find the end of the current group of identical digits
		groupEnd := i + 1
		for groupEnd < len(prevTerm) && prevTerm[groupEnd] == prevTerm[i] {
			groupEnd++
		}

		// Append count followed by the digit
		digitCount := strconv.Itoa(groupEnd - i)
		result.WriteString(digitCount)
		result.WriteByte(prevTerm[i])

		// Move to the next group
		i = groupEnd
	}

	return result.String()
}

// CountAndSay1 generates the nth term of the count-and-say sequence using iteration.
// The sequence starts with "1" and each subsequent term describes the previous term
// by counting consecutive identical digits.
// Example: "1" -> "11" (one 1) -> "21" (two 1s) -> "1211" (one 2, one 1)
//
// Approach: Iteratively builds the sequence starting from "1" and generating each
// subsequent term in order until reaching the nth term. This approach avoids recursion
// overhead and is more space-efficient as it only keeps track of the current term.
//
// Time Complexity: O(2^n) - each term can be roughly twice the length of the previous term
// Space Complexity: O(2^n) - for storing the current and previous terms
func CountAndSay1(n int) string {
	currentTerm := "1"

	// Build each term iteratively from term 1 to term n
	for term := 1; term < n; term++ {
		var nextTerm strings.Builder
		i := 0

		// Process each group of consecutive identical digits in the current term
		for i < len(currentTerm) {
			// Find the end of the current group of identical digits
			groupEnd := i + 1
			for groupEnd < len(currentTerm) && currentTerm[i] == currentTerm[groupEnd] {
				groupEnd++
			}

			// Append count followed by the digit
			digitCount := strconv.Itoa(groupEnd - i)
			nextTerm.WriteString(digitCount)
			nextTerm.WriteByte(currentTerm[i])

			// Move to the next group
			i = groupEnd
		}

		// Update current term for the next iteration
		currentTerm = nextTerm.String()
	}

	return currentTerm
}
