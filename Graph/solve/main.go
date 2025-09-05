package main

import "fmt"

// solveSurroundedRegions solves the "Surrounded Regions" problem by capturing all 'O' regions
// that are completely surrounded by 'X'. The approach uses reverse thinking:
// 1. Find all 'O' cells that are connected to the border (these are "safe" and cannot be captured)
// 2. Mark all safe regions using DFS traversal from border 'O' cells
// 3. Convert all remaining 'O' cells to 'X' since they are surrounded
func Solve(board [][]byte) {
	totalRows, totalCols := len(board), len(board[0])
	borderConnectedCells := map[string]bool{}

	// Check left and right borders for 'O' cells and mark their connected regions as safe
	for currentRow := range totalRows {
		// Left border
		if board[currentRow][0] == 'O' {
			markBorderConnectedRegion(board, currentRow, 0, borderConnectedCells)
		}
		// Right border
		if board[currentRow][totalCols-1] == 'O' {
			markBorderConnectedRegion(board, currentRow, totalCols-1, borderConnectedCells)
		}
	}

	// Check top and bottom borders for 'O' cells and mark their connected regions as safe
	for currentCol := range totalCols {
		// Top border
		if board[0][currentCol] == 'O' {
			markBorderConnectedRegion(board, 0, currentCol, borderConnectedCells)
		}
		// Bottom border
		if board[totalRows-1][currentCol] == 'O' {
			markBorderConnectedRegion(board, totalRows-1, currentCol, borderConnectedCells)
		}
	}

	// Convert all 'O' cells that are not connected to border (surrounded regions) to 'X'
	for currentRow := 1; currentRow < totalRows-1; currentRow++ {
		for currentCol := 1; currentCol < totalCols-1; currentCol++ {
			cellKey := fmt.Sprintf("%d-%d", currentRow, currentCol)
			// If cell is 'O' and not in safe region, it's surrounded - convert to 'X'
			if board[currentRow][currentCol] == 'O' && !borderConnectedCells[cellKey] {
				board[currentRow][currentCol] = 'X'
			}
		}
	}
}

// Direction vectors for 4-directional movement (right, left, down, up)
var movementDirections = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// markBorderConnectedRegion performs DFS to mark all 'O' cells that are connected to border cells.
// These cells form regions that cannot be captured since they have a path to the border.
// Parameters:
//   - board: the game board
//   - currentRow, currentCol: current position being explored
//   - borderConnectedCells: map to track all cells connected to border
func markBorderConnectedRegion(board [][]byte, currentRow int, currentCol int, borderConnectedCells map[string]bool) {
	cellKey := fmt.Sprintf("%d-%d", currentRow, currentCol)

	// Base cases: stop if out of bounds, not an 'O' cell, or already visited
	if currentRow < 0 || currentCol < 0 ||
		currentRow >= len(board) || currentCol >= len(board[0]) ||
		board[currentRow][currentCol] != 'O' ||
		borderConnectedCells[cellKey] {
		return
	}

	// Mark current cell as connected to border (safe from capture)
	borderConnectedCells[cellKey] = true

	// Explore all 4 adjacent cells recursively
	for _, direction := range movementDirections {
		nextRow, nextCol := currentRow+direction[0], currentCol+direction[1]
		markBorderConnectedRegion(board, nextRow, nextCol, borderConnectedCells)
	}
}
