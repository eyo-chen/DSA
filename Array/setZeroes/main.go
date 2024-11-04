package main

import "math"

type point struct {
	row int
	col int
}

// Create Array To Store Zero Position
func SetZeroes(matrix [][]int) {
	zeroPosition := []point{}

	// Find all zero positions
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				zeroPosition = append(zeroPosition, point{r, c})
			}
		}
	}

	// Set the row and col to zero for each zero position
	for _, p := range zeroPosition {
		row, col := p.row, p.col

		// Set the row to zero
		for r := 0; r < len(matrix); r++ {
			matrix[r][col] = 0
		}

		// Set the column to zero
		for c := 0; c < len(matrix[row]); c++ {
			matrix[row][c] = 0
		}
	}
}

// Create Two Arrays(First Row and First Col) To Store Zero Position
func SetZeroes2(matrix [][]int) {
	// firstRow[r] represents whether the r-th row need to be set to zero
	firstRow := make([]bool, len(matrix))
	// firstCol[c] represents whether the c-th column need to be set to zero
	firstCol := make([]bool, len(matrix[0]))

	// Find all zero positions
	// If matrix[r][c] is 0, set firstRow[r] and firstCol[c] to true
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				firstRow[r] = true
				firstCol[c] = true
			}
		}
	}

	// If firstRow[r] is true, set all element in row r to zero
	for r, isZero := range firstRow {
		if !isZero {
			continue
		}

		for c := range matrix[r] {
			matrix[r][c] = 0
		}
	}

	// If firstCol[c] is true, set all element in col c to zero
	for c, isZero := range firstCol {
		if !isZero {
			continue
		}

		for r := range matrix {
			matrix[r][c] = 0
		}
	}
}

// Using First Row and First Col To Store Zero Position
func SetZeroes3(matrix [][]int) {
	// isZeroAtFirstRow represents whether the first row need to be set to zero
	isZeroAtFirstRow := false
	// isZeroAtFirstCol represents whether the first col need to be set to zero
	isZeroAtFirstCol := false

	// Check if the first col need to be set to zero
	for r := range matrix {
		if matrix[r][0] == 0 {
			isZeroAtFirstCol = true
			break
		}
	}

	// Check if the first row need to be set to zero
	for c := range matrix[0] {
		if matrix[0][c] == 0 {
			isZeroAtFirstRow = true
			break
		}
	}

	// Check each cell in the matrix
	// If matrix[r][c] is 0, set matrix[r][0] and matrix[0][c] to 0
	// This is to indicate that the r-th row and c-th column need to be set to zero
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	// Traverse the inner matrix(excluding the first row and first col)
	// If either matrix[r][0] or matrix[0][c] is 0, set matrix[r][c] to 0
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// If isZeroAtFirstRow is true, set all element in first row to zero
	if isZeroAtFirstRow {
		for c := range matrix[0] {
			matrix[0][c] = 0
		}
	}

	// If isZeroAtFirstCol is true, set all element in first col to zero
	if isZeroAtFirstCol {
		for r := range matrix {
			matrix[r][0] = 0
		}
	}
}

// Mutate The Matrix To Special Value
func SetZeroes4(matrix [][]int) {
	// Traverse the matrix
	// If matrix[r][c] is 0, set all element in col c to a special value(math.MaxInt)
	for r := range matrix {
		for c := range matrix[r] {

			// matrix[r][c] is 0, set all element in col c to a special value(math.MaxInt)
			if matrix[r][c] == 0 {
				// Set all element in col c to a special value(math.MaxInt)
				for rr := range matrix {
					if matrix[rr][c] != 0 {
						matrix[rr][c] = math.MaxInt
					}
				}

				// Set all element in row r to a special value(math.MaxInt)
				for cc := range matrix[r] {
					if matrix[r][cc] != 0 {
						matrix[r][cc] = math.MaxInt
					}
				}
			}
		}
	}

	// Traverse the matrix again
	// If matrix[r][c] is a special value(math.MaxInt), set it to 0
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == math.MaxInt {
				matrix[r][c] = 0
			}
		}
	}
}
