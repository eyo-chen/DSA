package main

import "fmt"

func MinimumTotal(triangle [][]int) int {
	memo := map[string]int{}
	return helper(triangle, memo, 0, 0)
}

func helper(triangle [][]int, memo map[string]int, depth, index int) int {
	if depth >= len(triangle) {
		return 0
	}

	key := fmt.Sprintf("%d-%d", depth, index)
	if v, ok := memo[key]; ok {
		return v
	}

	val := triangle[depth][index]
	left := helper(triangle, memo, depth+1, index)
	right := helper(triangle, memo, depth+1, index+1)

	memo[key] = val + min(left, right)
	return memo[key]
}

func MinimumTotal1(triangle [][]int) int {
	l := len(triangle)
	table := make([]int, l)

	for i := 0; i < l; i++ {
		table[i] = triangle[l-1][i]
	}

	for i := l - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			v := triangle[i][j]
			m := min(table[j], table[j+1])
			table[j] = v + m
		}
	}

	return table[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
