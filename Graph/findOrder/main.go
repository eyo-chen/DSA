package main

// Solution 1: Use BFS to find the topological sort
func FindOrder(numCourses int, prerequisites [][]int) []int {
	// adjacency list is to help us to represent the graph
	adj := make(map[int][]int, numCourses)

	// dependency is to help us to count the number of dependencies for each course
	dependency := make([]int, numCourses)

	// ans is to help us to store the result
	ans := make([]int, 0, numCourses)

	// build the adjacency list and the dependency count for each course
	for _, preq := range prerequisites {
		cur, pre := preq[0], preq[1]
		adj[pre] = append(adj[pre], cur)
		dependency[cur]++
	}

	// queue is to help us to traverse the graph
	queue := []int{}

	// find all the courses that have no dependencies
	// we can only start processing the courses that have no dependencies
	for node, dep := range dependency {
		if dep == 0 {
			queue = append(queue, node)
		}
	}

	// traverse the graph using BFS
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// add the course to the result list
		ans = append(ans, node)

		// traverse the neighbors of the current course
		for _, neighbor := range adj[node] {
			// decrease the dependency count for the neighbor
			dependency[neighbor]--
			// if the dependency count is 0, add the neighbor to the queue
			if dependency[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// if the length of the result list is equal to the number of courses, return the result list
	if len(ans) == numCourses {
		return ans
	}

	// if the length of the result list is not equal to the number of courses, return an empty list
	return []int{}
}

// State is to help us to check if the course has been visited
type State int

const (
	NoVisit State = iota
	Visiting
	Visited
)

// Solution 2: Use DFS to find the topological sort
func FindOrder2(numCourses int, prerequisites [][]int) []int {
	// adjacency list is to help us to represent the graph
	adj := make(map[int][]int, numCourses)

	// state is to help us to check if the course has been visited, visiting or visited
	state := make([]State, numCourses)

	// ans is to help us to store the result
	ans := make([]int, 0, numCourses)

	// build the adjacency list
	for _, preq := range prerequisites {
		cur, pre := preq[0], preq[1]
		adj[pre] = append(adj[pre], cur)
	}

	// traverse the graph using DFS
	for i := 0; i < numCourses; i++ {
		if hasCycle(adj, state, &ans, i) {
			return []int{}
		}
	}

	// reverse the answer list
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}

// hasCycle is to help us to check if there's a cycle in the graph
func hasCycle(adj map[int][]int, state []State, ans *[]int, node int) bool {
	if state[node] == Visited {
		return false
	}

	if state[node] == Visiting {
		return true
	}

	// set the state of the current course to visiting
	state[node] = Visiting

	// traverse the neighbors of the current course
	for _, neighbor := range adj[node] {
		if hasCycle(adj, state, ans, neighbor) {
			return true
		}
	}

	// add the current course to the result list after all the neighbors have been processed
	*ans = append(*ans, node)

	// set the state of the current course to visited
	state[node] = Visited

	// return false if there's no cycle
	return false
}
