package main

// Brute Force
func MaxArea(height []int) int {
	ans := 0

	for i := 0; i < len(height); i++ {
		minHeight := height[i]

		for k := i + 1; k < len(height); k++ {
			minHeight = min(minHeight, height[k])
			ans = max(ans, minHeight*(k-i))
		}
	}

	return ans
}

// Two Pointers
func MaxArea2(height []int) int {
	left, right := 0, len(height)-1
	ans := 0

	for left < right {
		// Calculate the area with the two pointers
		minHeight := min(height[left], height[right])
		ans = max(ans, minHeight*(right-left))

		// Move the pointer with the smaller height inward
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return ans
}
