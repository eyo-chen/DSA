package main

import (
	"slices"
)

// LadderLengthBruteForce finds the shortest transformation sequence using brute force approach.
// For each word in the current level, it checks all words in wordList to find valid neighbors.
//
// Approach: Level-by-level BFS where we check every word in wordList for single character difference
// Time Complexity: O(N² × M) where N = number of words in wordList, M = length of each word
// Space Complexity: O(N) for visited map and queue storage
func LadderLengthBruteForce(beginWord string, endWord string, wordList []string) int {
	// Early return if target word doesn't exist in dictionary
	if !slices.Contains(wordList, endWord) {
		return 0
	}

	currentLevel := 1
	visitedWords := map[string]bool{}
	wordsToProcess := []string{beginWord}

	// Process words level by level using BFS
	for len(wordsToProcess) > 0 {
		currentLevelSize := len(wordsToProcess)

		// Process all words at current level before moving to next level
		for range currentLevelSize {
			currentWord := wordsToProcess[0]
			wordsToProcess = wordsToProcess[1:]

			// Check if we reached the target word
			if currentWord == endWord {
				return currentLevel
			}

			// Find all valid neighbors from wordList
			for _, candidateWord := range wordList {
				// Skip if already visited
				if visitedWords[candidateWord] {
					continue
				}

				// Check if candidate word differs by exactly one character
				if hasSingleCharDifference(currentWord, candidateWord) {
					visitedWords[candidateWord] = true
					wordsToProcess = append(wordsToProcess, candidateWord)
				}
			}
		}
		currentLevel++
	}

	return 0 // No transformation sequence found
}

// hasSingleCharDifference checks if two words differ by exactly one character.
// Returns true if words have same length and exactly one character difference.
func hasSingleCharDifference(word1, word2 string) bool {
	// Words must have same length for valid transformation
	if len(word1) != len(word2) {
		return false
	}

	differenceCount := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			differenceCount++
			// Early return if more than one difference found
			if differenceCount > 1 {
				return false
			}
		}
	}

	return differenceCount == 1
}

// QueueItem represents a word and its level in the BFS traversal
type QueueItem struct {
	word  string
	level int
}

// LadderLengthOptimized finds the shortest transformation sequence using character replacement.
// For each word, it generates all possible single-character variations and checks if they exist in wordSet.
//
// Approach: BFS with character replacement to generate neighbors efficiently
// Time Complexity: O(N × M² × 26) = O(N × M²) where N = number of words, M = word length
// Space Complexity: O(N) for wordSet, visited map and queue
func LadderLengthOptimized(beginWord string, endWord string, wordList []string) int {
	// Convert wordList to hash set for O(1) lookup
	wordSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordSet[word] = true
	}

	// Early return if target word doesn't exist in dictionary
	if !wordSet[endWord] {
		return 0
	}

	// Initialize BFS data structures
	visitedWords := make(map[string]bool, len(wordList)+1)
	queue := []QueueItem{{word: beginWord, level: 1}}
	visitedWords[beginWord] = true

	// Process queue using BFS
	for len(queue) > 0 {
		currentItem := queue[0]
		queue = queue[1:]

		currentWord, currentLevel := currentItem.word, currentItem.level

		// Check if we reached the target word
		if currentWord == endWord {
			return currentLevel
		}

		// Generate all possible single-character variations
		wordChars := []byte(currentWord)
		for charIndex := range wordChars {
			originalChar := wordChars[charIndex]

			// Try replacing current character with all 26 lowercase letters
			for newChar := byte('a'); newChar <= byte('z'); newChar++ {
				// Skip if same character
				if newChar == originalChar {
					continue
				}

				wordChars[charIndex] = newChar
				neighborWord := string(wordChars)

				// If neighbor exists in dictionary and not visited, add to queue
				if wordSet[neighborWord] && !visitedWords[neighborWord] {
					visitedWords[neighborWord] = true
					queue = append(queue, QueueItem{word: neighborWord, level: currentLevel + 1})
				}
			}

			// Restore original character for next iteration
			wordChars[charIndex] = originalChar
		}
	}

	return 0 // No transformation sequence found
}
