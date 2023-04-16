//////////////////////////////////////////////////////
// *** Surrounded Regions ***
//////////////////////////////////////////////////////
/*
Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example 1:
Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation: Surrounded regions should not be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.

Example 2:
Input: board = [["X"]]
Output: [["X"]]
*/
/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
/*
This problem is a little bit tricky

The question aksed us
capture all surrounded region
What is surrounded region?
It's the region is that are 4-directionally surrounded by 'X'
If we use this mindset to solve the problem
It seems that we need to do some works from inside to outside
But it's actually very hard to do

So we can change the mindset
Instead of thinking capture all surrounded region,
capture all region which is not surrounded by "X"(aka, safe region)
Why this is simpler?
Because it's easier to find the safe region
The safe region is the region connect with the border

So we can just simply iterate through all the border
Once finding the "O", we do DFS or BFS to capture all the connected region
Add it to the seenSafe object, which means these regions are safe
aka, don't need to change to "X"

After that, all we need to do is simply iterate through all matrix
If it's "O" and not in the seenSafe
We know it's the region is that are 4-directionally surrounded by 'X'
So change to "X"

The is very important concept to solve this problem
Reverse thinking

Also, don't always think a whole matrix as a single graph
We can also think a whole matrix as mutiple small graphs
Like this problem,
All border is "O" is a small graphs, so that we use DFS or BFS
For example,
  ['X', 'X', 'O', 'O', 'O'],
  ['X', 'X', 'O', 'X', 'X'],
  ['X', 'X', 'O', 'X', 'X'],
  ['X', 'X', 'X', 'O', 'X'],
  ['X', 'O', 'O', 'X', 'X'],
  ['X', 'X', 'O', 'X', 'X']
 
Small graph
              O   O   O
              O
              O

Another one
        O    O
             O

************************************************************
r = row, c = col
Time: O(r * c)

Space: O(r * c)
=> seenSafe object, and queue
*/
var solve = function (board) {
  const rowLength = board.length;
  const colLength = board[0].length;
  const seenSafe = {};

  // iterate through border
  // if it's "O", does DFS
  for (let i = 0; i < rowLength; i++) {
    if (board[i][0] === 'O') findSafeVertices(board, i, 0, seenSafe);

    if (board[i][colLength - 1] === 'O')
      findSafeVertices(board, i, colLength - 1, seenSafe);
  }

  // iterate through border
  // if it's "O", does DFS
  for (let i = 0; i < colLength; i++) {
    if (board[0][i] === 'O') findSafeVertices(board, 0, i, seenSafe);

    if (board[rowLength - 1][i] === 'O')
      findSafeVertices(board, rowLength - 1, i, seenSafe);
  }

  // iterate through all matrix
  for (let r = 0; r < rowLength; r++) {
    for (let c = 0; c < colLength; c++) {
      const key = `${r}-${c}`;

      // if it's "O" and not in the seenSafe, change to "X"
      if (board[r][c] === 'O' && !seenSafe[key]) board[r][c] = 'X';
    }
  }
};

function findSafeVertices(board, row, col, seenSafe) {
  const stack = [[row, col]];

  while (stack.length > 0) {
    const [r, c] = stack.pop();
    const key = `${r}-${c}`;

    // If it's invalid or has been seen or it's "X"
    // just skip it
    if (
      r < 0 ||
      r >= board.length ||
      c < 0 ||
      c >= board[0].length ||
      board[r][c] === 'X' ||
      seenSafe[key]
    )
      continue;

    seenSafe[key] = true;

    stack.push([r + 1, c]);
    stack.push([r - 1, c]);
    stack.push([r, c + 1]);
    stack.push([r, c - 1]);
  }
}
