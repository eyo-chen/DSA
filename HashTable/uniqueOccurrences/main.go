package main

import "sort"

// Using Hash Table
// Time Complexity O(n)
// Space Complexity O(n)
func UniqueOccurrences(arr []int) bool {
	// First, we count the frequency of each element in the array
	// freqs is the first hash table to store the frequency of each element
	freqs := map[int]int{}
	for _, v := range arr {
		freqs[v]++
	}

	// Then, we use another hash table to check if the frequency of each element is unique
	// hashTable is the second hash table to store the frequency of each element(check for uniqueness)
	hashTable := map[int]bool{}
	for _, v := range freqs {
		if hashTable[v] {
			return false
		}

		hashTable[v] = true
	}

	// If we can successfully iterate through the hash table, it means all frequencies are unique
	return true
}

// Using Sorting
// Time Complexity O(nlogn)
// Space Complexity O(n)
func UniqueOccurrences2(arr []int) bool {
	// First, we sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	// Then, we count the frequency of each element in the array
	// counts is the array to store the frequency of each element
	counts := []int{}
	count := 1
	for i := 0; i < len(arr); i++ {
		// If the current element is the same as the next element, we increment the count
		if i+1 < len(arr) && arr[i] == arr[i+1] {
			count++
		} else {
			// Otherwise, we append the count to the counts array and reset the count
			// which means that we start counting the frequency of the next element
			counts = append(counts, count)
			count = 1
		}
	}

	// Then, we sort the counts array to check for uniqueness
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	// Finally, we check if there are any duplicate frequencies in the counts array
	for i := 1; i < len(counts); i++ {
		if counts[i] == counts[i-1] {
			return false
		}
	}

	return true
}
