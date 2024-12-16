package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(h)
// In the worst case(skewed tree), the time complexity is O(n)
// In the best case(balanced tree), the time complexity is O(logn)
// Space Complexity: O(h)
// In the worst case(skewed tree), the space complexity is O(n)
// In the best case(balanced tree), the space complexity is O(logn)
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return SearchBST(root.Left, val)
	}

	return SearchBST(root.Right, val)
}

// Time Complexity: O(h)
// In the worst case(skewed tree), the time complexity is O(n)
// In the best case(balanced tree), the time complexity is O(logn)
// Space Complexity: O(1)
func SearchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}
