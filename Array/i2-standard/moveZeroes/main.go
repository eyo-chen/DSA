package main

// MoveZeroes moves all zeros to the end while maintaining the relative order
// of non-zero elements using extra space.
// Approach: Collect non-zero elements first, then append zeros.
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(n) - requires additional array for non-zero elements
func MoveZeroes(nums []int) {
	// Count how many zeros exist in the array
	zeroCount := 0

	// Collect all non-zero elements while preserving their order
	nonZeroElements := []int{}

	// Iterate through the array to separate zeros from non-zeros
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			nonZeroElements = append(nonZeroElements, num)
		}
	}

	// Create an array filled with zeros
	zeros := make([]int, zeroCount)

	// Append all zeros to the end of non-zero elements
	nonZeroElements = append(nonZeroElements, zeros...)

	// Copy the rearranged elements back to the original array
	copy(nums, nonZeroElements)
}

// MoveZeroesInPlace moves all zeros to the end while maintaining the relative
// order of non-zero elements using a two-pass approach.
// Approach: First pass fills non-zeros from the start, second pass fills zeros at the end.
// Time Complexity: O(n) - two passes through the array
// Space Complexity: O(1) - modifies array in-place with only a pointer variable
func MoveZeroesInPlace(nums []int) {
	// Track the position where the next non-zero element should be placed
	writeIndex := 0

	// First pass: move all non-zero elements to the front
	for _, num := range nums {
		if num != 0 {
			nums[writeIndex] = num
			writeIndex++
		}
	}

	// Second pass: fill remaining positions with zeros
	for i := writeIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}

// MoveZeroesOptimal moves all zeros to the end while maintaining the relative
// order of non-zero elements using an optimal single-pass swap approach.
// Approach: Use two pointers - one scanning the array, one tracking where to place non-zeros.
// When a non-zero is found, swap it with the element at the placement position.
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(1) - modifies array in-place with only a pointer variable
func MoveZeroesOptimal(nums []int) {
	// Track the boundary between non-zero and zero elements
	// All elements before this index are non-zero
	nonZeroBoundary := 0

	for currentIndex := range len(nums) {
		// When we find a non-zero element, swap it to the non-zero section
		if nums[currentIndex] != 0 {
			// Swap current element with the element at the boundary
			nums[currentIndex], nums[nonZeroBoundary] = nums[nonZeroBoundary], nums[currentIndex]
			// Expand the non-zero section by moving the boundary forward
			nonZeroBoundary++
		}
	}
}
