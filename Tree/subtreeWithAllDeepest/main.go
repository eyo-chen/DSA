package main

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Result holds the depth and the root of subtree containing deepest nodes
type Result struct {
	depth int
	node  *TreeNode
}

func SubtreeWithAllDeepest(root *TreeNode) *TreeNode {
	return dfs(root).node
}

func dfs(node *TreeNode) Result {
	// Base case: if node is nil, return depth -1 and nil node
	if node == nil {
		return Result{depth: -1, node: nil}
	}

	// Recursively get results from left and right subtrees
	leftResult := dfs(node.Left)
	rightResult := dfs(node.Right)

	// Current node's depth is 1 + max depth from children
	currentDepth := max(leftResult.depth, rightResult.depth) + 1

	// Determine which subtree contains the deepest nodes
	if leftResult.depth > rightResult.depth {
		// Deepest nodes are only in left subtree
		return Result{depth: currentDepth, node: leftResult.node}
	} else if rightResult.depth > leftResult.depth {
		// Deepest nodes are only in right subtree
		return Result{depth: currentDepth, node: rightResult.node}
	} else {
		// Both subtrees have same depth, so current node is the LCA
		return Result{depth: currentDepth, node: node}
	}
}
