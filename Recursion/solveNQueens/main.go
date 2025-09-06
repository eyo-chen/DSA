package main

// SolveNQueensWithBoolMatrix solves the N-Queens puzzle using a boolean matrix approach.
// This method uses a 2D boolean array to track queen positions, then converts to string format.
//
// Algorithm:
// 1. Create an n×n boolean matrix initialized to false
// 2. Use backtracking to try placing queens row by row
// 3. For each position, validate no conflicts (column, diagonal, anti-diagonal)
// 4. If valid, place queen (set to true) and recurse to next row
// 5. If invalid or backtracking, remove queen (set to false)
// 6. Convert final valid configurations to string representation
//
// Time Complexity: O(N!) - exploring all possible queen placements
// Space Complexity: O(N²) - for the board matrix
func SolveNQueensWithBoolMatrix(n int) [][]string {
	// Initialize n×n boolean matrix to track queen positions
	// false = empty square, true = queen placed
	queenPositions := make([][]bool, n)
	for row := range queenPositions {
		queenPositions[row] = make([]bool, n)
	}

	// Store all valid solutions
	solutions := [][]string{}

	// Start backtracking from row 0
	placeQueensRecursively(n, 0, queenPositions, &solutions)
	return solutions
}

// placeQueensRecursively attempts to place queens using backtracking approach.
// It tries to place one queen per row, validating each placement before proceeding.
//
// Parameters:
//   - boardSize: size of the n×n chessboard
//   - currentRow: the row we're currently trying to place a queen in
//   - queenPositions: 2D boolean matrix tracking current queen placements
//   - solutions: pointer to slice storing all valid board configurations
func placeQueensRecursively(boardSize int, currentRow int, queenPositions [][]bool, solutions *[][]string) {
	// Base case: if we've successfully placed queens in all rows
	if currentRow >= boardSize {
		// Convert current valid configuration to string format and save
		*solutions = append(*solutions, convertBoolMatrixToStrings(queenPositions))
		return
	}

	// Try placing a queen in each column of the current row
	for currentCol := range boardSize {
		// Check if placing queen at (currentRow, currentCol) is valid
		if isPositionSafe(queenPositions, currentRow, currentCol) {
			// Place queen at this position
			queenPositions[currentRow][currentCol] = true

			// Recursively try to place queens in subsequent rows
			placeQueensRecursively(boardSize, currentRow+1, queenPositions, solutions)

			// Backtrack: remove queen to try next position
			queenPositions[currentRow][currentCol] = false
		}
	}
}

// convertBoolMatrixToStrings converts a 2D boolean matrix to string representation.
// true values become 'Q' (queen), false values become '.' (empty square).
func convertBoolMatrixToStrings(queenPositions [][]bool) []string {
	boardRows := make([]string, len(queenPositions))

	for rowIdx := range queenPositions {
		// Build current row as byte array for efficiency
		currentRowBytes := make([]byte, len(queenPositions[rowIdx]))

		for colIdx := range queenPositions[rowIdx] {
			if queenPositions[rowIdx][colIdx] {
				currentRowBytes[colIdx] = 'Q' // Queen position
			} else {
				currentRowBytes[colIdx] = '.' // Empty square
			}
		}

		// Convert byte array to string
		boardRows[rowIdx] = string(currentRowBytes)
	}

	return boardRows
}

// isPositionSafe checks if placing a queen at (targetRow, targetCol) conflicts with existing queens.
// Validates three attack patterns: same column, main diagonal, and anti-diagonal.
//
// Queens can attack in 8 directions, but since we place row by row from top to bottom,
// we only need to check upward directions (positions already filled).
func isPositionSafe(queenPositions [][]bool, targetRow int, targetCol int) bool {
	boardSize := len(queenPositions)

	// Check 1: Column conflict - no queen should be in the same column above
	for row := range targetRow {
		if queenPositions[row][targetCol] {
			return false // Found queen in same column
		}
	}

	// Check 2: Main diagonal conflict (top-left ↖︎ bottom-right)
	// Move diagonally up-left from target position
	for row, col := targetRow-1, targetCol-1; row >= 0 && col >= 0; row, col = row-1, col-1 {
		if queenPositions[row][col] {
			return false // Found queen on main diagonal
		}
	}

	// Check 3: Anti-diagonal conflict (top-right ↗ bottom-left)
	// Move diagonally up-right from target position
	for row, col := targetRow-1, targetCol+1; row >= 0 && col < boardSize; row, col = row-1, col+1 {
		if queenPositions[row][col] {
			return false // Found queen on anti-diagonal
		}
	}

	return true // No conflicts found, position is safe
}

// SolveNQueensWithByteMatrix solves the N-Queens puzzle using a byte matrix approach.
// This method directly uses characters ('Q' for queens, '.' for empty) for more efficient string conversion.
//
// Algorithm:
// 1. Create an n×n byte matrix initialized with '.' characters
// 2. Use backtracking to try placing queens ('Q') row by row
// 3. For each position, validate no conflicts with existing queens
// 4. If valid, place 'Q' and recurse; if invalid, backtrack by restoring '.'
// 5. Convert byte matrix directly to strings (more efficient than boolean approach)
//
// Time Complexity: O(N!) - exploring all possible queen placements
// Space Complexity: O(N²) - for the board matrix
// Performance: Faster than boolean approach due to direct string conversion
func SolveNQueensWithByteMatrix(n int) [][]string {
	// Initialize n×n byte matrix with '.' representing empty squares
	chessboard := make([][]byte, n)
	for row := range chessboard {
		chessboard[row] = make([]byte, n)
		// Fill each row with '.' characters
		for col := range chessboard[row] {
			chessboard[row][col] = '.'
		}
	}

	// Store all valid solutions
	solutions := [][]string{}

	// Start backtracking from row 0
	solveQueenPlacement(n, 0, chessboard, &solutions)
	return solutions
}

// solveQueenPlacement uses backtracking to find all valid queen placements.
// It places queens row by row, ensuring no conflicts at each step.
//
// Parameters:
//   - boardSize: size of the n×n chessboard
//   - currentRow: the row we're currently trying to place a queen in
//   - chessboard: 2D byte matrix representing the current board state
//   - solutions: pointer to slice storing all valid board configurations
func solveQueenPlacement(boardSize int, currentRow int, chessboard [][]byte, solutions *[][]string) {
	// Base case: successfully placed queens in all rows
	if currentRow >= boardSize {
		// Convert current board state to string format and save
		*solutions = append(*solutions, convertByteMatrixToStrings(chessboard))
		return
	}

	// Try placing a queen in each column of the current row
	for currentCol := range chessboard[currentRow] {
		// Validate that this position doesn't conflict with existing queens
		if isQueenPlacementValid(chessboard, currentRow, currentCol) {
			// Place queen at current position
			chessboard[currentRow][currentCol] = 'Q'

			// Recursively solve for the next row
			solveQueenPlacement(boardSize, currentRow+1, chessboard, solutions)

			// Backtrack: remove queen to explore other possibilities
			chessboard[currentRow][currentCol] = '.'
		}
	}
}

// convertByteMatrixToStrings efficiently converts byte matrix to string slice.
// Since each row is already a byte array, conversion is direct and fast.
func convertByteMatrixToStrings(chessboard [][]byte) []string {
	boardRepresentation := make([]string, len(chessboard))

	for rowIdx := range chessboard {
		// Direct conversion from byte slice to string - very efficient
		boardRepresentation[rowIdx] = string(chessboard[rowIdx])
	}

	return boardRepresentation
}

// isQueenPlacementValid checks if placing a queen at (targetRow, targetCol) is safe.
// Examines three potential conflict directions: column, main diagonal, and anti-diagonal.
//
// Since queens are placed row by row from top to bottom, we only check upward directions
// where queens have already been placed.
func isQueenPlacementValid(chessboard [][]byte, targetRow int, targetCol int) bool {
	// Check 1: Column safety - scan upward in the same column
	for row := range targetRow {
		if chessboard[row][targetCol] == 'Q' {
			return false // Column conflict detected
		}
	}

	// Check 2: Main diagonal safety (↖ direction)
	// Check all positions diagonally up-left from target
	for row, col := targetRow-1, targetCol-1; row >= 0 && col >= 0; row, col = row-1, col-1 {
		if chessboard[row][col] == 'Q' {
			return false // Main diagonal conflict detected
		}
	}

	// Check 3: Anti-diagonal safety (↗ direction)
	// Check all positions diagonally up-right from target
	for row, col := targetRow-1, targetCol+1; row >= 0 && col < len(chessboard[targetRow]); row, col = row-1, col+1 {
		if chessboard[row][col] == 'Q' {
			return false // Anti-diagonal conflict detected
		}
	}

	return true // No conflicts found - safe to place queen
}
