//////////////////////////////////////////////////////
// *** Set Matrix Zeroes ***
//////////////////////////////////////////////////////
/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.

Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 
Constraints:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1
 
Follow up:
A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/
/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
/*
This is non-optimize solution, but easy to understand

1. Create two array
   One is for row, and the other one is for col
   Note that do NOT confused with them
   For example, input matrix
     1  1  1  1 -> firstRow
    1 [1, 2, 3, 4],
    1 [5, 0, 7, 8],
    1 [0, 10, 11, 12],
    1 [13, 14, 15, 0],
    -> firstCol
    
    These two arrays help us to remember which row and col has to be set to 0
    So that we don't need to do repeated work
    For example, matrix[1][1] = 0
    Then we set it's col and row to 0
    Later, matrix[2][1] = 0,
    then we have to set col and row to 0
    Do you see that?
    col1 set to 0 twice, which is bad

    We just need these two arrays to help us to remember which row and col has to be set to 0
    After iteration, we just respecitvely loop through two arrays to set it's row and col to 0

2. Iterate through all the element in the matrix
   If matrix[r][c] === 0, 
   set firstRow[c] = 0, which means entire col c has to be set to 0
   set firstCol[r] = 0, which means entire row r has to be set to 0 

3. Respecitvely loop through two arrays to set it's row and col to 0

************************************************************
Time compelxity: O(n * m)
Space comelxity: O(n + m)
*/
var setZeroes = function (matrix) {
  // firstRow indicate which col has to be to set 0
  // image firstRow is the row above the true first row
  const firstRow = new Array(matrix[0].length).fill(1);

  // firstCol indicate which row has to be to set 0
  // image firstCol is the col lefter the true first col
  const firstCol = new Array(matrix.length).fill(1);
  const ROW = matrix.length;
  const COL = matrix[0].length;

  for (let r = 0; r < ROW; r++) {
    for (let c = 0; c < COL; c++) {
      if (matrix[r][c] === 0) {
        firstRow[c] = 0;
        firstCol[r] = 0;
      }
    }
  }

  // If firstRow[c] === 0, then entire col(c) has to be set to 0
  for (let c = 0; c < COL; c++) {
    if (firstRow[c] === 0) {
      for (let r = 0; r < ROW; r++) {
        matrix[r][c] = 0;
      }
    }
  }

  // If firstCol[r] === 0, then entire row(r) has to be set to 0
  for (let r = 0; r < ROW; r++) {
    if (firstCol[r] === 0) {
      for (let c = 0; c < COL; c++) {
        matrix[r][c] = 0;
      }
    }
  }
};

/*
This is optimize solution

The main idea is very similar to previous one
But we're not gonna create two arrays
Instead, we use first row and first col as our array
To help us to remember which col and row have to set to 0
But tow things to note
1. The matrix[0][0] indicate both first col and first row
So that we need to another variable to hold
    [1, 2, 3, 4], -> firstRow, indicate which col has to be to set 0
    [5, 0, 7, 8],
    [0, 10, 11, 12],
    [13, 14, 15, 0],
    -> firstCol, indicate which row has to be to set 0

    variable firstCol, indicate if firstCol has to be set to 0

2. After first iteration, we need another iteration
   Because now first row and col is our array to help us to remember
   So that we start at matrix[1][1]
   Then we iterate through remaining matrix
   But after this iteration, we haven't check the first row and first col
   So that we need another two for-loops to check if we need to another work
   In this case,
   matrix[0][0] indicates first row
   If it's 0, then we're gonna set entrie first row to 0
   variable firstCol, indicates first col
   If it's true, then we're gonna set entire first col to 0

The main idea is 
1. Iterte through all element in the matrix
   If it's 0, then we're gonna set it's [0][c] and [r][0] to 0
   For example, if matrix[3][2] = 0,
   then we're gonna set 
   matrix[3][0] = 0, which means entire row has to be 0
   matrix[0][2] = 0, which means entire col has to be 0
   In short, we just set the border cell to zero to indicate which col and row have to be 0 after the iteration

2. Iterate through all element in the matrix, expect for first col an row
   Both starts at 1
   Then we check if we have to set this element to 0
   matrix[r][0] === 0 || matrix[0][c] === 0

3. We also need to check if first row need to be set to 0
   We're using first element as index
   matrix[0][0]

4. Note that the matrix[0][0] will both indicate row 0 and col 0
   So that we need another variable
   In this case, we're using firstCol to tell us if first col need to be set to 0
   If it's true, then we need to see entire first col to 0

Both leetcode solution and https://www.youtube.com/watch?v=T41rL0L3Pnw
Have great detail
************************************************************
Time compelxity: O(n * m)
Space comelxity: O(1)
*/
var setZeroes = function (matrix) {
  const ROW = matrix.length;
  const COL = matrix[0].length;
  let firstCol = false;

  for (let r = 0; r < ROW; r++) {
    /*
    Note that c is always 0
    it means if there's any [row][0] is 0
    then we're gonna set entire first col to 0 later
    */
    if (matrix[r][0] === 0) {
      firstCol = true;
    }
    for (let c = 1; c < COL; c++) {
      if (matrix[r][c] === 0) {
        matrix[0][c] = 0;
        matrix[r][0] = 0;
      }
    }
  }

  for (let r = 1; r < ROW; r++) {
    for (let c = 1; c < COL; c++) {
      if (matrix[r][0] === 0 || matrix[0][c] === 0) {
        matrix[r][c] = 0;
      }
    }
  }

  // matrix[0][0] indicate first row
  // if it's 0, then we're gonna set entire first row to 0
  if (matrix[0][0] === 0) {
    for (let c = 0; c < COL; c++) {
      matrix[0][c] = 0;
    }
  }

  // if firstCol is true, then we're gonna set entire first col to 0
  if (firstCol) {
    for (let r = 0; r < ROW; r++) {
      matrix[r][0] = 0;
    }
  }
};

/*
Another solution I came up with by myself

The idea is simple,
1. Iterate through all element in the matrix
   If it's 0, then invoke mutateToZero

2. mutateToZero only does two things
   Set it's entire row to "."
   Set it's entire col to "."
   The reason we don't set to 0 is that 
   If martrix[1][1] = 0, and matrix[1][2] is not
   If we just set entire first row to 0
   Later, when we hit matrix[1][2], it's 0
   Then we're gonna set entire col 2 to 0, which is not the problem asks
   Just try to set to 0, and see the result

3. After first iteration, we just iterate through all the element again
   If it's ".", just set to 0

************************************************************
Time compelxity: O(n * m)
Space comelxity: O(1)
*/
var setZeroes = function (matrix) {
  for (let r = 0; r < matrix.length; r++) {
    for (let c = 0; c < matrix[0].length; c++) {
      if (matrix[r][c] === 0) {
        mutateToZero(matrix, r, c);
      }
    }
  }

  for (let r = 0; r < matrix.length; r++) {
    for (let c = 0; c < matrix[0].length; c++) {
      if (matrix[r][c] === '.') {
        matrix[r][c] = 0;
      }
    }
  }
};

function mutateToZero(matrix, row, col) {
  for (let r = 0; r < matrix.length; r++) {
    if (matrix[r][col] !== 0) {
      matrix[r][col] = '.';
    }
  }

  for (let c = 0; c < matrix[0].length; c++) {
    if (matrix[row][c] !== 0) {
      matrix[row][c] = '.';
    }
  }
}
