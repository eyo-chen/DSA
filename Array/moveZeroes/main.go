package main

func MoveZeroes(nums []int) {
	// zeroCounts is the number of zero elements
	zeroCounts := 0

	// tmp is the array that stores the non-zero elements in order
	tmp := []int{}

	// iterate through the array
	for _, n := range nums {
		// if the element is zero, increment the zeroCounts
		if n == 0 {
			zeroCounts++
		} else {
			// if the element is not zero, append it to the tmp array
			tmp = append(tmp, n)
		}
	}

	// create a new array that contains all zeros
	zeros := make([]int, zeroCounts)

	// append the zeros to the tmp array
	// now, tmp array contains all non-zero elements in order, followed by all zeros
	tmp = append(tmp, zeros...)

	// copy the elements from the tmp array to the original array
	/*
		for i, v := range tmp {
			nums[i] = v
		}
	*/
	copy(nums, tmp)
}

func MoveZeroes2(nums []int) {
	// ptr is the pointer that points to the last place of non-zero elements
	ptr := 0

	// iterate through the array
	for _, n := range nums {
		// if the element is not zero, put it in the front of the array, and increment the pointer
		if n != 0 {
			nums[ptr] = n
			ptr++
		}
	}

	// fill the rest of the array with zeros
	for i := ptr; i < len(nums); i++ {
		nums[i] = 0
	}
}

func MoveZeroes3(nums []int) {
	// swapIndex is a pointer to represent the index that should be swapped with when we find a non-zero element
	swapIndex := 0

	for i := 0; i < len(nums); i++ {
		// if the element is not zero, swap it with the element at swapIndex index
		// and then increment the swapIndex by 1
		if nums[i] != 0 {
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			swapIndex++
		}
	}
}
