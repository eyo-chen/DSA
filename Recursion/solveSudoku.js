//////////////////////////////////////////////////////
// *** Sudoku Solver ***
//////////////////////////////////////////////////////
/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.

Example 1:
Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
Explanation: The input board is shown above and the only valid solution is shown below:

Constraints:
board.length == 9
board[i].length == 9
board[i][j] is a digit or '.'.
It is guaranteed that the input board has only one solution.
*/
/*
The though process is 
1. Traverse all cell of the board(2D array)
2. If cell is "."(empty cell), then we know we have to try to fill this cell up
3. Try 1 ~ 9 in this cell
4. Before we try out, we have to valid this value can be filled in this cell
   (Have to pass these three constraints(rulls))
5. If it's all pass, then we can temporarily fill this value in the cell
6. Keep trying next cell (keep recursing)
7. If can't find value after trying 1 ~ 9, then we know the value temporarily have been filled is wrong, so we have to go back to modifiy(backtrack)

************************************************************
Time complexity: O(1)
Space complexity: O(1)
=> Sudoku is always 9x9 grid
*/
/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solveSudoku = function (board) {
  recursiveHelper(board, 0, 0);

  function recursiveHelper(board, row, col) {
    for (let i = row; i < 9; i++, col = 0) {
      for (let j = col; j < 9; j++) {
        // skip if it's not emtpy
        if (board[i][j] !== '.') continue;

        // for each non-empty cell, try 1 ~ 9
        for (let num = 1; num <= 9; num++) {
          const tmpVal = String(num);

          if (checkValidPlacement(board, i, j, tmpVal)) {
            board[i][j] = tmpVal;

            if (recursiveHelper(board, i, j + 1)) return true;
          }
        }
        board[i][j] = '.'; // backtrack, reset to empty cell
        return false;
      }
    }

    return true;
  }
};

function checkValidPlacement(board, row, col, val) {
  const subRowStart = Math.floor(row / 3) * 3;
  const subColStart = Math.floor(col / 3) * 3;

  // check Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      if (board[subRowStart + i][subColStart + j] === val) return false;
    }
  }

  /*
  check
  Each of the digits 1-9 must occur exactly once in each row.
  Each of the digits 1-9 must occur exactly once in each column.
  */
  for (let i = 0; i < 9; i++) {
    if (board[row][i] === val || board[i][col] === val) return false;
  }

  return true;
}

const board = [
  ['5', '3', '.', '.', '7', '.', '.', '.', '.'],
  ['6', '.', '.', '1', '9', '5', '.', '.', '.'],
  ['.', '9', '8', '.', '.', '.', '.', '6', '.'],
  ['8', '.', '.', '.', '6', '.', '.', '.', '3'],
  ['4', '.', '.', '8', '.', '3', '.', '.', '1'],
  ['7', '.', '.', '.', '2', '.', '.', '.', '6'],
  ['.', '6', '.', '.', '.', '.', '2', '8', '.'],
  ['.', '.', '.', '4', '1', '9', '.', '.', '5'],
  ['.', '.', '.', '.', '8', '.', '.', '7', '9'],
];

// console.log(solveSudoku(board));
// console.log(board);
// console.log(canPlaceValue(board, 2, 5, '1'));
