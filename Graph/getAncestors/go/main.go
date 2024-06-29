package main

func GetAncestors(n int, edges [][]int) [][]int {
	adj := map[int][]int{}
	result := make([][]int, n)

	// convert edges to adjacency list
	for _, e := range edges {
		from, to := e[0], e[1]
		adj[from] = append(adj[from], to)
	}

	// loop through each node to find ancestors, treat each node as root node
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		helper(adj, result, visited, i, i)
	}

	return result
}

func helper(adj map[int][]int, result [][]int, visited []bool, cur int, parent int) {
	visited[cur] = true
	for _, node := range adj[cur] {
		if visited[node] {
			continue
		}

		result[node] = append(result[node], parent)
		helper(adj, result, visited, node, parent)
	}
}
