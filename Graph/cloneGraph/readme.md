# Problem Explanation

My initial idea to solve this problem is very straightforward.<br>
Traverse the graph using BFS or DFS, and an array to keep tracke of the visited nodes to avoid revisiting the same node because the graph is undirected.<br>
Along the way, we build the clone graph.<br>

However, this approach has a problem.<br>
Let's walk through the example below to understand the problem<br>
```
1   -   2
|       |
4   -   3
```
We start at node 1<br>
We visit it's neighbors 2 and 4<br>
Add 2 and 4 to 1's neighbors, and also add 2 and 4 to the queue<br>
neighbors: {1: [2, 4]}<br>
queue: [2, 4]<br>

We visit node 2<br>
Now the problem arises<br>
When we're at node 2, we again try to add it's neighbors<br>
In the previous step, we've already created a clone node for 1<br>
Now, when we try to add 1 to 2's neighbors, we lose the reference to 1<br>
Note that we can't just create another node for 1 because that's the different memory location(reference)<br>
That's the problem<br>

Now, let's correct our approach, but the problem is still the same<br>
Let's see the thought process<br>
When we visit a node,<br>
We not only add it's neighbors, but we also add it's neighbors of neighbors<br>
What does that mean?<br>
Suppose we are at node 1<br>
We know that 1's neighbors are 2 and 4<br>
We add 2 and 4 to 1's neighbors(1: [2, 4])<br>
Also, we add 1 to 2's and 4's neighbors(2: [1], 4: [1])<br>

Let's walk through again<br>
```
1   -   2
|       |
4   -   3
```
We start at node 1<br>
We visit it's neighbors 2 and 4<br>
Add 2 and 4 to 1's neighbors, and also add 2 and 4 to the queue<br>
Add 1 to 2's neighbors, and 1 to 4's neighbors<br>
neighbors: {1: [2, 4], 2: [1], 4: [1]}<br>
queue: [2, 4]<br>
has visited: [1]<br>

We visit node 2<br>
We visit it's neighbors 1 and 3<br>
We skip 1 because we have visited it<br>
Add 3 to 2's neighbors<br>
Add 2 to 3's neighbors<br>
neighbors: {1: [2, 4], 2: [1, 3], 3: [2], 4: [1]}<br>
queue: [4, 3]<br>
has visited: [1, 2]<br>

We visit node 4<br>
We visit it's neighbors 1 and 3<br>
We skip 1 because we have visited it<br>
Add 3 to 4's neighbors<br>
Now, we have the similar problem as before<br>
We have to add 3 to 4's neighbors, but we lose the reference to 3<br>


So, how we can solve this problem?<br>
The problem is that we lose the reference to the clone node<br>
Therefore, the solution is to keep a reference of the clone node by using a hashmap<br>
key: the value of the node or the node(pointer) itself<br>
value: the clone node<br>

When we visit a node, we check if the node is in the hashmap<br>
If it is, we don't need to create a new node, we can simply reference the clone node<br>
If it is not, we create a new node, and add it to the hashmap<br>

Let's summarize the steps:<br>
(Assume using BFS)<br>
1. Create a queue, and add starting node(input node) to the queue<br>
2. Create a hashmap, and add the starting node, and it's clone node to the hashmap<br>
3. While the queue is not empty<br>
    1. Pop the node from the queue<br>
    2. For each neighbor of the node<br>
        1. If the neighbor is not in the hashmap<br>
            1. Create a new node, and add it to the hashmap<br>
            2. Add the neighbor to the queue<br>
        2. Add the neighbor to the clone node's neighbors<br>

Let's walk through the example again<br>
```
1   -   2
|       |
4   -   3
```
c: clone node<br>
o: original node<br>
- Starting point
  - queue: [1o]
  - hashmap: {1o: 1c}
  - result: []

- First iteration
  - visit 1o
  - get copy node 1c from hashmap
  - visit 2o, 4o
    - visit 2o
      - haven't visited 2o
      - create 2c
      - add {2o: 2c} to hashmap
      - add 2o into queue
      - add 2c into 1c's neighbors
    - visit 4o
      - haven't visited 4o
      - create 4c
      - add {4o: 4c} to hashmap
      - add 4o into queue
      - add 4c into 1c's neighbors
   - queue: [2o, 4o]
   - hashmap: {1o: 1c, 2o: 2c, 4o: 4c}
   - result: [[2c, 4c]]

- Second iteration
  - visit 2o
  - get copy node 2c from hashmap
  - visit 1o, 3o
    - visit 1o
      - have visited 1o
      - get copy node 1c from hashmap
      - add 1c into 2c's neighbors
    - visit 3o
      - haven't visited 3o
      - create 3c
      - add {3o: 3c} to hashmap
      - add 3o into queue
      - add 3c into 2c's neighbors
   - queue: [4o, 3o]
   - hashmap: {1o: 1c, 2o: 2c, 3o: 3c, 4o: 4c}
   - result: [[2c, 4c], [1c, 3c]]

- Third iteration
  - visit 4o
  - get copy node 4c from hashmap
  - visit 1o, 3o
    - visit 1o
      - have visited 1o
      - get copy node 1c from hashmap
      - add 1c into 4c's neighbors
    - visit 3o
      - have visited 3o
      - get copy node 3c from hashmap
      - add 3c into 4c's neighbors
   - queue: [3o]
   - hashmap: {1o: 1c, 2o: 2c, 3o: 3c, 4o: 4c}
   - result: [[2c, 4c], [1c, 3c], [], [1c, 3c]]

- Fourth iteration
  - visit 3o
  - get copy node 3c from hashmap
  - visit 2o, 4o
    - visit 2o
      - have visited 2o
      - get copy node 2c from hashmap
      - add 2c into 3c's neighbors
    - visit 4o
      - have visited 4o
      - get copy node 4c from hashmap
      - add 4c into 3c's neighbors
   - queue: []
   - hashmap: {1o: 1c, 2o: 2c, 3o: 3c, 4o: 4c}
   - result: [[2c, 4c], [1c, 3c], [2c, 4c], [1c, 3c]]

- Done

## Caveat
- Note that we can use either value or node(pointer) as the key of the hashmap because we can guarantee that the value of the node is unique
- However, if that's not the case, we need to use the node(pointer) as the key of the hashmap to make sure that the key is unique
- We only use pointer as key in the second solution

# Complexity Analysis
## Time Complexity (O(|V| + |E|))
- We visit each node once
- We visit each edge once


## Space Complexity (O(|V|))
- We use a hashmap to keep track of the clone nodes
- We use a queue to traverse the graph
