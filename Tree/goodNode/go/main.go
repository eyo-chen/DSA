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
