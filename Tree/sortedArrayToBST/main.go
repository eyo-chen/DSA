package sortedarraytobst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST converts a sorted array into a height-balanced binary search tree.
//
// Approach: Use a divide-and-conquer strategy by recursively selecting the middle
// element as the root to ensure the tree remains balanced. The left half of the array
// forms the left subtree, and the right half forms the right subtree.
//
// Time Complexity: O(n) - we visit each element exactly once to create a node
// Space Complexity: O(log n) - recursion stack depth for a balanced tree
func SortedArrayToBST(nums []int) *TreeNode {
	return buildBalancedBST(nums, 0, len(nums)-1)
}

// buildBalancedBST recursively constructs a balanced BST from a subarray.
// It takes the array and the current subarray boundaries (start and end indices).
//
// Time Complexity: O(n) where n is the number of elements in the subarray
// Space Complexity: O(log n) for the recursion call stack
func buildBalancedBST(nums []int, start int, end int) *TreeNode {
	// Base case: empty subarray means no node to create
	if start > end {
		return nil
	}

	// Find the middle index to ensure balanced partitioning
	// Using start + (end - start) / 2 to avoid potential integer overflow
	mid := start + (end-start)/2

	// Create the root node with the middle element
	root := &TreeNode{Val: nums[mid]}

	// Recursively build the left subtree from elements before mid
	root.Left = buildBalancedBST(nums, start, mid-1)

	// Recursively build the right subtree from elements after mid
	root.Right = buildBalancedBST(nums, mid+1, end)

	return root
}
