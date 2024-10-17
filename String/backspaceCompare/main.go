package main

import (
	"bytes"
)

// Reading Backwards
func BackspaceCompare(s string, t string) bool {
	return remove(s) == remove(t)
}

func remove(s string) string {
	skip := 0
	var buf bytes.Buffer

	// Loop through the string from the end to the beginning
	for i := len(s) - 1; i >= 0; i-- {
		// If the current character is '#', increase the skip counter
		if s[i] == '#' {
			skip++
			continue
		}

		// If the skip counter is greater than 0, decrease the skip counter
		// It means the current character should be ignored
		if skip > 0 {
			skip--
			continue
		}

		// Add the current character to the buffer
		buf.WriteByte(s[i])
	}

	// We should reverse the string because we are iterating from the end to the beginning
	// However, because we ultimately are comparing two strings, it's not necessary
	// If two strings are the same after processing, the order does not matter
	return buf.String()
}

// Using Stack
func BackspaceCompare1(s string, t string) bool {
	sk := genStack(s)
	tk := genStack(t)

	// If the length of two stacks are not the same, return false
	if len(sk) != len(tk) {
		return false
	}

	// Compare the characters in the stack
	for i := 0; i < len(sk); i++ {
		if sk[i] != tk[i] {
			return false
		}
	}

	return true
}

func genStack(s string) []byte {
	stack := []byte{}

	// Loop through the string
	for i := 0; i < len(s); i++ {
		// If the current character is '#', pop the stack
		if s[i] == '#' {
			// Only pop the stack when the stack is NOT empty
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		// If the current character is not '#', push the character to the stack
		stack = append(stack, s[i])
	}

	return stack
}

// Two Pointers
func BackspaceCompare2(s string, t string) bool {
	ptrS := len(s) - 1
	ptrT := len(t) - 1

	// Loop until both pointers are less than 0
	for ptrS >= 0 || ptrT >= 0 {
		// If pointerS is still within the bound of the string
		// And the current character is '#'
		// We need to update the pointerS
		if ptrS >= 0 && s[ptrS] == '#' {
			// Initialize the skip counter
			// It means we have one character to ignore
			skip := 1

			// Move the pointer to the left
			// It means we are done considering the current character('#')
			ptrS--

			// While
			// (1) the pointer is still within the bound of the string
			// (2) the skip counter is greater than 0 OR the current character is '#'
			//     - skip > 0 means we have characters to ignore
			//     - s[ptrS] == '#' means there are still characters to ignore, so we need to go inside loop to increase the skip counter
			for ptrS >= 0 && (skip > 0 || s[ptrS] == '#') {
				// If the current character is '#'
				// We increase the skip counter
				// Otherwise, we decrease the skip counter
				if s[ptrS] == '#' {
					skip++
				} else {
					skip--
				}

				// Move the pointer to the left
				ptrS--
			}
		}

		// Same logic as above
		if ptrT >= 0 && t[ptrT] == '#' {
			skip := 1
			ptrT--

			for ptrT >= 0 && (skip > 0 || t[ptrT] == '#') {
				if t[ptrT] == '#' {
					skip++
				} else {
					skip--
				}
				ptrT--
			}
		}

		// If both pointers are less than 0, return true
		// It means both strings are done considering
		if ptrT < 0 && ptrS < 0 {
			return true
		}

		// If one of the pointers is less than 0, return false
		// It means one of the strings is done considering but the other string is not
		if ptrT < 0 || ptrS < 0 {
			return false
		}

		// If the current character is not the same, return false
		if s[ptrS] != t[ptrT] {
			return false
		}

		// Move both pointers to the left
		ptrS--
		ptrT--
	}

	return true
}
