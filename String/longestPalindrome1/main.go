package main

// Brute Force
// Time Complexity: O(n^3)
// Space Complexity: O(1)
func LongestPalindrome(s string) string {
	ans := ""
	maxLen := 0

	// Check all the possible substrings
	for i := 0; i < len(s); i++ {
		for k := i; k < len(s); k++ {
			// Get the current substring
			curStr := s[i : k+1]

			// Check if the current substring is palindrome
			if len(curStr) > maxLen && isPalindrome(curStr) {
				// Update the longest palindrome substring
				ans = curStr
				maxLen = len(curStr)
			}
		}
	}

	return ans
}

// Check if the given string is palindrome
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// Expand Around Center
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func LongestPalindrome2(s string) string {
	start, maxLen := 0, 0

	for i := 0; i < len(s); i++ {
		// Check the string from the center to both sides (outward)
		// Case 1: The length of the string is odd
		left, right := i, i

		// Check if the characters before and after are the same
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// Update the longest palindrome substring
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}

		// Case 2: The length of the string is even
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// Update the longest palindrome substring
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	// Return the longest palindrome substring
	return s[start : start+maxLen]
}
