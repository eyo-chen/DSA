package main

import "fmt"

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	explored := map[string]bool{}

	helper(image, sr, sc, color, image[sr][sc], explored)

	return image
}

func helper(image [][]int, r int, c int, color int, startColor int, explored map[string]bool) {
	// check if the location is valid
	if r < 0 || c < 0 || r >= len(image) || c >= len(image[0]) {
		return
	}

	// check if the location is already explored
	key := fmt.Sprintf("%d-%d", r, c)
	if _, ok := explored[key]; ok {
		return
	}

	// check if the location has the same color as the startColor
	if image[r][c] != startColor {
		return
	}

	// fill the location with the new color
	image[r][c] = color

	// mark the location as explored
	explored[key] = true

	// explore the neighbors
	helper(image, r-1, c, color, startColor, explored)
	helper(image, r+1, c, color, startColor, explored)
	helper(image, r, c-1, color, startColor, explored)
	helper(image, r, c+1, color, startColor, explored)
}
