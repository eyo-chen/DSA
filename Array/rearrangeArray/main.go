package main

func RearrangeArray(nums []int) []int {
	// Initialize slices to store positive and negative numbers
	positive := []int{}
	negative := []int{}
	// Initialize result slice with capacity equal to input length
	res := make([]int, 0, len(nums))

	// Separate numbers based on sign
	for _, num := range nums {
		if num < 0 {
			// Append negative numbers to negative slice
			negative = append(negative, num)
		} else {
			// Append positive numbers to positive slice
			positive = append(positive, num)
		}
	}

	// Interleave positive and negative numbers in result
	for i := 0; i < len(positive); i++ {
		// Append one positive and one negative number per iteration
		res = append(res, positive[i], negative[i])
	}

	return res
}

// Using two pointers
func RearrangeArray1(nums []int) []int {
	// Initialize pointers for next positive and negative numbers
	positivePtr, negativePtr := 0, 0
	// Initialize result slice with capacity equal to input length
	res := make([]int, 0, len(nums))

	// Skip negative numbers to find first positive
	for positivePtr < len(nums) && nums[positivePtr] < 0 {
		positivePtr++
	}

	// Skip positive numbers to find first negative
	for negativePtr < len(nums) && nums[negativePtr] > 0 {
		negativePtr++
	}

	// Alternate between positive and negative numbers
	for positivePtr < len(nums) && negativePtr < len(nums) {
		// Append current positive and negative numbers
		res = append(res, nums[positivePtr], nums[negativePtr])

		// Find next positive number
		positivePtr++
		for positivePtr < len(nums) && nums[positivePtr] < 0 {
			positivePtr++
		}

		// Find next negative number
		negativePtr++
		for negativePtr < len(nums) && nums[negativePtr] > 0 {
			negativePtr++
		}
	}

	return res
}

// Using two indexes
func RearrangeArray2(nums []int) []int {
	// Initialize result array with full length
	res := make([]int, len(nums))
	// Initialize indices: even for positives, odd for negatives
	positiveIdx, negativeIdx := 0, 1

	// Iterate through input array once
	for _, num := range nums {
		if num < 0 {
			// Place negative number at next odd index
			res[negativeIdx] = num
			negativeIdx += 2
		} else {
			// Place positive number at next even index
			res[positiveIdx] = num
			positiveIdx += 2
		}
	}

	return res
}
