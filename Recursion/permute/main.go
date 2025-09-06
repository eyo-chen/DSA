package main

// Permute generates all unique permutations of the given integer slice.
//
// Approach: Uses backtracking with a boolean array to track used elements.
// For each position in the permutation, we try each unused number from the
// input array, mark it as used, recursively build the rest of the permutation,
// then backtrack by marking it as unused for the next iteration.
//
// Time Complexity: O(n! * n) where n is the length of nums
// - There are n! permutations to generate
// - Each permutation takes O(n) time to copy into the result
//
// Space Complexity: O(n) for recursion stack depth and the used array
// (not counting the output space which is O(n! * n))
func Permute(nums []int) [][]int {
	result := [][]int{}
	used := make([]bool, len(nums)) // Track which elements are already used
	buildPermutations(nums, used, &result, []int{})
	return result
}

// buildPermutations recursively constructs all permutations using backtracking
func buildPermutations(nums []int, used []bool, result *[][]int, currentPermutation []int) {
	// Base case: if current permutation is complete, add it to result
	if len(currentPermutation) == len(nums) {
		// Create a copy to avoid reference issues when currentPermutation is modified
		completedPermutation := make([]int, len(currentPermutation))
		copy(completedPermutation, currentPermutation)
		*result = append(*result, completedPermutation)
		return
	}

	// Try each number in the input array
	for index, number := range nums {
		// Skip if this number is already used in current permutation
		if used[index] {
			continue
		}

		// Choose: mark this number as used and add it to current permutation
		used[index] = true

		// Explore: recursively build the rest of the permutation
		buildPermutations(nums, used, result, append(currentPermutation, number))

		// Backtrack: unmark this number so it can be used in other permutations
		used[index] = false
	}
}

// permute generates all unique permutations of the given integer slice.
//
// Approach: Uses backtracking with frequency counting to handle duplicates.
// We count the frequency of each unique number, then recursively build
// permutations by selecting available numbers (with frequency > 0) and
// backtracking when we reach the target length.
//
// Time Complexity: O(n! * n) where n is the length of nums
// - There are at most n! unique permutations
// - Each permutation takes O(n) time to copy into the result
//
// Space Complexity: O(n) for recursion stack depth and temporary arrays
// (not counting the output space which is O(n! * n))
func Permute1(nums []int) [][]int {
	// Count frequency of each number to handle duplicates
	numFrequency := map[int]int{}
	for _, num := range nums {
		numFrequency[num]++
	}

	var result [][]int
	generatePermutations(numFrequency, &result, []int{}, len(nums))
	return result
}

// generatePermutations recursively builds all unique permutations using backtracking
func generatePermutations(numFrequency map[int]int, result *[][]int, currentPath []int, targetLength int) {
	// Base case: if current permutation reaches target length, add it to result
	if len(currentPath) == targetLength {
		// Create a copy to avoid reference issues when backtracking modifies currentPath
		permutation := make([]int, targetLength)
		copy(permutation, currentPath)
		*result = append(*result, permutation)
		return
	}

	// Try each unique number that still has available count
	for value, count := range numFrequency {
		// Skip numbers that are exhausted
		if count == 0 {
			continue
		}

		// Choose: add current value to path and decrease its frequency
		numFrequency[value]--

		// Explore: recursively generate permutations with updated state
		generatePermutations(numFrequency, result, append(currentPath, value), targetLength)

		// Backtrack: restore the frequency count for next iteration
		numFrequency[value]++
	}
}
