package main

func LongestMountain(arr []int) int {
	ans := 0
	up := make([]int, len(arr))
	down := make([]int, len(arr))

	// Build up array (from left to right)
	for i := 1; i < len(arr); i++ {
		// Only update the value if the current value is greater than the previous value
		// Means that the value is increasing
		if arr[i] > arr[i-1] {
			up[i] = up[i-1] + 1
		}
	}

	// Build down array (from right to left)
	for i := len(arr) - 2; i > 0; i-- {
		// Only update the value if the current value is greater than the next value
		// Means that the value is increasing
		if arr[i] > arr[i+1] {
			down[i] = down[i+1] + 1
		}
	}

	// Loop through the array to find the longest mountain
	for i := 0; i < len(arr); i++ {
		if up[i] > 0 && down[i] > 0 {
			ans = max(ans, up[i]+down[i]+1)
		}
	}

	return ans
}

func LongestMountain2(arr []int) int {
	ans := 0

	if len(arr) < 3 {
		return 0
	}

	for i := 1; i < len(arr)-1; i++ {
		// If the current value is not a peak, skip
		// A peak is the value that is greater than the previous and next value
		// e.g. [a,b,c]
		// b > a && b > c
		if arr[i-1] >= arr[i] || arr[i] <= arr[i+1] {
			continue
		}

		// Find the left-most side of the mountain
		// Keep moving left pointer to the left until the value is not decreasing
		left := i - 1
		for left > 0 && arr[left] > arr[left-1] {
			left--
		}

		// Find the right-most side of the mountain
		// Keep moving right pointer to the right until the value is not decreasing
		right := i + 1
		for right < len(arr)-1 && arr[right] > arr[right+1] {
			right++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}

func LongestMountain3(arr []int) int {
	ptr := 0
	ans := 0

	for ptr < len(arr) {
		// Set the starting point of the mountain
		start := ptr

		// Keep moving the pointer to the right until we find a peak
		// If the value is not increasing, it means we find a peak
		for ptr < len(arr)-1 && arr[ptr+1] > arr[ptr] {
			ptr++
		}

		// If the starting point is still the same as the pointer
		// It means the current value is equal to the next value
		// So it's not a mountain, we need to move the pointer to the right
		if ptr == start {
			ptr++
			continue
		}

		// Set the peak of the mountain
		peak := ptr

		// Keep moving the pointer to the right until we find the end of the mountain
		// If the value is not decreasing, it means we find the end of the mountain
		for ptr < len(arr)-1 && arr[ptr+1] < arr[ptr] {
			ptr++
		}

		// If the pointer is still the same as the peak
		// It means the current value is equal to the next value
		// So it's not a mountain, we need to move the pointer to the right
		if ptr == peak {
			ptr++
			continue
		}

		// Calculate the length of the mountain and update the longest length
		ans = max(ans, ptr-start+1)
	}

	return ans
}

func LongestMountain4(arr []int) int {
	ans := 0

	for i := 0; i < len(arr); i++ {
		end := i
		hasDecrease := false
		hasIncrease := false

		// Find the left-most side of the mountain
		// Keep moving the end pointer to the right until the value is not increasing
		for end < len(arr)-1 && arr[end] < arr[end+1] {
			end++
			hasIncrease = true
		}

		// Find the right-most side of the mountain
		// Keep moving the end pointer to the right until the value is not decreasing
		for end < len(arr)-1 && arr[end] > arr[end+1] {
			end++
			hasDecrease = true
		}

		// If the mountain has both increasing and decreasing
		// And the length of the mountain is greater than 3
		// We can update the longest length
		if hasIncrease && hasDecrease && end-i+1 >= 3 {
			ans = max(ans, end-i+1)
		}
	}

	return ans
}
