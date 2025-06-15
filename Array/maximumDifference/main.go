package main

func MaximumDifference(nums []int) int {
	res := -1
	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[k] > nums[i] {
				res = max(res, nums[k]-nums[i])
			}
		}
	}

	return res
}

func MaximumDifference1(nums []int) int {
	res := -1
	minVal := nums[0]

	for i := 1; i < len(nums); i++ {
		// Only get the difference when nums[i] > minVal
		if nums[i] > minVal {
			res = max(nums[i]-minVal, res)
		}

		minVal = min(nums[i], minVal)
	}

	return res
}
