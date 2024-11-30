package main

func Permute(nums []int) [][]int {
	ans := [][]int{}
	dfs(nums, []int{}, make([]bool, len(nums)), &ans)
	return ans
}

func dfs(nums []int, curr []int, hashTable []bool, ans *[][]int) {
	if len(curr) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return
	}

	for i, n := range nums {
		if hashTable[i] {
			continue
		}

		curr = append(curr, n)
		hashTable[i] = true

		dfs(nums, curr, hashTable, ans)

		curr = curr[:len(curr)-1]
		hashTable[i] = false
	}
}
