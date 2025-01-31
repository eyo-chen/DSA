package main

import "sort"

// Sorting
// Time Complexity: O(nlog(n))
// Space Complexity: O(n)
func TopKFrequent(nums []int, k int) []int {
	// Create a hash table to store the frequency of each number
	// Key: number, Value: frequency
	freqMap := map[int]int{}
	for _, n := range nums {
		freqMap[n]++
	}

	// Create a slice of freqPairs to store the number and its frequency
	// [number, frequency]
	freqPairs := make([][2]int, 0, len(freqMap))
	for key, val := range freqMap {
		freqPairs = append(freqPairs, [2]int{key, val})
	}

	// Sort the freqPairs by frequency in descending order
	sort.Slice(freqPairs, func(i, j int) bool {
		return freqPairs[i][1] > freqPairs[j][1]
	})

	// Create a slice to store the result
	res := make([]int, 0, k)
	for i := 0; i < k; i++ { // Only need to iterate k times
		res = append(res, freqPairs[i][0])
	}

	return res
}

// Bucket Sort
// Time Complexity: O(n)
// Space Complexity: O(n)
func TopKFrequent1(nums []int, k int) []int {
	// Create a hash table to store the frequency of each number
	// Key: number, Value: frequency
	hashTable := map[int]int{}
	for _, n := range nums {
		hashTable[n]++
	}

	// Create a freqArr of size n+1 (each index represents a frequency)
	// The value of each index is a slice of numbers that have this frequency
	// Note that the value of each index is a slice, because there might be multiple numbers that have the same frequency
	freqArr := make([][]int, len(nums)+1)
	for num, freq := range hashTable {
		freqArr[freq] = append(freqArr[freq], num)
	}

	// Create a slice to store the result
	// The outer loop iterates from the end of the freqArr to the beginning(most frequent to least frequent)
	// The inner loop iterates through the numbers in the current frequency
	// Note that for these two loops, we immediately stop when we have k numbers
	ans := make([]int, 0, k)
	for i := len(freqArr) - 1; i >= 0 && len(ans) < k; i-- {
		for j := 0; j < len(freqArr[i]) && len(ans) < k; j++ {
			ans = append(ans, freqArr[i][j])
		}
	}

	return ans
}

type node struct {
	val      int
	priority int
}

// Min Heap
// Time Complexity: O(nlog(k))
// Space Complexity: O(n)
func TopKFrequent2(nums []int, k int) []int {
	// Create a hash table to store the frequency of each number
	// Key: number, Value: frequency
	freqMap := map[int]int{}
	for _, n := range nums {
		freqMap[n]++
	}

	// Create a min heap with a size of k
	// Loop through the hash table, and insert the value into the heap
	minHeap := NewMinHeap(k)
	for num, freq := range freqMap { // O(n)
		minHeap.Insert(node{val: num, priority: freq}) // O(log(k))

		// Check if the heap is overflowed
		// If that's the case, pop the root from the heap
		// This ensures that the heap is always a min heap with a size of k
		if minHeap.Size() > k {
			minHeap.Pop() // O(log(k))
		}
	}

	// Pop the root from the heap k times, and the numbers we pop are the most frequent numbers
	res := make([]int, 0, k)
	for i := 0; i < k; i++ { // O(k)
		res = append(res, minHeap.Pop().val) // O(log(k))
	}

	return res
}

// Go to Heap/findKthLargest/minheap.go for the detailed implementation of the min heap
type MinHeap struct {
	val []node
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		val: make([]node, 0, size),
	}
}

func (m *MinHeap) Insert(n node) {
	m.val = append(m.val, n)
	m.bubbleUp(len(m.val) - 1)
}

func (m *MinHeap) Pop() node {
	root := m.val[0]

	lastIdx := len(m.val) - 1
	m.swap(lastIdx, 0)
	m.val = m.val[:lastIdx]

	if len(m.val) > 0 {
		m.bubbleDown(0)
	}

	return root
}

func (m *MinHeap) bubbleDown(idx int) {
	for {
		left, right := 2*idx+1, 2*idx+2
		smallIdx := idx

		if left < len(m.val) && m.val[left].priority < m.val[smallIdx].priority {
			smallIdx = left
		}
		if right < len(m.val) && m.val[right].priority < m.val[smallIdx].priority {
			smallIdx = right
		}

		if smallIdx == idx {
			break
		}

		m.swap(idx, smallIdx)
		idx = smallIdx
	}
}

func (m *MinHeap) bubbleUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if m.val[idx].priority >= m.val[parentIdx].priority {
			break
		}

		m.swap(parentIdx, idx)
		idx = parentIdx
	}
}

func (m *MinHeap) swap(i, j int) {
	m.val[i], m.val[j] = m.val[j], m.val[i]
}

func (m *MinHeap) Size() int {
	return len(m.val)
}
