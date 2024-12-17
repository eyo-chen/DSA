package main

// Build Two Arrays
// Time Complexity: O(n)
// Space Complexity: O(n)
func PivotIndex(nums []int) int {
	arr1 := make([]int, len(nums))
	arr2 := make([]int, len(nums))

	for i, acc := 0, 0; i < len(nums); i++ {
		acc += nums[i]
		arr1[i] = acc
	}

	for i, acc := len(nums)-1, 0; i >= 0; i-- {
		acc += nums[i]
		arr2[i] = acc
	}

	for i := 0; i < len(nums); i++ {
		if arr1[i] == arr2[i] {
			return i
		}
	}

	return -1
}

// Calculate Sum and Iterate
// Time Complexity: O(n)
// Space Complexity: O(1)
func PivotIndex2(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	leftSum := 0
	for i, n := range nums {
		rightSum := sum - n - leftSum

		if leftSum == rightSum {
			return i
		}

		leftSum += n
	}

	return -1
}
