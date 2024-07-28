# Problem Explanation

Because the problem asks for the shortest path, we know that we should use BFS to solve this problem.<br>
When we do the BFS, we just need to keep track of the steps, and when we reach the destination, we immediately return the steps because we know that this would be the shortest path.<br>
Therefore, the initial approach is following:<br>
1. Create a queue and add the start location to the queue.
2. Start the BFS loop.
   - (1) Pop the first element from the queue.
   - (2) Check if the current location is the destination. If it is, return the steps.
   - (3) Loop through the four directions.
      - (a) Check if the neighbor is out of the boundary
      - (b) Check if the neighbor is obstacle and we have no k to eliminate the obstacle.
      - (c) Add neighbor to the queue. If it's obstacle, we need to decrease the k by 1.

This approach does have one problem, which is that we may visit the same location multiple times.<br>
Therefore, we need to keep track of the visited locations.<br>
At first, the intuition is to use a 2D array(bool) to keep track of the visited locations.<br>
But this still has a problem.<br>
Let's see the following example:<br>
```
[
  [0, 0],
  [1, 0],
  [1, 0],
  [0, 0],
  [0, 1],
  [0, 0]
]
```
k = 1<br>
From a high level, we know the correct answer is 6<br>
Path: (0, 0) -> (0, 1) -> (1, 1) -> (2, 1) -> (3, 1) -> (4, 1) -> (5, 1)<br>

However, this path won't reach if we use the above approach.<br>
Let's walk through the example:<br>
- Start Node: (0, 0)
  - neighbors: [(1, 0), (0, 1)]
  - queue: [(1, 0), (0, 1)]
  - visited: {(0, 0): true, (1, 0): true, (0, 1): true}

- Node: (1, 0)
  - neighbors: [(0, 0), (1, 1)]
  - queue: [(0, 1), (1, 1)]
  - visited: {(0, 0): true, (1, 0): true, (0, 1): true, (1, 1): true}
  - k = 0 because we have obstacle at (1, 0)

- Node: (0, 1)
Now, we have a problem.<br>
Reacll that we said the correct path is (0, 0) -> (0, 1) -> (1, 1) -> ....<br>
Now, we should move to (1, 1) from (0, 1), but we can't because we have visited (1, 1) before.<br>

What's the problem?<br>
The problem is that the state of the node is not only determined by the location, but also the k.<br>
When node is (1, 0), we mark (1, 1) as visited<br>
However, to be more precise, we should mark (1, 1) with k = 0 as visited<br>
When we reach (0, 1), k should still be 1, so the state of node is (1, 1) with k = 1, which is different from the previous state.<br>

Therefore, we not only need to consider the location, but also the k.<br>
We should consider k as part of the state of the node.<br>
How can we do that?<br>
We can use a 3D array to keep track of the visited locations.<br>
The first two dimensions are the location, and the third dimension is the k.<br>
`[row][column][k]`<br>
```
[                    -> row
  [                  -> column
    [false, false]   -> k
  ],
]
```

Let's summarize our corrected approach:
1. Create a queue and add the start location to the queue.
2. Create a 3D array to keep track of the visited locations.
3. Start the BFS loop.
   - (1) Pop the first element from the queue.
   - (2) Check if the current location is the destination. If it is, return the steps.
   - (3) Loop through the four directions.
      - (a) Check if the neighbor is out of the boundary
      - (b) Check if the neighbor is obstacle and we have no k to eliminate the obstacle.
      - (c) Check if the neighbor is visited with location and k.
      - (d) Add neighbor to the queue. If it's obstacle, we need to decrease the k by 1.
      - (e) Mark the neighbor as visited with location and k.


Note that we can still use DFS to solve this problem, but it's less efficient than BFS because we need explore all the paths to find the shortest path.<br>
The time complexity of DFS is O(4^(m * n)), where m is the number of rows and n is the number of columns.<br>
Because for each cell, we have 4 directions to explore.<br>
And the depth of the tree is m * n.<br>

# Complexity Analysis
## Time Complexity O(m * n * k)
- m is the number of rows in the grid
- n is the number of columns in the grid
- k is the number of obstacles that we can eliminate
- We know the there are m * n cells in the grid we have to visit
- Additionally, we can visit each cell with k different states
  - For example, suppose k = 2, we can visit the cell(1, 0) twice with k = 0 and k = 1

## Space Complexity O(m * n * k)
- We need to keep track of the visited locations with location and k
- Therefore, the space complexity is O(m * n * k)
