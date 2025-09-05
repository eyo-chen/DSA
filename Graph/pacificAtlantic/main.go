package main

import "fmt"

// Movement directions: right, left, down, up
var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// PacificAtlanticDFS solves the problem using DFS approach with explicit marking of starting cells
// Approach: Start from ocean boundaries and use DFS to find all reachable cells in reverse direction
// Time Complexity: O(m * n) where m is rows and n is columns - each cell visited at most twice
// Space Complexity: O(m * n) for the hash tables storing reachable cells
func PacificAtlanticDFS(heights [][]int) [][]int {
	pacificReachable := map[string]bool{}
	atlanticReachable := map[string]bool{}

	// Process all rows: left column (Pacific) and right column (Atlantic)
	for row := range heights {
		// Mark Pacific boundary cell and explore from it
		pacificKey := fmt.Sprintf("%d-%d", row, 0)
		pacificReachable[pacificKey] = true
		exploreFromCell(heights, row, 0, pacificReachable)

		// Mark Atlantic boundary cell and explore from it
		atlanticKey := fmt.Sprintf("%d-%d", row, len(heights[0])-1)
		atlanticReachable[atlanticKey] = true
		exploreFromCell(heights, row, len(heights[0])-1, atlanticReachable)
	}

	// Process all columns: top row (Pacific) and bottom row (Atlantic)
	for col := range heights[0] {
		// Mark Pacific boundary cell and explore from it
		pacificKey := fmt.Sprintf("%d-%d", 0, col)
		pacificReachable[pacificKey] = true
		exploreFromCell(heights, 0, col, pacificReachable)

		// Mark Atlantic boundary cell and explore from it
		atlanticKey := fmt.Sprintf("%d-%d", len(heights)-1, col)
		atlanticReachable[atlanticKey] = true
		exploreFromCell(heights, len(heights)-1, col, atlanticReachable)
	}

	// Find intersection: cells reachable from both oceans
	result := [][]int{}
	for row := range heights {
		for col := range heights[0] {
			cellKey := fmt.Sprintf("%d-%d", row, col)
			if pacificReachable[cellKey] && atlanticReachable[cellKey] {
				result = append(result, []int{row, col})
			}
		}
	}

	return result
}

// exploreFromCell performs DFS to find all cells reachable from current position
// Water can flow to cells with higher or equal height (reverse flow logic)
func exploreFromCell(heights [][]int, row int, col int, reachableCells map[string]bool) {
	// Try all four directions
	for _, direction := range directions {
		newRow, newCol := row+direction[0], col+direction[1]
		cellKey := fmt.Sprintf("%d-%d", newRow, newCol)

		// Skip if: out of bounds, water can't flow (height decreases), or already visited
		if newRow < 0 || newCol < 0 || newRow >= len(heights) || newCol >= len(heights[0]) ||
			heights[row][col] > heights[newRow][newCol] || reachableCells[cellKey] {
			continue
		}

		// Mark cell as reachable and continue exploration
		reachableCells[cellKey] = true
		exploreFromCell(heights, newRow, newCol, reachableCells)
	}
}

// PacificAtlanticDFSWithPrevHeight solves using DFS with previous height parameter
// The code is more concise
// Approach: Start from ocean boundaries and use DFS with height comparison parameter
// Time Complexity: O(m * n) where m is rows and n is columns
// Space Complexity: O(m * n) for the hash tables and recursion stack
func PacificAtlanticDFSWithPrevHeight(heights [][]int) [][]int {
	pacificReachable := map[string]bool{}
	atlanticReachable := map[string]bool{}

	// Process boundary cells for both oceans
	for row := range heights {
		// Explore from Pacific boundary (left column)
		exploreWithHeightCheck(heights, row, 0, heights[row][0], pacificReachable)
		// Explore from Atlantic boundary (right column)
		exploreWithHeightCheck(heights, row, len(heights[0])-1, heights[row][len(heights[0])-1], atlanticReachable)
	}

	for col := range heights[0] {
		// Explore from Pacific boundary (top row)
		exploreWithHeightCheck(heights, 0, col, heights[0][col], pacificReachable)
		// Explore from Atlantic boundary (bottom row)
		exploreWithHeightCheck(heights, len(heights)-1, col, heights[len(heights)-1][col], atlanticReachable)
	}

	// Collect cells reachable from both oceans
	result := [][]int{}
	for row := range heights {
		for col := range heights[0] {
			cellKey := fmt.Sprintf("%d-%d", row, col)
			if pacificReachable[cellKey] && atlanticReachable[cellKey] {
				result = append(result, []int{row, col})
			}
		}
	}

	return result
}

// exploreWithHeightCheck performs DFS with explicit previous height comparison
// prevHeight represents the height we're coming from (for flow validation)
func exploreWithHeightCheck(heights [][]int, row int, col int, prevHeight int, reachableCells map[string]bool) {
	cellKey := fmt.Sprintf("%d-%d", row, col)

	// Skip if: out of bounds, water can't flow (height decreases), or already visited
	if row < 0 || col < 0 || row >= len(heights) || col >= len(heights[0]) ||
		prevHeight > heights[row][col] || reachableCells[cellKey] {
		return
	}

	// Mark current cell as reachable
	reachableCells[cellKey] = true

	// Explore all neighboring cells with current height as previous height
	for _, direction := range directions {
		newRow, newCol := row+direction[0], col+direction[1]
		exploreWithHeightCheck(heights, newRow, newCol, heights[row][col], reachableCells)
	}
}

// PacificAtlanticBFS solves the problem using BFS approach
// Approach: Start from ocean boundaries and use BFS to find all reachable cells
// Time Complexity: O(m * n) where m is rows and n is columns
// Space Complexity: O(m * n) for the hash tables and BFS queue
func PacificAtlanticBFS(heights [][]int) [][]int {
	pacificReachable := map[string]bool{}
	atlanticReachable := map[string]bool{}

	// Start BFS from all boundary cells
	for row := range heights {
		// BFS from Pacific boundary (left column)
		breadthFirstSearch(heights, row, 0, pacificReachable)
		// BFS from Atlantic boundary (right column)
		breadthFirstSearch(heights, row, len(heights[0])-1, atlanticReachable)
	}

	for col := range heights[0] {
		// BFS from Pacific boundary (top row)
		breadthFirstSearch(heights, 0, col, pacificReachable)
		// BFS from Atlantic boundary (bottom row)
		breadthFirstSearch(heights, len(heights)-1, col, atlanticReachable)
	}

	// Find cells reachable from both oceans
	result := [][]int{}
	for row := range heights {
		for col := range heights[0] {
			cellKey := fmt.Sprintf("%d-%d", row, col)
			if pacificReachable[cellKey] && atlanticReachable[cellKey] {
				result = append(result, []int{row, col})
			}
		}
	}

	return result
}

// breadthFirstSearch performs BFS to find all cells reachable from starting position
// Uses queue to explore cells level by level
func breadthFirstSearch(heights [][]int, startRow int, startCol int, reachableCells map[string]bool) {
	queue := [][]int{{startRow, startCol}}
	startKey := fmt.Sprintf("%d-%d", startRow, startCol)
	reachableCells[startKey] = true

	// Process queue until empty
	for len(queue) > 0 {
		// Dequeue front element
		currentCell := queue[0]
		queue = queue[1:]

		currentRow, currentCol := currentCell[0], currentCell[1]

		// Explore all neighboring cells
		for _, direction := range directions {
			newRow, newCol := currentRow+direction[0], currentCol+direction[1]
			cellKey := fmt.Sprintf("%d-%d", newRow, newCol)

			// Skip if: out of bounds, water can't flow, or already visited
			if newRow < 0 || newCol < 0 || newRow >= len(heights) || newCol >= len(heights[0]) ||
				heights[currentRow][currentCol] > heights[newRow][newCol] || reachableCells[cellKey] {
				continue
			}

			// Mark as reachable and add to queue for further exploration
			reachableCells[cellKey] = true
			queue = append(queue, []int{newRow, newCol})
		}
	}
}
