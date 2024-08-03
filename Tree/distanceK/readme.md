# Problem Explanation

Even though it's a tree problem, we can't solve it with the usual tree traversal methods.<br>
Because we can't access the parent node from the child node.<br>
So, we need to convert the tree into a graph and then solve the problem.<br>

Once we convert the tree into a graph, we can use the BFS algorithm to find the nodes at distance K from the target node.<br>

Let's summarize the solution:<br>
1. Convert the tree into a graph (adjacency list)
2. BFS from the target node
3. Return the current level nodes when the distance is K

Note two things:<br>
1. We need to convert to undirected graph
2. We need to keep track of the visited nodes
   - because it's undirected graph, we have to avoid going back to the explored nodes
3. When building the graph(genAdj), the logic is a little bit verbose
   - when we at a node, we need to do two things
     - add the child to the parent's adjacency list
     - add the parent to the child's adjacency list

For example, let's say we have a tree like this:<br>
```
    0
   / \
  1   2
     / \
    3   4
```
The adjacency list will look like this:<br>
```
{
  0: [1, 2],
  1: [0],
  2: [0, 3, 4],
  3: [2],
  4: [2]
}
```
When we're at node 2, we need to add 0, 3 and 4 to the adjacency list of 2.<br>
Also, we need to add 2 to the adjacency list of 0, 3 and 4.<br>

# Complexity Analysis
## Time Complexity O(n)
- We need to traverse the tree to build the graph
- We need to traverse the graph to find the nodes at distance K

## Space Complexity O(n)
- We need to store the adjacency list
- We need to store the visited nodes
- We need to store the nodes when BFS