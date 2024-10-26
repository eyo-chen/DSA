package main

func CombinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	temp := []int{}

	// We can sort the candidates first to optimize the performance
	// So that we can skip the elements that are larger than the target
	dfs(candidates, &ans, temp, 0, target)

	return ans
}

func dfs(candidates []int, ans *[][]int, temp []int, idx int, target int) {
	if target == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		*ans = append(*ans, t)
		return
	}

	if target < 0 {
		return
	}

	for i := idx; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		dfs(candidates, ans, temp, i, target-candidates[i])
		temp = temp[:len(temp)-1]
	}
}
