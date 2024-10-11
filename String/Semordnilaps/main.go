package main

import (
	"bytes"
	"strings"
)

func Semordnilaps(strs []string) [][]string {
	// Use a hash table to store the strings we have seen
	hashTable := make(map[string]bool, len(strs))
	ans := [][]string{}

	// Loop through the strings
	for _, s := range strs {
		// Reverse the string
		r := reverse(s)

		// If the reversed string is in the hash table and it's not the same as the original string
		// And the original string is not in the hash table
		// Then we have found a semordnilap pair
		if hashTable[r] && r != s && !hashTable[s] {
			ans = append(ans, []string{r, s})
		}

		// Mark the original string as seen
		hashTable[s] = true
	}

	return ans
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// This method uses strings.Builder, which is more efficient for string concatenation. It preallocates the required capacity and builds the reversed string by appending characters from the end to the beginning.
func Reverse2(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

// This method uses bytes.Buffer, which is similar to strings.Builder but doesn't require preallocation of capacity. It's slightly less efficient than strings.Builder but still performs well.
func Reverse3(s string) string {
	var buf bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
	}
	return buf.String()
}

/*
Both of these methods are more efficient than the original implementation for larger strings because they avoid creating intermediate string allocations. The original method using runes is still good, especially when dealing with Unicode strings, as it correctly handles multi-byte characters.

Choose the method that best fits your use case:
- If you're dealing with ASCII strings or single-byte encodings, the strings.Builder or bytes.Buffer methods are more efficient.
- If you need to handle Unicode strings correctly, stick with the original rune-based method.
Remember that for very short strings, the performance difference between these methods is negligible, and readability should be prioritized.
*/
