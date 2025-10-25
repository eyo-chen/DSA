package main

// node represents a single character position in the trie.
// Each node contains pointers to its children and a flag indicating word completion.
type node struct {
	children  [26]*node // Array of 26 pointers for lowercase letters 'a'-'z'
	isWordEnd bool      // True if this node marks the end of a valid word
}

// Trie is a prefix tree data structure for efficient string storage and retrieval.
type Trie struct {
	root *node
}

// Constructor initializes and returns a new empty Trie.
// Time Complexity: O(1)
// Space Complexity: O(1)
func Constructor() Trie {
	return Trie{root: &node{}}
}

// Insert adds a word to the trie.
// Approach: Traverse the trie character by character, creating new nodes as needed.
// Mark the final node to indicate the end of a valid word.
// Time Complexity: O(m) where m is the length of the word
// Space Complexity: O(m) in worst case when all characters need new nodes
func (t *Trie) Insert(word string) {
	currentNode := t.root

	// Traverse through each character in the word
	for i := 0; i < len(word); i++ {
		char := word[i]
		// Convert character to array index: 'a' -> 0, 'b' -> 1, etc.
		index := char - 'a'

		// Create a new node if the path doesn't exist
		if currentNode.children[index] == nil {
			currentNode.children[index] = &node{}
		}

		// Move to the next node in the path
		currentNode = currentNode.children[index]
	}

	// Mark the end of the word
	currentNode.isWordEnd = true
}

// Search returns true if the word exists in the trie.
// Approach: Follow the path corresponding to each character. The word exists only if
// the complete path exists AND the final node is marked as a word ending.
// Time Complexity: O(m) where m is the length of the word
// Space Complexity: O(1)
func (t *Trie) Search(word string) bool {
	currentNode := t.root

	// Traverse through each character in the word
	for i := 0; i < len(word); i++ {
		char := word[i]
		// Convert character to array index
		index := char - 'a'

		// If the path doesn't exist, the word is not in the trie
		if currentNode.children[index] == nil {
			return false
		}

		// Move to the next node in the path
		currentNode = currentNode.children[index]
	}

	// Word exists only if we reached a node marked as word ending
	return currentNode.isWordEnd
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
// Approach: Follow the path corresponding to each character. The prefix exists if
// the complete path exists (we don't check the isWordEnd flag).
// Time Complexity: O(m) where m is the length of the prefix
// Space Complexity: O(1)
func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root

	// Traverse through each character in the prefix
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		// Convert character to array index
		index := char - 'a'

		// If the path doesn't exist, no word has this prefix
		if currentNode.children[index] == nil {
			return false
		}

		// Move to the next node in the path
		currentNode = currentNode.children[index]
	}

	// Prefix exists if we successfully traversed all characters
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
