//////////////////////////////////////////////////////
// *** Rotate Image ***
//////////////////////////////////////////////////////
/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.


Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Constraints:
n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/
/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
/*
The idea of this problem is not that hard
The hardest part is how to get the right index

5  1  9  11
2  4  8  10
13 3  6  7
15 14 12 16

There are two layers in this matrix
5  1  9  11
2        10
13       7
15 14 12 16
=> first layer
=> In this layer, we have to do 3 times rotation
=> 1. 5 -> 11, 11 -> 16, 16 -> 15
=> 2. 1 -> 10, 10 -> 12, 12 -> 13
=> 3. 9 -> 7, 7 -> 14, 14 -> 2

   4  8  
   3  6  
=> second layer
=> In this layer, we only do 1 time rotation

************************************************************
Time compelxity: O(n ** 2)
Space comelxity: O(1)
*/
var rotate = function (matrix) {
  const SIZE = matrix.length - 1;

  // just have to see the pattern
  const LAYER = Math.trunc(matrix.length / 2);

  // the outer for-loop is loop through layer by layer (inward)
  for (let i = 0; i < LAYER; i++) {
    /*
    the inner for-loop is loop through four points to rotate
    In first layer, we have three points to move(0 ~ 2)
    In second layer, we only have one point to move(1 ~ 1)

    SIZE - i is kind of hard to come up with
    */
    for (let j = i; j < SIZE - i; j++) {
      const top = matrix[i][j];
      const right = matrix[j][SIZE - i];
      const bottom = matrix[SIZE - i][SIZE - j];
      const left = matrix[SIZE - j][i];

      matrix[i][j] = left;
      matrix[j][SIZE - i] = top;
      matrix[SIZE - i][SIZE - j] = right;
      matrix[SIZE - j][i] = bottom;
    }
  }
};
