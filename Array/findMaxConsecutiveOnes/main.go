package main

// Time Complexity O(n^2)
// Space Complexity O(1)
func FindMaxConsecutiveOnes(nums []int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				break
			}

			tmp++
		}

		ans = max(ans, tmp)
	}

	return ans
}

// Time Complexity O(n)
// Space Complexity O(1)
func FindMaxConsecutiveOnes2(nums []int) int {
	ans, tmp := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
			ans = max(ans, tmp)
		} else {
			tmp = 0
		}
	}

	return ans
}
