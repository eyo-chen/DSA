package main

import "fmt"

// Brute Force Approach
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	return helper(obstacleGrid, 0, 0)
}

func helper(obstacleGrid [][]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(obstacleGrid) || c >= len(obstacleGrid[0]) {
		return 0
	}

	if obstacleGrid[r][c] == 1 {
		return 0
	}

	if r == len(obstacleGrid)-1 && c == len(obstacleGrid[0])-1 {
		return 1
	}

	right := helper(obstacleGrid, r+1, c)
	left := helper(obstacleGrid, r, c+1)

	return right + left
}

// Top-Down Approach with Memoization
func UniquePathsWithObstacles1(obstacleGrid [][]int) int {
	memo := map[string]int{}
	return helper1(obstacleGrid, memo, 0, 0)
}

func helper1(obstacleGrid [][]int, memo map[string]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(obstacleGrid) || c >= len(obstacleGrid[0]) {
		return 0
	}

	if obstacleGrid[r][c] == 1 {
		return 0
	}

	if r == len(obstacleGrid)-1 && c == len(obstacleGrid[0])-1 {
		return 1
	}

	key := fmt.Sprintf("%d-%d", r, c)
	if v, ok := memo[key]; ok {
		return v
	}

	right := helper1(obstacleGrid, memo, r+1, c)
	left := helper1(obstacleGrid, memo, r, c+1)

	memo[key] = right + left
	return memo[key]
}

// Dynamic Programming Approach (Bottom-Up) With 2D Array
func UniquePathsWithObstacles2(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	table := make([][]int, row)
	for r := 0; r < row; r++ {
		table[r] = make([]int, col)
	}

	table[0][0] = 1
	for r := 1; r < row; r++ {
		// if the current cell is an obstacle, break the loop
		// leave the rest of the cells in that row as 0
		if obstacleGrid[r][0] == 1 {
			break
		}

		// if the current cell is not an obstacle, set the current cell to be equal to the cell above it
		table[r][0] = table[r-1][0]
	}

	for c := 1; c < col; c++ {
		// if the current cell is an obstacle, break the loop
		// leave the rest of the cells in that column as 0
		if obstacleGrid[0][c] == 1 {
			break
		}

		// if the current cell is not an obstacle, set the current cell to be equal to the cell to the left of it
		table[0][c] = table[0][c-1]
	}

	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			// if the current cell is an obstacle, set the current cell to be equal to 0
			if obstacleGrid[r][c] == 1 {
				table[r][c] = 0
				continue
			}

			// if the current cell is not an obstacle, set the current cell to be equal to the sum of the cell above it and the cell to the left of it
			table[r][c] = table[r-1][c] + table[r][c-1]
		}
	}

	return table[row-1][col-1]
}

// Dynamic Programming Approach (Bottom-Up) With 1D Array
func UniquePathsWithObstacles3(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	row, col := len(obstacleGrid), len(obstacleGrid[0])

	table := make([]int, col)
	table[0] = 1
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if obstacleGrid[r][c] == 1 {
				table[c] = 0
				continue
			}

			if c > 0 {
				table[c] += table[c-1]
			}
		}
	}

	return table[col-1]
}
