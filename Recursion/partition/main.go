package partition

// Partition finds all possible ways to partition a string into palindromic substrings.
// It uses a backtracking approach to explore all possible partitions:
// 1. For each position, try all possible substrings starting from that position
// 2. If a substring is a palindrome, recursively partition the remaining string
// 3. Backtrack by removing the current substring and trying the next possibility
// Time complexity: O(N * 2^N) where N is the length of string
// Space complexity: O(N) for recursion depth and temporary storage
func Partition(s string) [][]string {
	result := [][]string{}
	partitionHelper(s, &result, []string{}, 0)
	return result
}

// partitionHelper performs the recursive backtracking to find all palindromic partitions
// s: input string to partition
// result: pointer to slice that accumulates all valid partitions
// currentPartition: current partition being built
// startIndex: starting index for the current substring exploration
func partitionHelper(s string, result *[][]string, currentPartition []string, startIndex int) {
	// Base case: if we've processed all characters, we have a complete partition
	if startIndex >= len(s) {
		// Create a copy of currentPartition to avoid reference issues
		partitionCopy := make([]string, len(currentPartition))
		copy(partitionCopy, currentPartition)
		*result = append(*result, partitionCopy)
		return
	}

	// Try all possible substrings starting from startIndex
	for endIndex := startIndex; endIndex < len(s); endIndex++ {
		// Extract substring from startIndex to endIndex (inclusive)
		substring := s[startIndex : endIndex+1]

		// If the substring is a palindrome, include it in the current partition
		if isPalindrome(substring) {
			// Add substring to current partition and recurse for remaining string
			partitionHelper(s, result, append(currentPartition, substring), endIndex+1)
			// Backtracking happens automatically when function returns
			// as currentPartition is passed by value (slice header is copied)
		}
	}
}

// isPalindrome checks if a given string reads the same forwards and backwards
// Uses two pointers approaching from both ends towards the center
func isPalindrome(s string) bool {
	// Compare characters from both ends moving towards center
	for leftPointer, rightPointer := 0, len(s)-1; leftPointer < rightPointer; leftPointer, rightPointer = leftPointer+1, rightPointer-1 {
		if s[leftPointer] != s[rightPointer] {
			return false
		}
	}
	return true
}
