# Problem Explanation

The problem is essentially asking if a graph has a cycle or not. If a cycle is present, then it is not possible to finish all courses.<br>
Consider the following example:<br>
input: numCourses = 2, prerequisites = [[1,0],[0,1]]<br>
output: false<br>
```
0 -> 1
1 -> 0
```
As we can see, there is a cycle present in the graph. Hence, it is not possible to finish all courses.<br>

So, now the problem is how to detect a cycle in a graph.<br>
***A path has a cycle when we visit a node that is already visited in the current path.***<br>
For example, when we explore 0 -> 1, we've visited 0 and 1.<br>
Then we visit 1 -> 0, when we hit 0, we know that there is a cycle.<br>
Because 0 has been visited in the current path.<br>

Therefore, we simply do the DFS on each node.<br>
Along with the process, we also maintain the state of nodes.<br>
The state of a node can be:<br>
- UnVisited
  - It means this node is allowed to visit.
- Visiting
  - It means this node is currently being visited.
  - When we visit a node with state is Visiting, it means there is a cycle.
- Visited
  - It means this node has been visited.
  - It also means that there's no cycle on this node.

DFS is a great tool to sovle this problem<br>
Because we're essentially exploring every path in the graph.<br>
We're not searching the region or finding the shortest path.<br>

Let's summarize the steps:<br>
1. Create a adjacency list from the input.
2. Create a state list to maintain the state of each node.
   - Initialize the state list with UnVisited.
3. Do the DFS on each node.(simple for-loop)
   - If the node is NOT UnVisited, skip it. (It means we've visited this node, and there's no cycle.)
   - Do the DFS on this code
4. In the DFS function:
   - If the node is Visiting, return True.
     - We visited a node that is already visited in the current path, which means there's a cycle.
   - If the node is Visited, return False.
     - It means we've visited this node, and there's no cycle.
   - Set the state of the node to Visiting.
   - Do the DFS on the neighbors.
   - Set the state of the node to Visited.


Let's walk through the example:<br>
numCourses = 2, prerequisites = [[1,0],[0,1]]<br>
1. Create a adjacency list from the input.
   ```
   0 -> 1
   1 -> 0
   ```
2. Create a state list to maintain the state of each node.
    - Initialize the state list with UnVisited.
      ```
      [UnVisited, UnVisited]
      ```
3. Do the DFS on each node.
    - DFS(0)
      - Set the state of 0 to Visiting.
      - DFS(1)
        - Set the state of 1 to Visiting.
        - DFS(0)
          - 0 is Visiting, return True.
          - Bubble up to the parent recursive call.

# Complexity Analysis
## Time Complexity O(V + E)

1. **Building the Adjacency List:**
   - We iterate over the `prerequisites` array once, which has \(E\) elements, where \(E\) is the number of edges (prerequisites).
   - Each iteration involves appending to a map, which is \(O(1)\) on average.
   - Therefore, constructing the adjacency list takes \(O(E)\) time.

2. **Cycle Detection:**
   - We iterate over each course (node) once, which is \(O(V)\), where \(V\) is the number of vertices (courses).
   - For each course, if it hasn't been visited yet, we perform a DFS to check for cycles.
   - In the worst case, the DFS will visit each node and edge once.
   - The DFS for each node \(i\) runs in \(O(V + E)\) time in total because each node and edge is visited exactly once.

Therefore, the overall time complexity is:
\[ O(E) + O(V + E) = O(V + E) \]

## Space Complexity O(V + E)

1. **Adjacency List:**
   - We store each edge in the adjacency list, which requires \(O(E)\) space.

2. **State Array:**
   - We use an array `states` of size \(V\) to keep track of the state of each node, which requires \(O(V)\) space.

3. **Call Stack:**
   - The depth of the recursion stack for DFS can go up to \(V\) in the worst case.

Therefore, the overall space complexity is:
\[ O(V + E) \]