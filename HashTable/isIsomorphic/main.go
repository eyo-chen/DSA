package main

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashTableS := map[byte]byte{}
	hashTableT := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		if v, ok := hashTableS[s[i]]; ok && v != t[i] {
			return false
		}
		if v, ok := hashTableT[t[i]]; ok && v != s[i] {
			return false
		}

		hashTableS[s[i]] = t[i]
		hashTableT[t[i]] = s[i]
	}

	return true
}
