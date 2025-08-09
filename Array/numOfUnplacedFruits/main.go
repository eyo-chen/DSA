package main

import "fmt"

// numOfUnplacedFruits places fruits into baskets following specific rules and returns
// the count of fruits that cannot be placed.
//
// Algorithm:
// - Process fruits in order from left to right
// - For each fruit, find the leftmost available basket with sufficient capacity
// - Mark used baskets as unavailable (-1) to prevent reuse
// - Count successful placements and return the difference
//
// Time Complexity: O(n²) where n is the length of arrays
//   - Outer loop: O(n) to process each fruit
//   - Inner loop: O(n) to find suitable basket in worst case
//
// Space Complexity: O(1) excluding input arrays
//   - Uses constant extra space, modifies input array in-place
func NumOfUnplacedFruits(fruits []int, baskets []int) int {
	placedFruitsCount := 0

	// Process each fruit in the given order
	for _, fruitQuantity := range fruits {
		// Find the leftmost available basket that can fit this fruit
		for basketIndex, basketCapacity := range baskets {
			// Check if basket is available (not marked as used) and has enough capacity
			if basketCapacity > 0 && basketCapacity >= fruitQuantity {
				placedFruitsCount++
				// Mark this basket as used by setting it to -1
				baskets[basketIndex] = -1
				break // Move to the next fruit
			}
		}
	}

	fmt.Println(baskets)

	// Return the number of unplaced fruits
	return len(baskets) - placedFruitsCount
}
