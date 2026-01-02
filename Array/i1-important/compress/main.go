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

// Updated at 2025-01-12
func Compress2(chars []byte) int {
	tmp := []byte{}
	ptr := 0

	for ptr < len(chars) {
		char := chars[ptr]
		tmp = append(tmp, char)
		count := 0

		for ptr < len(chars) && chars[ptr] == char {
			ptr++
			count++
		}

		if count > 1 {
			strCount := strconv.Itoa(count)
			for s := 0; s < len(strCount); s++ {
				tmp = append(tmp, strCount[s])
			}
		}
	}

	copy(chars, tmp)
	return len(tmp)
}

func Compress3(chars []byte) int {
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

// Updated at 2025-01-12
/*
1. init write and read at the beginning(0)
2. keep updating read until it's out of the bound
   (1) init count(to count the number of consecutive characters)
   (2) access current char(for comparison and update)
   (3) keep updating read until it's out of the bound OR the value is different
   (4) update chars[write] = char, write++
   (5) update chars[write] = count(handle the case when count is greater than 10), write++
*/
func Compress4(chars []byte) int {
	read, write := 0, 0

	for read < len(chars) {
		count := 0
		char := chars[read]

		for read < len(chars) && chars[read] == char {
			count++
			read++
		}

		chars[write] = char
		write++

		if count > 1 {
			countStr := strconv.Itoa(count)
			for s := 0; s < len(countStr); s++ {
				chars[write] = countStr[s]
				write++
			}
		}
	}

	return write
}
