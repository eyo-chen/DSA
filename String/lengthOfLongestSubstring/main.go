package main

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	ans := 1
	for i := 0; i < len(s); i++ {
		hashTable := map[byte]bool{
			s[i]: true,
		}

		for k := i + 1; k < len(s); k++ {
			if _, ok := hashTable[s[k]]; ok {
				// !!!!WRONG!!!!
				// "au"
				// ans = max(ans, k-i)
				break
			}

			ans = max(ans, k-i+1)
			hashTable[s[k]] = true
		}
	}

	return ans
}

func LengthOfLongestSubstring1(s string) int {
	ans := 0
	left, right := 0, 0
	hashTable := make(map[byte]bool, len(s))

	for right < len(s) {
		// take the character at the right pointer
		char := s[right]

		// if current character is already in the hash table, we need to move the left pointer to the right until there's no duplicate character in the current substring
		for hashTable[char] {
			leftChar := s[left]
			hashTable[leftChar] = false
			left++
		}

		// mark the current character as true in the hash table
		hashTable[char] = true

		// move the right pointer to the right
		right++

		// update the answer with the maximum length of the current substring
		ans = max(ans, right-left)
	}

	return ans
}

func LengthOfLongestSubstring2(s string) int {
	ans := 0
	left, right := 0, 0
	hashTable := make(map[byte]int, len(s))

	for right < len(s) {
		char := s[right]

		// if the character is already in the hash table, we need to move the left pointer to the right until there's no duplicate character in the current substring
		if val, ok := hashTable[char]; ok {
			left = max(left, val+1)
		}

		hashTable[char] = right
		right++
		ans = max(ans, right-left)
	}

	return ans
}
