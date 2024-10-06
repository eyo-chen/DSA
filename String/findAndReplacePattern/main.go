package main

import (
	"math"
)

// Convert To Universal Pattern
func FindAndReplacePattern(words []string, pattern string) []string {
	// convert pattern to universal pattern(int)
	patternInt := convertToInt(pattern)

	// convert words to universal pattern(int)
	wordsInt := make([]int, len(words))
	for i, w := range words {
		wordsInt[i] = convertToInt(w)
	}

	// compare the patternInt with wordsInt
	// see if both patterns are the same
	ans := make([]string, 0)
	for i, w := range wordsInt {
		if w == patternInt {
			ans = append(ans, words[i])
		}
	}

	return ans
}

func convertToInt(s string) int {
	hashTable := map[byte]int{}
	ans := 0
	idx := 0
	count := 1

	for i := len(s) - 1; i >= 0; i-- {
		// if the character is already in the hash table,
		// we can use the value in the hash table to represent the character
		if v, ok := hashTable[s[i]]; ok {
			ans += int(math.Pow(10, float64(idx))) * v
		} else {
			// if the character is not in the hash table,
			// we need to assign a new value to the character
			ans += int(math.Pow(10, float64(idx))) * count
			hashTable[s[i]] = count
			count++
		}
		idx++
	}

	return ans
}

func FindAndReplacePattern1(words []string, pattern string) []string {
	ans := make([]string, 0)
	for _, w := range words {
		if checkPattern(w, pattern) {
			ans = append(ans, w)
		}
	}

	return ans
}

func checkPattern(str, pattern string) bool {
	if len(str) != len(pattern) {
		return false
	}

	// sToPattern is used to store the relationship from str(word) to pattern
	sToPattern := make([]byte, 256)
	// patternToS is used to store the relationship from pattern to str(word)
	patternToS := make([]byte, 256)

	for i := 0; i < len(str); i++ {
		s := str[i]
		p := pattern[i]

		// if the current character in str(word) is not mapped to a character in pattern,
		// we map it to the current character in pattern
		if sToPattern[s] == 0 {
			sToPattern[s] = p
		}
		// if the current character in pattern is not mapped to a character in str(word),
		// we map it to the current character in str(word)
		if patternToS[p] == 0 {
			patternToS[p] = s
		}

		// check if the relationship is consistent
		if sToPattern[s] != p || patternToS[p] != s {
			return false
		}
	}

	return true
}
