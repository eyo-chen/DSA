package main

import (
	"strconv"
	"strings"
)

// Use Two Hash Table And Generate Key
func FindAnagrams(s string, p string) []int {
	ans := []int{}

	// Build the hash table and the key for the string p
	hashTableP := make([]int, 26)
	for i := 0; i < len(p); i++ {
		hashTableP[p[i]-'a']++
	}
	key := genKey(hashTableP)

	// Build the hash table and the key for the current window of the string s
	hashTableS := make([]int, 26)
	left, right := 0, 0
	for right < len(s) {
		// Add the new character to the hash table
		hashTableS[s[right]-'a']++

		// If we have a window, we generate the key for the current window
		if right-left+1 == len(p) {
			curKey := genKey(hashTableS)

			// If the key of the current window is the same as the key of the pattern string,
			// we add the start index of the current window to the result
			if curKey == key {
				ans = append(ans, left)
			}

			// Remove the leftmost character from the hash table
			hashTableS[s[left]-'a']--
			left++
		}

		// Move the right pointer to the next character
		right++
	}

	return ans
}

func genKey(arr []int) string {
	sb := strings.Builder{}
	for _, v := range arr {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString("-")
	}

	return sb.String()
}

// Use Two Hash Table And Directly Compare
func FindAnagrams2(s string, p string) []int {
	ans := []int{}
	if len(s) < len(p) {
		return ans
	}

	// Build the hash table for the pattern string p
	hashTableP := make([]int, 26)
	for i := 0; i < len(p); i++ {
		hashTableP[p[i]-'a']++
	}

	// Build the hash table for the current window of the string s
	// Note that we loop through the length of the pattern string p
	// because we need to build the hash table for the first window
	hashTableS := make([]int, 26)
	for i := 0; i < len(p); i++ {
		// Add the new character of string s to the hash table
		hashTableS[s[i]-'a']++
	}

	// If the hash table of the current window is the same as the hash table of the pattern string,
	// we add the start index of the current window to the result
	if arrayEqual(hashTableP, hashTableS) {
		ans = append(ans, 0)
	}

	// Keep Moving The Window
	// Note that we start from the length of the pattern string p
	// because we have already built the hash table for the first window
	for i := len(p); i < len(s); i++ {
		// Add the new character of string s to the hash table
		hashTableS[s[i]-'a']++

		// Remove the leftmost character from the hash table
		hashTableS[s[i-len(p)]-'a']--

		// If the hash table of the current window is the same as the hash table of the pattern string,
		// we add the start index of the current window to the result
		if arrayEqual(hashTableP, hashTableS) {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}

func arrayEqual(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// Use Match Count
func FindAnagrams3(s string, p string) []int {
	ans := []int{}

	if len(s) < len(p) {
		return ans
	}

	// Build the hash table for the pattern string p
	hashTable := make([]int, 26)
	for i := 0; i < len(p); i++ {
		hashTable[p[i]-'a']++
	}

	// Build the match count for the first window
	// Note that we loop through the length of the pattern string p
	// because we need to build the match count for the first window
	// match variable is used to count how many characters in the current window are correctly matched
	match := 0
	for i := 0; i < len(p); i++ {
		key := s[i] - 'a'

		// When adding a character to the current window, we decrease the count of the character in the hash table
		hashTable[key]--

		// If the count is greater than 0, it means the character is in the pattern string, so we increase the match count
		if hashTable[key] >= 0 {
			match++
		}
	}

	// If the match count is equal to the length of the pattern string,
	// we add the start index of the current window to the result
	if match == len(p) {
		ans = append(ans, 0)
	}

	// Keep Moving The Window
	for i := len(p); i < len(s); i++ {
		leftKey := s[i-len(p)] - 'a'

		// When removing a character from the current window, we increase the count of the character in the hash table
		hashTable[leftKey]++

		// If the count is greater than 0, it means the character is in the pattern string, so we decrease the match count
		if hashTable[leftKey] > 0 {
			match--
		}

		// When adding a character to the current window, we decrease the count of the character in the hash table
		rightKey := s[i] - 'a'
		hashTable[rightKey]--

		// If the count is greater than 0, it means the character is in the pattern string, so we increase the match count
		if hashTable[rightKey] >= 0 {
			match++
		}

		// If the match count is equal to the length of the pattern string,
		// we add the start index of the current window to the result
		if match == len(p) {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}
