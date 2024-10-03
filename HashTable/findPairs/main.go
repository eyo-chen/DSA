package main

// Brute force
func FindPairs(nums []int, k int) int {
	// uniquePairs is used to store unique pairs
	uniquePairs := make(map[[2]int]bool)

	// Iterate through the array
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// If the absolute difference is not equal to k, continue
			if abs(nums[i]-nums[j]) != k {
				continue
			}

			// Find the smaller and larger number
			// It's for storing the pair in a unique way
			smaller, larger := nums[i], nums[j]
			if nums[i] > nums[j] {
				smaller, larger = nums[j], nums[i]
			}

			// Store the pair in the map
			uniquePairs[[2]int{smaller, larger}] = true
		}
	}

	// Return the number of unique pairs
	return len(uniquePairs)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Hash table(2 passes)
func FindPairs2(nums []int, k int) int {
	// hashTable is used to store the number of times a number appears in the array
	hashTable := make(map[int]int)
	ans := 0

	// Store the number of times a number appears in the array
	for _, num := range nums {
		hashTable[num]++
	}

	// Iterate through the hash table
	for num := range hashTable {
		// If k is 0, we only need to check if the number appears more than once
		// It's kind of special case
		if k == 0 {
			if hashTable[num] > 1 {
				ans++
			}
			continue
		}

		// If the number plus k exists in the hash table, add it to the unique pairs
		// Here, we don't need to care about [num - k] because
		// If (a, b) is a valid pair and a < b, we'll find it when we process a (looking for a+k).
		// If (a, b) is a valid pair and a > b, we'll find it when we process b (looking for b+k).
		if hashTable[num+k] > 0 {
			ans++
		}
	}

	// Return the number of unique pairs
	return ans
}

// Hash table(1 pass)
func FindPairs3(nums []int, k int) int {
	hashTable := map[int]int{}
	ans := 0

	// Iterate through the array
	for _, n := range nums {
		// If the number exists in the hash table
		// That means we've checked it's (+k) and (-k) pair already
		// So we only need to check if k is 0 and the number appears more than once
		if v, ok := hashTable[n]; ok {
			if k == 0 && v == 1 {
				ans++
			}

			hashTable[n]++
			continue
		}
		// If the number haven't been checked before, check its (+k) and (-k) pair

		// If the number minus k exists in the hash table, update the answer
		if _, ok := hashTable[n-k]; ok {
			ans++
		}

		// If the number plus k exists in the hash table, update the answer
		if _, ok := hashTable[n+k]; ok {
			ans++
		}

		// Store the number in the hash table
		hashTable[n] = 1
	}

	return ans
}
