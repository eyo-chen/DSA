package main

// CountComponents counts the number of connected components in an undirected graph.
// Approach:
// 1. Build an adjacency list representation of the undirected graph
// 2. Use DFS to traverse each unvisited node and mark all nodes in its component as visited
// 3. Each DFS traversal represents one connected component
// 4. Return the total count of connected components found
//
// Time Complexity: O(V + E) where V is number of vertices and E is number of edges
// Space Complexity: O(V + E) for adjacency list and O(V) for visited array
func CountComponents(n int, edges [][]int) int {
	// Build adjacency list to represent the undirected graph
	adjacencyList := map[int][]int{}
	visitedNodes := make([]bool, n)

	// Process each edge to build bidirectional connections
	for _, edge := range edges {
		nodeA, nodeB := edge[0], edge[1]
		adjacencyList[nodeA] = append(adjacencyList[nodeA], nodeB)
		adjacencyList[nodeB] = append(adjacencyList[nodeB], nodeA) // Essential for undirected graph
	}

	componentCount := 0

	// Iterate through all nodes to find unvisited components
	for nodeIndex := range n {
		// Skip nodes that are already part of a discovered component
		if visitedNodes[nodeIndex] {
			continue
		}

		// Found a new component - increment counter and explore it
		componentCount++
		exploreComponent(adjacencyList, visitedNodes, nodeIndex)
	}

	return componentCount
}

// exploreComponent performs depth-first search to mark all nodes in a connected component as visited.
// This ensures that all nodes reachable from the starting node are marked, representing one complete component.
func exploreComponent(adjacencyList map[int][]int, visitedNodes []bool, currentNode int) {
	// Base case: if node is already visited, stop recursion
	if visitedNodes[currentNode] {
		return
	}

	// Mark current node as visited to avoid revisiting
	visitedNodes[currentNode] = true

	// Recursively visit all unvisited neighbors
	for _, neighborNode := range adjacencyList[currentNode] {
		exploreComponent(adjacencyList, visitedNodes, neighborNode)
	}
}
