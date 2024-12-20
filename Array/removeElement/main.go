package main

// Using extra space
// Time Complexity: O(n)
// Space Complexity: O(n)
func RemoveElement(nums []int, val int) int {
	arr := make([]int, 0, len(nums))

	for _, n := range nums {
		if n == val {
			continue
		}
		arr = append(arr, n)
	}

	copy(nums, arr)

	return len(arr)
}

// In-place
// Time Complexity: O(n)
// Space Complexity: O(1)
func RemoveElement2(nums []int, val int) int {
	writeIdx := 0
	for _, n := range nums {
		if n == val {
			continue
		}

		nums[writeIdx] = n
		writeIdx++
	}

	return writeIdx
}
