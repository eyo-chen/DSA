package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	dfs(root, targetSum, []int{}, &ans)
	return ans
}

func dfs(node *TreeNode, targetSum int, path []int, ans *[][]int) {
	if node == nil {
		return
	}

	path = append(path, node.Val)
	remainingSum := targetSum - node.Val

	if node.Left == nil && node.Right == nil {
		if remainingSum == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			*ans = append(*ans, tmp)
		}

		return
	}

	dfs(node.Left, remainingSum, path, ans)
	dfs(node.Right, remainingSum, path, ans)
}
