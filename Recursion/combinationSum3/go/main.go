package main

func CombinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	helper(k, n, 1, &ans, []int{})
	return ans
}

func helper(k int, n int, index int, ans *[][]int, cur []int) {
	if n < 0 {
		return
	}

	if n == 0 && len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ans = append(*ans, tmp)
		return
	}

	if len(cur) == k {
		return
	}

	for i := index; i <= 9; i++ {
		helper(k, n-i, i+1, ans, append(cur, i))
	}
}
