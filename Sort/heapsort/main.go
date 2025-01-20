package main

// heapSort sorts an array using the Heap Sort algorithm.
// It first builds a max heap and then extracts elements one by one to sort the array.
func HeapSort(arr []int) {
	// Build a max heap.
	// Start from the last non-leaf node(middle of the array) and heapify each node.
	// The core idea is to ensure the given node(i) is at the correct position where it's greater than its children.
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

	// One by one extract elements from the heap.
	for i := len(arr) - 1; i >= 0; i-- {
		// Move current root (maximum) to the end of the array.
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify on the reduced heap to maintain the max heap property.
		heapify(arr, i, 0)
	}
}

// heapify ensures that the subtree rooted at index i obeys the max heap property.
// size is the size of the heap.
// i is the index of the root of the subtree.
func heapify(arr []int, size int, i int) {
	for {
		// Initialize the largest index to the current node.
		largestIdx := i

		// Calculate the indices of the left and right children.
		left, right := 2*i+1, 2*i+2

		// Check if the left child exists and is greater than the current node.
		if left < size && arr[left] > arr[largestIdx] {
			largestIdx = left
		}

		if right < size && arr[right] > arr[largestIdx] {
			largestIdx = right
		}

		// If the largest index is the current node, the heap property is already satisfied.
		if largestIdx == i {
			return
		}

		// Swap the current node with the largest child to maintain the max heap property.
		arr[i], arr[largestIdx] = arr[largestIdx], arr[i]

		// Move to the largest child to continue heapifying.
		i = largestIdx
	}
}
