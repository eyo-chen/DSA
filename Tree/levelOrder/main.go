package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use Queue (BFS)
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	ans := [][]int{}

	for len(queue) > 0 {
		l := len(queue)
		curLevel := make([]int, l)

		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			curLevel[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, curLevel)
	}

	return ans
}

// Use DFS (recursion)
func LevelOrder2(root *TreeNode) [][]int {
	ans := [][]int{}
	dfs(root, &ans, 0)

	return ans
}

func dfs(node *TreeNode, ans *[][]int, level int) {
	if node == nil {
		return
	}

	if len(*ans) == level {
		*ans = append(*ans, []int{})
	}

	(*ans)[level] = append((*ans)[level], node.Val)

	dfs(node.Left, ans, level+1)
	dfs(node.Right, ans, level+1)
}

type nodeInfo struct {
	node  *TreeNode
	level int
}

// Use Queue (BFS) with nodeInfo
func LevelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*nodeInfo{{node: root, level: 0}}
	ans := [][]int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if len(ans) == node.level {
			ans = append(ans, []int{})
		}

		ans[node.level] = append(ans[node.level], node.node.Val)

		if node.node.Left != nil {
			queue = append(queue, &nodeInfo{node: node.node.Left, level: node.level + 1})
		}

		if node.node.Right != nil {
			queue = append(queue, &nodeInfo{node: node.node.Right, level: node.level + 1})
		}
	}

	return ans
}
