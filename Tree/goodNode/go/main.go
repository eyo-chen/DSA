package goodnode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GoodNodes(root *TreeNode) int {
	return helper(root, root.Val)
}

// Time complexity: O(n)
// Space complexity: O(h)
func helper(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	// check if the current node is a good node
	// if the current node's value is greater than or equal to maxVal,
	// it means it is a good node
	// we do two things here:
	// 1. update the current max value (for the next recursive call)
	// 2. set the isGoodNode to 1 (for the return value)
	curMaxVal := maxVal
	isGoodNode := 0
	if node.Val >= maxVal {
		curMaxVal = node.Val
		isGoodNode = 1
	}

	leftCount := helper(node.Left, curMaxVal)
	rightCount := helper(node.Right, curMaxVal)

	return leftCount + rightCount + isGoodNode
}

// nodeInfo is used to store the node and the target value
type nodeInfo struct {
	node   *TreeNode
	target int
}

// BFS Approach
func GoodNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*nodeInfo{{node: root, target: root.Val}}
	ans := 0

	for len(queue) > 0 {
		info := queue[0]
		queue = queue[1:]

		node, target := info.node, info.target

		// check if the current node is a good node
		// if the current node's value is greater than or equal to target,
		// it means it is a good node
		if node.Val >= target {
			ans++
		}

		// update the target value for the child nodes
		// the new target value is the maximum value between the current node's value and the target value
		newTarget := max(node.Val, target)

		if node.Left != nil {
			queue = append(queue, &nodeInfo{node: node.Left, target: newTarget})
		}
		if node.Right != nil {
			queue = append(queue, &nodeInfo{node: node.Right, target: newTarget})
		}
	}

	return ans
}
