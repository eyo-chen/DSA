package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Using DFS(recursion) with slice path
func SumNumbers(root *TreeNode) int {
	return helper(root, []int{})
}

func helper(node *TreeNode, path []int) int {
	if node == nil {
		return 0
	}

	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil {
		return sum(path)
	}

	left := helper(node.Left, path)
	right := helper(node.Right, path)

	return left + right
}

func sum(arr []int) int {
	ans := 0
	idx := 0
	for i := len(arr) - 1; i >= 0; i-- {
		ans += arr[i] * int(math.Pow(10, float64(idx)))
		idx++
	}

	return ans
}

// Using DFS(recursion) with int sum
func SumNumbers2(root *TreeNode) int {
	return helper2(root, 0)
}

func helper2(node *TreeNode, accVal int) int {
	if node == nil {
		return 0
	}

	accVal = accVal*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return accVal
	}

	left := helper2(node.Left, accVal)
	right := helper2(node.Right, accVal)

	return left + right
}

// Using BFS(queue) with int sum
type nodeInfo struct {
	node   *TreeNode
	accVal int
}

func SumNumbers3(root *TreeNode) int {
	ans := 0
	queue := []*nodeInfo{}
	queue = append(queue, &nodeInfo{node: root, accVal: 0})

	for len(queue) > 0 {
		ni := queue[0]
		queue = queue[1:]

		if ni.node == nil {
			continue
		}

		sum := ni.accVal*10 + ni.node.Val
		if ni.node.Left == nil && ni.node.Right == nil {
			ans += sum
		}

		queue = append(queue, &nodeInfo{node: ni.node.Left, accVal: sum})
		queue = append(queue, &nodeInfo{node: ni.node.Right, accVal: sum})
	}

	return ans
}
