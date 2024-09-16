package main

func ClimbStairs(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	amount := 0
	for i := 1; i <= 2; i++ {
		amount += ClimbStairs(n - i)
	}

	return amount
}

func ClimbStairs1(n int) int {
	memo := map[int]int{}
	return helper(n, memo)
}

func helper(n int, memo map[int]int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if v, ok := memo[n]; ok {
		return v
	}

	amount := 0
	for i := 1; i <= 2; i++ {
		amount += helper(n-i, memo)
	}

	memo[n] = amount
	return amount
}

func ClimbStairs2(n int) int {
	table := make([]int, n+1)
	table[0] = 1

	for i := 1; i <= n; i++ {
		for step := 1; step <= 2; step++ {
			if i >= step {
				table[i] += table[i-step]
			}
		}
	}

	return table[n]
}
