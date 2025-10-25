package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree constructs a binary tree from preorder and inorder traversal arrays.
//
// Approach:
// - The first element in preorder is always the root
// - Find the root in inorder to divide left and right subtrees
// - Recursively build left subtree with elements before root in inorder
// - Recursively build right subtree with elements after root in inorder
// - This version creates new slices on each recursion (less efficient)
//
// Time Complexity: O(n²) - O(n) for finding root in inorder * O(n) recursive calls
// Space Complexity: O(n²) - Creates new slices at each level
func BuildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeRecursive(preorder, inorder, 0)
}

// buildTreeRecursive is a helper that builds the tree using slice operations.
// preorderIndex tracks the current position in the preorder array.
func buildTreeRecursive(preorder []int, inorder []int, preorderIndex int) *TreeNode {
	// Base case: empty arrays or invalid index
	if len(preorder) == 0 || len(inorder) == 0 || preorderIndex >= len(preorder) {
		return nil
	}

	// The current preorder element is the root of this subtree
	rootVal := preorder[preorderIndex]
	root := &TreeNode{Val: rootVal}

	// Find the root's position in inorder array
	// Elements to the left are in left subtree, elements to the right are in right subtree
	rootIndexInInorder := -1
	for i, val := range inorder {
		if val == rootVal {
			rootIndexInInorder = i
			break
		}
	}

	// Build left subtree with elements before root in inorder
	// Next preorder element (preorderIndex + 1) is the left subtree's root
	root.Left = buildTreeRecursive(preorder, inorder[:rootIndexInInorder], preorderIndex+1)

	// Build right subtree with elements after root in inorder
	// Skip past all left subtree elements in preorder (preorderIndex + 1 + rootIndexInInorder)
	root.Right = buildTreeRecursive(preorder, inorder[rootIndexInInorder+1:], preorderIndex+1+rootIndexInInorder)

	return root
}

// BuildTreeOptimized constructs a binary tree from preorder and inorder traversal arrays.
//
// Approach:
// - Use a hash map to store inorder indices for O(1) lookup
// - Pass array bounds instead of creating new slices
// - Track the current preorder index and inorder left/right bounds
// - Recursively build left and right subtrees by adjusting bounds
//
// Time Complexity: O(n) - Visit each node once, O(1) lookup for root position
// Space Complexity: O(n) - Hash map storage + O(h) recursion stack where h is tree height
func BuildTreeOptimized(preorder []int, inorder []int) *TreeNode {
	// Create hash map for O(1) lookup of any value's index in inorder
	inorderIndexMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}

	return buildTreeOptimizedRecursive(preorder, inorderIndexMap, 0, 0, len(inorder)-1)
}

// buildTreeOptimizedRecursive builds the tree using index boundaries instead of slicing.
// preorderIndex: current position in preorder array
// inorderLeft: left boundary of current subtree in inorder array
// inorderRight: right boundary of current subtree in inorder array
func buildTreeOptimizedRecursive(preorder []int, inorderIndexMap map[int]int, preorderIndex int, inorderLeft int, inorderRight int) *TreeNode {
	// Base case: invalid range means no subtree exists
	if inorderLeft > inorderRight {
		return nil
	}

	// Current preorder element is the root
	rootVal := preorder[preorderIndex]
	root := &TreeNode{Val: rootVal}

	// O(1) lookup of root's position in inorder
	rootIndexInInorder := inorderIndexMap[rootVal]

	// Calculate the size of left subtree
	leftSubtreeSize := rootIndexInInorder - inorderLeft

	// Build left subtree: next element in preorder, constrain inorder to left of root
	root.Left = buildTreeOptimizedRecursive(preorder, inorderIndexMap, preorderIndex+1, inorderLeft, rootIndexInInorder-1)

	// Build right subtree: skip past left subtree in preorder, constrain inorder to right of root
	root.Right = buildTreeOptimizedRecursive(preorder, inorderIndexMap, preorderIndex+1+leftSubtreeSize, rootIndexInInorder+1, inorderRight)

	return root
}
