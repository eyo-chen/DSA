package main

import "fmt"

// MaxHeap is a generic max-heap data structure that maintains the heap property
// where each parent node is greater than or equal to its children based on the compare function.
type MaxHeap[T any] struct {
	elements    []T
	greaterThan func(T, T) bool // Returns true if first element should be prioritized over second
}

// NewMaxHeap creates a new max-heap with the specified initial capacity and comparison function.
// The compare function should return true if the first argument has higher priority (is "greater than") the second.
// Time complexity: O(1)
// Space complexity: O(capacity)
func NewMaxHeap[T any](capacity int, greaterThan func(T, T) bool) *MaxHeap[T] {
	return &MaxHeap[T]{
		elements:    make([]T, 0, capacity),
		greaterThan: greaterThan,
	}
}

// IsFull checks if the heap has reached its capacity.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MaxHeap[T]) IsFull() bool {
	return len(h.elements) == cap(h.elements)
}

// Insert adds a new element to the heap and maintains the heap property.
// Approach: Add element at the end and bubble it up to its correct position.
// Time complexity: O(log n) where n is the number of elements
// Space complexity: O(1) amortized
func (h *MaxHeap[T]) Insert(value T) {
	// Add the new element at the end of the heap
	h.elements = append(h.elements, value)

	// Restore heap property by moving the element up if needed
	h.bubbleUp(len(h.elements) - 1)
}

// Remove extracts and returns the maximum element (root) from the heap.
// Approach: Replace root with last element, remove last element, then bubble down.
// Time complexity: O(log n) where n is the number of elements
// Space complexity: O(1)
func (h *MaxHeap[T]) Remove() (T, error) {
	var zeroValue T

	// Check if heap is empty
	if len(h.elements) == 0 {
		return zeroValue, fmt.Errorf("heap is empty")
	}

	// Store the maximum value (root) to return
	maxValue := h.elements[0]

	lastIndex := len(h.elements) - 1

	// Move the last element to the root position
	h.swap(0, lastIndex)

	// Remove the last element (original root)
	h.elements = h.elements[:lastIndex]

	// Restore heap property by moving the new root down if needed
	if len(h.elements) > 0 {
		h.bubbleDown(0)
	}

	return maxValue, nil
}

// bubbleUp moves an element up the heap until the heap property is satisfied.
// Time complexity: O(log n)
// Space complexity: O(1)
func (h *MaxHeap[T]) bubbleUp(index int) {
	// Continue until we reach the root or find the correct position
	for index > 0 {
		parentIndex := h.parentIndex(index)

		// If current element has higher priority than parent, swap them
		if h.greaterThan(h.elements[index], h.elements[parentIndex]) {
			h.swap(index, parentIndex)
			index = parentIndex
		} else {
			// Heap property is satisfied
			break
		}
	}
}

// bubbleDown moves an element down the heap until the heap property is satisfied.
// Time complexity: O(log n)
// Space complexity: O(1)
func (h *MaxHeap[T]) bubbleDown(index int) {
	heapSize := len(h.elements)

	for {
		// Assume current index is the highest priority
		highestPriorityIndex := index
		leftChild, rightChild := h.childIndices(index)

		// Check if left child has higher priority
		if leftChild < heapSize && h.greaterThan(h.elements[leftChild], h.elements[highestPriorityIndex]) {
			highestPriorityIndex = leftChild
		}

		// Check if right child has higher priority
		if rightChild < heapSize && h.greaterThan(h.elements[rightChild], h.elements[highestPriorityIndex]) {
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
func (h *MaxHeap[T]) parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// childIndices returns the indices of the left and right children.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MaxHeap[T]) childIndices(parentIndex int) (leftChild int, rightChild int) {
	leftChild = 2*parentIndex + 1
	rightChild = 2*parentIndex + 2
	return
}

// swap exchanges two elements in the heap.
// Time complexity: O(1)
// Space complexity: O(1)
func (h *MaxHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// ============================================================================
// Example Usage
// ============================================================================

func main() {
	// MaxHeap example with integers
	maxHeap := NewMaxHeap(10, func(a, b int) bool {
		return a > b
	})

	maxHeap.Insert(5)
	maxHeap.Insert(3)
	maxHeap.Insert(7)
	maxHeap.Insert(1)

	max, _ := maxHeap.Remove()
	fmt.Printf("MaxHeap - Removed: %d\n", max) // Output: 7
}
