package main

// Use HashTable to count the number of duplicates
func RemoveDuplicates(nums []int) int {
	hashTable := map[int]int{}

	// writeIdx is the index where we'll write the next number
	writeIdx := 0

	// Iterate through the array
	for i := 0; i < len(nums); i++ {
		// If the number of duplicates is greater than 2, we skip it
		if hashTable[nums[i]] >= 2 {
			continue
		}

		// Otherwise, we write it to the array
		hashTable[nums[i]]++
		nums[writeIdx] = nums[i]
		writeIdx++
	}

	return writeIdx
}

// Use a write index to write the result to the array
func RemoveDuplicates2(nums []int) int {
	// If the array has less than 2 elements, return the length of the array
	if len(nums) < 2 {
		return len(nums)
	}

	// duplicate is the number of duplicates of the current number
	duplicate := 1
	// curNum is the current number
	curNum := nums[0]
	// writeIdx is the index where we'll write the next number
	writeIdx := 1

	for i := 1; i < len(nums); i++ {
		// If the current number is equal to the next number, we increment the duplicate counter
		if curNum == nums[i] {
			duplicate++
		} else {
			// Otherwise, we reset the duplicate counter, and update the current number
			curNum = nums[i]
			duplicate = 1
		}

		if duplicate <= 2 {
			nums[writeIdx] = nums[i]
			writeIdx++
		}
	}

	return writeIdx
}

func RemoveDuplicates3(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// k is the position where we'll write the next number
	k := 2

	// Start from the third element
	for i := 2; i < len(nums); i++ {
		// If current number is different from number two positions back,
		// it's safe to include (we can't have more than 2 duplicates)
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
