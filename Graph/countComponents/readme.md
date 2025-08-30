# Problem Statement (Number of Connected Components in an Undirected Graph)
There is an undirected graph with n nodes. There is also an edges array, where edges[i] = [a, b] means that there is an edge between node a and node b in the graph.<br>

The nodes are numbered from 0 to n - 1.<br>

Return the total number of connected components in that graph.<br>

Example 1:<br>
```
Input:
n=3
edges=[[0,1], [0,2]]
Output:
1
```


Example 2:<br>
```
Input:
n=6
edges=[[0,1], [1,2], [2,3], [4,5]]
Output:
2
```

Constraints:<br>
1 <= n <= 100<br>
0 <= edges.length <= n * (n - 1) / 2<br>

# Typical DFS
The idea to solve this problem is easy to come up with.<br>
We can use Depth First Search (DFS) to explore all the nodes in a connected component. We start from an unvisited node, mark it as visited, and then recursively visit all its neighbors. Each time we start a new DFS from an unvisited node, we have found a new connected component.<br>

Here's the summarize:<br>
1. Build an adjacency list from the edges.
2. Initialize a visited set to keep track of visited nodes.
3. Iterate through each node, and if it's unvisited, start a DFS from that node.
4. Each time a new DFS is started, increment the connected components count.

## Caveat
The only caveat of this problem is that the graph is undirected, meaning that if there is an edge from node a to node b, there is also an edge from node b to node a.<br>
That means we have to ***build two way connections in our adjacency list.***

For example, if the edge is [0,1]<br>
The adjacency list should include both directions:<br>
```
0: [1]
1: [0]
```
It's very important to remember this when building the adjacency list.

### Walk through example to see the issues
Let's walk through an example to see how it causes issues.
**Input:**
- `n = 3` (nodes: 0, 1, 2)
- `edges = [[1, 0], [1, 2]]`
- **Expected result:** 1 component (all nodes are connected through node 1)

#### Step 1: Build Adjacency List

The code with issue does:
```go
for _, edge := range edges {
    cur, next := edge[0], edge[1]
    adj[cur] = append(adj[cur], next)  // Only one direction!
}
```

Processing edges:
- Edge `[1, 0]`: `adj[1] = append(adj[1], 0)` → `adj[1] = [0]`
- Edge `[1, 2]`: `adj[1] = append(adj[1], 2)` → `adj[1] = [0, 2]`

**Final adjacency list:**
- `adj[0] = []` (empty - node 0 has no outgoing edges)
- `adj[1] = [0, 2]` (node 1 can reach 0 and 2)  
- `adj[2] = []` (empty - node 2 has no outgoing edges)

#### Step 2: DFS Traversal

Initialize: `visited = [false, false, false]`, `ans = 0`

**Loop iteration i = 0:**
- `visited[0] = false`, so continue with DFS
- `ans++` → `ans = 1`
- Call `dfs(adj, visited, 0)`

**DFS from node 0:**
- `visited[0] = true` → `visited = [true, false, false]`
- Loop through `adj[0]` which is `[]` (empty)
- **DFS ends - only visited node 0**

**Loop iteration i = 1:**  
- `visited[1] = false`, so continue with DFS
- `ans++` → `ans = 2`
- Call `dfs(adj, visited, 1)`

**DFS from node 1:**
- `visited[1] = true` → `visited = [true, true, false]`
- Loop through `adj[1] = [0, 2]`:
  - `dfs(adj, visited, 0)`: `visited[0] = true`, so return immediately
  - `dfs(adj, visited, 2)`: `visited[2] = false`, so continue:
    - `visited[2] = true` → `visited = [true, true, true]`
    - Loop through `adj[2]` which is `[]` (empty)
    - Return

**Loop iteration i = 2:**
- `visited[2] = true`, so continue (skip)

#### Final Result
**Your function returns: `ans = 2`** ❌
**Correct answer should be: `1`** ✅

#### Why It Failed
The problem is that when we start DFS from node 0, we can't reach nodes 1 and 2 because our adjacency list only has one-way connections. Node 0 has no outgoing edges in our representation, even though it's actually connected to node 1.

In an undirected graph, if there's an edge between nodes A and B, you should be able to traverse from A to B AND from B to A. Our current code only allows traversal in the direction the edge was listed in the input.

# Union Find (TBD)