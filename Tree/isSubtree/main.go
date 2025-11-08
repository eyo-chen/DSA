package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSubtree checks if subRoot is a subtree of root using recursive tree traversal.
// Approach: For each node in root, check if the tree rooted at that node is identical to subRoot.
// Time Complexity: O(m * n) where m is the number of nodes in root and n is the number of nodes in subRoot.
//
//	In the worst case, we check tree equality at every node.
//
// Space Complexity: O(h) where h is the height of root, due to recursion stack.
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Base case: if root is nil, subRoot cannot be a subtree
	if root == nil {
		return false
	}

	// Check if the tree rooted at current node matches subRoot exactly
	if areIdenticalTrees(root, subRoot) {
		return true
	}

	// Recursively check left and right subtrees
	// Return true if subRoot is found in either subtree
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

// areIdenticalTrees checks if two trees are structurally identical with same values.
// Time Complexity: O(min(m, n)) where m and n are the number of nodes in each tree.
// Space Complexity: O(min(h1, h2)) where h1 and h2 are the heights, due to recursion stack.
func areIdenticalTrees(tree1 *TreeNode, tree2 *TreeNode) bool {
	// Both trees are empty - they are identical
	if tree1 == nil && tree2 == nil {
		return true
	}

	// One tree is empty but the other is not - they are different
	if tree1 == nil || tree2 == nil {
		return false
	}

	// Current nodes have different values - trees are different
	if tree1.Val != tree2.Val {
		return false
	}

	// Recursively check if left subtrees and right subtrees are identical
	return areIdenticalTrees(tree1.Left, tree2.Left) &&
		areIdenticalTrees(tree1.Right, tree2.Right)
}

// IsSubtree1 checks if subRoot is a subtree of root using tree serialization.
// Approach: Serialize both trees to strings and check if subRoot's serialization is a substring of root's serialization.
//
// Time Complexity: O(m + n) where m is the number of nodes in root and n in subRoot.
//
//	Serialization is O(m + n) and substring search is O(m + n).
//
// Space Complexity: O(m + n) for storing the serialized strings.
func IsSubtree1(root *TreeNode, subRoot *TreeNode) bool {
	var rootSerialized, subRootSerialized strings.Builder

	// Serialize both trees to string representations
	serializeTree(root, &rootSerialized)
	serializeTree(subRoot, &subRootSerialized)

	// Check if subRoot's serialization appears in root's serialization
	return strings.Contains(rootSerialized.String(), subRootSerialized.String())
}

// serializeTree converts a binary tree to a unique string representation.
// Uses pre-order traversal with delimiters to ensure uniqueness.
func serializeTree(node *TreeNode, builder *strings.Builder) {
	// Represent nil nodes explicitly to maintain tree structure
	if node == nil {
		builder.WriteString(",nil")
		return
	}

	// Add delimiter before the value to prevent false matches
	// Example: Without leading delimiter, tree [12] would serialize as "12,nil,nil"
	//          and tree [2] as "2,nil,nil", causing [12] to incorrectly contain [2]
	// With delimiter: ",12,nil,nil" and ",2,nil,nil" - no false match
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(node.Val))

	// Recursively serialize left and right subtrees (pre-order traversal)
	serializeTree(node.Left, builder)
	serializeTree(node.Right, builder)
}
