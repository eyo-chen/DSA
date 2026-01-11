package main

import (
	"sort"
)

/*
NestedLoopJoin performs a classic nested-loop equi-join.

Approach:
- For each element in the outer table, scan the entire inner table
- Emit a result whenever the join keys are equal
- Most straightforward join algorithm to implement

Time Complexity:
- O(N * M), where N = len(outer), M = len(inner)
- This is the least efficient approach for large datasets
- Every outer row causes a full scan of the inner table

Space Complexity:
- O(K), where K is the number of join results
- No additional data structures needed beyond the result set
- Minimal memory footprint

When to Use:
- Small tables where N * M is acceptable (e.g., both tables < 1000 rows)
- No indexes available on join keys
- Inner table fits entirely in memory
- One table is very small (can cache in CPU cache)

Pros:
- Simple to implement and understand
- No preprocessing required (no sorting, no hash table building)
- Works with any comparison operator (=, <, >, etc.)
- Minimal memory overhead
- No risk of hash collisions or sort overhead

Cons:
- Extremely poor performance on large tables (quadratic complexity)
- Does not leverage indexes or data structures
- Inefficient CPU cache usage due to repeated full scans
- Not scalable - 10x data means 100x slower
- Inner table is scanned repeatedly (wasteful I/O in databases)
*/
func NestedLoopJoin(outer []int, inner []int) []int {
	results := []int{}

	// Iterate through each row in the outer table
	for _, outerKey := range outer {
		// For each outer row, scan all rows in the inner table
		for _, innerKey := range inner {
			// When keys match, emit the join result
			if outerKey == innerKey {
				results = append(results, outerKey)
			}
		}
	}

	return results
}

/*
SortMergeJoin performs a sort-merge equi-join.

Approach:
1. Sort both tables by the join key
2. Use two pointers to scan both tables in parallel
3. When keys match, collect all duplicates on both sides
4. Emit the cross product of matching groups
5. Advance pointers based on comparison results

Time Complexity:
- Sorting phase: O(N log N + M log M)
- Merge phase: O(N + M + K), where K is the number of results
- Overall: O(N log N + M log M)
- More efficient than nested-loop when N and M are large
- Dominates when tables are already sorted or have indexes

Space Complexity:
- O(N + M) for defensive copies of input tables
- O(K) for join results
- Total: O(N + M + K)
- Can be O(K) if tables are pre-sorted and in-place modification is allowed

When to Use:
- Both tables are already sorted on the join key (e.g., clustered indexes)
- Large tables where sorting cost is justified
- Output needs to be sorted (get sorted results "for free")
- Memory is available for sorting but not for large hash tables
- Database has efficient external sort algorithms for data that doesn't fit in RAM

Pros:
- Predictable performance - O(N log N) sorting is well-understood
- Efficient for pre-sorted data (no sorting needed, just O(N + M) merge)
- Handles duplicate keys elegantly
- Good CPU cache locality during merge phase (sequential access)
- Works well with disk-based data (sequential I/O is fast)
- Can leverage existing indexes on join columns
- Stable performance regardless of data distribution

Cons:
- Sorting overhead if data is not pre-sorted
- Requires O(N + M) extra space for sorting (or in-place modification)
- Not optimal if only small portion of data will match
- Less efficient than hash join for unsorted data
- Sorting can be expensive for very large datasets that don't fit in memory
*/
func SortMergeJoin(outer []int, inner []int) []int {
	// Create defensive copies to avoid mutating the caller's input
	sortedOuter := append([]int(nil), outer...)
	sortedInner := append([]int(nil), inner...)

	// Sort both tables on the join key
	sort.Ints(sortedOuter)
	sort.Ints(sortedInner)

	results := []int{}

	// Initialize pointers for both sorted tables
	outerIdx := 0
	innerIdx := 0

	// Merge phase: walk through both sorted tables simultaneously
	for outerIdx < len(sortedOuter) && innerIdx < len(sortedInner) {
		if sortedOuter[outerIdx] < sortedInner[innerIdx] {
			// Outer key is smaller — no match possible, advance outer pointer
			outerIdx++

		} else if sortedOuter[outerIdx] > sortedInner[innerIdx] {
			// Inner key is smaller — no match possible, advance inner pointer
			innerIdx++

		} else {
			// Keys are equal — handle all duplicates of this join key
			matchKey := sortedOuter[outerIdx]

			// Find the range of duplicate keys in the outer table
			outerStart := outerIdx
			for outerIdx < len(sortedOuter) && sortedOuter[outerIdx] == matchKey {
				outerIdx++
			}
			outerEnd := outerIdx

			// Find the range of duplicate keys in the inner table
			innerStart := innerIdx
			for innerIdx < len(sortedInner) && sortedInner[innerIdx] == matchKey {
				innerIdx++
			}
			innerEnd := innerIdx

			// Emit cross product: for each outer match, join with each inner match
			// If outer has 3 duplicates and inner has 2, we emit 3 × 2 = 6 results
			for i := outerStart; i < outerEnd; i++ {
				for j := innerStart; j < innerEnd; j++ {
					results = append(results, matchKey)
				}
			}
		}
	}

	return results
}

/*
HashJoin performs a hash-based equi-join.

Approach:
1. Build phase: scan the outer table and build a hash table
  - Key: join key value
  - Value: count of occurrences

2. Probe phase: scan the inner table and probe the hash table
  - For each inner key, emit results based on the count in the hash table

3. This is a "hash semi-join" variant optimized for equality checks

Time Complexity:
- Build phase: O(N) to build hash table
- Probe phase: O(M + K), where K is the number of results
- Overall: O(N + M + K)
- Most efficient for large datasets with good hash distribution
- Best case: linear time in total input size

Space Complexity:
- O(N) for the hash table storing outer table keys and counts
- O(K) for join results
- Total: O(N + K)
- Space depends on number of distinct keys in outer table

When to Use:
- Large tables where sorting is expensive
- No existing indexes on join columns
- Sufficient memory available for hash table
- Unsorted data that would require expensive sorting
- Join key has good hash distribution (minimal collisions)
- One table significantly smaller than the other (use smaller as build input)

Pros:
- Fastest join algorithm for large unsorted datasets
- O(N + M) time complexity - linear in input size
- Single pass over each table (no repeated scans)
- Excellent CPU cache performance during probe phase
- No sorting required
- Scales well with data size
- Most commonly used in modern database systems

Cons:
- Requires O(N) additional memory for hash table
- Performance degrades with hash collisions (poor hash distribution)
- Hash table may not fit in memory for very large outer tables
- Doesn't preserve any order in results
- Only works for equality joins (=), not range joins (<, >, etc.)
- Memory-intensive if outer table has many distinct keys
- Requires good hash function to avoid collisions

Database Optimization Notes:
- DBs often use "hash partitioning" to handle large tables
- Hybrid hash join: partition data when hash table doesn't fit in RAM
- Choose smaller table as build input to minimize hash table size
- Can use Bloom filters to skip non-matching inner rows early
*/
func HashJoin(outer []int, inner []int) []int {
	results := []int{}

	// Build phase: create hash table from outer table
	// Map stores: key -> count of occurrences
	hashTable := map[int]int{}
	for _, key := range outer {
		hashTable[key]++
	}

	// Probe phase: scan inner table and lookup in hash table
	for _, key := range inner {
		count := hashTable[key]

		// If key exists in hash table, emit one result per occurrence
		// Example: if outer has 3 occurrences of key=5, and inner has key=5,
		// we emit 3 results
		for i := 0; i < count; i++ {
			results = append(results, key)
		}
	}

	return results
}
