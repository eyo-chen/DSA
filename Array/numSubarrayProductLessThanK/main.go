package main

import "fmt"

func NumSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		curNum := nums[i]
		if curNum < k {
			ans++
		} else {
			continue
		}

		acc := curNum
		for j := i + 1; j < len(nums); j++ {
			acc *= nums[j]
			if acc < k {
				ans++
			} else {
				break
			}
		}
	}

	return ans
}

func NumSubarrayProductLessThanK2(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	left, right := 0, 0
	curProduct := 1
	ans := 0

	for right < len(nums) {
		curProduct *= nums[right]

		// e.g. [57,44,92,28,66,60,37,33,52,38,29,76,8,75,22], k = 18
		// output: 18, expected: 1
		for curProduct >= k && left <= right {
			curProduct /= nums[left]
			left++
		}

		ans += right - left + 1
		right++
	}

	return ans
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	left, right := 0, 0
	curProduct := 1
	ans := 0

	for right < len(nums) {
		curProduct *= nums[right]

		if curProduct < k {
			ans++
			right++
			continue
		}

		for curProduct >= k && left <= right {
			curProduct /= nums[left]
			left++
		}
		ans++
	}

	return ans
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
