//////////////////////////////////////////////////////
// *** Pacific Atlantic Water Flow ***
//////////////////////////////////////////////////////
/*
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.

Example 1:
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]

Example 2:
Input: heights = [[2,1],[1,2]]
Output: [[0,0],[0,1],[1,0],[1,1]]

Constraints:
m == heights.length
n == heights[r].length
1 <= m, n <= 200
0 <= heights[r][c] <= 105
*/
/**
 * @param {number[][]} heights
 * @return {number[][]}
 */
/*
The main idea of solving this problem is very similar to 
Surrounded Regions(slove)
There are two main way to solve this problem
1. Iterate through all element in the matrix
   Do DFS or BFS on every element
   Test if it can reach both Pacific Ocean and Atlantic Ocean 
   This is kind of un-efficient
   A lot of repeated work

2. Set two visited object
   Iterate all the border
   Why?
   Because we can gurantee the border can reach one of the ocean
         top
    left      right
        bottom
    We know that top and left is connected with Pacific Ocean, and bottom and right is connected with  Atlantic Ocean
    We just iterate each element on the border of top and left
    Do DFS, 
    Check others elements can reach from these border elements, and marked it as Pacific Ocean
    Same things for right and bottom
    Now we have both visited object
    We just iterate through all the element in the matrix
    If this element is on both ocean, add it to the result

Second idea is kind of like reverse thinking
*/
var pacificAtlantic = function (heights) {
  const pacVisited = {};
  const atlVisited = {};
  const res = [];
  const ROWS = heights.length;
  const COLS = heights[0].length;

  // loop from top ~ bottom
  for (let i = 0; i < ROWS; i++) {
    // left border(Pacific Ocean)
    DFS(heights, i, 0, heights[i][0], pacVisited);

    // right border(Atlantic Ocean)
    DFS(heights, i, COLS - 1, heights[i][COLS - 1], atlVisited);
  }

  // loop from left ~ right
  for (let i = 0; i < COLS; i++) {
    // top border(Pacific Ocean)
    DFS(heights, 0, i, heights[0][i], pacVisited);

    // bottom border(Atlantic Ocean)
    DFS(heights, ROWS - 1, i, heights[ROWS - 1][i], atlVisited);
  }

  // iterate all element, check if it's both on two visited object
  for (let r = 0; r < ROWS; r++) {
    for (let c = 0; c < COLS; c++) {
      const key = `${r}-${c}`;
      if (pacVisited[key] && atlVisited[key]) res.push([r, c]);
    }
  }

  return res;
};

function DFS(grid, r, c, preHeight, visited) {
  const key = `${r}-${c}`;

  if (
    r < 0 ||
    r >= grid.length ||
    c < 0 ||
    c >= grid[0].length ||
    visited[key] ||
    grid[r][c] < preHeight // we can only go outward when current height is greater than preHeight
    /*
    4  5
    In this example, 5 is current height, 4 is preHeight,
    In this case, 5 can flood into 4
    so when we start at 4, we have to go outward to 5
    we use reverse thinking
    */
  )
    return;

  visited[key] = true;
  DFS(grid, r + 1, c, grid[r][c], visited);
  DFS(grid, r - 1, c, grid[r][c], visited);
  DFS(grid, r, c + 1, grid[r][c], visited);
  DFS(grid, r, c - 1, grid[r][c], visited);
}

/*
First idea, not efficient
*/
var pacificAtlantic = function (heights) {
  const res = [];

  for (let r = 0; r < heights.length; r++) {
    for (let c = 0; c < heights[0].length; c++) {
      if (BFS(heights, r, c)) {
        res.push([r, c]);
      }
    }
  }

  return res;
};

function BFS(grid, r, c) {
  const seen = {};
  const queue = [[r, c]];
  let pacific = false;
  let atlantic = false;

  while (queue.length > 0) {
    const [row, col] = queue.shift();
    const key = `${row}-${col}`;
    const height = grid[row][col];

    if (seen[key]) continue;
    else seen[key] = true;

    if (row === 0 || col === 0) pacific = true;

    if (row === grid.length - 1 || col === grid[0].length - 1) atlantic = true;

    if (pacific && atlantic) return true;

    if (row - 1 >= 0 && height >= grid[row - 1][col])
      queue.push([row - 1, col]);

    if (row + 1 <= grid.length - 1 && height >= grid[row + 1][col])
      queue.push([row + 1, col]);

    if (col - 1 >= 0 && height >= grid[row][col - 1])
      queue.push([row, col - 1]);

    if (col + 1 <= grid[0].length && height >= grid[row][col + 1])
      queue.push([row, col + 1]);
  }

  return false;
}
