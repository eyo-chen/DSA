package main

import "strings"

func WordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	for _, word := range wordDict {
		if !strings.HasPrefix(s, word) {
			continue
		}

		remainingStr := s[len(word):]
		if WordBreak(remainingStr, wordDict) {
			return true
		}
	}

	return false
}

func WordBreak2(s string, wordDict []string) bool {
	memo := map[string]bool{}
	return helper(s, wordDict, memo)
}

func helper(s string, wordDict []string, memo map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	if val, ok := memo[s]; ok {
		return val
	}

	for _, word := range wordDict {
		if !strings.HasPrefix(s, word) {
			continue
		}

		remainingStr := s[len(word):]
		if helper(remainingStr, wordDict, memo) {
			memo[s] = true
			return true
		}
	}

	memo[s] = false
	return false
}

func WordBreak3(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s)+1; i++ {
		if !dp[i] {
			continue
		}

		for _, word := range wordDict {
			if i+len(word) <= len(s) && s[i:len(word)+i] == word {
				dp[len(word)+i] = true
			}
		}
	}

	return dp[len(s)]
}
