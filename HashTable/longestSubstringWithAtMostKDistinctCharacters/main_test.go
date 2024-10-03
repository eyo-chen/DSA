package main

import (
	"testing"

	"math/rand"
)

func TestRandomInputs(t *testing.T) {
	const numTests = 1000
	const maxLen = 5
	const maxK = 3

	for i := 0; i < numTests; i++ {
		s := generateRandomString(maxLen)
		k := rand.Intn(maxK) + 1

		result1 := LongestSubstringWithAtMostKDistinctCharacters2(s, k)
		result2 := Solution(s, k)

		if result1 != result2 {
			t.Errorf("Mismatch for s=%q, k=%d: longestSubstringWithAtMostKDistinctCharacters=%d, solution=%d", s, k, result1, result2)
		}
	}
}

func generateRandomString(maxLen int) string {
	length := rand.Intn(maxLen) + 1
	result := make([]byte, length)
	for i := range result {
		result[i] = byte(rand.Intn(26) + 'a')
	}
	return string(result)
}

// longestSubstringWithAtMostKDistinctCharacters finds the length of the longest
// substring that contains at most k distinct characters using brute force
func Solution(s string, k int) int {
	maxLen := 0
	n := len(s)

	// Check all possible substrings
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			substring := s[i : j+1]
			if countDistinctChars(substring) <= k {
				if len(substring) > maxLen {
					maxLen = len(substring)
				}
			}
		}
	}

	return maxLen
}
