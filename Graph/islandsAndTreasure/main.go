package main

import "fmt"

// IslandsAndTreasureOptimized calculates the shortest distance from each empty room to the nearest treasure gate.
// Uses multi-source BFS approach: starts BFS from all treasure gates (0s) simultaneously,
// which is more efficient than running separate BFS from each gate.
// Time complexity: O(m*n), Space complexity: O(m*n)
func IslandsAndTreasureOptimized(grid [][]int) {
	treasureGates := [][]int{}

	// Find all treasure gates (cells with value 0) and add them to the queue
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 0 {
				treasureGates = append(treasureGates, []int{row, col})
			}
		}
	}

	// Perform multi-source BFS from all treasure gates simultaneously
	performMultiSourceBFS(grid, treasureGates)
}

// performMultiSourceBFS performs breadth-first search from multiple starting points (treasure gates)
// to find the shortest distance to all reachable empty rooms
func performMultiSourceBFS(grid [][]int, initialQueue [][]int) {
	queue := initialQueue
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // right, left, down, up

	// Process all nodes level by level
	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:] // dequeue first element

		currentRow, currentCol := currentPosition[0], currentPosition[1]

		// Explore all 4 adjacent directions
		for _, direction := range directions {
			neighborRow := currentRow + direction[0]
			neighborCol := currentCol + direction[1]

			// Check bounds and if the cell is an empty room (INF value)
			if neighborRow < 0 || neighborCol < 0 ||
				neighborRow >= len(grid) || neighborCol >= len(grid[0]) ||
				grid[neighborRow][neighborCol] != 2147483647 {
				continue
			}

			// Update distance: current distance + 1
			grid[neighborRow][neighborCol] = grid[currentRow][currentCol] + 1
			// Add neighbor to queue for further exploration
			queue = append(queue, []int{neighborRow, neighborCol})
		}
	}
}

// IslandsAndTreasureBruteForce calculates shortest distances using individual BFS from each treasure gate.
// This approach is less efficient as it performs separate BFS for each gate.
// Time complexity: O(k*m*n) where k is number of gates, Space complexity: O(m*n)
func IslandsAndTreasureBruteForce(grid [][]int) {
	// Find each treasure gate and run BFS from it
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 0 {
				runBFSFromTreasureGate(grid, row, col)
			}
		}
	}
}

var movementDirections = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// runBFSFromTreasureGate performs BFS from a single treasure gate to update distances
func runBFSFromTreasureGate(grid [][]int, startRow int, startCol int) {
	queue := [][]int{{startRow, startCol}}
	currentDistance := 1
	visitedCells := map[string]bool{} // track visited cells to avoid cycles

	// Process nodes level by level
	for len(queue) > 0 {
		currentLevelSize := len(queue)

		// Process all nodes at current distance level
		for range currentLevelSize {
			currentNode := queue[0]
			queue = queue[1:] // dequeue

			// Explore all 4 directions from current node
			for _, direction := range movementDirections {
				neighborRow := currentNode[0] + direction[0]
				neighborCol := currentNode[1] + direction[1]
				cellKey := fmt.Sprintf("%d,%d", neighborRow, neighborCol)

				// Check bounds, cell type, and if already visited
				if neighborRow < 0 || neighborCol < 0 ||
					neighborRow >= len(grid) || neighborCol >= len(grid[0]) ||
					grid[neighborRow][neighborCol] <= 0 || visitedCells[cellKey] {
					continue
				}

				// Update cell with minimum distance found so far
				grid[neighborRow][neighborCol] = min(grid[neighborRow][neighborCol], currentDistance)
				queue = append(queue, []int{neighborRow, neighborCol})
				visitedCells[cellKey] = true // mark as visited
			}
		}
		currentDistance++ // increment distance for next level
	}
}
