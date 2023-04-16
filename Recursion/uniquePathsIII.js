//////////////////////////////////////////////////////
// *** Unique Paths III ***
//////////////////////////////////////////////////////
/*
ou are given an m x n integer array grid where grid[i][j] could be:

1 representing the starting square. There is exactly one starting square.
2 representing the ending square. There is exactly one ending square.
0 representing empty squares we can walk over.
-1 representing obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.

Example 1:
Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
Output: 2
Explanation: We have the following two paths: 
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)

Example 2:
Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
Output: 4
Explanation: We have the following four paths: 
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)

Example 3:
Input: grid = [[0,1],[2,0]]
Output: 0
Explanation: There is no path that walks over every empty square exactly once.
Note that the starting and ending square can be anywhere in the grid.

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 20
1 <= m * n <= 20
-1 <= grid[i][j] <= 2
There is exactly one starting cell and one ending cell.
*/
/*
Look at the comment, the solution should be pretty straightforward

************************************************************
row, col = grid
Time complexity: O(4 ^ (m*n))
=> The max bracnhing is four because we'll explore top, right, down and left
=> The deepest recursive tree is m * n
=> Why?
=> Because the worst case is travasring all the cell of the grid
=> the max amount of cell is m * n
=> so we have to go m * n deep to search every single possibilities

Space complexity: O(m * n)
=> The deepest recursive tree is m * n


*/
/**
 * @param {number[][]} grid
 * @return {number}
 */
var uniquePathsIII = function (grid) {
  let availableSteps = 1;
  let sRow, sCol;

  // find the starting position and total steps to traverse all available cell
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[0].length; j++) {
      // 0 means it's available cell to traverse
      if (grid[i][j] === 0) availableSteps++;

      // find the starting position
      if (grid[i][j] === 1) {
        sRow = i;
        sCol = j;
      }
    }
  }

  return recursiveHelper(grid, sRow, sCol, availableSteps);
};

function recursiveHelper(grid, row, col, availableSteps) {
  // base case 1: out of the range or invalid cell
  if (
    row < 0 || // out of the range
    col < 0 || // out of the range
    row === grid.length || // out of the range
    col === grid[0].length || // out of the range
    grid[row][col] === -1 // invalid cell (can't traverse)
  )
    return 0;

  // find the ending position
  if (grid[row][col] === 2) {
    // only return 1 if finishing traversing all available cell
    if (availableSteps === 0) return 1;
    return 0;
  }

  // choose (use -1 as traversed cell)(same as obstacles)
  grid[row][col] = -1;

  // explore
  const res =
    recursiveHelper(grid, row + 1, col, availableSteps - 1) +
    recursiveHelper(grid, row, col + 1, availableSteps - 1) +
    recursiveHelper(grid, row - 1, col, availableSteps - 1) +
    recursiveHelper(grid, row, col - 1, availableSteps - 1);

  // unchoose
  grid[row][col] = 0;

  return res;
}
