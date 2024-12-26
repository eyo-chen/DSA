package main

import "unicode"

// Brute Force (3 Passes)
// Time Complexity: O(n)
// Space Complexity: O(n)
func IsPalindrome(s string) bool {
	str := make([]rune, 0, len(s))

	// Filter out non-alphanumeric characters
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			str = append(str, c)
		}
	}

	// Convert to uppercase
	for i, c := range str {
		str[i] = unicode.ToUpper(c)
	}

	// Check if the string is a palindrome
	for left, right := 0, len(str)-1; left < right; left, right = left+1, right-1 {
		if str[left] != str[right] {
			return false
		}
	}

	return true
}

// Two Pointers (1 Pass)
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsPalindrome2(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		// Skip non-alphanumeric characters
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		// Check if the characters are the same
		if !isSameChar(s[left], s[right]) {
			return false
		}
	}

	return true
}

func isAlphaNumeric(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func isSameChar(b1, b2 byte) bool {
	return unicode.ToLower(rune(b1)) == unicode.ToLower(rune(b2))
}
