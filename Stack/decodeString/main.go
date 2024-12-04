package main

import (
	"strings"
	"unicode"
)

func DecodeString(s string) string {
	var numStack []int
	var strStack []string
	var currentNum int
	var currentStr string

	for _, char := range s {
		if unicode.IsDigit(char) {
			// Build number
			currentNum = currentNum*10 + int(char-'0')
		} else if char == '[' {
			// Push current number and string to respective stacks
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)
			// Reset current values
			currentNum = 0
			currentStr = ""
		} else if char == ']' {
			// Pop number and string from stacks
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// Repeat current string num times and append to previous string
			currentStr = prevStr + strings.Repeat(currentStr, num)
		} else {
			// Add character to current string
			currentStr += string(char)
		}
	}

	return currentStr
}
