package main

func CanConstruct(ransomNote string, magazine string) bool {
	hashTable := map[string]int{}

	for _, s := range magazine {
		hashTable[string(s)]++
	}

	for _, s := range ransomNote {
		key := string(s)
		// return false if the character is not in the hash table
		// or the character is 0, which means there's no enough characters in the magazine
		if v, ok := hashTable[key]; !ok || v == 0 {
			return false
		}
		hashTable[key]--
	}

	return true
}

func CanConstruct1(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	table := make([]int, 26)

	for _, s := range magazine {
		table[s-'a']++
	}

	for _, s := range ransomNote {
		key := s - 'a'
		if table[key] == 0 {
			return false
		}
		table[key]--
	}

	return true
}
