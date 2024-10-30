package main

// Using Hash Table
func SortColors(nums []int) {
	hashTable := make([]int, 3)

	// Count the frequency of each number
	for _, n := range nums {
		hashTable[n]++
	}

	// idx is the index where we should fill the number
	idx := 0

	// Fill the array with the numbers based on the frequency
	for v, freq := range hashTable {
		// As long as the frequency is greater than 0, we fill the array with the number
		for freq > 0 {
			nums[idx] = v
			idx++
			freq--
		}
	}
}

// Loop through the array three times
func SortColors2(nums []int) {
	// ptr is the pointer where we should fill the number
	ptr := 0

	// Loop through the array three times
	// e.g. first time we aim to fill 0s, second time we aim to fill 1s, third time we aim to fill 2s
	for i := 0; i < 3; i++ {
		// Loop through the array from the pointer to the end
		for k := ptr; k < len(nums); k++ {
			// If the number is the same as the current number we're aiming to fill, we swap it with the pointer
			if nums[k] == i {
				nums[k], nums[ptr] = nums[ptr], nums[k]
				ptr++
			}
		}
	}
}

func SortColors3(nums []int) {
	ptr := 0
	left := 0
	right := len(nums) - 1

	for ptr <= right {
		// If the number is 0, we swap it with the number at the left pointer
		// Then, we update both the left and pointer
		if nums[ptr] == 0 {
			nums[left], nums[ptr] = nums[ptr], nums[left]
			left++
			ptr++
			continue
		}

		// If the number is 2, we swap it with the number at the right pointer
		if nums[ptr] == 2 {
			nums[right], nums[ptr] = nums[ptr], nums[right]
			right--
			continue
		}

		// If the number is 1, we just update the pointer
		ptr++
	}
}
