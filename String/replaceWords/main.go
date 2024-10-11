package main

import (
	"sort"
	"strings"
)

// Compare each word with dictionary
func ReplaceWords(dictionary []string, sentence string) string {
	// Sort the dictionary by the length of the word
	// Make sure the shortest word is in the front
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	// Loop through each word in the sentence
	words := strings.Split(sentence, " ")
	for i, w := range words {
		// Loop through each word in the dictionary
		for _, d := range dictionary {
			// If the word in the sentence and the word in the dictionary match the prefix pattern
			// Replace the word in the sentence with the word in the dictionary
			if compare(d, w) {
				words[i] = d
				break
			}
		}
	}

	return strings.Join(words, " ")
}

// Compare two strings to see if they match the prefix pattern
// s1 is the dictionary word, s2 is the word in the sentence
func compare(s1, s2 string) bool {
	// If the length of the dictionary word is greater than the word in the sentence
	// Then they can't be the prefix pattern
	if len(s1) > len(s2) {
		return false
	}

	// Loop through each character in the dictionary word
	// We only need to loop through the length of the dictionary word
	// Because we only need to check if the dictionary word is the prefix of the word in the sentence
	for i := 0; i < len(s1); i++ {
		// If the character in the dictionary word is not the same as the character in the word in the sentence
		// Then they can't be the prefix pattern
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// Use hash table to store the dictionary words
func ReplaceWords1(dictionary []string, sentence string) string {
	// Create a hash table to store the dictionary words
	hashTable := make(map[string]bool, len(dictionary))
	for _, d := range dictionary {
		hashTable[d] = true
	}

	// Split the sentence into words
	words := strings.Split(sentence, " ")
	// Loop through each word in the sentence
	for idx, w := range words {
		// Loop through each character in the word
		for i := range w {
			// Get the prefix of the word
			// Check if the prefix is in the hash table
			target := w[:i+1]
			if _, ok := hashTable[target]; ok {
				// If it's in the hash table, we replace the word in the sentence with the word in the hash table
				words[idx] = target
				break
			}
		}
	}

	return strings.Join(words, " ")
}
