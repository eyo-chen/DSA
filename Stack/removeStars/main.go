package main

// Using Stack
// Time Complexity: O(n)
// Space Complexity: O(n)
func RemoveStars(s string) string {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '*' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, c)
	}

	return string(stack)
}

// Using Two Pointers
// Time Complexity: O(n)
// Space Complexity: O(1)
func RemoveStars2(s string) string {
	// Convert the string to a slice of bytes
	result := []byte(s)
	index := 0

	// Iterate through the string
	for i := 0; i < len(s); i++ {
		c := s[i]

		// If the character is a star, decrement the index
		// It means we are removing the closest non-star character to its left
		// This works because later we will overwrite the index with the next character
		if c == '*' {
			index--
		} else {
			// If the character is not a star, overwrite the index with the character
			result[index] = c
			index++
		}
	}

	// Convert the slice back to a string up to the new length index
	return string(result[:index])
}
