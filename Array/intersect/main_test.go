package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "Example 1 - Multiple duplicates (sorted)",
			nums1: []int{1, 1, 2, 2},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			name:  "Example 2 - Mixed elements (sorted)",
			nums1: []int{4, 5, 9},
			nums2: []int{4, 4, 8, 9, 9},
			want:  []int{4, 9},
		},
		{
			name:  "No intersection - non-overlapping ranges",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  []int{},
		},
		{
			name:  "No intersection - gaps between sorted arrays",
			nums1: []int{1, 3, 5, 7},
			nums2: []int{2, 4, 6, 8},
			want:  []int{},
		},
		{
			name:  "Complete intersection - identical sorted arrays",
			nums1: []int{1, 2, 3},
			nums2: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "One element arrays - match",
			nums1: []int{1},
			nums2: []int{1},
			want:  []int{1},
		},
		{
			name:  "One element arrays - no match",
			nums1: []int{1},
			nums2: []int{2},
			want:  []int{},
		},
		{
			name:  "First array empty",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  []int{},
		},
		{
			name:  "Second array empty",
			nums1: []int{1, 2, 3},
			nums2: []int{},
			want:  []int{},
		},
		{
			name:  "Both arrays empty",
			nums1: []int{},
			nums2: []int{},
			want:  []int{},
		},
		{
			name:  "Different frequencies - sorted",
			nums1: []int{1, 1, 1, 2, 2},
			nums2: []int{1, 1, 2, 2, 2, 3},
			want:  []int{1, 1, 2, 2},
		},
		{
			name:  "All duplicates in first array - sorted",
			nums1: []int{5, 5, 5, 5},
			nums2: []int{5, 5},
			want:  []int{5, 5},
		},
		{
			name:  "All duplicates in second array - sorted",
			nums1: []int{7, 7},
			nums2: []int{7, 7, 7, 7, 7},
			want:  []int{7, 7},
		},
		{
			name:  "Large numbers at boundary",
			nums1: []int{0, 1000},
			nums2: []int{0, 500, 1000},
			want:  []int{0, 1000},
		},
		{
			name:  "All zeros",
			nums1: []int{0, 0, 0},
			nums2: []int{0, 0},
			want:  []int{0, 0},
		},
		{
			name:  "Subset relationship - sorted",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{2, 3},
			want:  []int{2, 3},
		},
		{
			name:  "Reverse subset - sorted",
			nums1: []int{8, 9},
			nums2: []int{7, 8, 9, 10},
			want:  []int{8, 9},
		},
		{
			name:  "Partial overlap with duplicates - sorted",
			nums1: []int{1, 2, 2, 3, 4, 4, 4},
			nums2: []int{2, 2, 2, 4, 4, 5},
			want:  []int{2, 2, 4, 4},
		},
		{
			name:  "Single common element with many duplicates - sorted",
			nums1: []int{9, 9, 9, 9, 9},
			nums2: []int{9, 9, 9},
			want:  []int{9, 9, 9},
		},
		{
			name:  "Alternating pattern - sorted",
			nums1: []int{1, 3, 5, 7, 9},
			nums2: []int{2, 3, 4, 5, 6},
			want:  []int{3, 5},
		},
		{
			name:  "First array much smaller - sorted",
			nums1: []int{5, 10},
			nums2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:  []int{5, 10},
		},
		{
			name:  "Second array much smaller - sorted",
			nums1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			nums2: []int{3, 7},
			want:  []int{3, 7},
		},
		{
			name:  "Consecutive duplicates at start - sorted",
			nums1: []int{1, 1, 1, 5, 6, 7},
			nums2: []int{1, 1, 8, 9},
			want:  []int{1, 1},
		},
		{
			name:  "Consecutive duplicates at end - sorted",
			nums1: []int{1, 2, 3, 9, 9, 9},
			nums2: []int{4, 5, 9, 9},
			want:  []int{9, 9},
		},
		{
			name:  "Multiple groups of duplicates - sorted",
			nums1: []int{1, 1, 3, 3, 5, 5},
			nums2: []int{1, 3, 3, 5, 5, 5},
			want:  []int{1, 3, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make copies to ensure function doesn't modify inputs
			nums1Copy := make([]int, len(tt.nums1))
			nums2Copy := make([]int, len(tt.nums2))
			copy(nums1Copy, tt.nums1)
			copy(nums2Copy, tt.nums2)

			got := Intersect1(nums1Copy, nums2Copy)

			// Since arrays are already sorted and output should be sorted,
			// we can compare directly
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}

			// Verify original arrays weren't modified
			if !reflect.DeepEqual(tt.nums1, nums1Copy) {
				t.Errorf("nums1 was modified: original %v, after %v", tt.nums1, nums1Copy)
			}
			if !reflect.DeepEqual(tt.nums2, nums2Copy) {
				t.Errorf("nums2 was modified: original %v, after %v", tt.nums2, nums2Copy)
			}
		})
	}
}
