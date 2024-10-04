package main

import (
	"math"
	"sort"
	"strconv"
)

// Brute Force
func FindMinDifference(timePoints []string) int {
	minVal, maxVal := 24*60, 0
	ans := 24 * 60

	for i := 0; i < len(timePoints); i++ {
		iVal := toMinute(timePoints[i])
		for k := i + 1; k < len(timePoints); k++ {
			kVal := toMinute(timePoints[k])

			ans = min(ans, int(math.Abs(float64(iVal-kVal))))
		}

		// Keep track of the minimum and maximum time points
		// Because we need these two values to handle the wraparound difference
		minVal = min(minVal, iVal)
		maxVal = max(maxVal, iVal)
	}

	// Handle the wraparound difference
	ans = min(ans, 24*60+minVal-maxVal)
	return ans
}

// Sorting
func FindMinDifference1(timePoints []string) int {
	// Convert all time points to minutes
	minutes := make([]int, len(timePoints))
	for i, t := range timePoints {
		minutes[i] = toMinute(t)
	}

	// Sort the time points
	sort.Slice(minutes, func(i, j int) bool {
		return minutes[i] < minutes[j]
	})

	// Calculate the difference between each time points
	ans := 24 * 60
	for i := 0; i < len(minutes)-1; i++ {
		ans = min(ans, minutes[i+1]-minutes[i])
	}

	// Handle the wraparound difference
	// The first time point is the smallest, and the last time point is the largest
	ans = min(ans, 24*60+minutes[0]-minutes[len(minutes)-1])
	return ans
}

// Hash Table
func FindMinDifference2(timePoints []string) int {
	hashTable := make([]bool, 60*24)

	// Build the hash table
	// If a time point is already in the hash table, return 0
	// because we find the minimum difference
	// Otherwise, we set the hash table to true to indicate that this time point is in the hash table
	for _, t := range timePoints {
		m := toMinute(t)
		if hashTable[m] {
			return 0
		}

		hashTable[m] = true
	}

	// Keep track of the first and previous time points
	// first is for handling the wraparound difference
	// prev is for calculating the difference between the current time point and the previous time point
	first, prev := -1, -1
	ans := 24*60 + 1

	for i, seen := range hashTable {
		// If the time point is not in the hash table, continue
		if !seen {
			continue
		}

		// If the previous time point is not set, it means this is the first time point
		// So we set the first time point to the current time point
		if prev == -1 {
			first = i
		} else {
			// Otherwise, calculate the difference between the current time point and the previous time point
			ans = min(ans, i-prev)
		}

		// Update the previous time point to the current time point
		prev = i
	}

	// Handle the wraparound difference
	// The first time point is the smallest, and the prev is the largest
	ans = min(ans, 24*60+first-prev)
	return ans
}

func toMinute(s string) int {
	hour, _ := strconv.Atoi(s[0:2])
	minute, _ := strconv.Atoi(s[3:5])

	return hour*60 + minute
}
