package main

func Jump(nums []int) int {
	memo := make([]int, len(nums))
	return helper(nums, 0, memo)
}

func helper(nums []int, index int, memo []int) int {
	if index == len(nums)-1 {
		return 0
	}

	if memo[index] != 0 {
		return memo[index]
	}

	steps := nums[index]
	if steps == 0 {
		return -1
	}

	// Initialize minStep to the maximum possible value
	// The maximum number of jumps required to reach the end of the array is the length of the array
	minStep := len(nums)

	// Iterate over the possible steps from 1 to steps
	// also make sure that we don't go out of the bounds of the array
	for i := 1; i <= steps && index+i < len(nums); i++ {
		curStep := helper(nums, index+i, memo)

		// If the current step is not -1, it means that we can reach the end of the array from the current index
		// so we update the minStep
		if curStep != -1 {
			minStep = min(minStep, curStep+1)
		}
	}

	memo[index] = minStep
	return minStep
}

func Jump1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps, leftPtr, rightPtr := 0, 0, 0
	for i := 0; i < n-1; i++ {
		steps := nums[i]
		// keep track of the maximum steps we can take
		rightPtr = max(rightPtr, i+steps)

		// If we're still within the current interval, continue
		if i != leftPtr {
			continue
		}

		// If we're at the end of the current interval, we need to make a jump
		// and update the leftPtr to the rightPtr
		// and increment the jumps
		jumps++
		leftPtr = rightPtr
		if leftPtr >= n-1 {
			break
		}
	}

	return jumps
}
