package main

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  bool
	}{
		// Basic cases from examples
		{
			name:      "overlapping meetings",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  false,
		},
		{
			name:      "non-overlapping meetings",
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  true,
		},

		// Edge cases
		{
			name:      "empty intervals",
			intervals: [][]int{},
			expected:  true,
		},
		{
			name:      "single meeting",
			intervals: [][]int{{1, 5}},
			expected:  true,
		},

		// Adjacent meetings (no overlap)
		{
			name:      "back-to-back meetings - no overlap",
			intervals: [][]int{{0, 5}, {5, 10}, {10, 15}},
			expected:  true,
		},
		{
			name:      "meetings ending when another starts",
			intervals: [][]int{{1, 5}, {5, 9}},
			expected:  true,
		},

		// Overlap cases
		{
			name:      "two meetings overlap",
			intervals: [][]int{{1, 5}, {3, 7}},
			expected:  false,
		},
		{
			name:      "one meeting completely inside another",
			intervals: [][]int{{1, 10}, {3, 5}},
			expected:  false,
		},
		{
			name:      "exact same meeting times",
			intervals: [][]int{{1, 5}, {1, 5}},
			expected:  false,
		},
		{
			name:      "multiple overlaps",
			intervals: [][]int{{1, 5}, {2, 6}, {3, 7}},
			expected:  false,
		},

		// Unsorted input
		{
			name:      "unsorted non-overlapping",
			intervals: [][]int{{10, 15}, {0, 5}, {5, 10}},
			expected:  true,
		},
		{
			name:      "unsorted with overlap",
			intervals: [][]int{{15, 20}, {5, 15}, {0, 10}},
			expected:  false,
		},

		// Large gaps
		{
			name:      "meetings with large gaps",
			intervals: [][]int{{0, 5}, {100, 200}, {300, 400}},
			expected:  true,
		},

		// Boundary cases
		{
			name:      "zero duration meeting (start equals end would be invalid per problem)",
			intervals: [][]int{{1, 2}, {3, 4}},
			expected:  true,
		},
		{
			name:      "meetings at time zero",
			intervals: [][]int{{0, 1}, {1, 2}},
			expected:  true,
		},

		// Many meetings
		{
			name:      "many non-overlapping meetings",
			intervals: [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			expected:  true,
		},
		{
			name:      "many meetings with one overlap",
			intervals: [][]int{{0, 1}, {2, 3}, {4, 5}, {4, 6}, {7, 8}},
			expected:  false,
		},

		// Reverse sorted
		{
			name:      "reverse sorted non-overlapping",
			intervals: [][]int{{20, 25}, {15, 19}, {10, 14}, {5, 9}, {0, 4}},
			expected:  true,
		},
		{
			name:      "reverse sorted with overlap",
			intervals: [][]int{{20, 25}, {15, 21}, {10, 14}},
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanAttendMeetings(tt.intervals)
			if result != tt.expected {
				t.Errorf("canAttendMeetings(%v) = %v; expected %v",
					tt.intervals, result, tt.expected)
			}
		})
	}
}
