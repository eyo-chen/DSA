package main

import (
	"reflect"
	"testing"
)

func TestFindMissingRanges1(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		lower    int
		upper    int
		expected [][]int
	}{
		{
			name:     "Example 1 - Multiple missing ranges",
			nums:     []int{0, 1, 3, 50, 75},
			lower:    0,
			upper:    99,
			expected: [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}},
		},
		{
			name:     "Example 2 - No missing ranges (single element)",
			nums:     []int{-1},
			lower:    -1,
			upper:    -1,
			expected: [][]int{},
		},
		{
			name:     "Empty array - entire range missing",
			nums:     []int{},
			lower:    1,
			upper:    5,
			expected: [][]int{{1, 5}},
		},
		{
			name:     "All elements present - no missing ranges",
			nums:     []int{1, 2, 3, 4, 5},
			lower:    1,
			upper:    5,
			expected: [][]int{},
		},
		{
			name:     "Missing at the beginning",
			nums:     []int{3, 4, 5},
			lower:    1,
			upper:    5,
			expected: [][]int{{1, 2}},
		},
		{
			name:     "Missing at the end",
			nums:     []int{1, 2, 3},
			lower:    1,
			upper:    5,
			expected: [][]int{{4, 5}},
		},
		{
			name:     "Missing in the middle",
			nums:     []int{1, 2, 5},
			lower:    1,
			upper:    5,
			expected: [][]int{{3, 4}},
		},
		{
			name:     "Single element in range - missing before and after",
			nums:     []int{3},
			lower:    1,
			upper:    5,
			expected: [][]int{{1, 2}, {4, 5}},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-5, -3, -1},
			lower:    -7,
			upper:    0,
			expected: [][]int{{-7, -6}, {-4, -4}, {-2, -2}, {0, 0}},
		},
		{
			name:     "Mix of negative and positive",
			nums:     []int{-2, 0, 2},
			lower:    -3,
			upper:    3,
			expected: [][]int{{-3, -3}, {-1, -1}, {1, 1}, {3, 3}},
		},
		{
			name:     "Single missing number",
			nums:     []int{1, 3},
			lower:    1,
			upper:    3,
			expected: [][]int{{2, 2}},
		},
		{
			name:     "Two consecutive numbers - missing ends",
			nums:     []int{2, 3},
			lower:    1,
			upper:    4,
			expected: [][]int{{1, 1}, {4, 4}},
		},
		{
			name:     "Large gap in middle",
			nums:     []int{1, 100},
			lower:    1,
			upper:    100,
			expected: [][]int{{2, 99}},
		},
		{
			name:     "Zero in range",
			nums:     []int{0},
			lower:    -1,
			upper:    1,
			expected: [][]int{{-1, -1}, {1, 1}},
		},
		{
			name:     "All negative numbers",
			nums:     []int{-100, -50, -25},
			lower:    -100,
			upper:    -1,
			expected: [][]int{{-99, -51}, {-49, -26}, {-24, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMissingRanges1(tt.nums, tt.lower, tt.upper)

			// Handle nil vs empty slice comparison
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findMissingRanges1(%v, %d, %d) = %v, want %v",
					tt.nums, tt.lower, tt.upper, result, tt.expected)
			}
		})
	}
}

// Test with edge case: maximum constraint values (smaller version for practical testing)
func TestFindMissingRangesLargeRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		lower    int
		upper    int
		expected [][]int
	}{
		{
			name:     "Large range with few elements",
			nums:     []int{0, 1000000},
			lower:    0,
			upper:    1000000,
			expected: [][]int{{1, 999999}},
		},
		{
			name:     "Empty array with large range",
			nums:     []int{},
			lower:    -1000,
			upper:    1000,
			expected: [][]int{{-1000, 1000}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMissingRanges1(tt.nums, tt.lower, tt.upper)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findMissingRanges1(%v, %d, %d) = %v, want %v",
					tt.nums, tt.lower, tt.upper, result, tt.expected)
			}
		})
	}
}
