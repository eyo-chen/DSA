package main

func IsMonotonic(nums []int) bool {
	// order: 0 means we haven't found the pattern, 1 means increasing, 2 means descending
	order := 0

	for i := 1; i < len(nums); i++ {
		// If order is 0, it means we haven't found the pattern
		if order == 0 {
			// If nums[i] > nums[i-1], it means the pattern is increasing
			if nums[i] > nums[i-1] {
				order = 1
			}
			// If nums[i] < nums[i-1], it means the pattern is descending
			if nums[i] < nums[i-1] {
				order = 2
			}
			continue
		}

		// If order is 1, it means the pattern is increasing
		if order == 1 && nums[i] < nums[i-1] {
			return false
		}

		// If order is 2, it means the pattern is descending
		if order == 2 && nums[i] > nums[i-1] {
			return false
		}
	}

	return true
}
