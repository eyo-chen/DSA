package main

import "slices"

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}

// Brute Force
// Time Complexity: O(n * k)
// Space Complexity: O(1)
func MaxVowels(s string, k int) int {
	ans := 0

	for i := 0; i <= len(s)-k; i++ {
		vowelCounts := 0
		for j := 0; j < k; j++ {
			if slices.Contains(vowels, s[i+j]) {
				vowelCounts++
			}
		}

		ans = max(ans, vowelCounts)
	}

	return ans
}

// Sliding Window (reference from findMaxAverage)
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxVowels2(s string, k int) int {
	ans, vowelCounts := 0, 0

	// Calculate initial window vowel counts
	for i := 0; i < k; i++ {
		if slices.Contains(vowels, s[i]) {
			vowelCounts++
		}
	}
	ans = max(ans, vowelCounts)

	// Slide the window
	for i := k; i < len(s); i++ {
		// Add the new element
		if slices.Contains(vowels, s[i]) {
			vowelCounts++
		}

		// Remove the old element
		if slices.Contains(vowels, s[i-k]) {
			vowelCounts--
		}

		ans = max(ans, vowelCounts)
	}

	return ans
}

// Sliding Window Optimized (Only use one loop, code is more concise)
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxVowels3(s string, k int) int {
	ans, vowelCounts := 0, 0

	for i := 0; i < len(s); i++ {
		// Add the new element
		if slices.Contains(vowels, s[i]) {
			vowelCounts++
		}

		// Remove the old element (when the window is full)
		if i >= k && slices.Contains(vowels, s[i-k]) {
			vowelCounts--
		}

		ans = max(ans, vowelCounts)
	}

	return ans
}
