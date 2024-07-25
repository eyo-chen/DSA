package main

type state uint8

const (
	stateUnVisited state = iota
	stateVisiting
	stateVisited
)

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Initialize the adjacency list
	adj := make(map[int][]int)
	for _, p := range prerequisites {
		cur, pre := p[0], p[1]
		adj[pre] = append(adj[pre], cur)
	}

	// Initialize the state of each node
	states := make([]state, numCourses)

	// Check for cycles in each component of the graph
	for i := 0; i < numCourses; i++ {
		if states[i] == stateVisited {
			continue
		}

		if hasCycle(adj, states, i) {
			return false
		}
	}

	return true
}

func hasCycle(adj map[int][]int, states []state, node int) bool {
	// Check if the node is being visited
	// If it is, then there is a cycle
	if states[node] == stateVisiting {
		return true
	}

	// Check if the node has been visited
	// If it has, then there is no cycle
	if states[node] == stateVisited {
		return false
	}

	// Mark the node as visiting
	states[node] = stateVisiting

	// Visit all adjacent nodes
	for _, neighbor := range adj[node] {
		if hasCycle(adj, states, neighbor) {
			return true
		}
	}

	// Mark the node as visited
	states[node] = stateVisited

	return false
}
