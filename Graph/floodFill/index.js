//////////////////////////////////////////////////////
// *** Flood Fill ***
//////////////////////////////////////////////////////
/*
An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.

Return the modified image after performing the flood fill.

Example 1:
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

Example 2:
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
Output: [[2,2,2],[2,2,2]]

Constraints:
m == image.length
n == image[i].length
1 <= m, n <= 50
0 <= image[i][j], newColor < 216
0 <= sr < m
0 <= sc < n
*/
/**
 * @param {number[][]} image
 * @param {number} sr
 * @param {number} sc
 * @param {number} newColor
 * @return {number[][]}
 */
/*
This problem is pretty straightforward
It's just DFS, using queue
Just see the code

************************************************************
r = row, c = col
Time: O(r * c)
Space: O(r * c)
=> seen object, and queue
*/
var floodFill = function (image, sr, sc, newColor) {
  const seen = {};
  const queue = [[sr, sc]];
  const index = image[sr][sc];

  while (queue.length > 0) {
    const [row, col] = queue.shift();
    const key = `${row}-${col}`;

    // check if the cur element is valid
    // 1) cannot out of the bound
    // 2) cannot be seen before
    // 3) the color is as same as input
    if (!checkValid(image, row, col, index, seen, key)) continue;

    image[row][col] = newColor;
    seen[key] = true;

    // just add all the adjacent element, checkValid will help us in later iteration
    queue.push([row + 1, col]);
    queue.push([row - 1, col]);
    queue.push([row, col + 1]);
    queue.push([row, col - 1]);
  }

  return image;
};

function checkValid(arr, row, col, index, seen, key) {
  return (
    row >= 0 &&
    row <= arr.length - 1 &&
    col >= 0 &&
    col <= arr[0].length - 1 &&
    arr[row][col] === index &&
    !seen[key]
  );
}

/*
Same logic, but using DFS, iterative, stack
*/
var floodFill = function (image, sr, sc, newColor) {
  const seen = {};
  const stack = [[sr, sc]];
  const index = image[sr][sc];

  while (stack.length > 0) {
    const [row, col] = stack.pop();
    const key = `${row}-${col}`;

    if (!checkValid(image, row, col, index, seen, key)) continue;

    image[row][col] = newColor;
    seen[key] = true;

    stack.push([row + 1, col]);
    stack.push([row - 1, col]);
    stack.push([row, col + 1]);
    stack.push([row, col - 1]);
  }

  return image;
};

/*
DFS, recursive
*/
var floodFill = function (image, sr, sc, newColor) {
  recursiveHelper(image, sr, sc, image[sr][sc], {}, newColor);
  return image;
};

function recursiveHelper(image, row, col, index, seen, newColor) {
  const key = `${row}-${col}`;

  if (!checkValid(image, row, col, index, seen, key)) return;

  image[row][col] = newColor;
  seen[key] = true;
  recursiveHelper(image, row + 1, col, index, seen, newColor);
  recursiveHelper(image, row - 1, col, index, seen, newColor);
  recursiveHelper(image, row, col + 1, index, seen, newColor);
  recursiveHelper(image, row, col - 1, index, seen, newColor);
}
