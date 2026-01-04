package main

import "slices"

// MaxConsecutiveFloorsBruteForce finds the maximum consecutive floors without special floors
// by iterating through every floor from bottom to top.
// Time Complexity: O(n + (top - bottom)) where n is the length of special array
// Space Complexity: O(n) for the hash table
// Note: This approach will cause Time Limit Exceeded for large ranges (top - bottom can be up to 10^9)
func MaxConsecutiveFloorsBruteForce(bottom int, top int, special []int) int {
	maxConsecutive := 0

	// Build a hash table for O(1) lookup of special floors
	specialFloors := map[int]bool{}
	for _, floor := range special {
		specialFloors[floor] = true
	}

	// Iterate through every floor and count consecutive non-special floors
	currentConsecutive := 0
	for floor := bottom; floor <= top; floor++ {
		if specialFloors[floor] {
			// Found a special floor, update max and reset counter
			maxConsecutive = max(maxConsecutive, currentConsecutive)
			currentConsecutive = 0
		} else {
			// Regular floor, increment counter
			currentConsecutive++
		}
	}

	// Don't forget to check the final consecutive sequence
	maxConsecutive = max(maxConsecutive, currentConsecutive)

	return maxConsecutive
}

// MaxConsecutiveFloorsOptimal finds the maximum consecutive floors without special floors
// by calculating gaps between sorted special floors.
// Time Complexity: O(n log n) where n is the length of special array (dominated by sorting)
// Space Complexity: O(1) if we don't count the space used by sorting (or O(n) if we do)
func MaxConsecutiveFloorsOptimal(bottom int, top int, special []int) int {
	// Sort special floors to process gaps in order
	slices.Sort(special)

	maxConsecutive := 0

	// Check gap between bottom and first special floor
	maxConsecutive = max(maxConsecutive, special[0]-bottom)

	// Check gaps between consecutive special floors
	for i := 1; i < len(special); i++ {
		gapSize := special[i] - special[i-1] - 1
		maxConsecutive = max(maxConsecutive, gapSize)
	}

	// Check gap between last special floor and top
	maxConsecutive = max(maxConsecutive, top-special[len(special)-1])

	return maxConsecutive
}
