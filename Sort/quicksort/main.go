package main

func QuicKSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

// quickSortHelper is the recursive helper function for quickSort
func quickSortHelper(arr []int, left int, right int) {
	// Base case: if the left index is greater than or equal to the right index, return
	if left >= right {
		return
	}

	// Partition the array and get the pivot index
	pivotIdx := partition(arr, left, right)

	// Recursively sort the left and right subarrays
	quickSortHelper(arr, left, pivotIdx-1)
	quickSortHelper(arr, pivotIdx+1, right)

}

// partition takes last element as pivot, places the pivot element at its correct position,
// and places all smaller elements to the left of pivot and all greater elements to the right.
func partition(arr []int, left int, right int) int {
	// Choose the last element as pivot
	pivot := arr[right]

	// Initialize the index of the smaller element
	// The purpose of this index is to keep track where's the last smaller element than pivot
	// After the loop, it will be the index of the pivot element(swapped with the last element)
	idxSmallerThanPivot := left

	// Loop through the array from left to right-1(exclude the pivot)
	for i := left; i < right; i++ {
		// If the current element is greater than pivot, continue
		if arr[i] > pivot {
			continue
		}

		// Swap the current element with the last smaller element than pivot
		arr[idxSmallerThanPivot], arr[i] = arr[i], arr[idxSmallerThanPivot]

		// Move the index of the last smaller element to the right
		idxSmallerThanPivot++
	}

	// Swap the pivot element with the last smaller element than pivot
	arr[idxSmallerThanPivot], arr[right] = arr[right], arr[idxSmallerThanPivot]

	return idxSmallerThanPivot
}
