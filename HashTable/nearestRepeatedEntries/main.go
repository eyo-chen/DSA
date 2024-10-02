package main

func NearestRepeatedEntries(sentence []string) int {
	hashTable := make(map[string]int, len(sentence))

	// set the answer to the length of the sentence(max possible distance)
	ans := len(sentence)

	// iterate through the sentence
	for i, s := range sentence {
		// if the word is already in the hash table
		if idx, ok := hashTable[s]; ok {
			// calculate the distance between the current index and the index of the word in the hash table
			ans = min(ans, i-idx)
		}

		// set the index of the word in the hash table
		hashTable[s] = i
	}

	// if no word is repeated, return -1
	if ans == len(sentence) {
		return -1
	}

	return ans
}
