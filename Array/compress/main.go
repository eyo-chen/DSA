package main

import (
	"strconv"
)

func Compress(chars []byte) int {
	tmp := []byte{}

	for i := 0; i < len(chars); i++ {
		// Initialize the count for each character
		// Start from 1 because we count the current character
		count := 1

		// Count the number of consecutive characters (start from i + 1)
		for k := i + 1; k < len(chars); k++ {
			// If the current character is the same as the next one, we increment the count
			// Also, update the pointer i to the next one
			if chars[i] == chars[k] {
				count++
				i++
			} else {
				// If the current character is not the same as the next one, we break the loop
				break
			}
		}

		// Add the current character to the temporary array
		tmp = append(tmp, chars[i])

		// If the count is greater than 1, we add the count to the temporary array
		if count > 1 {
			// Convert the count to string
			countStr := strconv.Itoa(count)

			// Add each character in the count string to the temporary array
			for i := 0; i < len(countStr); i++ {
				tmp = append(tmp, countStr[i])
			}
		}
	}

	// Copy the temporary array to the original array
	copy(chars, tmp)

	return len(tmp)
}

func Compress2(chars []byte) int {
	// Initialize the pointer
	ptr := 0

	// Iterate through the array
	for i := 0; i < len(chars); i++ {
		// Initialize the count for each character
		// Start from 1 because we count the current character
		count := 1

		// Count the number of consecutive characters (start from i + 1)
		for k := i + 1; k < len(chars); k++ {
			// If the current character is the same as the next one, we increment the count
			// Also, update the pointer i to the next one
			if chars[i] == chars[k] {
				count++
				i++
			} else {
				break
			}
		}

		// Add the current character to the original array
		chars[ptr] = chars[i]
		ptr++

		// If the count is greater than 1, we add the count to the original array
		if count > 1 {
			countStr := strconv.Itoa(count)

			// Add each character in the count string to the original array
			for i := 0; i < len(countStr); i++ {
				chars[ptr] = countStr[i]
				ptr++
			}
		}
	}

	return ptr
}
