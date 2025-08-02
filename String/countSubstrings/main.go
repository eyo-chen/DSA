package main

// It's basically the same problem as String/longestPalindrome1/main.go
func CountSubstrings(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			ans++
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			ans++
			left--
			right++
		}
	}

	return ans
}
