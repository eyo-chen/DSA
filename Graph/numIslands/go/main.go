package main

func NumIslands(grid [][]byte) int {
	result := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' {
				helper(grid, r, c)
				result++
			}
		}
	}

	return result
}

func helper(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0'
	helper(grid, r+1, c)
	helper(grid, r-1, c)
	helper(grid, r, c+1)
	helper(grid, r, c-1)
}
