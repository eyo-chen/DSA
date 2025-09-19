package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := helper(node.Left)
	if left == -1 {
		return -1
	}

	right := helper(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) > 1 {
		return -1
	}

	return max(left, right) + 1
}

type nodeInfo struct {
	level      int
	isBalanced bool
}

func IsBalanced1(root *TreeNode) bool {
	res := helper1(root)
	return res.isBalanced
}

func helper1(root *TreeNode) nodeInfo {
	if root == nil {
		return nodeInfo{level: 0, isBalanced: true}
	}

	right := helper1(root.Right)
	if !right.isBalanced {
		return nodeInfo{isBalanced: false}
	}

	left := helper1(root.Left)
	if !left.isBalanced {
		return nodeInfo{isBalanced: false}
	}

	isBalanced := math.Abs(float64(right.level-left.level)) <= 1
	level := max(right.level, left.level) + 1

	return nodeInfo{level: level, isBalanced: isBalanced}
}
