package main

/*
Solution Concept:
1. Since tiles form a CIRCLE, we need to check ALL possible groups of 3 consecutive tiles
2. For array of length n, there are exactly n such groups (each starting position creates one group)
3. Use modulo (%) to handle circular wrapping: when we reach the end, wrap back to beginning
4. An alternating group = middle tile has DIFFERENT color from both left and right neighbors
5. Pattern: 0-1-0 or 1-0-1 (middle differs from both sides)

Example: [0,1,0,0,1] has groups:
- Start 0: [0,1,0] ✓ alternating (middle=1, neighbors=0,0)
- Start 1: [1,0,0] ✗ not alternating (middle=0, right=0)
- Start 2: [0,0,1] ✗ not alternating (middle=0, left=0)
- Start 3: [0,1,0] ✓ alternating (middle=1, neighbors=0,0)
- Start 4: [1,0,1] ✓ alternating (middle=0, neighbors=1,1)
Total: 3 alternating groups
*/
func NumberOfAlternatingGroupsV2(colors []int) int {
	n := len(colors)
	count := 0

	// For each tile, check if it can be the middle of an alternating group
	for i := 0; i < n; i++ {
		// Get the neighbors (with circular indexing)
		leftNeighbor := colors[(i-1+n)%n] // Previous tile (wraps around)
		middle := colors[i]               // Current tile as middle
		rightNeighbor := colors[(i+1)%n]  // Next tile (wraps around)

		// Check if current tile forms alternating pattern with its neighbors
		if middle != leftNeighbor && middle != rightNeighbor {
			count++
		}
	}

	return count
}
