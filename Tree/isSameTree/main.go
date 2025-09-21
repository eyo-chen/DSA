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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// node is used to store the pair of nodes to be compared
type node struct {
	p *TreeNode
	q *TreeNode
}

func IsSameTree1(p *TreeNode, q *TreeNode) bool {
	// use BFS to compare the two trees
	queue := []*node{{p: p, q: q}}

	for len(queue) > 0 {
		// get the first pair of nodes from the queue
		p, q := queue[0].p, queue[0].q
		queue = queue[1:]

		// if both are nil, it means they are same, continue to the next pair
		if p == nil && q == nil {
			continue
		}

		// if one of them is nil, it means they are not same
		if p == nil || q == nil {
			return false
		}

		// if the value is not same, it means they are not same
		if p.Val != q.Val {
			return false
		}

		queue = append(queue, &node{p: p.Left, q: q.Left}, &node{p: p.Right, q: q.Right})
	}

	return true
}
