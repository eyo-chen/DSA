package main

// Time Complexity: O(n)
// Space Complexity: O(1)
func LargestAltitude(gain []int) int {
	ans := 0
	prefixSum := 0
	for _, g := range gain {
		prefixSum += g
		ans = max(ans, prefixSum)
	}

	return ans
}
