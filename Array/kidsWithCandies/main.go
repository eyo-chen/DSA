package main

// Time Complexity: O(n)
// Space Complexity: O(1)
func KidsWithCandies(candies []int, extraCandies int) []bool {
	// Find the maximum number of candies in the array
	maxCandies := candies[0]
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	// Create result array and check if each kid could have the maximum
	result := make([]bool, len(candies))
	for i, candy := range candies {
		// Check if current candy + extra candies would be >= maximum
		result[i] = candy+extraCandies >= maxCandies
	}

	return result
}
