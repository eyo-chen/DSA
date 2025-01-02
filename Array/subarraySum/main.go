package main

// Brute force
func SubarraySum(nums []int, k int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		target := k - nums[i]
		if target == 0 {
			ans++
		}

		for k := i + 1; k < len(nums); k++ {
			target -= nums[k]
			if target == 0 {
				ans++
			}
		}
	}

	return ans
}

// This is more concise version of the brute force solution
func SubarraySum2(nums []int, k int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}

	return ans
}

// Using hash table
func SubarraySum3(nums []int, k int) int {
	hashTable := map[int]int{0: 1}
	sum := 0
	ans := 0

	for i := 0; i < len(nums); i++ {
		// Calculate the current sum of subarray
		sum += nums[i]

		// Check if the difference between the current sum and k exists in the hash table
		if v, ok := hashTable[sum-k]; ok {
			ans += v
		}

		// Store the current sum in the hash table
		hashTable[sum]++
	}

	return ans
}
