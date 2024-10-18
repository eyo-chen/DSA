package main

func ValidPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Find the mismatch
		// Remove either one and check if the rest of the string is palindrome
		if s[left] != s[right] {
			// Remove left or right pointer
			// If either one is palindrome, return true
			// It means we can make the string a palindrome by removing at most one character
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}

		left++
		right--
	}

	return true
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// This is the solution I came up with
// It didn't pass the test case in leetcode
// But it did pass the test case I generated randomly
func ValidPalindrome1(s string) bool {
	left, right := 0, len(s)-1

	// Skip is used to check if we have removed a character
	skip := false

	for left < right {
		// If the characters are the same, move the pointers
		if s[left] == s[right] {
			left++
			right--
			continue
		}

		// If the characters are different and we have already removed a character
		if skip {
			return false
		}

		// After removing one character, if left and right pointer meet, it means we have found a palindrome
		// e.g. "bc", after removing 'c', "b" is a palindrome. After removing 'b', "c" is a palindrome
		if left+1 == right || right-1 == left {
			return true
		}

		// Try to remove left pointer and check if it matches right pointer
		// If yes, move the left pointer and set skip to true
		if s[left+1] == s[right] {
			left++
			skip = true
			continue
		}

		// Try to remove right pointer and check if it matches left pointer
		// If yes, move the right pointer and set skip to true
		if s[right-1] == s[left] {
			right--
			skip = true
			continue
		}

		// Up to this point, it means we can't move the pointers to find match
		// So we return false
		return false
	}

	return true
}
