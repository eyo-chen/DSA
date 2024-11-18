package main

// Mutate the array
func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		// If the current value is 1, we can just update the pointer for two steps
		if flowerbed[i] == 1 {
			// Update the pointer here one time, the next loop will update second time
			i++
			continue
		}

		// If the value before and after the current one is 0, we can place a flower here
		// i - 1 < 0 is to check if the current index is the first one
		// e.g. [0, 0, 1], here the first index is able to place a flower
		// i + 1 >= len(flowerbed) is to check if the current index is the last one
		// e.g. [1, 0, 0], here the last index is able to place a flower
		if (i-1 < 0 || flowerbed[i-1] == 0) &&
			(i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
			// We can place a flower here
			n--
			flowerbed[i] = 1

			// Update the pointer here one time, the next loop will update second time
			i++
		}

		if n == 0 {
			return true
		}
	}

	return false
}

// Without mutate the array
func CanPlaceFlowers2(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		// If the current value is not 0, we can just update the pointer
		if flowerbed[i] != 0 {
			// Update the pointer here one time, the next loop will update second time
			i++
			continue
		}

		// If the value before and after the current one is 0, we can place a flower here
		// i - 1 < 0 is to check if the current index is the first one
		// e.g. [0, 0, 1], here the first index is able to place a flower
		// i + 1 >= len(flowerbed) is to check if the current index is the last one
		// e.g. [1, 0, 0], here the last index is able to place a flower
		if (i-1 < 0 || flowerbed[i-1] == 0) &&
			(i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
			// We can place a flower here
			count++

			// Update the pointer here one time, the next loop will update second time
			i++
		}
	}

	return count >= n
}

// Using Math
func CanPlaceFlowers3(flowerbed []int, n int) bool {
	zeros := 1 // Start with 1 to handle edge case at the beginning
	count := 0

	for _, f := range flowerbed {
		if f == 0 {
			zeros++
		} else {
			count += (zeros - 1) / 2
			zeros = 0
		}
	}

	// Handle the last group of zeros
	zeros++ // Add one more zero to handle edge case at the end
	count += (zeros - 1) / 2

	return count >= n
}
