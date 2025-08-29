# Problem Statement
You are given an m×n 2D grid initialized with these three possible values:
- -1 - A water cell that cannot be traversed.
- 0 - A treasure chest.
- INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.

Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest, then the value should remain INF.

Assume the grid can only be traversed up, down, left, or right.

Modify the grid in-place.

Example 1:
```
Input: [
  [2147483647,-1,0,2147483647],
  [2147483647,2147483647,2147483647,-1],
  [2147483647,-1,2147483647,-1],
  [0,-1,2147483647,2147483647]
]

Output: [
  [3,-1,0,1],
  [2,2,1,-1],
  [1,-1,2,-1],
  [0,-1,3,4]
]
```

Example 2:
```
Input: [
  [0,-1],
  [2147483647,2147483647]
]

Output: [
  [0,-1],
  [1,2]
]
```

Constraints:
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 100
- grid[i][j] is one of {-1, 0, 2147483647}

# Typical BFS Approach
The first intuition for solving this problem is using a typical BFS approach:
- Loop through all cells in the grid
- When a cell's value equals 0, start a BFS from this cell
- For the BFS approach, we use level-by-level exploration:
  - Get the current length of the queue
  - Start a loop that runs for the length of the queue:
    - Pop a node from the queue and try all possible directions
    - For each node, validate it:
      - Cannot be out of bounds
      - The value cannot be less than or equal to 0 (if the cell is 0 or -1, we can't traverse it)
      - We haven't visited it before
    - Modify the value with the current distance
  - Increment the distance

One important thing to note for this solution is that we must use a hash table to keep track of which cells we've explored in the current BFS. Without it, we might explore duplicate cells, causing an infinite loop.

The major problem with this approach is its inefficiency. Let's look at an example:

```
Initial State:
[  0][INF]
[INF][INF]
[INF][INF]
[INF][  0]
```

When we start from (0,0), we'll perform a BFS on the entire grid (except for the last cell).

This is the result after the first BFS:
```
After first BFS:
[0][1]
[1][2]
[2][3]
[3][0]
```

Then, we'll start a second BFS from (3,1).

This is the result after the second BFS:
```
After second BFS:
[0][1]
[1][2]
[2][1]
[1][0]
```

Even though we only update two cells, we actually need to explore the entire grid again. This is duplicate work, and we can do better.

# Multi-Source BFS
The idea is actually quite straightforward. All we need to do is first add all cells with value equal to 0 into the queue. This means we'll have multiple starting points. Whenever we find a cell that hasn't been processed (the value is still INF), we simply update the cell value to the current cell value plus 1.

Let's walk through an example:

```
[INF][ -1][  0][INF]
[INF][INF][INF][ -1]
[INF][ -1][INF][ -1]
[  0][ -1][INF][INF]
```

Queue starts with treasure locations: `[(0,2), (3,0)]`

**Step 1:** Process (0,2) - treasure at top
- Check neighbors of (0,2):
  - (0,1): wall (-1), skip
  - (0,3): INF → update to 1, add to queue
  - (-1,2): out of bounds, skip
  - (1,2): INF → update to 1, add to queue

Queue: `[(3,0), (0,3), (1,2)]`
```
[INF][ -1][  0][  1]
[INF][INF][  1][ -1]
[INF][ -1][INF][ -1]
[  0][ -1][INF][INF]
```

**Step 2:** Process (3,0) - treasure at bottom left
- Check neighbors of (3,0):
  - (2,0): INF → update to 1, add to queue
  - (3,1): wall (-1), skip
  - (4,0): out of bounds, skip
  - (3,-1): out of bounds, skip

Queue: `[(0,3), (1,2), (2,0)]`
```
[INF][ -1][  0][  1]
[INF][INF][  1][ -1]
[  1][ -1][INF][ -1]
[  0][ -1][INF][INF]
```

**Step 3:** Process (0,3)
- Check neighbors of (0,3):
  - (0,2): already 0 (treasure), skip
  - (0,4): out of bounds, skip
  - (-1,3): out of bounds, skip
  - (1,3): wall (-1), skip

Queue: `[(1,2), (2,0)]`

**Step 4:** Process (1,2)
- Check neighbors of (1,2):
  - (1,1): INF → update to 2, add to queue
  - (1,3): wall (-1), skip
  - (0,2): already 0 (treasure), skip
  - (2,2): INF → update to 2, add to queue

Queue: `[(2,0), (1,1), (2,2)]`
```
[INF][ -1][  0][  1]
[INF][  2][  1][ -1]
[  1][ -1][  2][ -1]
[  0][ -1][INF][INF]
```

**Step 5:** Process (2,0)
- Check neighbors of (2,0):
  - (1,0): INF → update to 2, add to queue
  - (2,1): wall (-1), skip
  - (3,0): already 0 (treasure), skip
  - (2,-1): out of bounds, skip

Queue: `[(1,1), (2,2), (1,0)]`
```
[INF][ -1][  0][  1]
[  2][  2][  1][ -1]
[  1][ -1][  2][ -1]
[  0][ -1][INF][INF]
```

**Step 6:** Process (1,1)
- Check neighbors of (1,1):
  - (1,0): already 2, skip
  - (1,2): already 1, skip
  - (0,1): wall (-1), skip
  - (2,1): wall (-1), skip

Queue: `[(2,2), (1,0)]`

**Step 7:** Process (2,2)
- Check neighbors of (2,2):
  - (2,1): wall (-1), skip
  - (2,3): wall (-1), skip
  - (1,2): already 1, skip
  - (3,2): INF → update to 3, add to queue

Queue: `[(1,0), (3,2)]`
```
[INF][ -1][  0][  1]
[  2][  2][  1][ -1]
[  1][ -1][  2][ -1]
[  0][ -1][  3][INF]
```

**Step 8:** Process (1,0)
- Check neighbors of (1,0):
  - (1,1): already 2, skip
  - (1,-1): out of bounds, skip
  - (0,0): INF → update to 3, add to queue
  - (2,0): already 1, skip

Queue: `[(3,2), (0,0)]`
```
[  3][ -1][  0][  1]
[  2][  2][  1][ -1]
[  1][ -1][  2][ -1]
[  0][ -1][  3][INF]
```

**Step 9:** Process (3,2)
- Check neighbors of (3,2):
  - (3,1): wall (-1), skip
  - (3,3): INF → update to 4, add to queue
  - (2,2): already 2, skip
  - (4,2): out of bounds, skip

Queue: `[(0,0), (3,3)]`
```
[  3][ -1][  0][  1]
[  2][  2][  1][ -1]
[  1][ -1][  2][ -1]
[  0][ -1][  3][  4]
```

**Step 10:** Process (0,0)
- All neighbors are either walls, out of bounds, or already have smaller values

Queue: `[(3,3)]`

**Step 11:** Process (3,3)
- All neighbors are either walls, out of bounds, or already have smaller values

Queue: `[]` - Done!

**Final Result:**
```
[  3][ -1][  0][  1]
[  2][  2][  1][ -1]
[  1][ -1][  2][ -1]
[  0][ -1][  3][  4]
```

The reason this solution works is because we use BFS. ***Whenever we find a cell that hasn't been processed, we know that our current cell value plus one is the shortest distance to that cell.***