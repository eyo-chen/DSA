package main

import "sort"

// Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func MinimumAbsDifference(arr []int) [][]int {
	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Find the smallest difference
	smallDiff := arr[len(arr)-1]
	for i := 1; i < len(arr); i++ {
		smallDiff = min(smallDiff, arr[i]-arr[i-1])
	}

	ans := [][]int{}

	// Find all pairs with the smallest difference
	for i := 0; i < len(arr); i++ {
		for k := i + 1; k < len(arr); k++ {
			if arr[k]-arr[i] == smallDiff {
				ans = append(ans, []int{arr[i], arr[k]})
			}
		}
	}

	return ans
}

// Optimized (Two Pass)
// Time Complexity: O(n)
// Space Complexity: O(1)
func MinimumAbsDifference2(arr []int) [][]int {
	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Find the smallest difference
	smallDiff := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		smallDiff = min(smallDiff, arr[i]-arr[i-1])
	}

	ans := [][]int{}

	// Find all pairs with the smallest difference in one pass
	// Why we can do this?
	// Because up to this point, we have already found the smallest difference, and get the sorted array
	// For each pair, the difference is only have TWO cases:
	// 1. The difference is equal to the smallest difference
	//    - We can directly append the pair to the answer
	//    - Also, we can update two pointers(i, i - 1) at a time
	// 2. The difference is greater than the smallest difference
	//    - We can directly skip this pair
	//    - Also, we can update two pointers(i, i - 1) at a time
	// For either case, we both have to update two pointers at a time
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == smallDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}

// Optimized (One Pass)
// Time Complexity: O(n)
// Space Complexity: O(1)
func MinimumAbsDifference3(arr []int) [][]int {
	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Loop through the array, find the smallest difference and build the pair at the same time
	// Why we can do this?
	// Because up to this point, we have already found the sorted array
	// For each pair, the difference is only have THREE cases:
	// 1. The current difference is equal to the smallest difference
	//    - We can directly append the pair to the answer
	//    - Also, we can update two pointers(i, i - 1) at a time
	// 2. The current difference is greater than the smallest difference
	//    - We can directly skip this pair
	//    - Also, we can update two pointers(i, i - 1) at a time
	// 3. The current difference is less than the smallest difference
	//    - In this case, we know that "All the previous pairs in the answer are not the smallest difference, so they are definitely NOT the answer"
	//    - So we can directly clear the answer and update the smallest difference
	ans := [][]int{}
	smallDiff := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		// Current difference is less than the smallest difference
		// So we know that "All the previous pairs in the answer are not the smallest difference, so they are definitely NOT the answer"
		// So we can directly clear the answer and update the smallest difference
		if diff < smallDiff {
			ans = [][]int{}
			smallDiff = diff
		}

		// Current difference is equal to the smallest difference
		// So we can directly append the pair to the answer
		if smallDiff == diff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}
