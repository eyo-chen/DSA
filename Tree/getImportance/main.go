package main

// Employee represents the employee data structure as specified in the problem
type Employee struct {
	Id           int   // The unique ID of the employee
	Importance   int   // The importance value of the employee
	Subordinates []int // List of direct subordinate IDs
}

// getImportance1 solves the problem using a graph-based approach with adjacency lists.
// This approach treats the problem as a graph traversal where we separate the structure
// (adjacency list) from the node values (importance map).
// Time Complexity: O(n) where n is the number of employees
// Space Complexity: O(n) for the maps plus O(h) for recursion stack where h is hierarchy height
func GetImportance1(employees []*Employee, id int) int {
	// Create adjacency list to represent the reporting structure
	// This maps each employee ID to their list of direct subordinates
	adjacencyList := map[int][]int{}

	// Create importance lookup map for O(1) access to each employee's value
	// This separates the graph structure from the node values
	importanceMap := map[int]int{}

	// Build both maps in a single pass through the employee data
	for _, employee := range employees {
		adjacencyList[employee.Id] = employee.Subordinates
		importanceMap[employee.Id] = employee.Importance
	}

	// Delegate to recursive helper that traverses the hierarchy
	return calculateTotalImportance1(adjacencyList, importanceMap, id)
}

// calculateTotalImportance1 recursively calculates the sum of importance values
// for an employee and all their direct and indirect subordinates
func calculateTotalImportance1(adjacencyList map[int][]int, importanceMap map[int]int, employeeId int) int {
	// Start with current employee's importance value
	totalSum := importanceMap[employeeId]

	// Get list of direct subordinates for this employee
	subordinateList := adjacencyList[employeeId]

	// Recursively add importance of each subordinate and their teams
	// This naturally handles the case where subordinateList is empty (leaf nodes)
	for _, subordinateId := range subordinateList {
		totalSum += calculateTotalImportance1(adjacencyList, importanceMap, subordinateId)
	}

	return totalSum
}

// getImportance2 solves the problem using Depth-First Search with recursion.
// This approach directly traverses the employee hierarchy using the original data structure,
// treating each employee as a node in a tree and recursively visiting all children.
// Time Complexity: O(n) where n is the number of employees
// Space Complexity: O(n) for the lookup map plus O(h) for recursion stack
func GetImportance2(employees []*Employee, id int) int {
	// Build lookup map for O(1) employee access by ID
	// Without this optimization, we'd have O(n²) complexity from repeated linear searches
	employeeLookup := make(map[int]*Employee)
	for _, employee := range employees {
		employeeLookup[employee.Id] = employee
	}

	// Start recursive DFS traversal from the target employee
	return dfsTraversal(employeeLookup, id)
}

// dfsTraversal recursively visits an employee and all their subordinates using DFS
// This follows the classic recursive pattern: process current node, then recurse on children
func dfsTraversal(employeeLookup map[int]*Employee, employeeId int) int {
	// Retrieve the current employee from our lookup map
	currentEmployee := employeeLookup[employeeId]

	// Initialize total with current employee's importance
	totalImportance := currentEmployee.Importance

	// Recursively process each direct subordinate
	// The recursive calls will handle all indirect subordinates automatically
	for _, subordinateId := range currentEmployee.Subordinates {
		totalImportance += dfsTraversal(employeeLookup, subordinateId)
	}

	return totalImportance
}

// getImportance3 solves the problem using Breadth-First Search with an iterative approach.
// This approach processes the hierarchy level by level using a queue, visiting all employees
// at depth 1, then depth 2, etc. While the order is different from DFS, the final sum is identical.
// Time Complexity: O(n) where n is the number of employees
// Space Complexity: O(n) for the lookup map plus O(w) for the queue where w is max width of hierarchy
func GetImportance3(employees []*Employee, id int) int {
	// Create lookup map for efficient employee retrieval by ID
	employeeLookup := make(map[int]*Employee)
	for _, employee := range employees {
		employeeLookup[employee.Id] = employee
	}

	// Initialize BFS with a queue containing the starting employee ID
	// We'll process employees level by level, adding subordinates to the queue
	processingQueue := []int{id}
	totalImportance := 0

	// Continue processing until queue is empty (all employees in hierarchy visited)
	for len(processingQueue) > 0 {
		// Dequeue the next employee to process (FIFO - First In, First Out)
		currentEmployeeId := processingQueue[0]
		processingQueue = processingQueue[1:]

		// Look up the employee object and add their importance to running total
		currentEmployee := employeeLookup[currentEmployeeId]
		totalImportance += currentEmployee.Importance

		// Enqueue all direct subordinates for future processing
		// This ensures we'll eventually visit every employee in the subtree
		processingQueue = append(processingQueue, currentEmployee.Subordinates...)
		// for _, subordinateId := range currentEmployee.Subordinates {
		// 	processingQueue = append(processingQueue, subordinateId)
		// }
	}

	return totalImportance
}
