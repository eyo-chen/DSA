package main

func LongestPalindrome(s string) int {
	hashTable := map[rune]int{}

	// count the number of each character
	for _, b := range s {
		hashTable[b]++
	}

	ans := 0
	hasOdd := false

	// calculate the longest palindrome
	for _, value := range hashTable {
		// find the odd count character
		if value%2 == 1 {
			hasOdd = true
		}

		// if the count is greater than or equal to 2
		// add both even and odd count to the result
		if value >= 2 {
			ans += (value / 2) * 2
		}
	}

	// if there is any odd count character
	// add 1 to the result
	if hasOdd {
		ans++
	}

	return ans
}

func LongestPalindrome2(s string) int {
	hashSet := map[rune]bool{}
	ans := 0

	for _, r := range s {
		// if the character is already in the hash set
		// we find a pair to form a palindrome
		if _, ok := hashSet[r]; ok {
			ans += 2
			delete(hashSet, r)
		} else {
			// if the character is not in the hash set
			// we add the character to the hash set
			hashSet[r] = true
		}
	}

	// if there is any character in the hash set
	// we add 1 to the result
	if len(hashSet) > 0 {
		ans++
	}

	return ans
}
