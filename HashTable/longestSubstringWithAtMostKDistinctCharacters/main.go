package main

import "fmt"

// Brute force solution
func LongestSubstringWithAtMostKDistinctCharacters(s string, k int) int {
	ans := 0
	n := len(s)

	// Check all possible substrings
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			substring := s[i : j+1]
			if countDistinctChars(substring) <= k {
				ans = max(ans, len(substring))
			}
		}
	}

	return ans
}

// countDistinctChars counts the number of distinct characters in a string
func countDistinctChars(s string) int {
	charSet := make(map[rune]bool)
	for _, char := range s {
		charSet[char] = true
	}
	return len(charSet)
}

// Sliding window solution (delete key from hash table)
func LongestSubstringWithAtMostKDistinctCharacters2(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	left, right := 0, 0
	hashTable := make(map[byte]int)
	ans := 0

	for right < len(s) {
		// Add the right pointer character as distinct character to the hash table
		hashTable[s[right]]++

		// Widen the window(updating right pointer)
		right++

		// Shrink the window from the left if we have more than k distinct characters
		// When there's more than k distinct characters, we need to shrink the window from the left
		for len(hashTable) > k {
			hashTable[s[left]]--

			// If the character count reaches 0, delete it from the hash table
			if hashTable[s[left]] == 0 {
				delete(hashTable, s[left])
			}
			left++
		}

		ans = max(ans, right-left)
	}

	return ans
}

// Sliding window solution (without deleting key from hash table)
// It's the same as the previous solution, but without deleting key from hash table
// Instead, we use a variable to count the number of distinct characters
// When it's the first time we see a character, we increment the distinct count
// When the count of a character reaches 0, we decrement the distinct count
func LongestSubstringWithAtMostKDistinctCharacters3(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	left, right := 0, 0
	hashTable := make(map[byte]int)
	ans := 0
	distinctChars := 0

	for right < len(s) {
		hashTable[s[right]]++
		right++

		// When the character count is 1, it means it's a new distinct character
		if hashTable[s[right]] == 1 {
			distinctChars++
		}

		// When there's more than k distinct characters, we need to shrink the window from the left
		for distinctChars > k {
			hashTable[s[left]]--

			// When the character count reaches 0, it means it's no longer in the window
			if hashTable[s[left]] == 0 {
				distinctChars--
			}
			left++
		}

		ans = max(ans, right-left)
	}

	return ans
}

func main() {
	s := "coffee"
	s1 := s[1 : 3+1]
	fmt.Println(s1)
}
