//////////////////////////////////////////////////////
// *** Island Perimeter ***
//////////////////////////////////////////////////////\
/*
You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example 1:
Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
Output: 16
Explanation: The perimeter is the 16 yellow stripes in the image above.

Example 2:
Input: grid = [[1]]
Output: 4

Example 3:
Input: grid = [[1,0]]
Output: 4
 
Constraints:
row == grid.length
col == grid[i].length
1 <= row, col <= 100
grid[i][j] is 0 or 1.
There is exactly one island in grid.
*/
/**
 * @param {number[][]} grid
 * @return {number}
 */
/*
The problem is fairly easy, just seen the code and comment

************************************************************
r = row, c = col
Time: O(r * c)
Space: O(r * c)
*/
var islandPerimeter = function (grid) {
  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      // find the first land(1), then return the result(perimeter)
      // because the problem said there's only one island
      if (grid[r][c] === 1) {
        return BFS(grid, r, c);
      }
    }
  }
};

// It's just normal BFS, but we have helper function countPerimeter to count perimeter
function BFS(grid, r, c) {
  const queue = [[r, c]];
  const seen = {};
  let perimeter = 0;

  while (queue.length > 0) {
    const [row, col] = queue.shift();
    const key = `${row}-${col}`;

    if (
      row < 0 ||
      row >= grid.length ||
      col < 0 ||
      col >= grid[0].length ||
      seen[key] ||
      grid[row][col] === 0
    )
      continue;

    perimeter += countPerimeter(grid, row, col);
    seen[key] = true;

    queue.push([row + 1, col]);
    queue.push([row - 1, col]);
    queue.push([row, col + 1]);
    queue.push([row, col - 1]);
  }

  return perimeter;
}

/*
Here, it makes use the power of coercion
We respectively count four perimeter
If it's out of the bound of it's water
set it to true
If it's island
set it to false

So that when we add all the booleans
It will convert to number
true -> 1
false -> 0

That's what we want
If it's adjacent is island, we don't want to count perimeter
*/
function countPerimeter(grid, r, c) {
  // top
  const top = r - 1 < 0 || grid[r - 1][c] === 0;

  // right
  const right = c + 1 >= grid[0].length || grid[r][c + 1] === 0;

  // bottom
  const bottom = r + 1 >= grid.length || grid[r + 1][c] === 0;

  // left
  const left = c - 1 < 0 || grid[r][c - 1] === 0;

  return top + right + bottom + left;
}

/*
Another solution from https://www.youtube.com/watch?v=fISIuAFRM2s
Use DFS
*/
var islandPerimeter = function (grid) {
  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      if (grid[r][c] === 1) {
        return DFS(grid, r, c, {});
      }
    }
  }
};

function DFS(grid, r, c, seen) {
  /*
  Logic is simlar to above
  if it's out of the bound of if it's water
  return 1 (count perimeter)
  */
  if (
    r < 0 ||
    r >= grid.length ||
    c < 0 ||
    c >= grid[0].length ||
    grid[r][c] === 0
  )
    return 1;

  const key = `${r}-${c}`;

  // if it's seen before, which means it's a island, so return 0
  if (seen[key]) {
    return 0;
  }

  seen[key] = true;

  return (
    DFS(grid, r + 1, c, seen) +
    DFS(grid, r - 1, c, seen) +
    DFS(grid, r, c + 1, seen) +
    DFS(grid, r, c - 1, seen)
  );
}
