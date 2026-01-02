package main

// Brute Force
func CheckInclusion(s1 string, s2 string) bool {
	// Count the frequency of each character in s1
	hashTable := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		hashTable[s1[i]-'a']++
	}

	// Loop through s2 string
	for i := 0; i < len(s2); i++ {
		// Count the frequency for each substring of s2
		freq := make([]int, 26)

		// Only need to loop through len(s1) times
		// Because the length of the substring has to be the same as s1
		// Also, k+i has to be less than len(s2)
		// So that we don't go out of bounds
		// k +i has to be within the length of s2, so that s2[k+i] is a valid character
		for k := 0; k < len(s1) && k+i < len(s2); k++ {
			freq[s2[k+i]-'a']++
		}

		// Check if the frequency of each character in the substring is the same as s1
		match := true
		for i := 0; i < 26; i++ {
			if hashTable[i] != freq[i] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

// Sliding Window
func CheckInclusion1(s1 string, s2 string) bool {
	// Create frequency hash table for s1
	freq1 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		freq1[s1[i]-'a']++
	}

	left, right := 0, 0
	// Create frequency hash table for s2
	freq2 := make([]int, 26)

	// Loop through s2 string
	for right < len(s2) {
		// Add the character of right pointer to freq2
		freq2[s2[right]-'a']++

		// If the window size is equal to s1's length
		// That means we find the full window
		if right-left+1 == len(s1) {
			// Check if freq2 is equal to freq1
			// At this point, freq2 represents the frequency of characters in the current window
			findMatch := true
			for i := 0; i < 26; i++ {
				if freq1[i] != freq2[i] {
					findMatch = false
					break
				}
			}

			// If we find a match, return true
			if findMatch {
				return true
			}
		}

		// If the window size is equal to s1's length
		// That means we have a full window, and the current window is not a permutation of s1
		// Update window by moving left pointer and right pointer to right
		// Decrement the frequency of the character of left pointer in freq2
		if right-left+1 == len(s1) {
			freq2[s2[left]-'a']--
			left++
			right++
		} else {
			// If the window size is not equal to s1's length
			// That means we haven't get a full window, move right pointer to right
			right++
		}
	}

	return false
}

// Sliding Window Optimized
func CheckInclusion2(s1 string, s2 string) bool {
	// Create a hash table to store the frequency of each character in s1
	hashTable := make(map[byte]int, 26)
	for i := 0; i < len(s1); i++ {
		hashTable[s1[i]-'a']++
	}

	left, right := 0, 0
	matches := 0
	for right < len(s2) {
		rightChar := s2[right] - 'a'

		// If the character is in s1, that means the current character is in s1
		if _, ok := hashTable[rightChar]; ok {
			hashTable[rightChar]--

			// If the frequency is still equal to or greater than 0, that means the character is in s1
			// And the frequency is still within the limit
			// For example, hashTable = {a: 1}, after decrement, it will be {a: 0}
			// That means there are still one 'a' in s1 we can match
			// Next time we find 'a' again, the frequency will be -1
			// That means we can't use this 'a' to match
			if hashTable[rightChar] >= 0 {
				matches++
			}
		}

		if matches == len(s1) {
			return true
		}

		// It means we have a full window, but it's not a permutation of s1
		// We need to move the left pointer to right to find the next window
		if right-left+1 == len(s1) {
			leftChar := s2[left] - 'a'

			// Check if the character is in s1
			if _, ok := hashTable[leftChar]; ok {
				hashTable[leftChar]++

				// We only want to decrement the matches when the current character is in s1
				// And it's the character we added before
				// How can we make sure that?
				// First, we know that the current character is in s1
				// Second, we know that the current character is the character we added before
				// Because if we haven't added it before, the frequency of the character in hashTable will be at least negative 1
				// Why?
				// Because we always first decrement the frequency of the character in hashTable in the above if statement
				// So, if the frequency is greater than 0, that means the frequency is at least 0
				// then, we increment the frequency before
				if hashTable[leftChar] > 0 {
					matches--
				}
			}

			left++
		}

		right++
	}

	return false
}

// Updated at 0128 2025
// Brute Force
func CheckInclusion3(s1 string, s2 string) bool {
	hashTable1 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		hashTable1[s1[i]-'a']++
	}

	// Only need to loop through len(s2) - len(s1) times
	for i := 0; i <= len(s2)-len(s1); i++ {
		hashTable2 := make([]int, 26)

		for k := 0; k < len(s1); k++ {
			hashTable2[s2[i+k]-'a']++
		}

		match := true
		for k := 0; k < 26; k++ {
			if hashTable1[k] != hashTable2[k] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

// Updated at 0128 2025
// Sliding Window
func CheckInclusion4(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	left, right := 0, 0
	hashTable1 := make([]int, 26)
	hashTable2 := make([]int, 26)

	// Init the hash table for s1
	// Also, init the hash table for s2 and the window size
	for i := 0; i < len(s1); i++ {
		hashTable1[s1[i]-'a']++
		hashTable2[s2[i]-'a']++
		right++
	}

	// Check if the first window is a permutation of s1
	if match(hashTable1, hashTable2) {
		return true
	}

	// Loop through s2
	for right < len(s2) {
		// Add the character of right pointer to hashTable2
		hashTable2[s2[right]-'a']++

		// Remove the character of left pointer from hashTable2
		hashTable2[s2[left]-'a']--

		// Move the left pointer to right
		left++
		right++

		// Check if the current window is a permutation of s1
		if match(hashTable1, hashTable2) {
			return true
		}
	}

	return false
}

func match(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CheckInclusion5 checks if s2 contains a permutation of s1.
//
// Approach: Uses a sliding window technique with frequency counting.
// We maintain a frequency map for s1 and slide a window of size len(s1)
// across s2, comparing character frequencies at each position.
//
// Time Complexity: O(n + m) where n = len(s1), m = len(s2)
// - O(n) to build frequency map for s1
// - O(m) to slide window across s2
// - O(26) for each frequency comparison (constant)
//
// Space Complexity: O(1) - uses two fixed-size arrays of 26 elements
func CheckInclusion5(s1 string, s2 string) bool {
	// Early exit: s1 can't be a permutation substring if it's longer than s2
	if len(s1) > len(s2) {
		return false
	}

	// Build frequency map for s1 (the pattern we're searching for)
	s1Freq := make([]int, 26)
	for _, char := range s1 {
		s1Freq[char-'a']++
	}

	// Build frequency map for the first window in s2
	windowFreq := make([]int, 26)
	for i := range len(s1) {
		windowFreq[s2[i]-'a']++
	}

	// Check if the first window matches
	if hasMatchingFrequencies(s1Freq, windowFreq) {
		return true
	}

	// Slide the window across s2, one character at a time
	// !!! Note that this loop is very important !!!
	/*
		We start the index from len(s1) because we have processed the first window (0~len(s1)-1) already.
		For example, if s1 = "ab" and s2 = "eidbaooo", the above for loop processes the first two characters "ei".
		The length of s1 is 2, so the next character to process is at index 2 (the third character in s2).
		Then, we add the character at index 2 ('d') to the window and remove the character at index 0 ('e').
		This way, we maintain a sliding window of size len(s1) as we move through s2.
	*/
	for i := len(s1); i < len(s2); i++ {
		// Add the new character entering the window
		newChar := s2[i] - 'a'
		windowFreq[newChar]++

		// Remove the old character leaving the window
		oldChar := s2[i-len(s1)] - 'a'
		windowFreq[oldChar]--

		// Check if current window matches s1's frequency
		if hasMatchingFrequencies(s1Freq, windowFreq) {
			return true
		}
	}

	// No matching permutation found
	return false
}

// hasMatchingFrequencies compares two frequency arrays for equality.
// Returns true if both arrays have identical values at each index.
func hasMatchingFrequencies(freq1 []int, freq2 []int) bool {
	for i, count := range freq1 {
		if freq2[i] != count {
			return false
		}
	}
	return true
}
