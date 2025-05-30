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

type location struct {
	row int
	col int
}

var directions = []location{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func NumIslands2(grid [][]byte) int {
	ans := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				bfs(grid, r, c)
				ans++
			}
		}
	}

	return ans
}

func bfs(grid [][]byte, startRow, startCol int) {
	queue := []location{{startRow, startCol}}
	// When a node is added to the queue, it is immediately marked as visited by setting it to '0'
	grid[startRow][startCol] = '0'

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			row, col := node.row+d.row, node.col+d.col

			if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
				continue
			}
			if grid[row][col] != '1' {
				continue
			}

			queue = append(queue, location{row, col})

			// When a node is added to the queue, it is immediately marked as visited by setting it to '0'
			grid[row][col] = '0'
		}
	}
}

type location3 struct {
	row int
	col int
}

var (
	directions3 = []location3{{row: 1, col: 0}, {row: -1, col: 0}, {row: 0, col: 1}, {row: 0, col: -1}}
)

func NumIslands3(grid [][]byte) int {
	ans := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				bfs3(grid, r, c)
				ans++
			}
		}
	}

	return ans
}

func bfs3(grid [][]byte, row, col int) {
	queue := []location3{{row, col}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.row < 0 || node.col < 0 || node.row >= len(grid) || node.col >= len(grid[0]) {
			continue
		}

		if grid[node.row][node.col] != '1' {
			continue
		}

		grid[node.row][node.col] = '2'
		for _, d := range directions3 {
			r, c := node.row+d.row, node.col+d.col
			queue = append(queue, location3{row: r, col: c})
		}
	}
}
