package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
// Time Complexity: O(N)
// Space Complexity: O(N)
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	ans := []int{}

	for len(queue) > 0 {
		size := len(queue)

		// Traverse the current level
		for i := 0; i < size; i++ {
			// Get the first node in the queue
			node := queue[0]
			queue = queue[1:]

			// Add the first node's value to the answer
			// Note that because we use BFS, we traverse the tree level by level
			// So the first node in the queue is the rightmost node of the current level
			if i == 0 {
				ans = append(ans, node.Val)
			}

			// Add the right child to the queue if it exists
			// It's important to add the right child first because we want to get the rightmost node first
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			// Add the left child to the queue if it exists
			// This ensures that we will process the leftmost node last
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}

	return ans
}

// DFS
// Time Complexity: O(N)
// Space Complexity: O(N)
func RightSideView2(root *TreeNode) []int {
	ans := []int{}
	dfs(root, &ans, 0)
	return ans
}

func dfs(node *TreeNode, ans *[]int, level int) {
	if node == nil {
		return
	}

	if level == len(*ans) {
		*ans = append(*ans, node.Val)
	}

	dfs(node.Right, ans, level+1)
	dfs(node.Left, ans, level+1)
}
