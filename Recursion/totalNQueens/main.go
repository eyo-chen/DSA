package totalnqueens

import "C"

func totalNQueens(n int) int {
	colPosition := make([]int, 0, n)

	return helper(&colPosition, 0, n)
}

func helper(colPosition *[]int, row int, n int) int {
	if n == row {
		return 1
	}

	count := 0

	for c := 0; c < n; c++ {
		*colPosition = append(*colPosition, c)

		if isValid(row, c, colPosition) {
			count += helper(colPosition, row+1, n)
		}

		*colPosition = (*colPosition)[:len(*colPosition)-1]
	}

	return count
}

func isValid(row int, col int, colPosition *[]int) bool {
	for r := 0; r < len(*colPosition)-1; r++ {
		curColPosition := (*colPosition)[r]
		absColPosition := abs(curColPosition - col)
		absRowPosition := abs(row - r)

		if curColPosition == col {
			return false
		}
		if absColPosition == absRowPosition {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
