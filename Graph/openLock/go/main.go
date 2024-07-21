package main

import (
	"slices"
	"strconv"
)

func OpenLock(deadends []string, target string) int {
	if slices.Contains(deadends, "0000") {
		return -1
	}

	// Create a hash table to store deadends
	// This is to check if a digit is a deadend in O(1) time
	deadendsHashTable := map[string]bool{}
	for _, deadend := range deadends {
		deadendsHashTable[deadend] = true
	}

	q := []string{"0000"}
	seenHashTable := map[string]bool{
		"0000": true,
	}
	steps := 0

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			digit := q[0]
			q = q[1:]

			if digit == target {
				return steps
			}

			// Generate the incremented and decremented values for each digit
			nextDigits := GenNextDigits(digit)
			for _, nextDigit := range nextDigits {
				// Check if the incremented and decremented values are not in the deadends
				// Check if the incremented and decremented values are not seen before
				if !deadendsHashTable[nextDigit] && !seenHashTable[nextDigit] {
					q = append(q, nextDigit)
					seenHashTable[nextDigit] = true
				}
			}
		}

		steps++
	}

	return -1
}

// GenNextDigits is a function to generate the incremented and decremented values of a digit
func GenNextDigits(digit string) []string {
	nextDigits := []string{}

	// for each digit, generate the incremented and decremented values
	for i := 0; i < 4; i++ {
		inDigit, deDigit := GenInAndDe(digit, i)
		nextDigits = append(nextDigits, inDigit, deDigit)
	}

	return nextDigits
}

// GenInAndDe is a function to generate the increment and decrement values of a digit
// It using string concatenation to create new strings with the modified digits
func GenInAndDe(digit string, index int) (string, string) {
	/*
		byte is an alias for uint8
		In other words, byte is essentially an number between 0 and 255
		Also, it can represent a character in ASCII
		Because '0' is 48 in ASCII, we can subtract '0' from a byte to get the integer value of the digit
		e.g. '2' has an ASCII value of 50. 50 - 48 = 2
	*/
	originalDigit := digit[index] - '0'

	/*
		Calculate incremented and decremented values with wrapping
		If originalDigit is 9,
		(9 + 1) % 10 = 0
		(9 + 9) % 10 = 8

		If originalDigit is 0,
		(0 + 1) % 10 = 1
		(0 + 9) % 10 = 9
	*/
	incrementedDigit := (originalDigit + 1) % 10
	decrementedDigit := (originalDigit + 9) % 10

	// Create new strings with the modified digits using string concatenation
	incrementedString := digit[:index] + strconv.Itoa(int(incrementedDigit)) + digit[index+1:]
	decrementedString := digit[:index] + strconv.Itoa(int(decrementedDigit)) + digit[index+1:]

	return incrementedString, decrementedString
}

// GenInAndDe is a function to generate the increment and decrement values of a digit
// It using byte slice manipulation to create new strings with the modified digits
func GenInAndDe1(digit string, index int) (string, string) {
	bytes := []byte(digit)

	inc := make([]byte, len(bytes))
	dec := make([]byte, len(bytes))
	copy(inc, bytes)
	copy(dec, bytes)

	// convert the ASCII value of a digit to an integer
	// for example, '2' has an ASCII value of 50. 50 - 48 = 2
	originalDigit := bytes[index] - '0'

	// handle business logic
	incrementedDigit := (originalDigit + 1) % 10
	decrementedDigit := (originalDigit + 9) % 10

	// convert the integer value of a digit to an ASCII value
	// for example, 2 + 48 = 50. 50 is the ASCII value of '2'
	inc[index] = incrementedDigit + '0'
	dec[index] = decrementedDigit + '0'

	return string(inc), string(dec)
}

// GenInAndDe is a function to generate the increment and decrement values of a digit
// It also using byte slice manipulation to create new strings with the modified digits
// It's more naive but readable approach
func GenInAndDe2(digit string, index int) (string, string) {
	bytes := []byte(digit)

	var inByte, deByte byte

	// increment
	if bytes[index] == '9' {
		inByte = '0'
	} else {
		inByte = bytes[index] + 1
	}

	// decrement
	if bytes[index] == '0' {
		deByte = '9'
	} else {
		deByte = bytes[index] - 1
	}

	inBytes := []byte(digit)
	inBytes[index] = inByte

	deBytes := []byte(digit)
	deBytes[index] = deByte

	return string(inBytes), string(deBytes)
}
