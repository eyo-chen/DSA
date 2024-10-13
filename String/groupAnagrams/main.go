package main

import (
	"strconv"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	ans := [][]string{}

	// use a hash table to store the words that have been seen
	// one a word is added to a group, we mark it as seen
	seen := map[string]bool{}

	// iterate through the array
	for i := 0; i < len(strs); i++ {
		word1 := strs[i]

		// if the word has been seen, skip it
		if _, ok := seen[word1]; ok {
			continue
		}

		// group is the current group of anagrams
		group := []string{word1}
		// mark the word as seen
		seen[word1] = true

		// iterate through the rest of the array
		for k := i + 1; k < len(strs); k++ {
			word2 := strs[k]
			// if the word is an anagram, add it to the current group
			if isAnagrams(word1, word2) {
				group = append(group, word2)
				// after adding the word to the group, mark it as seen
				// because we don't need to consider it again
				seen[word2] = true
			}
		}

		// add the current group to the answer array
		ans = append(ans, group)
	}

	return ans
}

// isAnagrams checks if two words are anagrams
func isAnagrams(s1, s2 string) bool {
	// if two words are not the same length, they are not anagrams
	if len(s1) != len(s2) {
		return false
	}

	// generate the frequency count array for both words
	freq1 := genFreq(s1)
	freq2 := genFreq(s2)

	// compare the frequency count array
	// if any element is different, the two words are not anagrams
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	// if all elements are the same, the two words are anagrams
	return true
}

// genFreq generates the frequency count array for a word
func genFreq(s string) []int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	return freq
}

func GroupAnagrams1(strs []string) [][]string {
	// use a hash table to store the key and the group of anagrams
	// key is the frequency count array converted to a string
	// value is the group of anagrams
	hashTable := map[string][]string{}
	ans := [][]string{}

	// iterate through the array
	for i := 0; i < len(strs); i++ {
		word := strs[i]
		// generate the frequency count key for the word
		// e.g. "eat" -> "1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0"
		key := genFreqKey(word)

		// add the word to the group that is associated with the key
		// we don't need to use if condition to check if the key exists
		// If the key doesn't exist, then we just simply append the word to an empty slice
		// If the key exists, then we append the word to the slice
		hashTable[key] = append(hashTable[key], word)
	}

	// iterate through the hash table and add the groups to the answer array
	for _, val := range hashTable {
		ans = append(ans, val)
	}

	return ans
}

// genFreqKey generates the frequency count key for a word
func genFreqKey(s string) string {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[int(s[i]-'a')]++
	}
	return join(freq)
}

// join joins the frequency count array to a string
func join(arr []int) string {
	strBuilder := strings.Builder{}
	for _, val := range arr {
		strBuilder.WriteString(strconv.Itoa(val))

		// it's important to add "," to differentiate the difference between "1,0,0,0" and "10,0,0"
		strBuilder.WriteString(",")
	}
	return strBuilder.String()
}
