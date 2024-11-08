package main

import (
	"sort"
	"strconv"
	"strings"
)

// Brute Force
func GroupAnagrams(strs []string) [][]string {
	// hashTable to keep track of the strings that we already checked
	hashTable := make([]bool, len(strs))
	ans := [][]string{}

	// iterate through the input array
	for i, s := range strs {
		// if the string is already checked, skip it
		if hashTable[i] {
			continue
		}

		// mark the string as checked
		hashTable[i] = true
		group := []string{s}

		// iterate through the rest of the strings
		for k := i + 1; k < len(strs); k++ {
			// if the string is already checked, skip it
			if hashTable[k] {
				continue
			}

			// if the string is an anagram of the current string, add it to the current group
			// and mark it as checked
			if isAnagrams(s, strs[k]) {
				group = append(group, strs[k])
				hashTable[k] = true
			}
		}

		// add the current group to the result
		ans = append(ans, group)
	}

	return ans
}

func isAnagrams(str1, str2 string) bool {
	// if the length of the two strings are not the same, they are not anagrams
	if len(str1) != len(str2) {
		return false
	}

	frequency := make([]int, 26)

	// iterate both strings at the same time
	// one string is used to increase the frequency of each character
	// the other string is used to decrease the frequency of each character
	for i := 0; i < len(str1); i++ {
		frequency[str1[i]-'a']++
		frequency[str2[i]-'a']--
	}

	// after the loop, if all the frequencies are 0, the two strings are anagrams
	for _, f := range frequency {
		if f != 0 {
			return false
		}
	}

	return true
}

// Sorting
func GroupAnagrams2(strs []string) [][]string {
	// hashTable to store the groups of anagrams
	// key is the sorted string
	// value is the list of strings that are anagrams of each other
	hashTable := map[string][]string{}
	ans := [][]string{}

	// iterate through the input array
	for _, s := range strs {
		// sort the string (In Go, we need to convert the string to a rune array first to sort it)
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool {
			return r[i] > r[j]
		})

		// use the sorted string as the key
		ss := string(r)
		hashTable[ss] = append(hashTable[ss], s)
	}

	// iterate through the hash table to get the result
	for _, group := range hashTable {
		ans = append(ans, group)
	}

	return ans
}

// Frequency As Key
func GroupAnagrams3(strs []string) [][]string {
	// hashTable to store the groups of anagrams
	// key is the frequency array converted to a string
	// value is the list of strings that are anagrams of each other
	hashTable := map[string][]string{}
	ans := [][]string{}

	// iterate through the input array
	for _, s := range strs {
		// convert the string to a frequency array and then to a string
		key := genKey(s)
		hashTable[key] = append(hashTable[key], s)
	}

	// iterate through the hash table to get the result
	for _, group := range hashTable {
		ans = append(ans, group)
	}

	return ans
}

func genKey(str string) string {
	// frequency array to store the frequency of each character
	frequency := make([]int, 26)

	// iterate through the string to get the frequency of each character
	for i := 0; i < len(str); i++ {
		frequency[str[i]-'a']++
	}

	// convert the frequency array to a string
	builder := strings.Builder{}
	for _, f := range frequency {
		builder.WriteString(strconv.Itoa(f))
		builder.WriteString(",")
	}

	return builder.String()
}
