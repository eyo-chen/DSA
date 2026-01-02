package main

import (
	"testing"
)

// TestNewMaxHeap tests the creation of a new max heap
func TestNewMaxHeap(t *testing.T) {
	t.Run("Create heap with zero capacity", func(t *testing.T) {
		heap := NewMaxHeap(0, func(a, b int) bool { return a > b })
		if heap == nil {
			t.Error("Expected non-nil heap")
		}
		if len(heap.elements) != 0 {
			t.Errorf("Expected length 0, got %d", len(heap.elements))
		}
	})

	t.Run("Create heap with positive capacity", func(t *testing.T) {
		heap := NewMaxHeap(10, func(a, b int) bool { return a > b })
		if cap(heap.elements) != 10 {
			t.Errorf("Expected capacity 10, got %d", cap(heap.elements))
		}
	})
}

// TestIsFull tests the IsFull method
func TestIsFull(t *testing.T) {
	t.Run("Empty heap is not full", func(t *testing.T) {
		heap := NewMaxHeap(3, func(a, b int) bool { return a > b })
		if heap.IsFull() {
			t.Error("Empty heap should not be full")
		}
	})

	t.Run("Partially filled heap is not full", func(t *testing.T) {
		heap := NewMaxHeap(3, func(a, b int) bool { return a > b })
		heap.Insert(1)
		heap.Insert(2)
		if heap.IsFull() {
			t.Error("Partially filled heap should not be full")
		}
	})

	t.Run("Heap at capacity is full", func(t *testing.T) {
		heap := NewMaxHeap(3, func(a, b int) bool { return a > b })
		heap.Insert(1)
		heap.Insert(2)
		heap.Insert(3)
		if !heap.IsFull() {
			t.Error("Heap at capacity should be full")
		}
	})
}

// TestInsertSingleElement tests inserting a single element
func TestInsertSingleElement(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })
	heap.Insert(5)

	if len(heap.elements) != 1 {
		t.Errorf("Expected length 1, got %d", len(heap.elements))
	}
	if heap.elements[0] != 5 {
		t.Errorf("Expected root to be 5, got %d", heap.elements[0])
	}
}

// TestInsertMultipleElements tests inserting multiple elements
func TestInsertMultipleElements(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	values := []int{5, 3, 7, 1, 9, 2, 8}
	for _, v := range values {
		heap.Insert(v)
	}

	// Root should be the maximum value
	if heap.elements[0] != 9 {
		t.Errorf("Expected root to be 9, got %d", heap.elements[0])
	}

	// Verify heap property: parent >= children
	if !verifyMaxHeapProperty(heap) {
		t.Error("Heap property violated after insertions")
	}
}

// TestInsertDuplicates tests inserting duplicate values
func TestInsertDuplicates(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	heap.Insert(5)
	heap.Insert(5)
	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(7)

	if len(heap.elements) != 5 {
		t.Errorf("Expected length 5, got %d", len(heap.elements))
	}

	if !verifyMaxHeapProperty(heap) {
		t.Error("Heap property violated with duplicates")
	}
}

// TestInsertNegativeNumbers tests inserting negative numbers
func TestInsertNegativeNumbers(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	heap.Insert(-5)
	heap.Insert(-3)
	heap.Insert(-7)
	heap.Insert(0)

	if heap.elements[0] != 0 {
		t.Errorf("Expected root to be 0, got %d", heap.elements[0])
	}

	if !verifyMaxHeapProperty(heap) {
		t.Error("Heap property violated with negative numbers")
	}
}

// TestRemoveFromEmptyHeap tests removing from an empty heap
func TestRemoveFromEmptyHeap(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	_, err := heap.Remove()
	if err == nil {
		t.Error("Expected error when removing from empty heap")
	}
}

// TestRemoveSingleElement tests removing the only element
func TestRemoveSingleElement(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })
	heap.Insert(5)

	val, err := heap.Remove()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 5 {
		t.Errorf("Expected removed value to be 5, got %d", val)
	}
	if len(heap.elements) != 0 {
		t.Errorf("Expected empty heap, got length %d", len(heap.elements))
	}
}

// TestRemoveMultipleElements tests removing multiple elements
func TestRemoveMultipleElements(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	values := []int{5, 3, 7, 1, 9, 2, 8}
	for _, v := range values {
		heap.Insert(v)
	}

	expectedOrder := []int{9, 8, 7, 5, 3, 2, 1}
	for i, expected := range expectedOrder {
		val, err := heap.Remove()
		if err != nil {
			t.Errorf("Unexpected error at index %d: %v", i, err)
		}
		if val != expected {
			t.Errorf("At index %d: expected %d, got %d", i, expected, val)
		}

		// Verify heap property after each removal
		if len(heap.elements) > 0 && !verifyMaxHeapProperty(heap) {
			t.Errorf("Heap property violated after removing element at index %d", i)
		}
	}

	// Heap should be empty now
	if len(heap.elements) != 0 {
		t.Errorf("Expected empty heap, got length %d", len(heap.elements))
	}
}

// TestRemoveAllThenInsert tests removing all elements then inserting again
func TestRemoveAllThenInsert(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	// Insert and remove all
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)

	heap.Remove()
	heap.Remove()
	heap.Remove()

	// Insert again
	heap.Insert(5)
	heap.Insert(4)

	val, _ := heap.Remove()
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}

// TestInsertAndRemoveMixed tests mixed insert and remove operations
func TestInsertAndRemoveMixed(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b int) bool { return a > b })

	heap.Insert(5)
	heap.Insert(3)

	val, _ := heap.Remove()
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	heap.Insert(7)
	heap.Insert(2)

	val, _ = heap.Remove()
	if val != 7 {
		t.Errorf("Expected 7, got %d", val)
	}

	if !verifyMaxHeapProperty(heap) {
		t.Error("Heap property violated after mixed operations")
	}
}

// TestCustomComparator tests using a custom comparator for min-heap behavior
func TestCustomComparator(t *testing.T) {
	// Create a min-heap by reversing the comparator
	minHeap := NewMaxHeap(10, func(a, b int) bool { return a < b })

	minHeap.Insert(5)
	minHeap.Insert(3)
	minHeap.Insert(7)
	minHeap.Insert(1)

	// Should return minimum value (1) first
	val, _ := minHeap.Remove()
	if val != 1 {
		t.Errorf("Expected 1 (min value), got %d", val)
	}

	val, _ = minHeap.Remove()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

// TestStringHeap tests heap with string type
func TestStringHeap(t *testing.T) {
	heap := NewMaxHeap(10, func(a, b string) bool { return a > b })

	heap.Insert("apple")
	heap.Insert("zebra")
	heap.Insert("banana")
	heap.Insert("mango")

	val, _ := heap.Remove()
	if val != "zebra" {
		t.Errorf("Expected 'zebra', got '%s'", val)
	}

	val, _ = heap.Remove()
	if val != "mango" {
		t.Errorf("Expected 'mango', got '%s'", val)
	}
}

// TestStructHeap tests heap with custom struct type
func TestStructHeap(t *testing.T) {
	type Task struct {
		Name     string
		Priority int
	}

	heap := NewMaxHeap(10, func(a, b Task) bool { return a.Priority > b.Priority })

	heap.Insert(Task{"Low", 1})
	heap.Insert(Task{"High", 10})
	heap.Insert(Task{"Medium", 5})

	val, _ := heap.Remove()
	if val.Priority != 10 {
		t.Errorf("Expected priority 10, got %d", val.Priority)
	}
}

// TestLargeHeap tests with a large number of elements
func TestLargeHeap(t *testing.T) {
	heap := NewMaxHeap(1000, func(a, b int) bool { return a > b })

	// Insert 1000 elements
	for i := 0; i < 1000; i++ {
		heap.Insert(i)
	}

	// Verify all elements come out in descending order
	prev := 1000
	for i := 0; i < 1000; i++ {
		val, err := heap.Remove()
		if err != nil {
			t.Errorf("Unexpected error at index %d: %v", i, err)
		}
		if val >= prev {
			t.Errorf("Elements not in descending order: %d followed by %d", prev, val)
		}
		prev = val
	}
}

// TestHeapWithSingleCapacity tests heap with capacity of 1
func TestHeapWithSingleCapacity(t *testing.T) {
	heap := NewMaxHeap(1, func(a, b int) bool { return a > b })

	heap.Insert(5)
	if !heap.IsFull() {
		t.Error("Heap should be full after one insertion")
	}

	val, _ := heap.Remove()
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	// Can insert again (heap grows beyond initial capacity)
	heap.Insert(3)
	val, _ = heap.Remove()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

// Helper function to verify max heap property
func verifyMaxHeapProperty(heap *MaxHeap[int]) bool {
	for i := 0; i < len(heap.elements); i++ {
		leftChild := 2*i + 1
		rightChild := 2*i + 2

		if leftChild < len(heap.elements) {
			if heap.greaterThan(heap.elements[leftChild], heap.elements[i]) {
				return false
			}
		}

		if rightChild < len(heap.elements) {
			if heap.greaterThan(heap.elements[rightChild], heap.elements[i]) {
				return false
			}
		}
	}
	return true
}

// Benchmark tests
func BenchmarkInsert(b *testing.B) {
	heap := NewMaxHeap(b.N, func(a, b int) bool { return a > b })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

func BenchmarkRemove(b *testing.B) {
	heap := NewMaxHeap(b.N, func(a, b int) bool { return a > b })
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Remove()
	}
}
