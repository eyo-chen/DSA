package main

import (
	"strconv"
	"strings"
)

// FindAnagrams finds all start indices of p's anagrams in s using brute force approach.
// For each possible window in s, it builds a character frequency map and compares with p's frequency map.
// Time Complexity: O(n * m) where n = len(s), m = len(p)
// Space Complexity: O(1) - frequency arrays are fixed size 26
func FindAnagrams(s string, p string) []int {
	result := []int{}

	// Build character frequency map for pattern string p
	patternFreq := make([]int, 26)
	for i := range p {
		patternFreq[p[i]-'a']++
	}

	// Check each window of size len(p) in string s
	windowCount := len(s) - len(p)
	for i := 0; i <= windowCount; i++ {
		// Build character frequency map for current window
		windowFreq := make([]int, 26)
		for k := range p {
			windowFreq[s[i+k]-'a']++
		}

		// If frequency maps match, current window is an anagram
		if areFrequenciesEqual(windowFreq, patternFreq) {
			result = append(result, i)
		}
	}

	return result
}

// areFrequenciesEqual checks if two character frequency arrays are identical.
func areFrequenciesEqual(freq1 []int, freq2 []int) bool {
	for i := range freq1 {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}

// FindAnagrams1 finds all start indices of p's anagrams in s using sliding window with string key comparison.
// Generates a unique string key from frequency maps and compares keys instead of arrays.
// Time Complexity: O(n * 26) = O(n) where n = len(s)
// Space Complexity: O(1) - frequency arrays are fixed size, keys are constant size
func FindAnagrams1(s string, p string) []int {
	result := []int{}

	// Build character frequency map for pattern string p
	patternFreq := make([]int, 26)
	for i := 0; i < len(p); i++ {
		patternFreq[p[i]-'a']++
	}
	patternKey := generateFrequencyKey(patternFreq)

	// Initialize sliding window with frequency map
	windowFreq := make([]int, 26)
	left, right := 0, 0

	// Expand window to the right
	for right < len(s) {
		// Add current character to window
		windowFreq[s[right]-'a']++

		// When window reaches target size, check for anagram
		if right-left+1 == len(p) {
			windowKey := generateFrequencyKey(windowFreq)

			// If keys match, we found an anagram
			if windowKey == patternKey {
				result = append(result, left)
			}

			// Shrink window from left to maintain size
			windowFreq[s[left]-'a']--
			left++
		}

		// Move right pointer forward
		right++
	}

	return result
}

// generateFrequencyKey creates a unique string representation of character frequencies.
// Example: [1,0,2,0,...] -> "1-0-2-0-..."
func generateFrequencyKey(frequencies []int) string {
	builder := strings.Builder{}
	for _, count := range frequencies {
		builder.WriteString(strconv.Itoa(count))
		builder.WriteString("-")
	}
	return builder.String()
}

// FindAnagrams2 finds all start indices of p's anagrams in s using optimized sliding window.
// Maintains a sliding window of size len(p) and updates frequency map incrementally.
// Time Complexity: O(n) where n = len(s)
// Space Complexity: O(1) - frequency arrays are fixed size 26
func FindAnagrams2(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}

	// Build character frequency maps for pattern and initial window
	patternFreq := make([]int, 26)
	windowFreq := make([]int, 26)
	for i := range p {
		patternFreq[p[i]-'a']++
		windowFreq[s[i]-'a']++
	}

	// Check if initial window is an anagram
	if areFrequenciesEqual(patternFreq, windowFreq) {
		result = append(result, 0)
	}

	// Slide window through rest of string
	for i := len(p); i < len(s); i++ {
		// Add new character entering window on right
		windowFreq[s[i]-'a']++

		// Remove old character leaving window on left
		windowFreq[s[i-len(p)]-'a']--

		// Check if current window is an anagram
		if areFrequenciesEqual(patternFreq, windowFreq) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

// FindAnagrams3 finds all start indices of p's anagrams in s using match counter optimization.
// Instead of comparing entire frequency arrays, tracks count of matched characters.
// Time Complexity: O(n) where n = len(s)
// Space Complexity: O(1) - frequency array is fixed size 26
func FindAnagrams3(s string, p string) []int {
	result := []int{}

	if len(s) < len(p) {
		return result
	}

	// Build character frequency map for pattern string p
	charNeeded := make([]int, 26)
	for i := 0; i < len(p); i++ {
		charNeeded[p[i]-'a']++
	}

	// Build match count for initial window
	// matchedChars counts how many characters in window correctly match pattern
	matchedChars := 0
	for i := 0; i < len(p); i++ {
		charIndex := s[i] - 'a'

		// Decrement needed count as we add character to window
		charNeeded[charIndex]--

		// If count is non-negative, this character contributes to a valid match
		// (it was needed by the pattern and hasn't exceeded pattern's frequency)
		if charNeeded[charIndex] >= 0 {
			matchedChars++
		}
	}

	// If all characters match, initial window is an anagram
	if matchedChars == len(p) {
		result = append(result, 0)
	}

	// Slide window through rest of string
	for i := len(p); i < len(s); i++ {
		// Remove leftmost character from window
		leftCharIndex := s[i-len(p)] - 'a'
		charNeeded[leftCharIndex]++

		// If count becomes positive, we lost a valid match
		// (this character was contributing to the anagram)
		if charNeeded[leftCharIndex] > 0 {
			matchedChars--
		}

		// Add new rightmost character to window
		rightCharIndex := s[i] - 'a'
		charNeeded[rightCharIndex]--

		// If count is non-negative, we gained a valid match
		if charNeeded[rightCharIndex] >= 0 {
			matchedChars++
		}

		// If all characters match, current window is an anagram
		if matchedChars == len(p) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
