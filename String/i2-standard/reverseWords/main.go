package main

import "strings"

// ReverseWords reverses the order of words in a string.
// It splits the input string into words, then reconstructs the string
// with words in reverse order, separated by single spaces.
// Leading, trailing, and multiple consecutive spaces are normalized to single spaces.
//
// Time Complexity: O(n) where n is the length of the input string
// Space Complexity: O(n) for storing the words and building the result
func ReverseWords(s string) string {
	words := [][]byte{}
	i := 0

	// Extract all words from the string, skipping spaces
	for i < len(s) {
		// Skip any spaces
		if s[i] == ' ' {
			i++
			continue
		}

		// Build the current word character by character
		currentWord := []byte{}
		for i < len(s) && s[i] != ' ' {
			currentWord = append(currentWord, s[i])
			i++
		}

		// Add the completed word to our collection
		words = append(words, currentWord)
	}

	// Build the result string with words in reverse order
	var result strings.Builder
	for j := len(words) - 1; j >= 0; j-- {
		result.Write(words[j])

		// Add space between words, but not after the last word
		if j != 0 {
			result.WriteString(" ")
		}
	}

	return result.String()
}
