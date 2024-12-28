package main

import (
	"sort"
	"strconv"
	"strings"
)

// Time: O(nlogn)
// Space: O(n)
func LargestNumber(nums []int) string {
	// convert the numbers to strings
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	// sort the strings based on the custom comparator
	// e.g. ["12", "9"] -> ["9", "12"] because "912" > "129"
	// In short, we want to sort the array in such a way that the concatenated string of the numbers is the largest
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// if the first string is "0", return "0"
	// Why we can directly return "0"?
	// Because the array is sorted in such a way that the concatenated string is the largest
	// If the first string is "0", it means all the remaining strings are also "0"
	// So the concatenated string is "0"
	if strs[0] == "0" {
		return "0"
	}

	// join the strings with an empty string
	return strings.Join(strs, "")
}
