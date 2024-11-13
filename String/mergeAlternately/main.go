package main

// This is the first solution that I came up with.
// The idea is the most straightforward to me, but the code is not the most elegant.
// The idea is to use two pointers to traverse the two words
// When the index is even, add the character from word1 to the result
// When the index is odd, add the character from word2 to the result
// Time complexity: O(n + m), where n is the length of word1 and m is the length of word2
// Space complexity: O(1)
func MergeAlternately(word1 string, word2 string) string {
	ptr1, ptr2 := 0, 0
	idx := 0
	ans := make([]byte, len(word1)+len(word2))

	for idx < len(ans) {
		// Note that this condition is a little bit tricky.
		// We need to first make sure that ptr1 is less than the length of word1, otherwise we will access the memory out of bounds.
		// Then we just check (1) ptr2 is out of bounds or (2) the index is even
		// In either case, we add the character from word1 to the result
		if ptr1 < len(word1) && (ptr2 >= len(word2) || idx%2 == 0) {
			ans[idx] = word1[ptr1]
			ptr1++
		} else {
			ans[idx] = word2[ptr2]
			ptr2++
		}

		idx++
	}

	return string(ans)
}

// This is the second solution that I referenced from the solution section.
// The main difference is that we only need to use one pointer to traverse the two words.
// Also, we both add the characters in word1 and word2 to the result slice at the same time in for loop.
// e.g. word1 = "abc", word2 = "pqr"
// At the first iteration, idx = 0, we first add word1[0] to the result, then add word2[0] to the result
// At the second iteration, idx = 1, we first add word1[1] to the result, then add word2[1] to the result
// So on and so forth.
// The only thing we need to check is that index is less than the length of word1 or word2.
// Time complexity: O(max(n, m)), where n is the length of word1 and m is the length of word2
// Space complexity: O(1)
// Time complexity is better than the first solution because we only traverse the longer word. In first solution, we traverse both short word and long word.
func MergeAlternately2(word1 string, word2 string) string {
	idx := 0
	ans := make([]byte, 0, len(word1)+len(word2))

	// Note that we use || here because we want to make sure that we traverse both words.
	// If we use && here, we will not add any character to the result if the index is out of bounds for one of the words.
	for idx < len(word1) || idx < len(word2) {
		if idx < len(word1) {
			ans = append(ans, word1[idx])
		}
		if idx < len(word2) {
			ans = append(ans, word2[idx])
		}
		idx++
	}

	return string(ans)
}
