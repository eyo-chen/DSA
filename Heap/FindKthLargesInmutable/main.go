package main

import (
	"fmt"
	"sort"
	"time"

	"math/rand"
)

// Use Simple Array
func FindKthLargestInmutable(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	candidatesIndex := []int{0}
	res := make([]int, k)

	for i := 0; i < k; i++ {
		curMax, maxIdx, maxCandidateIdx := nums[candidatesIndex[0]], candidatesIndex[0], 0
		for candidateIdx, idx := range candidatesIndex {

			// Find the largest element in the candidates
			if nums[idx] > curMax {
				curMax = nums[idx]
				maxIdx = idx
				maxCandidateIdx = candidateIdx
			}
		}

		// Add the largest element to the result
		res[i] = curMax

		// Remove the largest element from the candidates
		candidatesIndex = removeElement(candidatesIndex, maxCandidateIdx)

		// Add the left and right child to the candidates
		left := maxIdx*2 + 1
		if left < len(nums) {
			candidatesIndex = append(candidatesIndex, left)
		}
		right := maxIdx*2 + 2
		if right < len(nums) {
			candidatesIndex = append(candidatesIndex, right)
		}
	}

	return res
}

func removeElement(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Use Max Heap
func FindKthLargestInmutable1(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	candidate := NewMaxHeap(len(nums))
	candidate.Insert(node{nums[0], 0})
	res := make([]int, k)

	for i := 0; i < k; i++ {
		// Get the largest element from the heap
		nextMax := candidate.Pull()

		// Add the largest element to the result
		res[i] = nextMax.val

		// left
		leftIdx := nextMax.idx*2 + 1
		if leftIdx < len(nums) {
			candidate.Insert(node{nums[leftIdx], leftIdx})
		}
		rightIdx := nextMax.idx*2 + 2
		if rightIdx < len(nums) {
			candidate.Insert(node{nums[rightIdx], rightIdx})
		}
	}

	return res
}

// This is used for testing
// When using sort, we guarantee the result is correct
func FindKthLargest(nums []int, k int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[:k]
}

// This is used for testing
func TestFindKthLargest() {
	for i := 0; i < 5; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		idx := 100000

		for i := 0; i < idx; i++ {
			// Generate random slice of integers
			length := r.Intn(100) + 1 // Random length between 1 and 100
			rawNums := make([]int, length)
			for j := range rawNums {
				rawNums[j] = r.Intn(1000) // Random integers between 0 and 999
			}

			// Generate random k
			k := r.Intn(length) + 1 // Random k between 1 and length

			// The input is guaranteed to be a valid max heap
			nums := NewMaxHeapInt(length)
			for _, n := range rawNums {
				nums.Insert(n)
			}

			// Run all three functions
			result1 := FindKthLargestInmutable(nums.vals, k)
			result2 := FindKthLargestInmutable1(nums.vals, k)
			result3 := FindKthLargest(nums.vals, k)

			// Compare results
			if !compareSlices(result1, result3) {
				fmt.Println("error for result1")
				fmt.Println("Mismatch found in test case", i+1)
				fmt.Println("Input: nums=", nums.vals, "k=", k)
				fmt.Println("FindKthLargestInmutable:", result1)
				fmt.Println("FindKthLargest:", result3)
				return
			}

			if !compareSlices(result2, result3) {
				fmt.Println("error for result2")
				fmt.Println("Mismatch found in test case", i+1)
				fmt.Println("Input: nums=", nums.vals, "k=", k)
				fmt.Println("FindKthLargestInmutable1:", result2)
				fmt.Println("FindKthLargest:", result3)
				return
			}
		}

		fmt.Println("All 10000 test cases passed successfully!")
	}
}

// This is used for comparing two slices without considering the order
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences of each number
	countA := make(map[int]int)
	countB := make(map[int]int)

	// Count occurrences in slice a
	for _, num := range a {
		countA[num]++
	}

	// Count occurrences in slice b
	for _, num := range b {
		countB[num]++
	}

	// Compare the counts
	for num, count := range countA {
		if countB[num] != count {
			return false
		}
	}

	return true
}

func main() {
	TestFindKthLargest()
}
