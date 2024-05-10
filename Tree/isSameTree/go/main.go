package issametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time complexity: O(n)
// Space complexity: O(h)
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// if both are nil, it means they are same
	if p == nil && q == nil {
		return true
	}

	// if one of them is nil, it means they are not same
	if p == nil || q == nil {
		return false
	}

	// if the value is not same, it means they are not same
	if p.Val != q.Val {
		return false
	}

	// have to make sure left and right are same
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
