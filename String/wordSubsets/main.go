package main

func WordSubsets(words1 []string, words2 []string) []string {
	// hashTable is used to store the maximum frequency of each letter in words2
	hashTable := make([]byte, 26)
	for i := 0; i < len(words2); i++ {
		freq := genFrequency(words2[i])

		// For each letter in the word, update the maximum frequency in hashTable
		// This is O(1) because the size of the slice is constant (26)
		for j := 0; j < 26; j++ {
			if freq[j] > hashTable[j] {
				hashTable[j] = freq[j]
			}
		}
	}

	// Loop through each word in words1
	ans := []string{}
	for i := 0; i < len(words1); i++ {
		// Generate the frequency of each letter in the word
		freq := genFrequency(words1[i])

		// Check if the frequency of each letter in the word is greater than or equal to the maximum frequency of each letter in words2
		flag := true
		for j := 0; j < 26; j++ {
			// If the frequency of a letter in the word is less than the maximum frequency of the letter in words2,
			// then the word is not a subset of words2
			// So we break the loop and check the next word
			if freq[j] < hashTable[j] {
				flag = false
				break
			}
		}

		// If the word is a subset of words2, add it to the answer
		if flag {
			ans = append(ans, words1[i])
		}
	}
	return ans
}

// genFrequency is a helper function to generate the frequency of each letter in a word
func genFrequency(word string) []byte {
	count := make([]byte, 26)
	for i := 0; i < len(word); i++ {
		count[word[i]-'a']++
	}
	return count
}
