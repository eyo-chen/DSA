//////////////////////////////////////////////////////
// *** Spiral Matrix ***
//////////////////////////////////////////////////////
/*
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 
Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/
/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var spiralOrder = function (matrix) {
  const res = [];
  let rowBegin = 0;
  let rowEnd = matrix.length - 1;
  let colBegin = 0;
  let colEnd = matrix[0].length - 1;

  while (rowBegin <= rowEnd && colBegin <= colEnd) {
    for (let c = colBegin; c <= colEnd; c++) {
      res.push(matrix[rowBegin][c]);
    }
    rowBegin++;

    for (let r = rowBegin; r <= rowEnd; r++) {
      res.push(matrix[r][colEnd]);
    }
    colEnd--;

    if (rowBegin <= rowEnd) {
      for (let c = colEnd; c >= colBegin; c--) {
        res.push(matrix[rowEnd][c]);
      }

      rowEnd--;
    }

    if (colBegin <= colEnd) {
      for (let r = rowEnd; r >= rowBegin; r--) {
        res.push(matrix[r][colBegin]);
      }

      colBegin++;
    }
  }

  return res;
};
