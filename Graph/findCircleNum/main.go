package main

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	// Visited array to keep track of the cities that have been visited
	visited := make([]bool, n)

	// Counter for the number of provinces
	provinces := 0

	// Iterate through each city
	// The number of cities is the length of the `isConnected` matrix
	for city := 0; city < n; city++ {
		// If the city hasn't been visited, it's a new province
		// We use BFS to mark all the cities that are connected to the current city as visited
		// Then, we increment the `provinces` counter
		if !visited[city] {
			bfs(isConnected, city, visited)
			provinces++
		}
	}

	return provinces
}

func bfs(isConnected [][]int, city int, visited []bool) {
	// Queue to perform BFS
	queue := []int{city}

	// Mark the current city as visited
	visited[city] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Loop through all the cities
		for neighbor := 0; neighbor < len(isConnected); neighbor++ {
			// If there is a connection between the current city and the city, and the city hasn't been visited, we add it to the queue and mark it as visited
			// Note that `isConnected[curr][neighbor] == 1` means there is a direct connection between city `curr` and city `neighbor`
			// We don't need to check `isConnected[neighbor][curr] == 1` because the matrix is symmetric
			// Which means that if `isConnected[curr][neighbor] == 1`, then `isConnected[neighbor][curr] == 1`
			// So we only need to check one of them
			if isConnected[curr][neighbor] == 1 && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func FindCircleNum2(isConnected [][]int) int {
	ans := 0
	visited := make([]bool, len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		if visited[i] {
			continue
		}

		DFS(isConnected, visited, i)
		ans++
	}

	return ans
}

func DFS(isConnected [][]int, visited []bool, node int) {
	if visited[node] {
		return
	}

	visited[node] = true
	for i := 0; i < len(isConnected[node]); i++ {
		if node == i || isConnected[node][i] == 0 {
			continue
		}

		DFS(isConnected, visited, i)
	}
}

func DFSStack(isConnected [][]int, visited []bool, node int) {
	stack := []int{node}

	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[curNode] {
			continue
		}

		visited[curNode] = true
		for i := 0; i < len(isConnected[curNode]); i++ {
			if i == curNode || isConnected[curNode][i] == 0 {
				continue
			}

			stack = append(stack, i)
		}
	}
}

func BFS(isConnected [][]int, visited []bool, node int) {
	queue := []int{node}

	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		if visited[curNode] {
			continue
		}

		visited[curNode] = true

		for i := 0; i < len(isConnected[curNode]); i++ {
			if curNode == i || isConnected[curNode][i] == 0 {
				continue
			}

			queue = append(queue, i)
		}
	}
}
