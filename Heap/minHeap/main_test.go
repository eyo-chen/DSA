package main

import (
	"reflect"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	t.Run("Create heap with positive size", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		if heap == nil {
			t.Fatal("Expected non-nil heap")
		}
	})

	t.Run("Create heap with zero size", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(0, compare)

		if heap == nil {
			t.Fatal("Expected non-nil heap")
		}
	})
}

func TestMinHeap_Insert(t *testing.T) {
	t.Run("Insert single element", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		heap.Insert(5)
		val, err := heap.Remove()

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if val != 5 {
			t.Errorf("Expected 5, got %d", val)
		}
	})

	t.Run("Insert multiple elements in ascending order", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		for i := 1; i <= 5; i++ {
			heap.Insert(i)
		}

		for i := 1; i <= 5; i++ {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if val != i {
				t.Errorf("Expected %d, got %d", i, val)
			}
		}
	})

	t.Run("Insert multiple elements in descending order", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		for i := 5; i >= 1; i-- {
			heap.Insert(i)
		}

		for i := 1; i <= 5; i++ {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if val != i {
				t.Errorf("Expected %d, got %d", i, val)
			}
		}
	})

	t.Run("Insert duplicate elements", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		heap.Insert(3)
		heap.Insert(3)
		heap.Insert(3)

		for i := 0; i < 3; i++ {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if val != 3 {
				t.Errorf("Expected 3, got %d", val)
			}
		}
	})

	t.Run("Insert random order elements", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		elements := []int{7, 2, 9, 1, 5, 3, 8, 4, 6}
		for _, v := range elements {
			heap.Insert(v)
		}

		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i, exp := range expected {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error at index %d, got %v", i, err)
			}
			if val != exp {
				t.Errorf("At index %d: expected %d, got %d", i, exp, val)
			}
		}
	})

	t.Run("Insert beyond initial capacity", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(2, compare)

		for i := 1; i <= 10; i++ {
			heap.Insert(i)
		}

		for i := 1; i <= 10; i++ {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if val != i {
				t.Errorf("Expected %d, got %d", i, val)
			}
		}
	})
}

func TestMinHeap_Remove(t *testing.T) {
	t.Run("Remove from empty heap", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		_, err := heap.Remove()
		if err == nil {
			t.Error("Expected error when removing from empty heap")
		}
	})

	t.Run("Remove until empty", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		heap.Insert(1)
		heap.Insert(2)
		heap.Insert(3)

		heap.Remove()
		heap.Remove()
		heap.Remove()

		_, err := heap.Remove()
		if err == nil {
			t.Error("Expected error when removing from empty heap")
		}
	})

	t.Run("Alternating insert and remove", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		heap.Insert(5)
		heap.Insert(3)

		val, err := heap.Remove()
		if err != nil || val != 3 {
			t.Errorf("Expected 3, got %d (err: %v)", val, err)
		}

		heap.Insert(1)
		heap.Insert(7)

		val, err = heap.Remove()
		if err != nil || val != 1 {
			t.Errorf("Expected 1, got %d (err: %v)", val, err)
		}

		val, err = heap.Remove()
		if err != nil || val != 5 {
			t.Errorf("Expected 5, got %d (err: %v)", val, err)
		}
	})
}

func TestMinHeap_WithStrings(t *testing.T) {
	t.Run("String comparison", func(t *testing.T) {
		compare := func(a, b string) bool { return a < b }
		heap := NewMinHeap[string](10, compare)

		words := []string{"zebra", "apple", "mango", "banana"}
		for _, w := range words {
			heap.Insert(w)
		}

		expected := []string{"apple", "banana", "mango", "zebra"}
		for i, exp := range expected {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error at index %d, got %v", i, err)
			}
			if val != exp {
				t.Errorf("At index %d: expected %s, got %s", i, exp, val)
			}
		}
	})
}

func TestMinHeap_WithCustomStruct(t *testing.T) {
	type Task struct {
		Name     string
		Priority int
	}

	t.Run("Custom struct with priority", func(t *testing.T) {
		compare := func(a, b Task) bool { return a.Priority < b.Priority }
		heap := NewMinHeap[Task](10, compare)

		tasks := []Task{
			{"Low", 5},
			{"High", 1},
			{"Medium", 3},
			{"Critical", 0},
		}

		for _, task := range tasks {
			heap.Insert(task)
		}

		expected := []int{0, 1, 3, 5}
		for i, expPriority := range expected {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error at index %d, got %v", i, err)
			}
			if val.Priority != expPriority {
				t.Errorf("At index %d: expected priority %d, got %d", i, expPriority, val.Priority)
			}
		}
	})
}

func TestMinHeap_MaxHeapBehavior(t *testing.T) {
	t.Run("Use max heap comparison", func(t *testing.T) {
		compare := func(a, b int) bool { return a > b }
		heap := NewMinHeap(10, compare)

		elements := []int{1, 5, 3, 9, 2}
		for _, v := range elements {
			heap.Insert(v)
		}

		expected := []int{9, 5, 3, 2, 1}
		for i, exp := range expected {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error at index %d, got %v", i, err)
			}
			if val != exp {
				t.Errorf("At index %d: expected %d, got %d", i, exp, val)
			}
		}
	})
}

func TestMinHeap_NegativeNumbers(t *testing.T) {
	t.Run("Handle negative numbers", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		elements := []int{-5, 3, -10, 0, 7, -2}
		for _, v := range elements {
			heap.Insert(v)
		}

		expected := []int{-10, -5, -2, 0, 3, 7}
		for i, exp := range expected {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error at index %d, got %v", i, err)
			}
			if val != exp {
				t.Errorf("At index %d: expected %d, got %d", i, exp, val)
			}
		}
	})
}

func TestMinHeap_LargeDataset(t *testing.T) {
	t.Run("Handle large number of elements", func(t *testing.T) {
		compare := func(a, b int) bool { return a < b }
		heap := NewMinHeap(10, compare)

		n := 1000
		for i := n; i > 0; i-- {
			heap.Insert(i)
		}

		for i := 1; i <= n; i++ {
			val, err := heap.Remove()
			if err != nil {
				t.Fatalf("Expected no error at iteration %d, got %v", i, err)
			}
			if val != i {
				t.Errorf("At iteration %d: expected %d, got %d", i, i, val)
			}
		}
	})
}

// Helper for error checking
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

// 1. Test Standard Integer Sorting
func TestMinHeap_Ints(t *testing.T) {
	// Comparator for Ints
	comp := func(a, b int) bool { return a < b }

	// Create heap with small initial capacity to force resizing later
	h := NewMinHeap(2, comp)

	inputs := []int{10, 5, 30, 2, 5, 8}
	expectedOrder := []int{2, 5, 5, 8, 10, 30}

	// Insert all
	for _, v := range inputs {
		h.Insert(v)
	}

	// Remove all and verify order
	var result []int
	for i := 0; i < len(inputs); i++ {
		val, err := h.Remove()
		assertNoError(t, err)
		result = append(result, val)
	}

	if !reflect.DeepEqual(result, expectedOrder) {
		t.Errorf("Heap order incorrect.\nGot: %v\nWant: %v", result, expectedOrder)
	}
}

// 2. Test Strings (Generics Check)
func TestMinHeap_Strings(t *testing.T) {
	comp := func(a, b string) bool { return a < b }
	h := NewMinHeap(5, comp)

	inputs := []string{"banana", "apple", "cherry", "date"}
	expected := "apple"

	for _, v := range inputs {
		h.Insert(v)
	}

	val, err := h.Remove()
	assertNoError(t, err)

	if val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}

// 3. Test Complex Structs (Priority Queue Use Case)
type Task struct {
	Name     string
	Priority int
}

func TestMinHeap_Structs(t *testing.T) {
	// Compare based on Priority field
	comp := func(a, b Task) bool { return a.Priority < b.Priority }
	h := NewMinHeap(10, comp)

	h.Insert(Task{Name: "LowPrio", Priority: 100})
	h.Insert(Task{Name: "HighPrio", Priority: 1})
	h.Insert(Task{Name: "MedPrio", Priority: 50})

	// First removal should be HighPrio (1)
	val, err := h.Remove()
	assertNoError(t, err)

	if val.Name != "HighPrio" {
		t.Errorf("Expected task HighPrio, got %s", val.Name)
	}

	// Second removal should be MedPrio (50)
	val, err = h.Remove()
	assertNoError(t, err)
	if val.Name != "MedPrio" {
		t.Errorf("Expected task MedPrio, got %s", val.Name)
	}
}

// 4. Test Edge Case: Empty Heap
func TestMinHeap_Empty(t *testing.T) {
	comp := func(a, b int) bool { return a < b }
	h := NewMinHeap(10, comp)

	// Try removing from fresh heap
	_, err := h.Remove()
	assertError(t, err)

	// Insert 1, Remove 1, Remove again (Empty)
	h.Insert(1)
	_, _ = h.Remove()
	_, err = h.Remove()
	assertError(t, err)
}

// 5. Test Edge Case: Duplicates
func TestMinHeap_Duplicates(t *testing.T) {
	comp := func(a, b int) bool { return a < b }
	h := NewMinHeap(10, comp)

	h.Insert(5)
	h.Insert(1)
	h.Insert(5)
	h.Insert(1)

	expected := []int{1, 1, 5, 5}
	var result []int

	for i := 0; i < 4; i++ {
		val, _ := h.Remove()
		result = append(result, val)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Duplicates not handled correctly.\nGot: %v\nWant: %v", result, expected)
	}
}

// 6. Test Interleaved Operations (Insert mixed with Remove)
func TestMinHeap_Interleaved(t *testing.T) {
	comp := func(a, b int) bool { return a < b }
	h := NewMinHeap(10, comp)

	h.Insert(10)
	h.Insert(5)
	// Heap: [5, 10]

	val, _ := h.Remove()
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
	// Heap: [10]

	h.Insert(2)
	h.Insert(20)
	// Heap: [2, 10, 20] (logical structure)

	val, _ = h.Remove()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	val, _ = h.Remove()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	val, _ = h.Remove()
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}
}
