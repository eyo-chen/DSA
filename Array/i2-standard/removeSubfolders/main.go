package main

import (
	"sort"
	"strings"
)

// RemoveSubfolders removes all sub-folders from a list of folder paths using sorting approach
// Time Complexity: O(n log n + n*m) where n = number of folders, m = average path length
// Space Complexity: O(1) excluding result array
//
// Algorithm:
// 1. Sort folders lexicographically - this ensures parent folders appear before their subfolders
// 2. Iterate through sorted folders and only keep folders that aren't subfolders of the previously added folder
// 3. A folder is a subfolder if it starts with "parent/" pattern
//
// Example: ["/a", "/a/b", "/c"] -> after sorting: ["/a", "/a/b", "/c"]
// - Add "/a" (first folder)
// - Skip "/a/b" (starts with "/a/")
// - Add "/c" (doesn't start with "/a/")
func RemoveSubfolders(folder []string) []string {
	// Step 1: Sort folders lexicographically
	// After sorting, any subfolder will appear immediately after its parent
	// Example: ["/a/b", "/a", "/c"] becomes ["/a", "/a/b", "/c"]
	sort.Strings(folder)

	// Step 2: Initialize result with first folder (it can't be a subfolder of anything)
	ans := []string{folder[0]}

	// Step 3: Check each remaining folder
	for i := 1; i < len(folder); i++ {
		// Get the last added folder and append "/" to check for subfolder pattern
		// We need to add "/" to ensure exact parent-child relationship
		// Example: "/a" + "/" = "/a/"
		// This prevents "/ab" from being considered a subfolder of "/a"
		lastAddedFolder := ans[len(ans)-1] + "/"

		// Check if current folder is NOT a subfolder of the last added folder
		// If current folder doesn't start with "lastAddedFolder/", it's not a subfolder
		if !strings.HasPrefix(folder[i], lastAddedFolder) {
			ans = append(ans, folder[i])
		}
		// If it does start with lastAddedFolder/, it's a subfolder and we skip it
	}

	return ans
}

// RemoveSubfolders1 removes all sub-folders using hash table approach
// Time Complexity: O(n * m²) where n = number of folders, m = average path length
// Space Complexity: O(n * m) for storing all folders in hash table
//
// Algorithm:
// 1. Store all folders in a hash table for O(1) lookup
// 2. For each folder, check all possible parent paths by finding '/' characters
// 3. If any parent path exists in hash table, current folder is a subfolder
// 4. Only keep folders that aren't subfolders of any other folder
//
// Example: For folder "/a/b/c", check if "/a" and "/a/b" exist in hash table
func RemoveSubfolders1(folder []string) []string {
	ans := []string{}

	// Step 1: Create hash table for O(1) folder existence lookup
	// This allows us to quickly check if a potential parent folder exists
	hashTable := map[string]bool{}
	for _, f := range folder {
		hashTable[f] = true
	}

	// Step 2: Check each folder to see if it's a subfolder of any other folder
	for _, dir := range folder {
		isSubFolder := false

		// Step 3: Check all possible parent paths for current directory
		// We iterate through each character and look for '/' separators
		// Each '/' marks a potential parent folder boundary
		for i := 1; i < len(dir); i++ { // Start from 1 to skip the leading '/'
			if dir[i] == '/' {
				// Extract potential parent folder path up to current '/'
				// Example: For "/a/b/c" at position 2, parentFolder = "/a"
				parentFolder := dir[:i]

				// Check if this parent folder exists in our original folder list
				if hashTable[parentFolder] {
					// Found a parent folder! Current folder is a subfolder
					isSubFolder = true
					break // No need to check further - we found one parent
				}
			}
		}

		// Step 4: Only add folders that are NOT subfolders
		if !isSubFolder {
			ans = append(ans, dir)
		}
	}

	return ans
}
