package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS(recursion)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMaxDepth := MaxDepth(root.Left) + 1
	rightMaxDepth := MaxDepth(root.Right) + 1

	return max(leftMaxDepth, rightMaxDepth)
}

// BFS(iteration)
func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	ans := 0

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		ans++
	}

	return ans
}

type nodeInfo struct {
	node  *TreeNode
	level int
}

// BFS(iteration with nodeInfo)
func MaxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []nodeInfo{{node: root, level: 1}}
	ans := 0

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		ans = max(ans, n.level)

		if n.node.Left != nil {
			q = append(q, nodeInfo{node: n.node.Left, level: n.level + 1})
		}
		if n.node.Right != nil {
			q = append(q, nodeInfo{node: n.node.Right, level: n.level + 1})
		}
	}

	return ans
}
