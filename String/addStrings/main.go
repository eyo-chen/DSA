package main

func AddStrings(num1 string, num2 string) string {
	var ans []byte
	idx1 := len(num1) - 1
	idx2 := len(num2) - 1
	carry := 0

	// Loop through the numbers from the end to the beginning
	// We have to loop through until both indices are less than 0
	for idx1 >= 0 || idx2 >= 0 {
		// Initialize sum with the carry
		sum := carry

		// Add the number from num1 to the sum if idx1 is still in the range
		if idx1 >= 0 {
			// Convert the character to integer by subtracting '0' from it
			// e.g. '6' = 54, '0' = 48, so '6' - '0' = 54 - 48 = 6
			sum += int(num1[idx1] - '0')
			idx1--
		}

		// Add the number from num2 to the sum if idx2 is still in the range
		if idx2 >= 0 {
			sum += int(num2[idx2] - '0')
			idx2--
		}

		// Append the last digit of the sum to the answer
		// e.g. sum = 18, 18 % 10 = 8, so we append '8' to the answer
		// After getting the last digit, we convert it back to the ASCII value by adding '0'
		ans = append(ans, byte(sum%10)+'0')

		// Update the carry
		carry = sum / 10
	}

	// If there is a carry left, append it to the answer
	if carry > 0 {
		ans = append(ans, byte(carry)+'0')
	}

	// Reverse the answer because we have been appending the digits from the end to the beginning
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}

	return string(ans)
}
