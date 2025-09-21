package kthsmallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// KthSmallest finds the kth smallest element in a BST using in-order traversal.
// Approach: Perform complete in-order traversal to collect all values in sorted order,
// then return the element at index k-1.
// Time Complexity: O(n) - visits all nodes
// Space Complexity: O(n) - stores all values in array
func KthSmallest(root *TreeNode, k int) int {
	sortedValues := []int{}
	inOrderTraversal(root, &sortedValues)
	return sortedValues[k-1]
}

// inOrderTraversal performs in-order traversal of BST and collects values in sorted order
func inOrderTraversal(currentNode *TreeNode, sortedValues *[]int) {
	if currentNode == nil {
		return
	}

	// Traverse left subtree first (smaller values)
	inOrderTraversal(currentNode.Left, sortedValues)

	// Process current node (add to sorted array)
	*sortedValues = append(*sortedValues, currentNode.Val)

	// Traverse right subtree (larger values)
	inOrderTraversal(currentNode.Right, sortedValues)
}

// KthSmallestOptimized finds the kth smallest element in a BST using early termination.
// Approach: Perform in-order traversal but stop as soon as we find the kth element.
// This avoids visiting all nodes when k is small.
// Time Complexity: O(h + k) where h is height of tree
// Space Complexity: O(h) - recursion stack space only
func KthSmallestOptimized(root *TreeNode, k int) int {
	return findKthSmallest(root, &k)
}

// findKthSmallest performs in-order traversal with early termination when kth element is found
func findKthSmallest(currentNode *TreeNode, remainingCount *int) int {
	if currentNode == nil {
		return -1 // Indicates not found in this subtree
	}

	// Search in left subtree first (smaller values)
	leftResult := findKthSmallest(currentNode.Left, remainingCount)
	if leftResult != -1 {
		return leftResult // Found in left subtree
	}

	// Process current node - decrement counter
	*remainingCount = *remainingCount - 1
	if *remainingCount == 0 {
		return currentNode.Val // Found the kth smallest element
	}

	// Search in right subtree (larger values)
	rightResult := findKthSmallest(currentNode.Right, remainingCount)
	if rightResult != -1 {
		return rightResult // Found in right subtree
	}

	return -1 // Not found in this subtree
}
