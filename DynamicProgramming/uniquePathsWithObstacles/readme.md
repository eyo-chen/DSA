# Problem Explanation
This problem is a variation of the unique paths problem where we need to calculate the number of ways to reach the bottom-right corner of a grid, given that some cells are blocked (obstacles).<br>
The core idea is as same as unique path 1.

For top-down, recursive approach, the idea is the same as unique path 1.<br>
We just need to add a check to see if the current cell is an obstacle. If it is, we return 0.

For bottom-up, iterative approach, we need to take care more about the obstacle case.<br>
First, when we initialize the first row and first column, if we encounter an obstacle, the rest of the cells in that row or column are not accessible, so we need to set the value to 0.
In this case, we can use the following technique:
```go
		if obstacleGrid[r][0] == 1 {
			break
		}
		table[r][0] = table[r-1][0]
```
Note that the whole grid is initialized with 0, except the starting point.<br>
When the current cell is an obstacle, we know that the rest of the cells in that row or column are not accessible, so we just break the loop, and leave the rest of the cells in that row or column as 0.<br>
When the current cell is not an obstacle, we set the current cell to be equal to the previous cell.<br>
Why this works?<br>
Because we set the starting point to 1<br>
If the current cell is not an obstacle, it's also set to 1<br>


Second, when we iterate through the rest of the cells, we need to check if the current cell is an obstacle. If it is, we continue. If it's not, we set the current cell to be equal to the sum of the cell above it and the cell to the left of it.(same logic as unique path 1)


In this problem, we come up with a new optimization approach.<br>
We don't need to use a 2D array to store the number of ways to reach the current cell.<br>
1. Dependency pattern:
   In this problem, when we're calculating the number of paths to a cell, we only need two pieces of information:
   - The number of paths to the cell directly above it
   - The number of paths to the cell directly to its left

2. Row-by-row processing:
   We process the grid row by row, from left to right. At any point, we only need the information from the current row and the row above it.

3. Overwriting previous data:
   As we move through each row, we can overwrite the data from the previous row because we won't need it anymore.

Let's walk through an example:
```
[
  [0,0,0]
  [0,1,0]
  [0,0,0]
]
```

We initialize the table as following:
```
[1,0,0]
```
Scan first row:
- (0,0) is not an obstacle, so table[0] = 1 (because we set the starting point to 1)
- (0,1) is an obstacle, so table[1] = table[1] + table[0] = 0 + 1 = 1
  - table[1] represents the value of the cell above it. (In this case, it's 0 because it's the first row)
  - table[0] represents the value of the cell to the left of it.
  - It's still the sum of the cell above it and the cell to the left of it.
- (0,2) is not an obstacle, so table[2] = table[2] + table[1] = 0 + 1 = 1
  - same as above

Scan second row:
- (1,0) is not an obstacle, so table[0] = 1
  - table[0] represents the value of the cell above it.
  - because there's no cell to the left of it, so it's just the value of the cell above it.
- (1,1) is an obstacle, so table[1] = 0
- (1,2) is not an obstacle, so table[2] = table[2] + table[1] = 1 + 0 = 1
  - table[2] represents the the value of the cell above it.
  - Why is that?
  - Because it's the value we set at (0,2)
  - table[1] represents the value of the cell to the left of it.
  - It's still the sum of the cell above it and the cell to the left of it.

Scan third row: Same as above


As we can see, we only need to 1D array to store the number of ways to reach the current cell.<br>
Because the current cell table[i] represents the value of the cell above it, and table[i-1] represents the value of the cell to the left of it.<br>
It's important to recognize that table[i] is the value of the cell above it<br>

Key Points:
- The 1D array `table` represents the current row we're processing.
- `table[i]` represents the number of paths to reach the cell directly above the current cell.
- When we update `table[i]`, we add `table[i-1]`, which represents the number of paths to reach the cell to the left of the current cell.

# Complexity Analysis
- Time Complexity: O(m*n)
  - We need to iterate through the grid, so it's O(m*n)
- Space Complexity: O(n)
  - We only need to store the current row and the row above it, so it's O(n)
