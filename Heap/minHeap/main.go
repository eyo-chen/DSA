package main

import "fmt"

// MinHeap is a generic min-heap data structure that maintains the heap property
// where each parent node is smaller than or equal to its children based on the compare function.
type MinHeap[T any] struct {
	elements []T
	lessThan func(T, T) bool // Returns true if first element should be prioritized over second
}

// NewMinHeap creates a new min-heap with the specified initial capacity and comparison function.
// The compare function should return true if the first argument has higher priority (is "less than") the second.
// Time complexity: O(1)
// Space complexity: O(capacity)
func NewMinHeap[T any](capacity int, lessThan func(T, T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		elements: make([]T, 0, capacity),
		lessThan: lessThan,
	}
}

// IsFull checks if the heap has reached its capacity.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MinHeap[T]) IsFull() bool {
	return len(h.elements) == cap(h.elements)
}

// Insert adds a new element to the heap and maintains the heap property.
// Approach: Add element at the end and bubble it up to its correct position.
// Time complexity: O(log n) where n is the number of elements
// Space complexity: O(1) amortized (may trigger slice reallocation)
func (h *MinHeap[T]) Insert(value T) {
	// Add the new element at the end of the heap
	h.elements = append(h.elements, value)

	// Restore heap property by moving the element up if needed
	h.bubbleUp(len(h.elements) - 1)
}

// Remove extracts and returns the minimum element (root) from the heap.
// Approach: Replace root with last element, remove last element, then bubble down.
// Time complexity: O(log n) where n is the number of elements
// Space complexity: O(1)
func (h *MinHeap[T]) Remove() (T, error) {
	var zeroValue T

	// Check if heap is empty
	if len(h.elements) == 0 {
		return zeroValue, fmt.Errorf("heap is empty")
	}

	// Store the minimum value (root) to return
	minValue := h.elements[0]

	lastIndex := len(h.elements) - 1

	// Move the last element to the root position
	h.swap(0, lastIndex)

	// Remove the last element (original root)
	h.elements = h.elements[:lastIndex]

	// Restore heap property by moving the new root down if needed
	if len(h.elements) > 0 {
		h.bubbleDown(0)
	}

	return minValue, nil
}

// bubbleUp moves an element up the heap until the heap property is satisfied.
// Used after insertion to restore heap order.
// Time complexity: O(log n)
// Space complexity: O(1)
func (h *MinHeap[T]) bubbleUp(index int) {
	// Continue until we reach the root or find the correct position
	for index > 0 {
		parentIndex := h.parentIndex(index)

		// If current element has higher priority than parent, swap them
		if h.lessThan(h.elements[index], h.elements[parentIndex]) {
			h.swap(index, parentIndex)
			index = parentIndex
		} else {
			// Heap property is satisfied
			break
		}
	}
}

// bubbleDown moves an element down the heap until the heap property is satisfied.
// Used after removal to restore heap order.
// Time complexity: O(log n)
// Space complexity: O(1)
func (h *MinHeap[T]) bubbleDown(index int) {
	heapSize := len(h.elements)

	for {
		// Assume current index is the highest priority
		highestPriorityIndex := index
		leftChild, rightChild := h.childIndices(index)

		// Check if left child has higher priority
		if leftChild < heapSize && h.lessThan(h.elements[leftChild], h.elements[highestPriorityIndex]) {
			highestPriorityIndex = leftChild
		}

		// Check if right child has higher priority
		if rightChild < heapSize && h.lessThan(h.elements[rightChild], h.elements[highestPriorityIndex]) {
			highestPriorityIndex = rightChild
		}

		// If current position is correct, we're done
		if highestPriorityIndex == index {
			break
		}

		// Swap with the child that has higher priority
		h.swap(index, highestPriorityIndex)
		index = highestPriorityIndex
	}
}

// parentIndex returns the index of the parent node.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MinHeap[T]) parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// childIndices returns the indices of the left and right children.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MinHeap[T]) childIndices(parentIndex int) (leftChild int, rightChild int) {
	leftChild = 2*parentIndex + 1
	rightChild = 2*parentIndex + 2
	return
}

// swap exchanges two elements in the heap.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MinHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// ============================================================================
// Example Usage
// ============================================================================

func main() {
	// MinHeap example with integers
	minHeap := NewMinHeap(10, func(a, b int) bool {
		return a < b
	})

	minHeap.Insert(5)
	minHeap.Insert(3)
	minHeap.Insert(7)
	minHeap.Insert(1)

	min, _ := minHeap.Remove()
	fmt.Printf("MinHeap - Removed: %d\n", min) // Output: 1
}
