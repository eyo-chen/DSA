package main

func Construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for k := 0; k < n; k++ {
			ans[i][k] = original[k+n*i]
		}
	}

	return ans
}

func Construct2DArray2(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < len(original); i++ {
		ans[i/n][i%n] = original[i]
	}

	return ans
}
