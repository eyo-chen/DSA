package main

// The most intuitive way
func FindKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		maxIdx := findMax(nums, i)
		nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
	}

	return nums[k-1]
}

func findMax(nums []int, start int) int {
	maxIndex := start
	for i := start + 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

// Using a max heap
func FindKthLargest1(nums []int, k int) int {
	var res int
	maxHeap := NewMaxHeap(len(nums))
	for _, n := range nums {
		maxHeap.Insert(n)
	}

	for i := 0; i < k; i++ {
		res = maxHeap.Pull()
	}

	return res
}

// Using a min heap (the opposite of a max heap)
func FindKthLargest2(nums []int, k int) int {
	minHeap := NewMinHeap(k)
	for _, n := range nums {
		minHeap.Insert(n)

		if minHeap.Len() == k+1 {
			minHeap.Pull()
		}
	}

	return minHeap.Pull()
}
