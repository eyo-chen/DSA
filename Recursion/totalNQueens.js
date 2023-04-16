//////////////////////////////////////////////////////
// *** N-Queens II ***
//////////////////////////////////////////////////////
/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

Example 1:
Input: n = 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown.

Example 2:
Input: n = 1
Output: 1
 
Constraints:
1 <= n <= 9
*/
/**
 * @param {number} n
 * @return {number}
 */
/*
This is basically the exact same idea as N-Queens 1

The main difference is we don't need to generate the board
we only need to return the number

************************************************************
n = the input n
Time complexity: O(!n * n)
=> checkValidPlacement is O(n) work

Space complexity: O(n)
=> the deepest height of recursive tree is n
*/
var totalNQueens = function (n) {
  return recursiveHelper(n, 0, []);
};

function recursiveHelper(n, row, pos) {
  if (n === row) {
    return 1;
  }

  let res = 0;

  for (let i = 0; i < n; i++) {
    if (checkValidPlacement(row, i, pos)) {
      pos.push(i);
      res += recursiveHelper(n, row + 1, pos, res);
      pos.pop();
    }
  }

  return res;
}

function checkValidPlacement(row, col, pos) {
  for (let i = 0; i < pos.length; i++) {
    if (
      row === i ||
      col === pos[i] ||
      row + col === i + pos[i] ||
      row - col === i - pos[i]
    )
      return false;
  }

  return true;
}
