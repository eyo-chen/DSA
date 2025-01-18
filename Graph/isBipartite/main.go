package main

func IsBipartite(graph [][]int) bool {
	// group1 -> 1, group2 -> -1
	hashTable := make([]int, len(graph))

	for node := range graph {
		// Only start traversal if the node is not visited
		if hashTable[node] == 0 && !DFS(graph, hashTable, node, 1) {
			return false
		}
	}

	return true
}

func DFS(graph [][]int, hashTable []int, node int, expectedGroup int) bool {
	// If node is already visited, check if it's the expected group
	if hashTable[node] != 0 {
		return hashTable[node] == expectedGroup
	}

	// Mark the node as visited and set its group
	hashTable[node] = expectedGroup

	// Recursively visit all neighbors
	for _, neighbor := range graph[node] {
		if !DFS(graph, hashTable, neighbor, expectedGroup*-1) {
			return false
		}
	}

	return true
}

type nodeInfo struct {
	node          int
	expectedGroup int
}

func BFS(graph [][]int, hashTable []int, node int, expectedGroup int) bool {
	queue := []nodeInfo{{node: node, expectedGroup: expectedGroup}}

	// Mark the starting node immediately
	hashTable[node] = expectedGroup

	for len(queue) > 0 {
		ni := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[ni.node] {
			if hashTable[neighbor] == 0 {
				// Only add unvisited neighbors
				hashTable[neighbor] = ni.expectedGroup * -1
				queue = append(queue, nodeInfo{node: neighbor, expectedGroup: ni.expectedGroup * -1})
			} else if hashTable[neighbor] != ni.expectedGroup*-1 {
				// If neighbor is visited but has wrong color, graph is not bipartite
				return false
			}
		}
	}

	return true
}
