package main

// Brute Force
func TotalFruit(fruits []int) int {
	ans := 0

	// Start from each element, find the longest subarray with at most 2 distinct integers.
	for i := 0; i < len(fruits); i++ {
		// Use a hash table to store the distinct integers.
		hashTable := map[int]bool{
			fruits[i]: true,
		}

		// Count the current subarray length.
		count := 1

		// Start from the next element, add the distinct integers to the hash table.
		for k := i + 1; k < len(fruits); k++ {
			// If the current element is not in the hash table, it means it's a new distinct integer.
			// We need to check if the hash table has reached 2 distinct integers.
			// If it has, we immediately break the loop because we can't add more distinct integers.
			if _, ok := hashTable[fruits[k]]; !ok {
				if len(hashTable) == 2 {
					break
				}
			}

			// Add the current element to the hash table.
			hashTable[fruits[k]] = true

			// Increment the count of the current subarray.
			count++
		}

		// Update the answer with the maximum count.
		ans = max(ans, count)
	}

	return ans
}

// Two Pointers
func TotalFruit2(fruits []int) int {
	ans := 0
	left, right := 0, 0
	hashTable := map[int]int{}

	for right < len(fruits) {
		f := fruits[right]
		hashTable[f]++

		// If the hash table has more than 2 distinct integers, we move the left pointer to the right until the hash table has at most 2 distinct integers.
		for len(hashTable) > 2 {
			lf := fruits[left]

			// Decrement the frequency of the integer at the left pointer.
			hashTable[lf]--

			// If the frequency of the integer at the left pointer is 0, we remove it from the hash table.
			if hashTable[lf] == 0 {
				delete(hashTable, lf)
			}

			// Move the left pointer to the right.
			left++
		}

		// Update the answer with the maximum count.
		ans = max(ans, right-left+1)

		// Move the right pointer to the right.
		right++
	}

	return ans
}
