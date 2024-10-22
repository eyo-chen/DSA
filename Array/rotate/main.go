package main

func Rotate(matrix [][]int) {
	// Reverse the matrix upside down
	// Reverse column by column
	for top, down := 0, len(matrix)-1; top < down; top, down = top+1, down-1 {
		matrix[top], matrix[down] = matrix[down], matrix[top]
	}

	// Swap the symmetry elements
	for r := 0; r < len(matrix); r++ {
		for c := r + 1; c < len(matrix[r]); c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
}
