package main

// Brute Force Approach 1: Using Hash Map
//
// Approach: For each starting position, find the longest contiguous subarray
// that contains at most 2 distinct fruit types using a hash map to track types.
//
// Algorithm:
// 1. For each possible starting position i
// 2. Use a hash map to track distinct fruit types in current subarray
// 3. Extend the subarray from i until we encounter a third fruit type
// 4. Track the maximum length found
//
// Time Complexity: O(n²) - nested loops, outer loop runs n times, inner loop runs up to n times
// Space Complexity: O(1) - hash map stores at most 3 fruit types (2 valid + 1 that breaks the condition)
func TotalFruit(fruits []int) int {
	maxFruits := 0

	// Try each position as a potential starting point
	for startIndex := 0; startIndex < len(fruits); startIndex++ {
		// Track distinct fruit types in current subarray
		fruitTypes := map[int]bool{
			fruits[startIndex]: true,
		}

		// Count fruits in current subarray starting from startIndex
		currentCount := 1

		// Extend subarray from startIndex + 1 onwards
		for endIndex := startIndex + 1; endIndex < len(fruits); endIndex++ {
			currentFruit := fruits[endIndex]

			// Check if adding this fruit would exceed 2 distinct types
			if _, exists := fruitTypes[currentFruit]; !exists {
				// If we already have 2 types and this is a new type, stop
				if len(fruitTypes) == 2 {
					break
				}
			}

			// Add current fruit type to our tracking
			fruitTypes[currentFruit] = true
			currentCount++
		}

		// Update maximum fruits collected
		maxFruits = max(maxFruits, currentCount)
	}

	return maxFruits
}

// Brute Force Approach 2: Optimized Space (O(1) space)
//
// Approach: For each starting position, track exactly 2 fruit types using
// two variables instead of a hash map, making space usage constant.
//
// Algorithm:
// 1. For each starting position, initialize first fruit type
// 2. Use two variables (firstType, secondType) to track the 2 allowed types
// 3. Extend subarray until we encounter a third distinct type
// 4. Track maximum length found
//
// Time Complexity: O(n²) - nested loops, outer loop runs n times, inner loop runs up to n times
// Space Complexity: O(1) - only uses two variables to track fruit types
func TotalFruit3(fruits []int) int {
	if len(fruits) == 1 {
		return 1
	}

	maxFruits := 0

	// Try each position as a potential starting point
	for startIndex := 0; startIndex < len(fruits)-1; startIndex++ {
		// Initialize the two fruit types we can collect
		// firstType is set to the fruit at starting position
		// secondType is uninitialized (-1) until we find a different fruit
		firstType, secondType := fruits[startIndex], -1
		currentCount := 1

		// Extend subarray from current starting position
		for endIndex := startIndex + 1; endIndex < len(fruits); endIndex++ {
			currentFruit := fruits[endIndex]

			// If current fruit matches one of our two allowed types
			if currentFruit == firstType || currentFruit == secondType {
				currentCount++
				continue
			}

			// If we haven't set the second type yet, set it now
			if secondType == -1 {
				secondType = currentFruit
				currentCount++
				continue
			}

			// Current fruit is a third type, cannot continue this subarray
			break
		}

		// Update maximum fruits collected
		maxFruits = max(maxFruits, currentCount)
	}

	return maxFruits
}

// Optimized Approach: Sliding Window with Hash Map
//
// Approach: Use sliding window technique with two pointers (left, right) and
// a hash map to track fruit frequencies in the current window.
//
// Algorithm:
// 1. Expand window by moving right pointer and adding fruits to hash map
// 2. When window contains more than 2 fruit types, shrink from left
// 3. Track maximum window size throughout the process
// 4. This ensures we always maintain at most 2 distinct fruit types
//
// Time Complexity: O(n) - each element is visited at most twice (once by right pointer, once by left pointer)
// Space Complexity: O(1) - hash map stores at most 3 fruit types at any time
func TotalFruit2(fruits []int) int {
	maxFruits := 0
	leftPointer, rightPointer := 0, 0
	fruitFrequency := map[int]int{}

	for rightPointer < len(fruits) {
		currentFruit := fruits[rightPointer]

		// Add current fruit to our sliding window
		fruitFrequency[currentFruit]++

		// Shrink window from left while we have more than 2 distinct fruit types
		for len(fruitFrequency) > 2 {
			leftFruit := fruits[leftPointer]

			// Remove one occurrence of the leftmost fruit
			fruitFrequency[leftFruit]--

			// If frequency becomes 0, remove this fruit type completely
			if fruitFrequency[leftFruit] == 0 {
				delete(fruitFrequency, leftFruit)
			}

			// Move left boundary of window
			leftPointer++
		}

		// Update maximum fruits collected (current window size)
		maxFruits = max(maxFruits, rightPointer-leftPointer+1)

		// Expand window by moving right pointer
		rightPointer++
	}

	return maxFruits
}
