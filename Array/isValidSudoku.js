//////////////////////////////////////////////////////
// *** Valid Sudoku ***
//////////////////////////////////////////////////////
/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
 

Example 1:
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Example 2:
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Constraints:
board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
Accepted
703,655
Submissions
1,283,562
*/
/**
 * @param {character[][]} board
 * @return {boolean}
 */
/*
This problem is not that hard after understanding the main logic
Just iterate whole board
And using hash table to cache the data about row, column, block and value
For example, row = 3, col = 5, val = 6
We can do sth like this
hashTable["row-3-6"] = true
=> row 3 has a value 6
hashTable["col-5-6"] = true
=> colum 5 has a value 6
hashTable["block-2-6"] = true
=> block 2 has a value 6

So that if we have a duplicate at same row, colum and block
Then the key is same, so that we can find it

************************************************************
Time compelxity: O(1)
Space comelxity: O(1)
=> sudoku always 9 * 9
*/
var isValidSudoku = function (board) {
  const hashTable = {};

  for (let r = 0; r < 9; r++) {
    for (let c = 0; c < 9; c++) {
      const val = board[r][c];
      if (val === '.') {
        continue;
      }

      // key for row, col and block
      const rowKey = `row-${r}-${val}`;
      const colKey = `col-${c}-${val}`;
      const blockKey = `block-${whichBlock(r, c)}-${val}`;

      // find duplicate
      if (hashTable[rowKey] || hashTable[colKey] || hashTable[blockKey])
        return false;

      hashTable[rowKey] = true;
      hashTable[colKey] = true;
      hashTable[blockKey] = true;
    }
  }

  return true;
};

/*
This is just naive way to check which block the element locate in
For example, 
row = 4, col = 1, locate at block 3
row = 7, col = 8, locate at block 8
so on and so forth
*/
function whichBlock(row, col) {
  if (row <= 2) {
    if (col <= 2) {
      return 0;
    }

    if (col > 2 && col <= 5) {
      return 1;
    }

    if (col > 5) {
      return 2;
    }
  }

  if (row > 2 && row <= 5) {
    if (col <= 2) {
      return 3;
    }

    if (col > 2 && col <= 5) {
      return 4;
    }

    if (col > 5) {
      return 5;
    }
  }

  if (row > 5) {
    if (col <= 2) {
      return 6;
    }

    if (col > 2 && col <= 5) {
      return 7;
    }

    if (col > 5) {
      return 8;
    }
  }
}

/*
the solution without whichBlock helper function
The way we decide which block it locate on is simply divide by 3
Why?
Because block is 3 * 3
We just use the very first element as index
For example,
row = 3, col = 3 is the first element in the block one
row = 4, col = 5 is also in the block one
Because Trunc(4 / 3) = 1, and Trunc(5 / 3) = 1, which is as same as the first element in the block one
row = 3, col = 3 => 3 / 3 = 1, 3 / 3 = 1

row = 6, col = 3 is the first element in the block seven
row = 8, col = 4 is also in the block seven
Because Trunc(6 / 3) = 2, and Trunc(4 / 3) = 1, which is as same as the first element in the block one
row = 6, col = 3 => 6 / 3 = 2, 3 / 3 = 1
https://www.youtube.com/watch?v=TjFXEUCMqI8
*/
var isValidSudoku = function (board) {
  const hashTable = {};

  for (let r = 0; r < 9; r++) {
    for (let c = 0; c < 9; c++) {
      const val = board[r][c];
      if (val === '.') {
        continue;
      }

      const rowKey = `row-${r}-${val}`;
      const colKey = `col-${c}-${val}`;
      const blockKey = `block-${Math.trunc(r / 3)}-${Math.trunc(c / 3)}-${val}`;

      if (hashTable[rowKey] || hashTable[colKey] || hashTable[blockKey])
        return false;

      hashTable[rowKey] = true;
      hashTable[colKey] = true;
      hashTable[blockKey] = true;
    }
  }

  return true;
};
