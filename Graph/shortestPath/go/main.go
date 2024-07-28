package main

import (
	"math"
)

type location struct {
	row   int
	col   int
	k     int
	steps int
}

// Using BFS
func ShortestPath(grid [][]int, k int) int {
	// initialize the queue with the starting location
	q := []location{{0, 0, k, 0}}

	// offsets for the four directions
	offsets := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// visited locations
	// for each location, we need to keep track of the state of the key
	// for example, if we have visited location (1, 2) with 1 key, we should not visit the same location with the same key again
	// so we need a 3D array to keep track of the visited locations
	// [                          -> row
	//   [  										  -> col
	//     [true, false, false],  -> key
	//   ]
	// ]
	visited := make([][][]bool, len(grid))
	for r := range grid {
		visited[r] = make([][]bool, len(grid[0]))
		for c := range grid[0] {
			visited[r][c] = make([]bool, k+1)
		}
	}

	visited[0][0][k] = true

	for len(q) > 0 {
		l := q[0]
		q = q[1:]

		// check if we have reached the destination
		// note that the index is 0-based, so we need to subtract 1 from the length
		if l.row == len(grid)-1 && l.col == len(grid[0])-1 {
			return l.steps
		}

		// loop through the four directions(adjacent locations)
		for _, o := range offsets {
			r, c := l.row+o[0], l.col+o[1]

			// check if the location is out of the boundary
			if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
				continue
			}

			// check if the location is obstacle and we don't have any key left
			if grid[r][c] == 1 && l.k <= 0 {
				continue
			}

			// update the key (if we have encountered an obstacle)
			newK := l.k
			if grid[r][c] == 1 {
				newK--
			}

			// check if the location has been visited
			// if the location has been visited with the same key, we should not visit it again
			if visited[r][c][newK] {
				continue
			}

			// add the location to the queue
			q = append(q, location{
				row:   r,
				col:   c,
				k:     newK,
				steps: l.steps + 1,
			})

			// mark the location as visited
			visited[r][c][newK] = true
		}
	}

	return -1
}

// Using DFS(less efficient)
func ShortestPath1(grid [][]int, k int) int {
	return dfs(grid, 0, 0, 0, k)
}

func dfs(grid [][]int, r int, c int, steps int, k int) int {
	// check if the location is out of the boundary
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return -1
	}

	// check if the location is visited
	if grid[r][c] == 2 {
		return -1
	}

	// check if the location is obstacle and we don't have any key left
	if grid[r][c] == 1 && k == 0 {
		return -1
	}

	// check if we have reached the destination
	if r == len(grid)-1 && c == len(grid[0])-1 {
		return steps
	}

	// update the key (if we have encountered an obstacle)
	newK := k
	if grid[r][c] == 1 {
		newK = k - 1
	}

	state := grid[r][c]

	// mark the location as visited
	// we can directly update the state of the location to 2 as visited, which is incorrect in the BFS solution
	// because we explore each path in the DFS solution, so we guarantee that there's only one state for each location
	grid[r][c] = 2

	// explore the four directions
	upSteps := dfs(grid, r, c+1, steps+1, newK)
	downSteps := dfs(grid, r, c-1, steps+1, newK)
	rightSteps := dfs(grid, r+1, c, steps+1, newK)
	leftSteps := dfs(grid, r-1, c, steps+1, newK)

	// restore the state of the location
	grid[r][c] = state

	return getPositiveMin(upSteps, downSteps, rightSteps, leftSteps)
}

func getPositiveMin(is ...int) int {
	var result = math.MaxInt
	for _, i := range is {
		if i > 0 && i < result {
			result = i
		}
	}

	if result == math.MaxInt {
		return -1
	}

	return result
}
