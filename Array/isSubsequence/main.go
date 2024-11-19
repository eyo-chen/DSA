package main

// This is the brute force solution I came up with at first
// The idea is to iterate through the string s and for each character,
// iterate through the string t to find if the character exists
// if it does, we move to the next character in s and continue the search in t
// from the next character after the found character
// if we reach the end of s, we return true
// if we reach the end of t without finding all characters of s, we return false
// Time complexity: O(n * m) where n is the length of s and m is the length of t
// Space complexity: O(1)
// e.g. s = "abc", t = "xxxxxxxxx"
// for each character in s, we iterate through t to find if the character exists
// The total time we need to do this is O(len(s) * len(t))
func IsSubsequence(s string, t string) bool {
	// idx is used to track the position where we're gonna start the next search in t
	idx := 0

	// found is used to track the number of characters in s that we've found in t
	found := 0

	// iterate through each character in s
	for i := 0; i < len(s); i++ {
		// iterate through t to find the character
		for k := idx; k < len(t); k++ {
			// if we find the character, we increment the found counter and update the idx
			if s[i] == t[k] {
				found++

				// update the idx to the next character after the found character
				// so, for next iteration to find the next character in s, we'll start from the next character in t
				idx = k + 1
				break
			}
		}
	}

	return found == len(s)
}

// This is the optimized solution
// The idea is to use two pointers to iterate through the strings s and t
// We update the pointer of t in each iteration
// We move the pointer of s when we find the character in t
// If we reach the end of s, we return true
// If we reach the end of t without finding all characters of s, we return false
// Time complexity: O(m) where m is the length of t
// Space complexity: O(1)
// e.g. s = "abc", t = "xxxxxxxxx"
// At most, we iterate through t once
func IsSubsequence2(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	ptrS, ptrT := 0, 0

	for ptrS < len(s) && ptrT < len(t) {
		// if we find the character, we move the pointer of s
		if s[ptrS] == t[ptrT] {
			ptrS++
		}

		// always move the pointer of t
		ptrT++
	}

	return len(s) == ptrS
}
