package main

func FindPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	// Handle edge cases
	// (1) only two elements
	// e.g. [5,4], just compare which one is greater
	// (2) descending order
	// e.g. [5,4,3,2,1], return 0 because 5 is the peak element, also because nums[-1] = -∞(problem statement)
	if nums[0] > nums[1] {
		return 0
	}

	// Check middle elements
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
	}

	// If no peak element is found, return the last element's index
	// e.g. [1,2,3,4,5]
	// because the order is ascending, the for loop will not find any peak element
	// so we just return the last element's index
	// because nums[n] = -∞(problem statement)
	return len(nums) - 1
}

// Brute Force
// More straightforward solution
func FindPeakElement2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 && nums[i] > nums[i+1] {
			return i
		}
		if i == len(nums)-1 && nums[i] > nums[i-1] {
			return i
		}
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
	}

	return -1
}

func FindPeakElement3(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := ((right - left) / 2) + left

		// Cut the left part
		if nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			// Cut the right part
			right = mid
		}
	}

	return left
}
