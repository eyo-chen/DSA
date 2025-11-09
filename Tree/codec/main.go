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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// serialize converts a binary tree to a string using pre-order traversal.
// Approach: Traverse the tree in pre-order (root -> left -> right), converting each node
// to a string. Nil nodes are represented as "nil". Values are comma-separated.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) for the string builder and O(h) for recursion stack where h is tree height
func (c *Codec) Serialize(root *TreeNode) string {
	var result strings.Builder
	serializePreOrder(root, &result)
	return result.String()
}

// serializePreOrder performs pre-order traversal and appends node values to the string builder
func serializePreOrder(node *TreeNode, result *strings.Builder) {
	// Base case: if node is nil, append "nil," as a placeholder
	if node == nil {
		result.WriteString("nil,")
		return
	}

	// Append current node's value followed by a comma
	result.WriteString(strconv.Itoa(node.Val))
	result.WriteString(",")

	// Recursively serialize left subtree
	serializePreOrder(node.Left, result)

	// Recursively serialize right subtree
	serializePreOrder(node.Right, result)
}

// deserialize reconstructs a binary tree from a serialized string.
// Approach: Split the string by commas and reconstruct the tree using pre-order traversal.
// Each recursive call consumes one token from the array.
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(n) for the split array and O(h) for recursion stack
func (c *Codec) Deserialize(data string) *TreeNode {
	// Split the serialized string into tokens
	tokens := strings.Split(data, ",")
	index := 0
	return deserializePreOrder(tokens, &index)
}

// deserializePreOrder reconstructs the tree by consuming tokens in pre-order
func deserializePreOrder(tokens []string, index *int) *TreeNode {
	// Check bounds and handle nil nodes
	if *index >= len(tokens) || tokens[*index] == "nil" {
		*index++ // Move to next token
		return nil
	}

	// Parse current token to get node value
	val, _ := strconv.Atoi(tokens[*index])
	node := &TreeNode{Val: val}
	*index++ // Move to next token

	// Recursively build left subtree
	node.Left = deserializePreOrder(tokens, index)

	// Recursively build right subtree
	node.Right = deserializePreOrder(tokens, index)

	return node
}
