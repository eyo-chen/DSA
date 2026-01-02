package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestThreeSumSmaller(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example case 1",
			nums:   []int{-2, 0, 1, 3},
			target: 2,
			want:   2,
		},
		{
			name:   "All triplets valid",
			nums:   []int{1, 1, 1},
			target: 100,
			want:   1,
		},
		{
			name:   "No valid triplets",
			nums:   []int{5, 10, 15},
			target: 10,
			want:   0,
		},
		{
			name:   "Single valid triplet",
			nums:   []int{1, 2, 3, 4},
			target: 7,
			want:   1, // [1,2,3]
		},
		{
			name:   "Multiple valid triplets",
			nums:   []int{-1, 0, 1, 2},
			target: 3,
			want:   3, // [-1,0,1], [-1,0,2], [-1,1,2]
		},
		{
			name:   "Empty array",
			nums:   []int{},
			target: 0,
			want:   0,
		},
		{
			name:   "Array with less than 3 elements",
			nums:   []int{1, 2},
			target: 10,
			want:   0,
		},
		{
			name:   "Exactly 3 elements - valid",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   1,
		},
		{
			name:   "Exactly 3 elements - invalid",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   0,
		},
		{
			name:   "All negative numbers",
			nums:   []int{-5, -4, -3, -2, -1},
			target: -6,
			want:   4, // [-5,-4,-3], [-5,-4,-2], [-5,-3,-2], [-4,-3,-2]
		},
		{
			name:   "All positive numbers",
			nums:   []int{1, 2, 3, 4, 5},
			target: 7,
			want:   1, // [1,2,3]
		},
		{
			name:   "Mixed positive and negative",
			nums:   []int{-2, -1, 0, 1, 2},
			target: 0,
			want:   3, // [-2,-1,0], [-2,-1,1], [-2,0,1]
		},
		{
			name:   "Duplicate values",
			nums:   []int{0, 0, 0, 0},
			target: 1,
			want:   4, // C(4,3) = 4 combinations
		},
		{
			name:   "Large negative target",
			nums:   []int{-5, -4, -3, -2, -1},
			target: -100,
			want:   0,
		},
		{
			name:   "Zero target with zeros",
			nums:   []int{0, 0, 0},
			target: 0,
			want:   0, // sum equals target, not less than
		},
		{
			name:   "Target at boundary",
			nums:   []int{1, 1, 1},
			target: 3,
			want:   0, // sum equals 3, not less than 3
		},
		{
			name:   "Unsorted input",
			nums:   []int{3, 1, 0, -2},
			target: 2,
			want:   2, // Same as example 1, just unsorted
		},
		{
			name:   "Large array with many valid triplets",
			nums:   []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			target: -15,
			want:   120, // Many combinations
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to ensure function doesn't modify input
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := ThreeSumSmaller(numsCopy, tt.target)
			want := ThreeSumSmaller1(tt.nums, tt.target) // Using brute-force for verification

			if got != want {
				t.Errorf("ThreeSumSmaller(%v, %d) = %d, want %d",
					tt.nums, tt.target, got, want)
			}
		})
	}
}

func TestRandomizedThreeSumSmaller(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	numTests := 1000
	maxSize := 500
	maxValue := 500

	for i := 0; i < numTests; i++ {
		size := rand.Intn(maxSize) + 3 // Ensure at least 3 elements
		nums := make([]int, size)
		for j := 0; j < size; j++ {
			nums[j] = rand.Intn(2*maxValue) - maxValue // Values between -maxValue and maxValue
		}
		target := rand.Intn(2*maxValue) - maxValue

		// Make a copy to ensure function doesn't modify input
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)

		got := ThreeSumSmaller(numsCopy, target)
		want := ThreeSumSmaller1(nums, target) // Using brute-force for verification

		if got != want {
			t.Errorf("Random test %d failed: ThreeSumSmaller(%v, %d) = %d, want %d",
				i+1, nums, target, got, want)
		}
	}

}

// Test for efficiency - O(n^2) requirement
func TestThreeSumSmallerEfficiency(t *testing.T) {
	// Generate a large test case
	size := 1000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i - 500 // Range from -500 to 499
	}
	target := 1000

	start := time.Now()
	result := ThreeSumSmaller(nums, target)
	elapsed := time.Since(start)

	t.Logf("Array size: %d, Result: %d, Time: %v", size, result, elapsed)

	// For O(n^2) algorithm, 1000 elements should complete very quickly
	// If it takes more than 1 second, it's likely not O(n^2)
	if elapsed > time.Second {
		t.Errorf("Algorithm too slow for n=%d: took %v, expected < 1s (likely not O(n^2))",
			size, elapsed)
	}
}

// Benchmark test
func BenchmarkThreeSumSmaller(b *testing.B) {
	nums := []int{-2, 0, 1, 3, -1, 5, -3, 2, 4}
	target := 5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumSmaller(nums, target)
	}
}

// Benchmark with larger input
func BenchmarkThreeSumSmallerLarge(b *testing.B) {
	nums := make([]int, 500)
	for i := 0; i < 500; i++ {
		nums[i] = i - 250
	}
	target := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumSmaller(nums, target)
	}
}
