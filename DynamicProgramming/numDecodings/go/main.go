package main

import (
	"strconv"
)

func NumDecodings(s string) int {
	return helper(s, 0)
}

func helper(s string, index int) int {
	if index >= len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	ways := 0
	ways += helper(s, index+1)
	if index+2 <= len(s) {
		subStr := s[index : index+2]
		twoDigit, _ := strconv.Atoi(subStr)
		if twoDigit <= 26 {
			ways += helper(s, index+2)
		}
	}

	return ways
}

func NumDecodings1(s string) int {
	memo := map[int]int{}
	return helper1(s, 0, memo)
}

func helper1(s string, index int, memo map[int]int) int {
	if index >= len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if v, ok := memo[index]; ok {
		return v
	}

	ways := 0
	ways += helper1(s, index+1, memo)
	if index+2 <= len(s) {
		subStr := s[index : index+2]
		twoDigit, _ := strconv.Atoi(subStr)
		if twoDigit <= 26 {
			ways += helper1(s, index+2, memo)
		}
	}

	memo[index] = ways
	return ways
}
