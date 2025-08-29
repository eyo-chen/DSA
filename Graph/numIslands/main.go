package main

// Coordinate represents a position in the 2D grid with row and column indices
type Coordinate struct {
	Row int
	Col int
}

// Define the four cardinal directions for grid traversal: right, left, down, up
var cardinalDirections = []Coordinate{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// NumIslandsDFS counts the number of islands in a 2D binary grid using Depth-First Search (DFS).
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// The algorithm iterates through each cell, and when it finds a '1' (land), it performs DFS to mark
// all connected land cells as visited by changing them to '0', then increments the island count.
func NumIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islandCount := 0

	// Iterate through every cell in the grid
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			// When we find an unvisited land cell, start DFS to explore the entire island
			if grid[rowIndex][colIndex] == '1' {
				markIslandAsVisited(grid, rowIndex, colIndex)
				islandCount++
			}
		}
	}

	return islandCount
}

// markIslandAsVisited uses DFS to mark all connected land cells as visited by changing '1' to '0'.
// This prevents counting the same island multiple times.
func markIslandAsVisited(grid [][]byte, currentRow, currentCol int) {
	// Base case: check bounds and if current cell is water or already visited
	if currentRow < 0 || currentCol < 0 || currentRow >= len(grid) || currentCol >= len(grid[0]) || grid[currentRow][currentCol] != '1' {
		return
	}

	// Mark current land cell as visited by setting it to '0'
	grid[currentRow][currentCol] = '0'

	// Recursively explore all four adjacent cells (up, down, left, right)
	markIslandAsVisited(grid, currentRow+1, currentCol) // down
	markIslandAsVisited(grid, currentRow-1, currentCol) // up
	markIslandAsVisited(grid, currentRow, currentCol+1) // right
	markIslandAsVisited(grid, currentRow, currentCol-1) // left
}

// NumIslandsBFS counts the number of islands in a 2D binary grid using Breadth-First Search (BFS).
// This approach uses a queue to explore all connected land cells level by level, marking them as visited
// to avoid counting the same island multiple times.
func NumIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islandCount := 0

	// Iterate through every cell in the grid
	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for colIndex := 0; colIndex < len(grid[0]); colIndex++ {
			// When we find an unvisited land cell, start BFS to explore the entire island
			if grid[rowIndex][colIndex] == '1' {
				exploreIslandBFS(grid, rowIndex, colIndex)
				islandCount++
			}
		}
	}

	return islandCount
}

// exploreIslandBFS uses BFS to mark all connected land cells as visited.
// It uses a queue to process cells level by level, ensuring all connected land is found.
func exploreIslandBFS(grid [][]byte, startRow, startCol int) {
	// Initialize queue with starting position
	cellQueue := []Coordinate{{startRow, startCol}}
	// Mark starting cell as visited immediately to prevent revisiting
	grid[startRow][startCol] = '0'

	// Process queue until all connected land cells are visited
	for len(cellQueue) > 0 {
		// Dequeue the first element
		currentCell := cellQueue[0]
		cellQueue = cellQueue[1:]

		// Explore all four cardinal directions from current cell
		for _, direction := range cardinalDirections {
			nextRow := currentCell.Row + direction.Row
			nextCol := currentCell.Col + direction.Col

			// Check if next position is within grid boundaries
			if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
				continue
			}
			// Skip if cell is water or already visited
			if grid[nextRow][nextCol] != '1' {
				continue
			}

			// Add valid unvisited land cell to queue and mark as visited
			cellQueue = append(cellQueue, Coordinate{nextRow, nextCol})
			grid[nextRow][nextCol] = '0'
		}
	}
}

// NumIslandsBFSLazyValidation counts islands using BFS with lazy boundary validation.
// This version defers boundary checking until cells are dequeued, which can be less efficient
// but demonstrates an alternative approach.
func NumIslandsBFSLazyValidation(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islandCount := 0

	// Iterate through every cell in the grid
	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for colIndex := 0; colIndex < len(grid[0]); colIndex++ {
			if grid[rowIndex][colIndex] == '1' {
				exploreIslandLazy(grid, rowIndex, colIndex)
				islandCount++
			}
		}
	}

	return islandCount
}

// exploreIslandLazy uses BFS with lazy validation - boundary checks happen when cells are processed,
// not when they're added to the queue. Uses '2' as a visited marker instead of '0'.
func exploreIslandLazy(grid [][]byte, startRow, startCol int) {
	cellQueue := []Coordinate{{startRow, startCol}}

	for len(cellQueue) > 0 {
		// Dequeue the first element
		currentCell := cellQueue[0]
		cellQueue = cellQueue[1:]

		// Validate boundaries and cell state when processing (lazy validation)
		if currentCell.Row < 0 || currentCell.Col < 0 ||
			currentCell.Row >= len(grid) || currentCell.Col >= len(grid[0]) {
			continue
		}

		// Skip if cell is not unvisited land
		if grid[currentCell.Row][currentCell.Col] != '1' {
			continue
		}

		// Mark current cell as visited using '2' to distinguish from water '0'
		grid[currentCell.Row][currentCell.Col] = '2'

		// Add all four adjacent cells to queue (validation happens later)
		for _, direction := range cardinalDirections {
			nextRow := currentCell.Row + direction.Row
			nextCol := currentCell.Col + direction.Col
			cellQueue = append(cellQueue, Coordinate{Row: nextRow, Col: nextCol})
		}
	}
}

// NumIslandsBuggy counts the number of islands in a 2D binary grid using Breadth-First Search (BFS).
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
//
// !WARNING: This implementation contains a BUG - it marks cells as visited when they are dequeued
// rather than when they are enqueued. This can lead to the same cell being added to the queue
// multiple times, making the algorithm less efficient (though still correct).
//
// The algorithm iterates through each cell, and when it finds a '1' (land), it performs BFS to mark
// all connected land cells as visited by changing them to '0', then increments the island count.
func NumIslandsBuggy(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islandCount := 0

	// Iterate through every cell in the grid
	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for colIndex := 0; colIndex < len(grid[0]); colIndex++ {
			// When we find an unvisited land cell, start BFS to explore the entire island
			if grid[rowIndex][colIndex] == '1' {
				islandCount++
				exploreIslandIncorrectly(grid, rowIndex, colIndex)
			}
		}
	}

	return islandCount
}

// exploreIslandIncorrectly uses BFS to mark all connected land cells as visited.
//
// !BUG: This function marks cells as visited when they are DEQUEUED instead of when they are ENQUEUED.
// This means the same cell can be added to the queue multiple times by different neighbors before
// it gets processed, leading to unnecessary queue growth and redundant processing.
//
// The correct approach would be to mark cells as '0' immediately when adding them to the queue.
func exploreIslandIncorrectly(grid [][]byte, startRow int, startCol int) {
	// Initialize queue with starting position as [row, col] slice
	coordinateQueue := [][]int{{startRow, startCol}}

	// Process queue until all connected land cells are visited
	for len(coordinateQueue) > 0 {
		// Dequeue the first coordinate pair
		currentCoords := coordinateQueue[0]
		coordinateQueue = coordinateQueue[1:]

		currentRow, currentCol := currentCoords[0], currentCoords[1]
		// BUG: Marking cell as visited here (when dequeued) instead of when enqueued
		// This allows the same cell to be added to queue multiple times
		grid[currentRow][currentCol] = '0'

		// Explore all four cardinal directions from current cell
		for _, direction := range cardinalDirections {
			nextRow := currentRow + direction.Row
			nextCol := currentCol + direction.Col

			// Check if next position is within grid boundaries and is unvisited land
			if nextRow < 0 || nextCol < 0 || nextRow >= len(grid) ||
				nextCol >= len(grid[0]) || grid[nextRow][nextCol] == '0' {
				continue
			}

			// Add valid unvisited land cell to queue
			// BUG: Not marking as visited here allows duplicate queue entries
			coordinateQueue = append(coordinateQueue, []int{nextRow, nextCol})
		}
	}
}
