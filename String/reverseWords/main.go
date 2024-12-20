package main

import (
	"strings"
)

// Use Array to store the words and then reverse the array
// Time Complexity: O(n)
// Space Complexity: O(n)
func ReverseWords(s string) string {
	// arr is used to store the words
	arr := []string{}

	// curString is used to store the current word
	curString := []byte{}
	for i := 0; i < len(s); i++ {
		// If the current character is a space and the current word is empty, skip it
		// Remove the leading spaces or trailing spaces or multiple spaces between words
		if len(curString) == 0 && s[i] == ' ' {
			continue
		}

		// If the current character is a space and the current word is not empty, add the current word to the array
		if s[i] == ' ' {
			arr = append(arr, string(curString))
			curString = []byte{}
			continue
		}

		// Add the current character to the current word
		curString = append(curString, s[i])
	}

	// If the current word is not empty, add it to the array
	if len(curString) > 0 {
		arr = append(arr, string(curString))
	}

	// Build the result string from the array
	sb := strings.Builder{}

	// Start from the last word and add it to the result string
	for i := len(arr) - 1; i >= 0; i-- {
		sb.WriteString(arr[i])

		// If it is not the last word, add a space
		if i != 0 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
