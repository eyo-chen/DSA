package main

func CharacterReplacement(s string, k int) int {
	// Use to store the frequency of each character in the current window
	hashTable := make([]int, 26)
	ans := 0
	left, right := 0, 0

	for right < len(s) {
		// Add the current character of right pointer to the hash table
		char := s[right] - 'A'
		hashTable[char]++

		// Check if the current window is valid
		if isValid(right-left+1, k, hashTable) {
			// If it's valid, update the answer
			ans = max(ans, right-left+1)
		} else {
			// If it's not valid, move the left pointer to the right
			// Until the current window is valid
			// Or the left pointer is equal to right pointer
			for left <= right && !isValid(right-left+1, k, hashTable) {
				hashTable[s[left]-'A']--
				left++
			}
		}

		// Always move the right pointer to the right
		// Up to this point, we only have two scenarios:
		// 1. The current window is valid, and we've updated the answer. Let's update the right pointer to see if there's a longer valid window
		// 2. The current window is not valid, and we've moved the left pointer to the right until the current window is valid
		//    Or the left pointer is equal to right pointer
		//    In both cases, we know that the current window is either valid or left and right pointer are at the same position
		//    So, we can move the right pointer to the right
		right++
	}

	return ans
}

// isValid checks if the current window is valid
// The formula is
// (length of substring) - (largest frequency of character) <= k
func isValid(l int, k int, hashTable []int) bool {
	m := maxVal(hashTable)
	return l-m <= k
}

func maxVal(s []int) int {
	ans := s[0]

	for i := 1; i < len(s); i++ {
		ans = max(ans, s[i])
	}

	return ans
}

// Maintain maximum frequency in the while-loop process
func CharacterReplacement1(s string, k int) int {
	freq := map[byte]int{}
	maxFreq := 0
	ans := 0
	left, right := 0, 0

	for right < len(s) {
		freq[s[right]]++

		// update the maximum frequency
		maxFreq = max(maxFreq, freq[s[right]])

		// shrink the window size if it's not valid
		// Note that we use current length to validate(right - left + 1)
		for right-left+1-maxFreq > k && left < right {
			freq[s[left]]--
			left++
		}

		// update the answer
		ans = max(ans, right-left+1)
		right++
	}

	return ans
}
