package main

import "fmt"

// Brute Force Approach
func UniquePaths(m int, n int) int {
	return helper(m, n, 0, 0)
}

func helper(m, n, r, c int) int {
	if r < 0 || c < 0 || r >= m || c >= n {
		return 0
	}

	if r == m-1 && c == n-1 {
		return 1
	}

	right := helper(m, n, r, c+1)
	left := helper(m, n, r+1, c)

	return right + left
}

// Memoization Approach
func UniquePaths1(m int, n int) int {
	memo := map[string]int{}
	return helper1(m, n, 0, 0, memo)
}

func helper1(m, n, r, c int, memo map[string]int) int {
	if r < 0 || c < 0 || r >= m || c >= n {
		return 0
	}

	if r == m-1 && c == n-1 {
		return 1
	}

	key := fmt.Sprintf("%d,%d", r, c)
	if val, ok := memo[key]; ok {
		return val
	}

	right := helper1(m, n, r, c+1, memo)
	left := helper1(m, n, r+1, c, memo)

	memo[key] = right + left

	return right + left
}

// Dynamic Programming Approach (Bottom-Up)
func UniquePaths2(m int, n int) int {
	table := make([][]int, m)
	for i := range table {
		t := make([]int, n)
		table[i] = t
	}

	for i := 0; i < m; i++ {
		table[i][0] = 1
	}

	for i := 0; i < n; i++ {
		table[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			table[i][j] = table[i-1][j] + table[i][j-1]
		}
	}

	return table[m-1][n-1]
}
