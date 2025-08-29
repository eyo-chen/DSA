# Problem Explanation

The idea to solve this problem is straightforward.<br>

1. Iterate through the whole grid
2. If find 1, increment the count of islands and call the helper function to mark all the connected 1's to 0 or other number.

The helper function could be BFS or DFS. Here, I have used DFS.

# Caveat
Look at the solution `NumIslands4`, which has a small bug.<br>
The wrong part is **WHEN we mutate the value to ‘0’** <br>

**When the node is put into the queue, we should immediately mark it as ‘0’**<br>

The key difference is:
- In the wrong version, a cell might be added to the queue multiple times if multiple neighbors discover it before it's processed
- In the corrected version, we mark cells as visited ('0') as soon as we discover them, preventing them from being added to the queue multiple times

While your original solution will still give the correct answer (because we eventually mark all connected '1's as '0'), it's less efficient because:
- The queue might become larger than necessary
- We might process the same cell multiple times
- It uses more memory than needed

This is a common pattern in graph algorithms - mark nodes as visited when you discover them, not when you process them, to avoid duplicate work.

Let's walk through an example:<br>
Consider this 3x3 grid:

```
1 1 0
1 1 0
0 0 0
```

Let's see how wrong implementation processes it, starting from position (0,0):
**Wrong Implementation:**
```
Step 1: Start at (0,0)
Queue: [(0,0)]
Grid:
1 1 0
1 1 0
0 0 0

Step 2: Process (0,0), mark it as 0
Queue: [(0,1), (1,0)]  // Both right and down neighbors see '1' and get added
Grid:
0 1 0
1 1 0
0 0 0

Step 3: Process (0,1), mark it as 0
Queue: [(1,0), (1,1)]  // (1,1) gets added because it's still '1'
Grid:
0 0 0
1 1 0
0 0 0

Step 4: Process (1,0), mark it as 0
Queue: [(1,1), (1,1)]  // Notice (1,1) gets added AGAIN because it's still
'1'
Grid:
0 0 0
0 1 0
0 0 0

Step 5: Process (1,1) twice because it's in queue twice
...
```

**Corrected Implementation:**

```
Step 1: Start at (0,0), mark it as 0 immediately
Queue: [(0,0)]
Grid:
0 1 0
1 1 0
0 0 0

Step 2: Process (0,0), check neighbors
Queue: [(0,1), (1,0)]  // Add neighbors and mark them as 0 immediately
Grid:
0 0 0
0 1 0
0 0 0

Step 3: Process (0,1), check neighbors
Queue: [(1,0), (1,1)]  // (1,1) gets added once and marked as 0
Grid:
0 0 0
0 0 0
0 0 0

Step 4: Process remaining queue items
...
```

The key differences:

1. In the wrong implementation, position (1,1) could be added to the queue multiple times because it remains '1' until it's actually processed
2. In the corrected version, as soon as we discover a cell, we mark it as '0', ensuring it won't be added to the queue again by other neighboring cells

This is particularly important for larger grids where a single cell might be surrounded by multiple '1's. In the worst case, the wrong implementation might add the same cell to the queue as many times as it has neighbors marked '1', while the corrected version will always add each cell exactly once.

For example, consider a cell surrounded by '1's:

```
1 1 1
1 1 1
1 1 1
```

In the wrong implementation, the center cell (1,1) could potentially be added to the queue 4 times (once by each neighbor), while in the corrected version it would only be added once.

# Complexity Analysis
## Time Complexity: O(M*N)
- Where `M` is the number of rows and `N` is the number of columns in the grid.
- We are visiting each cell at most once.

## Space Complexity: O(M*N)
- Where `M` is the number of rows and `N` is the number of columns in the grid.
- The space complexity is used by the recursive stack.
- In the worst case, the depth of the recursive stack could be `M*N`.