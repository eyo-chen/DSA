//////////////////////////////////////////////////////
// *** Max Area of Island ***
//////////////////////////////////////////////////////
/*
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.


Example 1:
Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.

Example 2:
Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0
 
Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] is either 0 or 1.
*/
/**
 * @param {number[][]} grid
 * @return {number}
 */
/*
This is basically the same idea as numIslands
Just iterate the grid
If it's 1(island) and haven't been seen
Do the DFS on this node

************************************************************
r = row, c = col
Time: O(r * c)
=> Eventually, we'll only touch each element in the grid once

Space: O(r * c)
*/
var maxAreaOfIsland = function (grid) {
  const seen = {};
  let res = 0;

  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      const key = `${r}-${c}`;
      if (grid[r][c] === 1 && !seen[key]) {
        const area = DFS(grid, r, c, seen);
        res = Math.max(res, area);
      }
    }
  }

  return res;
};

// recursive DFS
function DFS(grid, r, c, seen) {
  const key = `${r}-${c}`;
  if (
    r < 0 ||
    r >= grid.length ||
    c < 0 ||
    c >= grid[0].length ||
    grid[r][c] === 0 ||
    seen[key]
  )
    return 0;

  seen[key] = true;

  return (
    1 +
    DFS(grid, r + 1, c, seen) +
    DFS(grid, r - 1, c, seen) +
    DFS(grid, r, c + 1, seen) +
    DFS(grid, r, c - 1, seen)
  );
}

// iterative DFS
function DFS(grid, r, c, seen) {
  const stack = [[r, c]];
  let area = 0;

  while (stack.length > 0) {
    const [row, col] = stack.pop();
    const key = `${row}-${col}`;

    if (
      row < 0 ||
      row >= grid.length ||
      col < 0 ||
      col >= grid[0].length ||
      grid[row][col] === 0 ||
      seen[key]
    )
      continue;

    area++;
    seen[key] = true;
    stack.push([row + 1, col]);
    stack.push([row - 1, col]);
    stack.push([row, col + 1]);
    stack.push([row, col - 1]);
  }

  return area;
}
