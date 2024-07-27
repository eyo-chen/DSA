package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// using BFS
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// create the new root node
	rootNode := &Node{
		Val: node.Val,
	}

	// create a queue to traverse the graph
	// note that the queue contains the original nodes
	q := []*Node{node}

	// create a hashmap to store the new nodes
	// it helps to reference the new nodes(pointer) to avoid creating duplicate nodes
	hashMap := map[int]*Node{
		node.Val: rootNode,
	}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		// get the new node from the hashmap
		// note that the queue contains the original nodes
		// hashMap contains the new copy nodes
		// so we can get the new copy node from the hashMap
		copyNode := hashMap[n.Val]

		// iterate over the neighbors of the original node
		for _, ne := range n.Neighbors {
			var neighborNode *Node
			var ok bool
			neighborNode, ok = hashMap[ne.Val]

			// if the neighbor node is not in the hashmap, which means it is not visited yet
			// then
			// 1. create a new node
			// 2. add the new node to the hashmap
			// 3. add the new node to the queue
			if !ok {
				neighborNode = &Node{
					Val: ne.Val,
				}
				hashMap[ne.Val] = neighborNode

				q = append(q, ne)
			}

			// add the neighbor node to the neighbors of the copy node
			copyNode.Neighbors = append(copyNode.Neighbors, neighborNode)
		}
	}

	return rootNode
}

// using BFS (using pointer as hashmap key)
func CloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}

	// create the new root node
	rootNode := &Node{
		Val: node.Val,
	}

	// create a queue to traverse the graph
	// note that the queue contains the original nodes
	q := []*Node{node}

	// create a hashmap to store the new nodes
	// it helps to reference the new nodes(pointer) to avoid creating duplicate nodes
	// note that the key is the pointer of the original node
	hashMap := map[*Node]*Node{
		node: rootNode,
	}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		// get the new node from the hashmap
		// note that the queue contains the original nodes
		// hashMap contains the new copy nodes
		// so we can get the new copy node from the hashMap
		copyNode := hashMap[n]

		// iterate over the neighbors of the original node
		for _, ne := range n.Neighbors {
			var neighborNode *Node
			var ok bool
			neighborNode, ok = hashMap[ne]

			// if the neighbor node is not in the hashmap, which means it is not visited yet
			// then
			// 1. create a new node
			// 2. add the new node to the hashmap
			// 3. add the new node to the queue
			if !ok {
				neighborNode = &Node{
					Val: ne.Val,
				}
				hashMap[ne] = neighborNode

				q = append(q, ne)
			}

			// add the neighbor node to the neighbors of the copy node
			copyNode.Neighbors = append(copyNode.Neighbors, neighborNode)
		}
	}

	return rootNode
}

// using DFS
func CloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}

	// create the new root node
	rootNode := &Node{
		Val: node.Val,
	}

	// create a hashmap to store the new nodes
	// it helps to reference the new nodes(pointer) to avoid creating duplicate nodes
	hashMap := map[int]*Node{
		node.Val: rootNode,
	}

	dfs(node, hashMap)

	return rootNode
}

func dfs(node *Node, hashMap map[int]*Node) {
	if node == nil {
		return
	}

	// get the new node from the hashmap
	// note that the parameter node is the original node
	// hashMap contains the new copy nodes
	// so we can get the new copy node from the hashMap
	copyNode := hashMap[node.Val]

	// iterate over the neighbors of the original node
	for _, ne := range node.Neighbors {
		var neighborNode *Node
		var ok bool
		neighborNode, ok = hashMap[ne.Val]

		// if the neighbor node is not in the hashmap, which means it is not visited yet
		// then
		// 1. create a new node
		// 2. add the new node to the hashmap
		// 3. call the dfs function recursively with the neighbor node(original node)
		if !ok {
			neighborNode = &Node{
				Val: ne.Val,
			}
			hashMap[ne.Val] = neighborNode
			dfs(ne, hashMap)
		}

		// add the neighbor node to the neighbors of the copy node
		copyNode.Neighbors = append(copyNode.Neighbors, neighborNode)
	}
}
