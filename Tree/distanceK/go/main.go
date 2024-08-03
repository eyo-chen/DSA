package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DistanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	adj := map[int][]int{}
	genAdj(root, adj)

	q := []int{target.Val}
	seen := map[int]bool{target.Val: true}

	for len(q) > 0 {
		// if k == 0, return the current layer's node
		if k == 0 {
			return q
		}
		k--

		// BFS
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			for _, nei := range adj[node] {
				if !seen[nei] {
					q = append(q, nei)
					seen[nei] = true
				}
			}
		}
	}

	return q
}

// genAdj generates adjacency list for the tree
func genAdj(root *TreeNode, adj map[int][]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		// add child node to parent node's adjacency list
		adj[root.Val] = append(adj[root.Val], root.Left.Val)

		// add parent node to child node's adjacency list
		adj[root.Left.Val] = append(adj[root.Left.Val], root.Val)
		genAdj(root.Left, adj)
	}

	if root.Right != nil {
		// add child node to parent node's adjacency list
		adj[root.Val] = append(adj[root.Val], root.Right.Val)

		// add parent node to child node's adjacency list
		adj[root.Right.Val] = append(adj[root.Right.Val], root.Val)
		genAdj(root.Right, adj)
	}
}
