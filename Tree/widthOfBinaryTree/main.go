package widthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Node *TreeNode
	Idx  int
}

func WidthOfBinaryTree(root *TreeNode) int {
	queue := []*NodeInfo{{root, 0}}
	res := 1

	for len(queue) > 0 {
		length := len(queue)
		minIdx, maxIdx := 0, 0

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == 0 {
				minIdx = node.Idx
			}
			if i == length-1 {
				maxIdx = node.Idx
			}

			if node.Node.Left != nil {
				queue = append(queue, &NodeInfo{node.Node.Left, node.Idx * 2})
			}
			if node.Node.Right != nil {
				queue = append(queue, &NodeInfo{node.Node.Right, node.Idx*2 + 1})
			}
		}

		if len(queue) > 0 {
			idx := queue[0].Idx
			for i := range queue {
				queue[i].Idx -= idx
			}
		}

		res = max(res, maxIdx-minIdx+1)
	}

	return res
}
