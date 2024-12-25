package main

func LengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		// If the character is not a space, increment the count
		if s[i] != ' ' {
			count++
			continue
		}

		// If the character is a space and the count is not 0, return the count
		// Which means we have found the last word
		if count != 0 {
			return count
		}
	}

	return count
}
