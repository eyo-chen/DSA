package buildtree1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree constructs a binary tree from inorder and postorder traversal arrays.
// Approach: Use postorder's last element as root, find it in inorder to split left/right subtrees.
// Recursively build right subtree first (since postorder processes right before root),
// then left subtree, tracking indices to avoid array slicing overhead.
func BuildTree(inorder []int, postorder []int) *TreeNode {
	// Create a map for O(1) lookup of any value's index in inorder array
	inorderIndexMap := map[int]int{}
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}

	// Start building from the last element of postorder (the root)
	return buildTreeRecursive(inorderIndexMap, postorder, len(postorder)-1, 0, len(postorder)-1)
}

// buildTreeRecursive constructs the tree using index boundaries instead of array slicing.
// Parameters:
//   - inorderIndexMap: maps values to their positions in the inorder array
//   - postorder: the postorder traversal array
//   - postorderRootIdx: current root's index in postorder array
//   - inorderLeft: left boundary of current subtree in inorder array
//   - inorderRight: right boundary of current subtree in inorder array
func buildTreeRecursive(inorderIndexMap map[int]int, postorder []int, postorderRootIdx int, inorderLeft int, inorderRight int) *TreeNode {
	// Base case: empty subtree
	if inorderLeft > inorderRight {
		return nil
	}

	// The current root is at postorderRootIdx in postorder array
	rootVal := postorder[postorderRootIdx]
	root := &TreeNode{Val: rootVal}

	// Find where this root splits the inorder array into left and right subtrees
	rootPosInInorder := inorderIndexMap[rootVal]

	// Calculate the size of the right subtree
	rightSubtreeSize := inorderRight - rootPosInInorder

	// Build right subtree first (postorder processes: left, right, root)
	// The right subtree's root is just before current root in postorder
	root.Right = buildTreeRecursive(inorderIndexMap, postorder, postorderRootIdx-1, rootPosInInorder+1, inorderRight)

	// Build left subtree
	// Skip past the right subtree elements to find left subtree's root in postorder
	leftSubtreeRootIdx := postorderRootIdx - rightSubtreeSize - 1
	root.Left = buildTreeRecursive(inorderIndexMap, postorder, leftSubtreeRootIdx, inorderLeft, rootPosInInorder-1)

	return root
}
