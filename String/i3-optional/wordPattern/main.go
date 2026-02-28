package wordpattern

import "strings"

// WordPattern checks if string s follows the same pattern as the given pattern string.
// It uses two hash maps to enforce a bijection (one-to-one mapping) between
// each pattern character and each word in s.
//
// Time complexity: O(n) where n is the length of the pattern
// Space complexity: O(n) for the two hash maps
func WordPattern(pattern string, s string) bool {
	charToWord := map[byte]string{}
	wordToChar := map[string]byte{}

	words := strings.Split(s, " ")

	// Pattern and words must have the same length for a valid mapping
	if len(pattern) != len(words) {
		return false
	}

	for i := range pattern {
		ch, word := pattern[i], words[i]

		// Check if character already maps to a different word
		if mappedWord, exists := charToWord[ch]; exists {
			if mappedWord != word {
				return false
			}
		} else {
			charToWord[ch] = word
		}

		// Check if word already maps to a different character
		if mappedChar, exists := wordToChar[word]; exists {
			if mappedChar != ch {
				return false
			}
		} else {
			wordToChar[word] = ch
		}
	}

	return true
}
