//////////////////////////////////////////////////////
// *** Nearest Exit from Entrance in Maze ***
//////////////////////////////////////////////////////
/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
 

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/
/**
 * @param {character[][]} grid
 * @return {number}
 */
/*
This problem is a lot easier, just using BFS, queue
The idea is loop through every single element in the grid
If it's "1" which means it's an island, and it haven't been visited yet
Then we do BFS on this element, mark this element and all the island connected with this one as visited
And increment the islands by 1
Keeping the process

************************************************************
r = row, c = col
Time: O(r * c)
Space: O(r * c)
*/
var numIslands = function (grid) {
  let islands = 0;
  const seen = {};

  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      const key = `${r}-${c}`;
      if (grid[r][c] === '1' && !seen[key]) {
        BFS(grid, r, c, seen);
        islands++;
      }
    }
  }

  return islands;
};

function BFS(grid, r, c, seen) {
  const queue = [[r, c]];

  while (queue.length > 0) {
    const [row, col] = queue.shift();
    const key = `${row}-${col}`;

    if (
      row < 0 ||
      row >= grid.length ||
      col < 0 ||
      col >= grid[0].length ||
      grid[row][col] === '0' ||
      seen[key]
    )
      continue;

    seen[key] = true;
    queue.push([row + 1, col]);
    queue.push([row - 1, col]);
    queue.push([row, col + 1]);
    queue.push([row, col - 1]);
  }
}
