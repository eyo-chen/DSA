package main

// FindPeakIndexInMountainArray finds the index of the peak element in a mountain array.
// A mountain array is defined as an array that:
//   - Has at least 3 elements
//   - Has elements that strictly increase up to a peak, then strictly decrease
//
// Approach: Binary search
//   - Compare the middle element with its neighbors
//   - If mid > mid-1, the peak is at mid or to the right (ascending slope)
//   - If mid < mid-1, the peak is to the left (descending slope)
//
// Time Complexity: O(log n) - binary search halves the search space each iteration
// Space Complexity: O(1) - only uses constant extra space
func FindPeakIndexInMountainArray(arr []int) int {
	leftBound, rightBound := 0, len(arr)-1

	for {
		// Calculate middle index to avoid integer overflow
		midIndex := leftBound + (rightBound-leftBound)/2

		// Check if current element is the peak
		// Peak condition: greater than both neighbors
		if arr[midIndex] > arr[midIndex-1] && arr[midIndex] > arr[midIndex+1] {
			return midIndex
		}

		// Determine which half contains the peak
		if arr[midIndex] > arr[midIndex-1] {
			// We're on the ascending slope, peak is to the right
			leftBound = midIndex
		} else {
			// We're on the descending slope, peak is to the left
			rightBound = midIndex
		}
	}
}
