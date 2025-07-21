package main

// makeFancyString removes minimum characters to ensure no three consecutive characters are identical
// A "fancy string" has no three consecutive equal characters
func MakeFancyString(s string) string {
	// Base case: strings with 2 or fewer characters can't have three consecutive chars
	if len(s) <= 2 {
		return s
	}

	// Initialize result with first two characters (they're always safe to include)
	result := []byte{s[0], s[1]}

	// Process each character starting from the third one (index 2)
	for currentIndex := 2; currentIndex < len(s); currentIndex++ {
		currentChar := s[currentIndex]
		previousChar := s[currentIndex-1]
		secondPreviousChar := s[currentIndex-2]

		// Skip current character if it would create three consecutive identical characters
		// This happens when: current == previous == second_previous
		if currentChar == previousChar && currentChar == secondPreviousChar {
			continue // Skip this character (delete it)
		}

		// Safe to add current character - it won't create three consecutive identical chars
		result = append(result, currentChar)
	}

	return string(result)
}

/*
Algorithm Explanation:
===================

The key insight is that we only need to remove characters when they would
create a third consecutive identical character.

Step-by-step approach:
1. Keep the first two characters (they can't create three consecutive)
2. For each subsequent character, check if adding it would create three
   consecutive identical characters
3. If yes, skip it (this is the "deletion")
4. If no, add it to the result

Example walkthrough with "aaabaaaa":
- Start: result = "aa" (first two chars)
- Check 'a' at pos 2: 'a' == 'a' == 'a' → skip (would make "aaa")
- Check 'b' at pos 3: 'b' != 'a' → add → result = "aab"
- Check 'a' at pos 4: 'a' != 'b' → add → result = "aaba"
- Check 'a' at pos 5: 'a' == 'a' but != 'b' → add → result = "aabaa"
- Check 'a' at pos 6: 'a' == 'a' == 'a' → skip (would make "aabaa" + "a" = three a's at end)
- Check 'a' at pos 7: 'a' == 'a' == 'a' → skip
- Final result: "aabaa"

Time Complexity: O(n) - single pass through string
Space Complexity: O(n) - for result string
*/
