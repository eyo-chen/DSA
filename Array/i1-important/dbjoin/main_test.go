package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

// Helper function to generate random slice with potential duplicates
func generateRandomSlice(size int, maxValue int, duplicateProbability float64) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		if i > 0 && rand.Float64() < duplicateProbability {
			// Create duplicate of previous value
			result[i] = result[i-1]
		} else {
			result[i] = rand.Intn(maxValue)
		}
	}
	return result
}

// Test helper to compare results (order-independent) for all three algorithms
func compareResults(t *testing.T, nested, sortMerge, hash []int, testName string) {
	// Sort all results for comparison since join order might differ
	nestedSorted := make([]int, len(nested))
	sortMergeSorted := make([]int, len(sortMerge))
	hashSorted := make([]int, len(hash))

	copy(nestedSorted, nested)
	copy(sortMergeSorted, sortMerge)
	copy(hashSorted, hash)

	slices.Sort(nestedSorted)
	slices.Sort(sortMergeSorted)
	slices.Sort(hashSorted)

	// Compare all three results
	allMatch := reflect.DeepEqual(nestedSorted, sortMergeSorted) &&
		reflect.DeepEqual(sortMergeSorted, hashSorted)

	if !allMatch {
		t.Errorf("%s FAILED:\nNested result:    %v\nSortMerge result: %v\nHash result:      %v\n",
			testName, nestedSorted, sortMergeSorted, hashSorted)
	} else {
		fmt.Printf("%s PASSED (all three algorithms matched %d results)\n", testName, len(nested))
	}
}

func TestJoinAlgorithms(t *testing.T) {
	// Test 1: Empty tables
	t.Run("Empty tables", func(t *testing.T) {
		outer := []int{}
		inner := []int{}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "Empty tables")
	})

	// Test 2: One empty table
	t.Run("One empty table", func(t *testing.T) {
		outer := []int{1, 2, 3}
		inner := []int{}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "One empty table")
	})

	// Test 3: No matches
	t.Run("No matches", func(t *testing.T) {
		outer := []int{1, 2, 3}
		inner := []int{4, 5, 6}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "No matches")
	})

	// Test 4: All match (same value)
	t.Run("All same value", func(t *testing.T) {
		outer := []int{5, 5, 5}
		inner := []int{5, 5, 5, 5}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "All same value")
	})

	// Test 5: Duplicates in outer table only
	t.Run("Duplicates in outer", func(t *testing.T) {
		outer := []int{100, 200, 200, 300}
		inner := []int{100, 200, 400}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "Duplicates in outer")
	})

	// Test 6: Duplicates in inner table only
	t.Run("Duplicates in inner", func(t *testing.T) {
		outer := []int{100, 200, 300}
		inner := []int{100, 100, 200, 200, 400}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "Duplicates in inner")
	})

	// Test 7: Duplicates in both tables (lecture example)
	t.Run("Duplicates in both", func(t *testing.T) {
		outer := []int{100, 200, 200, 300, 400}
		inner := []int{100, 100, 200, 200, 400}
		nested := NestedLoopJoin(outer, inner)
		sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
		hash := HashJoin(outer, inner)
		compareResults(t, nested, sortMerge, hash, "Duplicates in both")
	})

	// Test 8-17: Random tests with varying parameters
	testCases := []struct {
		name                 string
		outerSize            int
		innerSize            int
		maxValue             int
		duplicateProbability float64
	}{
		{"Small random", 10, 10, 20, 0.1},
		{"Medium random", 50, 50, 50, 0.2},
		{"Large random", 200, 200, 100, 0.15},
		{"High duplicates", 100, 100, 30, 0.5},
		{"Low duplicates", 100, 100, 200, 0.05},
		{"Unbalanced sizes 1", 10, 100, 50, 0.2},
		{"Unbalanced sizes 2", 100, 10, 50, 0.2},
		{"Very high duplicates", 50, 50, 10, 0.7},
		{"Large with few values", 300, 300, 20, 0.3},
		{"Stress test", 500, 500, 100, 0.25},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outer := generateRandomSlice(tc.outerSize, tc.maxValue, tc.duplicateProbability)
			inner := generateRandomSlice(tc.innerSize, tc.maxValue, tc.duplicateProbability)

			nested := NestedLoopJoin(outer, inner)
			sortMerge := SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
			hash := HashJoin(outer, inner)

			compareResults(t, nested, sortMerge, hash, tc.name)
		})
	}
}

// Benchmark tests
func BenchmarkNestedJoin(b *testing.B) {
	outer := generateRandomSlice(100, 50, 0.2)
	inner := generateRandomSlice(100, 50, 0.2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NestedLoopJoin(outer, inner)
	}
}

func BenchmarkSortMergeJoin(b *testing.B) {
	outer := generateRandomSlice(100, 50, 0.2)
	inner := generateRandomSlice(100, 50, 0.2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortMergeJoin(slices.Clone(outer), slices.Clone(inner))
	}
}

func BenchmarkHashJoin(b *testing.B) {
	outer := generateRandomSlice(100, 50, 0.2)
	inner := generateRandomSlice(100, 50, 0.2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashJoin(outer, inner)
	}
}
