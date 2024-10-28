package main

import (
	"fmt"
	"sort"
)

// Using hash table
// Time complexity: O(n) where n is the length of the string
// Space complexity: O(n)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashTable := make([]int, 26)
	for i := 0; i < len(s); i++ {
		hashTable[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		c := t[i] - 'a'
		if hashTable[c] == 0 {
			return false
		}

		hashTable[c]--
	}

	return true
}

// Using sort
// Time complexity: O(n log n) where n is the length of the string
// Space complexity: O(n) or O(1) depends on the sorting algorithm
// Note that string is immutable in Go, so we need to convert it to rune slice first, which is less efficient
// This solution is also useful for handling Unicode characters
func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert strings to rune slices
	sRunes := []rune(s)
	tRunes := []rune(t)

	// Sort both slices
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	// Compare the sorted strings
	return string(sRunes) == string(tRunes)
}

// IsAnagram3 handles Unicode characters
// Time complexity: O(n) where n is the length of the string
// Space complexity: O(n)
func IsAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Use map instead of fixed-size array since Unicode has many more characters
	charCount := make(map[rune]int)

	// Count characters in first string
	for _, ch := range s {
		charCount[ch]++
	}

	// Verify characters in second string
	for _, ch := range t {
		if charCount[ch] == 0 {
			return false
		}
		charCount[ch]--
	}

	return true
}

// Note that use IsAnagram2 and IsAnagram3 to handle Unicode characters
// IsAnagram only works for ASCII characters
// If we try to run IsAnagram("你好世界", "世界你好") it will cause a runtime error
// Because IsAnagram is using fixed-size array to store the character count
// It will only work for ASCII characters (0-127)
// Unicode characters have more than 127
func main() {
	fmt.Println(IsAnagram2("你好世界", "世界你好")) // true
	fmt.Println(IsAnagram3("你好世界", "世界你好")) // true

	fmt.Println("--------------------------------")

	fmt.Println(IsAnagram2("🌞🌝🌚", "🌚🌞🌝")) // true
	fmt.Println(IsAnagram3("🌞🌝🌚", "🌚🌞🌝")) // true

	fmt.Println("--------------------------------")

	fmt.Println(IsAnagram2("hello世界", "世界hello")) // true
	fmt.Println(IsAnagram3("hello世界", "世界hello")) // true

	fmt.Println("--------------------------------")

	fmt.Println(IsAnagram2("über", "reüb")) // true
	fmt.Println(IsAnagram3("über", "reüb")) // true
}
