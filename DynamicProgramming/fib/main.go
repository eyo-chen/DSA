package main

/*
Time Complexity: O(n)
- Each Fibonacci number from 0 to n is computed exactly once.
- The memoization ensures that once a value is computed, it is stored in the hashTable and reused for any subsequent calls, eliminating redundant calculations inherent in the naive recursive approach.

Space Complexity: O(n)
- The hashTable stores up to n key-value pairs, one for each Fibonacci number up to n.
- The call stack can also grow up to n in the worst case, contributing to the space used by the call stack.
*/
func Fib(n int) int {
	hashTable := map[int]int{}

	return helper(n, hashTable)
}

func helper(n int, hashTable map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, ok := hashTable[n]; ok {
		return val
	}

	result := helper(n-1, hashTable) + helper(n-2, hashTable)
	hashTable[n] = result

	return result
}

func Fib2(n int) int {
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}
