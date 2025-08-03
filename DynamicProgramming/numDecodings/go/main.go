package main

import "strconv"

// Solution 1: Substring-based memoization approach
//
// CORE IDEA: This solution thinks about the problem by progressively "eating" characters
// from the front of the string. At each step, we ask: "How many ways can I decode what's left?"
//
// The beauty of this approach is its intuitive nature - we literally work with the remaining
// portion of the string that needs to be decoded. When we see "226", we first consider:
// - Take "2" and figure out how to decode "26"
// - Take "22" and figure out how to decode "6"
//
// MEMOIZATION STRATEGY: We cache results using the actual substring as the key. This means
// if we encounter the same remaining substring again (like "26" appearing in multiple contexts),
// we can instantly return the cached result instead of recalculating.
//
// TRADE-OFFS: While this approach is very intuitive and mirrors how humans might think about
// the problem, it has higher memory overhead because:
// 1. Substring creation takes O(n) time and space for each recursive call
// 2. The cache uses strings as keys, which requires more memory than integers
// 3. Overall time complexity becomes O(n²) due to substring operations
//
// Time: O(n²), Space: O(n²) for memoization (stores substrings as keys)
func NumDecodings1(s string) int {
	// Use a map to cache results for each substring we've already processed
	// Key insight: if we see the same remaining substring again, we know the answer
	substringCache := map[string]int{}
	return decodeSubstring(s, substringCache)
}

func decodeSubstring(remainingString string, substringCache map[string]int) int {
	// Critical validation: A string starting with '0' cannot be decoded because
	// there's no letter mapped to 0, and leading zeros in multi-digit numbers
	// would create invalid mappings (like "01" trying to be "A")
	if len(remainingString) > 0 && remainingString[0] == '0' {
		return 0
	}

	// Base cases for recursion termination:
	// - Empty string: we've successfully consumed all characters, so this represents 1 valid decoding path
	// - Single character: if we got here, it must be 1-9 (we checked for 0 above), so it's 1 valid path
	if len(remainingString) < 2 {
		return 1
	}

	// Memoization lookup: if we've seen this exact remaining substring before,
	// return the cached result to avoid redundant calculation
	if cachedResult, exists := substringCache[remainingString]; exists {
		return cachedResult
	}

	totalWays := 0

	// Decision 1: Decode the first character as a single digit
	// This is always valid since we've already ruled out '0' at the start
	// We recursively solve for the substring after removing the first character
	remainingAfterOne := remainingString[1:]
	totalWays += decodeSubstring(remainingAfterOne, substringCache)

	// Decision 2: Decode the first two characters as a double digit (if valid)
	// We need to check if the two-digit number is in the valid range 10-26
	firstTwoDigits, _ := strconv.Atoi(remainingString[:2])
	if firstTwoDigits <= 26 {
		// If valid, recursively solve for the substring after removing the first two characters
		remainingAfterTwo := remainingString[2:]
		totalWays += decodeSubstring(remainingAfterTwo, substringCache)
	}

	// Cache this result before returning so future calls with the same substring can reuse it
	substringCache[remainingString] = totalWays
	return totalWays
}

// Solution 2: Index-based memoization approach
//
// CORE IDEA: Instead of creating new substrings, this solution uses an index to track
// our current position in the original string. Think of it as having a "pointer" that
// moves through the string as we make decoding decisions.
//
// This approach asks: "Starting from position i, how many ways can I decode the rest?"
// The elegance lies in avoiding substring creation entirely - we just pass around
// an integer index that tells us where we are in the decoding process.
//
// MEMOIZATION STRATEGY: We cache results using the starting position (integer) as the key.
// This is much more memory-efficient than storing substring keys, and integer comparison
// for cache lookups is faster than string comparison.
//
// PERFORMANCE ADVANTAGES: This approach is more efficient because:
// 1. No substring creation means no additional memory allocation per recursive call
// 2. Integer keys in the cache are smaller and faster to compare than strings
// 3. Overall time complexity is O(n) since each position is calculated exactly once
// 4. Space complexity is O(n) for the recursion stack and memoization cache
//
// Time: O(n), Space: O(n) for memoization
func NumDecodings2(s string) int {
	// Use a map to cache results for each starting position in the string
	// Key insight: the number of ways to decode from position i depends only on i and the string content
	positionCache := map[int]int{}
	return decodeFromPosition(s, 0, positionCache)
}

func decodeFromPosition(s string, currentIndex int, positionCache map[int]int) int {
	// Base case: we've processed all characters successfully
	// This represents one complete valid decoding path
	if currentIndex >= len(s) {
		return 1
	}

	// Invalid case: current character is '0'
	// Since there's no letter mapped to 0, and this isn't part of a valid two-digit code
	// (we would have consumed it in a previous step), this path is invalid
	if s[currentIndex] == '0' {
		return 0
	}

	// Memoization lookup: if we've already calculated the result for this position,
	// return it immediately to avoid redundant work
	if cachedResult, exists := positionCache[currentIndex]; exists {
		return cachedResult
	}

	totalWays := 0

	// Decision 1: Decode current character as single digit (1-9)
	// Since we've ruled out '0', any single digit at this position is valid
	// Move the index forward by 1 and solve the remaining subproblem
	totalWays += decodeFromPosition(s, currentIndex+1, positionCache)

	// Decision 2: Decode current and next character as two digits (10-26)
	// First, ensure we have at least 2 characters remaining
	if currentIndex+2 <= len(s) {
		// Extract the two-digit substring and convert to integer for validation
		twoDigitSubstring := s[currentIndex : currentIndex+2]
		twoDigitNumber, _ := strconv.Atoi(twoDigitSubstring)

		// Only proceed if the two-digit number represents a valid letter (10-26)
		// Note: numbers like 27, 28, 29, etc. don't map to any letters
		if twoDigitNumber <= 26 {
			// Move the index forward by 2 and solve the remaining subproblem
			totalWays += decodeFromPosition(s, currentIndex+2, positionCache)
		}
	}

	// Cache this result for the current position before returning
	// This ensures that if we encounter the same position again, we can return instantly
	positionCache[currentIndex] = totalWays
	return totalWays
}

/*
TABULAR DP SOLUTION FOR DECODE WAYS

The key insight for converting from recursion to tabulation is changing our perspective:
- Recursive: "How many ways to decode FROM position i to the end?"
- Tabular: "How many ways to decode UP TO position i?"

Think of it like filling a table from left to right, where each cell represents
"the number of ways to decode the string up to this point."

TABLE SETUP:
- dp[i] = number of ways to decode the first i characters of the string
- We need dp[0] through dp[n] where n is the length of the string
- dp[0] = 1 (empty string has one way to decode: do nothing)
- Our answer will be dp[n] (ways to decode all n characters)
*/

func NumDecodingsDP(s string) int {
	n := len(s)

	// Handle edge case: empty string
	if n == 0 {
		return 0
	}

	// Create DP table where dp[i] represents the number of ways to decode
	// the first i characters of the string
	// We need n+1 slots: dp[0] through dp[n]
	dp := make([]int, n+1)

	// Base case: empty string (first 0 characters) has exactly 1 way to decode
	// This might seem counterintuitive, but it's necessary for our recurrence to work
	// Think of it as "there's one way to decode nothing: by doing nothing"
	dp[0] = 1

	// Handle the first character specially
	// If the first character is '0', there's no way to decode it (no letter maps to 0)
	// If it's 1-9, there's exactly 1 way to decode it (as a single character)
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	// Fill the DP table for positions 2 through n
	// At each position i, we consider two possibilities:
	// 1. Decode s[i-1] as a single character (if it's 1-9)
	// 2. Decode s[i-2:i] as a two-character code (if it's 10-26)
	for i := 2; i <= n; i++ {
		// Extract the characters we're considering
		singleChar := s[i-1]        // The character at position i-1 (0-indexed)
		twoCharString := s[i-2 : i] // The two characters ending at position i-1

		// Option 1: Decode the current character as a single digit
		// This is valid if the character is '1' through '9' (not '0')
		if singleChar >= '1' && singleChar <= '9' {
			// If we can decode this character alone, then all the ways to decode
			// up to the previous position are still valid with this character added
			dp[i] += dp[i-1]
		}

		// Option 2: Decode the current and previous character as a two-digit code
		// This is valid if the two-digit number is between 10 and 26
		twoDigitNum, _ := strconv.Atoi(twoCharString)
		if twoDigitNum >= 10 && twoDigitNum <= 26 {
			// If we can decode these two characters together, then all the ways
			// to decode up to two positions ago are still valid with this pair added
			dp[i] += dp[i-2]
		}

		// Note: dp[i] might be 0 if neither option is valid, which correctly
		// represents that there's no way to decode up to position i
	}

	// Return the number of ways to decode the entire string
	return dp[n]
}

/*
STEP-BY-STEP EXAMPLE: Let's trace through s = "226"

Initial setup:
dp = [1, 0, 0, 0]  // dp[0]=1 (base case), others will be filled

After handling first character '2':
dp = [1, 1, 0, 0]  // dp[1]=1 because '2' is valid single digit

i=2 (considering first 2 characters "22"):
- singleChar = '2' (second '2')
- twoCharString = "22"
- Option 1: '2' is valid single → dp[2] += dp[1] = 0 + 1 = 1
- Option 2: "22" = 22, which is ≤ 26 → dp[2] += dp[0] = 1 + 1 = 2
dp = [1, 1, 2, 0]

i=3 (considering all 3 characters "226"):
- singleChar = '6'
- twoCharString = "26"
- Option 1: '6' is valid single → dp[3] += dp[2] = 0 + 2 = 2
- Option 2: "26" = 26, which is ≤ 26 → dp[3] += dp[1] = 2 + 1 = 3
dp = [1, 1, 2, 3]

Final answer: dp[3] = 3

The three ways are:
1. "2" + "2" + "6" = "BBF"
2. "22" + "6" = "VF"
3. "2" + "26" = "BZ"
*/

// SPACE-OPTIMIZED VERSION
// Since we only ever look back at most 2 positions, we can optimize space
// from O(n) to O(1) by keeping track of only the last two values
func NumDecodingsDPOptimized(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Instead of a full array, we only keep track of:
	// - prev2: dp[i-2] (ways to decode up to 2 positions ago)
	// - prev1: dp[i-1] (ways to decode up to 1 position ago)
	// - current: dp[i] (ways to decode up to current position)
	prev2 := 1 // This represents dp[0] = 1
	prev1 := 0 // This will become dp[1]

	// Handle first character
	if s[0] != '0' {
		prev1 = 1
	}

	// If string has only one character, return prev1
	if n == 1 {
		return prev1
	}

	// Process remaining characters
	for i := 2; i <= n; i++ {
		current := 0

		// Single character option
		if s[i-1] >= '1' && s[i-1] <= '9' {
			current += prev1
		}

		// Two character option
		twoDigit, _ := strconv.Atoi(s[i-2 : i])
		if twoDigit >= 10 && twoDigit <= 26 {
			current += prev2
		}

		// Shift values for next iteration
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
