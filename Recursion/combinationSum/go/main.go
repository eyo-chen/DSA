package main

import (
	"slices"
)

// combinationSum finds all unique combinations of candidates that sum to target.
// Each candidate can be used multiple times in a combination.
// Approach: Uses backtracking with sorted array to efficiently explore all valid combinations.
// The sorting optimization allows early termination when current candidate exceeds remaining target.
func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	// Sort candidates in ascending order to enable early termination optimization
	slices.Sort(candidates)

	// Start backtracking from index 0 with empty current combination
	findCombinations(candidates, &result, []int{}, 0, target)
	return result
}

// findCombinations uses backtracking to explore all valid combinations recursively.
// It builds combinations by including candidates and tracks remaining target sum.
func findCombinations(candidates []int, result *[][]int, currentCombination []int, startIndex int, remainingTarget int) {
	// Base case: invalid path - remaining target became negative
	if remainingTarget < 0 {
		return
	}

	// Base case: found valid combination - remaining target is exactly zero
	if remainingTarget == 0 {
		// Create a copy of current combination to avoid reference issues
		validCombination := make([]int, len(currentCombination))
		copy(validCombination, currentCombination)
		*result = append(*result, validCombination)
		return
	}

	// Explore all candidates starting from startIndex to avoid duplicate combinations
	for i := startIndex; i < len(candidates); i++ {
		currentCandidate := candidates[i]

		// Optimization: if current candidate exceeds remaining target,
		// all subsequent candidates will also exceed it (due to sorting)
		if currentCandidate > remainingTarget {
			return
		}

		// Recursively explore by including current candidate
		// Use same index 'i' to allow reuse of same candidate
		findCombinations(candidates, result, append(currentCombination, currentCandidate), i, remainingTarget-currentCandidate)
	}
}
