package main

import (
	"github.com/OYE0303/DSA/goutils/queue"
)

// Coordinate represents a position in the 2D grid with row and column indices
type Coordinate struct {
	row int
	col int
}

// OrangesRotting calculates the minimum time required for all fresh oranges to rot.
// Uses BFS (Breadth-First Search) approach where all rotten oranges at each time step
// simultaneously rot their adjacent fresh oranges. Returns -1 if some oranges cannot rot.
//
// Algorithm:
// 1. Initialize queue with all initially rotten oranges and count fresh oranges
// 2. For each time step, process all currently rotten oranges in the queue
// 3. Each rotten orange attempts to rot its 4-directional neighbors
// 4. Continue until no fresh oranges remain or queue is empty
// 5. Return elapsed time or -1 if fresh oranges are unreachable
func OrangesRotting(grid [][]int) int {
	rottenQueue := []Coordinate{}
	remainingFreshCount := 0
	elapsedMinutes := 0

	// Direction vectors for 4-directional movement (up, down, right, left)
	directions := []Coordinate{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// First pass: identify all rotten oranges and count fresh ones
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			// Add initially rotten oranges to processing queue
			if grid[rowIndex][colIndex] == 2 {
				rottenQueue = append(rottenQueue, Coordinate{rowIndex, colIndex})
			}

			// Count total fresh oranges that need to rot
			if grid[rowIndex][colIndex] == 1 {
				remainingFreshCount++
			}
		}
	}

	// BFS: process rotten oranges level by level (minute by minute)
	// Continue while there are rotten oranges to process AND fresh oranges remaining
	for len(rottenQueue) > 0 && remainingFreshCount > 0 {
		currentLevelSize := len(rottenQueue)

		// Process all oranges that became rotten in the current minute
		for range currentLevelSize {
			currentPosition := rottenQueue[0]
			rottenQueue = rottenQueue[1:] // Dequeue from front

			// Check all 4 adjacent cells for fresh oranges to rot
			currentRow, currentCol := currentPosition.row, currentPosition.col
			for _, direction := range directions {
				neighborRow := currentRow + direction.row
				neighborCol := currentCol + direction.col

				// Skip if neighbor is out of bounds or not a fresh orange
				if neighborRow < 0 || neighborCol < 0 ||
					neighborRow >= len(grid) || neighborCol >= len(grid[0]) ||
					grid[neighborRow][neighborCol] != 1 {
					continue
				}

				// Rot the fresh orange and add it to queue for next minute
				grid[neighborRow][neighborCol] = 2
				rottenQueue = append(rottenQueue, Coordinate{neighborRow, neighborCol})
				remainingFreshCount--
			}
		}

		// One minute has passed after processing current level
		elapsedMinutes++
	}

	// Return -1 if some fresh oranges couldn't be reached, otherwise return time elapsed
	if remainingFreshCount != 0 {
		return -1
	}

	return elapsedMinutes
}

// OrangesRottingWithQueue calculates the minimum time required for all fresh oranges to rot.
// Uses BFS (Breadth-First Search) approach with a custom queue data structure where all
// rotten oranges at each time step simultaneously rot their adjacent fresh oranges.
// Returns -1 if some oranges cannot rot.
//
// Algorithm:
// 1. Initialize queue with all initially rotten oranges and count fresh oranges
// 2. For each time step, process all currently rotten oranges in the queue
// 3. Each rotten orange attempts to rot its 4-directional neighbors
// 4. Continue until no fresh oranges remain or queue is empty
// 5. Return elapsed time or -1 if fresh oranges are unreachable
func OrangesRottingWithQueue(grid [][]int) int {
	rottenQueue := queue.Constructor[Coordinate]()
	elapsedMinutes := 0
	remainingFreshCount := 0

	// Direction vectors for 4-directional movement (up, down, right, left)
	directions := []Coordinate{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// First pass: identify all rotten oranges and count fresh ones
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			// Add initially rotten oranges to processing queue
			if grid[rowIndex][colIndex] == 2 {
				rottenQueue.Push(Coordinate{rowIndex, colIndex})
			}

			// Count total fresh oranges that need to rot
			if grid[rowIndex][colIndex] == 1 {
				remainingFreshCount++
			}
		}
	}

	// BFS: process rotten oranges level by level (minute by minute)
	// Continue while there are rotten oranges to process AND fresh oranges remaining
	for rottenQueue.Size() > 0 && remainingFreshCount > 0 {
		currentLevelSize := rottenQueue.Size()

		// Process all oranges that became rotten in the current minute
		for range currentLevelSize {
			currentPosition := rottenQueue.Front()
			rottenQueue.Pop() // Dequeue from front

			// Check all 4 adjacent cells for fresh oranges to rot
			currentRow, currentCol := currentPosition.row, currentPosition.col
			for _, direction := range directions {
				neighborRow := currentRow + direction.row
				neighborCol := currentCol + direction.col

				// Skip if neighbor is out of bounds or not a fresh orange
				if neighborRow < 0 || neighborCol < 0 ||
					neighborRow >= len(grid) || neighborCol >= len(grid[0]) ||
					grid[neighborRow][neighborCol] != 1 {
					continue
				}

				// Rot the fresh orange and add it to queue for next minute
				grid[neighborRow][neighborCol] = 2
				rottenQueue.Push(Coordinate{neighborRow, neighborCol})
				remainingFreshCount--
			}
		}

		// One minute has passed after processing current level
		elapsedMinutes++
	}

	// Return -1 if some fresh oranges couldn't be reached, otherwise return time elapsed
	if remainingFreshCount != 0 {
		return -1
	}

	return elapsedMinutes
}
