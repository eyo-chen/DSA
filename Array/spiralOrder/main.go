package main

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	direction := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	total := len(matrix) * len(matrix[0])
	seen := map[string]bool{}
	ans := make([]int, 0, total)
	idx := 0
	r, c := 0, 0

	ans = append(ans, matrix[r][c])
	seen[fmt.Sprintf("%d-%d", r, c)] = true

	for len(ans) < total {
		dirR, dirC := direction[idx][0], direction[idx][1]
		nextR, nextC := r+dirR, c+dirC
		key := fmt.Sprintf("%d-%d", nextR, nextC)

		if nextR < 0 || nextC < 0 || nextR >= len(matrix) || nextC >= len(matrix[0]) || seen[key] {
			idx++
			if idx == len(direction) {
				idx = 0
			}

			// OR
			// idx = (idx + 1) % len(direction)
			// This works because idx + 1 = 4, we mod 4, we get 0.
			// If idx + 1 = 5, we mod 4, we get 1.
			// If idx + 1 = 6, we mod 4, we get 2.
			// So, we don't need to check if idx is out of bound.
			continue
		}

		r, c = nextR, nextC
		ans = append(ans, matrix[r][c])
		seen[fmt.Sprintf("%d-%d", r, c)] = true
	}

	return ans
}

func SpiralOrder1(matrix [][]int) []int {
	rowStart, rowEnd := 0, len(matrix)-1
	colStart, colEnd := 0, len(matrix[0])-1
	total := len(matrix) * len(matrix[0])
	ans := make([]int, 0, total)

	for len(ans) < total {
		// go right
		for c := colStart; c <= colEnd; c++ {
			ans = append(ans, matrix[rowStart][c])
		}
		rowStart++

		// go down
		for r := rowStart; r <= rowEnd; r++ {
			ans = append(ans, matrix[r][colEnd])
		}
		colEnd--

		// go left
		if rowStart <= rowEnd {
			for c := colEnd; c >= colStart; c-- {
				ans = append(ans, matrix[rowEnd][c])
			}
			rowEnd--
		}

		// go top
		if colStart <= colEnd {
			for r := rowEnd; r >= rowStart; r-- {
				ans = append(ans, matrix[r][colStart])
			}
			colStart++
		}
	}

	return ans
}
