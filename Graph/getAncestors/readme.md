# Problem Explanation

edgeList = [[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]

The first intuition is that we depth first search on each node. When we reach a node has no children, we add the path to the result.<br>
This solution has one problem, which is it can't guarantee the order of the list of ancestors in the result<br>
For example, we know there are two paths to 5:
- 0 -> 3 -> 5
- 1 -> 3 -> 5

If we first add (0, 3) to the result of 5, then we explore the second path, we will add (1, 3) to the result of 5. Even though we can exclude the duplicate, the order of the list of ancestors is not guaranteed.<br>

The idea to solve this problem is that we first build the adjacency list of the graph<br>
Then we loop through all the nodes, we treat each node as the root of the tree, and do the depth first search on each node<br>
When traversing depth first search, we add the root node to the result of the current node<br>

For example, we have the following graph:
```
Graph:
0 -> 3, 4
1 -> 3
2 -> 4, 7
3 -> 5, 6, 7
4 -> 6
```
Result Set: [[], [], [], [], [], [], [], []]

Start from node 0,<br>
We hit 3, then we add 0 to the result of 3<br>
Result Set: [[], [], [], [0], [], [], [], []]<br>

We hit 5, then we add 0 to the result of 5<br>
Result Set: [[], [], [], [0], [], [0], [], []]<br>

We hit 6, then we add 0 to the result of 6<br>
Result Set: [[], [], [], [0], [], [0], [0], []]<br>

We hit 4, then we add 0 to the result of 4<br>
Result Set: [[], [], [], [0], [0], [0], [0], []]<br>

We hit 6 again, then we add 0 to the result of 6 (because there's path 0 -> 4 -> 6)<br>
Result Set: [[], [], [], [0], [0], [0], [0, 0], []]<br>
We have to prevent the duplicate, so we can use a set to store the visited node<br>

Start from node 1,<br>
We hit 3, then we add 1 to the result of 3<br>
We hit 5, then we add 1 to the result of 5<br>
Result Set: [[], [], [], [0, 1], [], [0, 1], [0], []]

Note two things
1. Along with the process of depth first search, we always only have one root node
   - When we start from node 0, we only add 0 to the result of 3 and 5
   - This guarantees the order of the list of ancestors in the result since we start from 0, then 1, then 2, ...
2. For each exploration, we have to prevent the duplicate, so we can use a set to store the visited node

Let's walk through the whole process<br>
edgeList = [[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]<br>
Result Set: [[], [], [], [], [], [], [], []]

Convert the edgeList to the adjacency list
```
Graph:
0 -> 3, 4
1 -> 3
2 -> 4, 7
3 -> 5, 6, 7
4 -> 6
```
Start from node 0, start the depth first search, and treat 0 as the root node<br>
visited = [F, F, F, F, F, F, F, F]<br>
First add 0 in visited, visited = [T, F, F, F, F, F, F, F]<br>

- 0 -> 3
  - add 3 as visited, visited = [T, F, F, T, F, F, F, F]
  - add 0 to the result of 3
  - Result Set: [[], [], [], [0], [], [], [], []]
  - 0 -> 3 -> 5
    - add 5 as visited, visited = [T, F, F, T, F, T, F, F]
    - add 0 to the result of 5
    - Result Set: [[], [], [], [0], [], [0], [], []]
  - 0 -> 3 -> 6
    - add 6 as visited, visited = [T, F, F, T, F, T, T, F]
    - add 0 to the result of 6
    - Result Set: [[], [], [], [0], [], [0], [0], []]
  - 0 -> 3 -> 7
    - add 7 as visited, visited = [T, F, F, T, F, T, T, T]
    - add 0 to the result of 7
    - Result Set: [[], [], [], [0], [], [0], [0], [0]]
- 0 -> 4
  - add 4 as visited, visited = [T, F, F, T, T, F, T, T]
  - add 0 to the result of 4
  - Result Set: [[], [], [], [0], [0], [0], [0], [0]]
  - 0 -> 4 -> 6
    - because 6 is visited, we don't need to explore 6 again

We're done with node 0, we can move to node 1<br>

Start from node 1, start the depth first search, and treat 1 as the root node<br>
visited = [F, F, F, F, F, F, F, F]<br>
First add 1 in visited, visited = [F, T, F, F, F, F, F, F]<br>

- 1 -> 3
  - add 3 as visited, visited = [F, T, F, T, F, F, F, F]
  - add 1 to the result of 3
  - Result Set: [[], [], [], [0, 1], [0], [0], [0], [0]]
  - 1 -> 3 -> 5
    - add 5 as visited, visited = [F, T, F, T, F, T, F, F]
    - add 1 to the result of 5
    - Result Set: [[], [], [], [0, 1], [0], [0, 1], [0], [0]]
  - 1 -> 3 -> 6
    - add 6 as visited, visited = [F, T, F, T, F, T, T, F]
    - add 1 to the result of 6
    - Result Set: [[], [], [], [0, 1], [0], [0, 1], [0, 1], [0]]
  - 1 -> 3 -> 7
    - add 7 as visited, visited = [F, T, F, T, F, T, T, T]
    - add 1 to the result of 7
    - Result Set: [[], [], [], [0, 1], [0], [0, 1], [0, 1], [0, 1]]

We're done with node 1, we can move to node 2<br>

Start from node 2, start the depth first search, and treat 2 as the root node<br>
visited = [F, F, F, F, F, F, F, F]<br>
First add 2 in visited, visited = [F, F, T, F, F, F, F, F]<br>

- 2 -> 4
  - add 4 as visited, visited = [F, F, T, F, T, F, F, F]
  - add 2 to the result of 4
  - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1], [0, 1], [0, 1]]
  - 2 -> 4 -> 6
    - add 6 as visited, visited = [F, F, T, F, T, T, F, F]
    - add 2 to the result of 6
    - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1], [0, 1, 2], [0, 1]]
  - 2 -> 4 -> 7
    - add 7 as visited, visited = [F, F, T, F, T, T, F, T]
    - add 2 to the result of 7
    - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1], [0, 1, 2], [0, 1, 2]]
- 2 -> 7
  - because 7 is visited, we don't need to explore 7 again

We're done with node 2, we can move to node 3<br>

Start from node 3, start the depth first search, and treat 3 as the root node<br>
visited = [F, F, F, F, F, F, F, F]<br>
First add 3 in visited, visited = [F, F, F, T, F, F, F, F]<br>

- 3 -> 5
  - add 5 as visited, visited = [F, F, F, T, F, T, F, F]
  - add 3 to the result of 5
  - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1, 3], [0, 1, 2], [0, 1, 2]]
- 3 -> 6
  - add 6 as visited, visited = [F, F, F, T, F, T, T, F]
  - add 3 to the result of 6
  - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1, 3], [0, 1, 2, 3], [0, 1, 2]]
- 3 -> 7
  - add 7 as visited, visited = [F, F, F, T, F, T, T, T]
  - add 3 to the result of 7
  - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1, 3], [0, 1, 2, 3], [0, 1, 2, 3]]

We're done with node 3, we can move to node 4<br>

Start from node 4, start the depth first search, and treat 4 as the root node<br>
visited = [F, F, F, F, F, F, F, F]<br>
First add 4 in visited, visited = [F, F, F, F, T, F, F, F]<br>

- 4 -> 6
  - add 6 as visited, visited = [F, F, F, F, T, F, T, F]
  - add 4 to the result of 6
  - Result Set: [[], [], [], [0, 1], [0, 2], [0, 1, 3], [0, 1, 2, 3, 4], [0, 1, 2, 3]]

Finish the whole process, we have the result set as follows<br>
Result Set: [[], [], [], [0, 1], [0, 2], [0, 1, 3], [0, 1, 2, 3, 4], [0, 1, 2, 3]]<br>

# Complexity Analysis
## Time Complexity O(n *(V + E))
- n and V is the number of nodes
- E is the number of edges
- For each depth first search, the time complexity is O(V + E)
- We have to do the depth first search for each node, so the time complexity is O(n * (V + E))

## Space Complexity O(V + E)
- We have to store the adjacency list, so the space complexity is O(V + E)