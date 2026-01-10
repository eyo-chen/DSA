package main

// Construct2DArray converts a 1D array into a 2D array with dimensions m x n.
// It fills the 2D array row by row from the original array.
//
// Approach: Create a 2D array and fill it by calculating the source index
// using row and column positions (index = col + n * row).
//
// Time Complexity: O(m * n) - we iterate through all m*n elements
// Space Complexity: O(m * n) - for the output 2D array (not counting input)
func Construct2DArray(original []int, m int, n int) [][]int {
	// Validate that the original array can form an m x n matrix
	if len(original) != m*n {
		return [][]int{}
	}

	// Initialize the 2D result array with m rows
	result := make([][]int, m)
	for i := range result {
		// Allocate n columns for each row
		result[i] = make([]int, n)
	}

	// Fill the 2D array by iterating through rows and columns
	for row := range m {
		for col := range n {
			// Calculate the index in the original array:
			// Each row contains n elements, so row i starts at index n*i
			// Then add the column offset
			sourceIndex := col + n*row
			result[row][col] = original[sourceIndex]
		}
	}

	return result
}

// Construct2DArrayOptimized converts a 1D array into a 2D array with dimensions m x n.
// It fills the 2D array by iterating through the original array once.
//
// Approach: Iterate through the original array sequentially and use division
// and modulo operations to determine the target position in the 2D array.
//
// Time Complexity: O(m * n) - single pass through all elements
// Space Complexity: O(m * n) - for the output 2D array (not counting input)
func Construct2DArrayOptimized(original []int, m int, n int) [][]int {
	// Validate that the original array can form an m x n matrix
	if len(original) != m*n {
		return [][]int{}
	}

	// Initialize the 2D result array with m rows
	result := make([][]int, m)
	for i := range result {
		// Allocate n columns for each row
		result[i] = make([]int, n)
	}

	// Fill the 2D array by iterating through the original array
	for idx := range original {
		// Calculate row: how many complete rows of n elements fit before this index
		row := idx / n
		// Calculate column: remainder gives position within the current row
		col := idx % n
		result[row][col] = original[idx]
	}

	return result
}
