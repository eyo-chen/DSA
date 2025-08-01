package main

func GenerateParenthesis(n int) []string {
	ans := []string{}
	helper(n, []byte{}, &ans, 0, 0)

	return ans
}

func helper(n int, str []byte, ans *[]string, countOpen int, countClose int) {
	if len(str) == n*2 {
		*ans = append(*ans, string(str))
		return
	}

	if countOpen < n {
		helper(n, append(str, '('), ans, countOpen+1, countClose)
	}

	if countClose < countOpen {
		helper(n, append(str, ')'), ans, countOpen, countClose+1)
	}
}

// Updated at 2025/01/10
func GenerateParenthesis2(n int) []string {
	ans := []string{}
	helper2(n, &ans, []byte{}, 0, 0)
	return ans
}

func helper2(n int, ans *[]string, curString []byte, open, close int) {
	if open == n && close == n {
		*ans = append(*ans, string(curString))
		return
	}

	if open > n || close > n {
		return
	}

	helper2(n, ans, append(curString, '('), open+1, close)
	if open > close {
		helper2(n, ans, append(curString, ')'), open, close+1)
	}
}

// Updated at 2025/07/31
func GenerateParenthesis3(n int) []string {
	result := []string{}

	// Start the recursive generation with:
	// - n opening parentheses remaining to place
	// - n closing parentheses remaining to place
	// - empty byte slice to build our current combination
	buildValidCombinations(n, n, n, &result, []byte{})

	return result
}

// buildValidCombinations recursively constructs all valid parentheses combinations
// by making careful decisions about when to place opening vs closing parentheses
func buildValidCombinations(totalPairs int, openingRemaining int, closingRemaining int, result *[]string, currentCombination []byte) {
	// Base case: we've successfully placed all opening and closing parentheses
	// This means we have a complete, valid combination
	if openingRemaining == 0 && closingRemaining == 0 {
		*result = append(*result, string(currentCombination))
		return
	}

	// Decision 1: Try placing an opening parenthesis '('
	// We can always place an opening parenthesis as long as we have some remaining
	// This is like saying "start a new nested structure"
	if openingRemaining > 0 {
		// Create a new slice with the opening parenthesis added
		// The append creates a new slice, so we don't modify the original
		buildValidCombinations(totalPairs, openingRemaining-1, closingRemaining, result, append(currentCombination, '('))
	}

	// Decision 2: Try placing a closing parenthesis ')'
	// We can only place a closing parenthesis if we have more opening parentheses
	// placed than closing ones so far. This ensures we never have more ')' than '('
	// at any point in our string, which would make it invalid
	if openingRemaining < closingRemaining {
		// The condition above works because:
		// - We started with equal counts of opening and closing parentheses to place
		// - If openingRemaining < closingRemaining, it means we've used more opening
		//   parentheses than closing ones so far, so we have unmatched '(' that need ')'
		buildValidCombinations(totalPairs, openingRemaining, closingRemaining-1, result, append(currentCombination, ')'))
	}
}
