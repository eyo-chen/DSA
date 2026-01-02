package main

// AsteroidCollision simulates asteroid collisions in a row.
// Approach: Use a stack to track surviving asteroids. Process each asteroid left-to-right:
//   - Right-moving asteroids (+) are added to stack
//   - Left-moving asteroids (-) collide with right-moving asteroids in stack
//   - Collision rules: larger survives, equal sizes both explode
//
// Time Complexity: O(n) - each asteroid is pushed/popped at most once
// Space Complexity: O(n) - stack stores surviving asteroids
func AsteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, current := range asteroids {
		survives := true

		// Handle collisions: current asteroid moving left meets right-moving asteroids in stack
		for len(stack) > 0 && current < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]

			// Case 1: Equal size - both asteroids explode
			if current == -top {
				stack = stack[:len(stack)-1]
				survives = false
				break
			}

			// Case 2: Current asteroid is smaller - it explodes
			if absInt(current) < absInt(top) {
				survives = false
				break
			}

			// Case 3: Current asteroid is larger - top explodes, continue checking
			stack = stack[:len(stack)-1]
		}

		// Add current asteroid to stack if it survived all collisions
		if survives {
			stack = append(stack, current)
		}
	}

	return stack
}

// absInt returns the absolute value of an integer
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
