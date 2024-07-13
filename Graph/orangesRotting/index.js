//////////////////////////////////////////////////////
// *** Minimum Passes Of Matrix *** (algoexpert) Rotting Oranges(leetcode)
//////////////////////////////////////////////////////
/*
Write a function that takes in an integer matrix of potentially unequal height and width and returns the minimum number of passes required to convert all negative integers in the matrix to positive integers.

A negative integer in the matrix can only be converted to a positive integer if one or more of its adjacent elements is positive. An adjacent element is an element that is to the left, to the right, above, or below the current element in the matrix. Converting a negative
to a positive simply involves multiplying it by -1

Note that the 0 value is neither positive nor negative, meaning that a 0 can't convert an adjacent negative to a positive.

A single pass through the matrix involves converting all the negative integers that can be converted at a particular point in time.

For example, input matrix is 
[
  [0, -1, -3, 2, 0],
  [1, -2, -5, -1, -3],
  [3, 0, 0, -4, -1]
]

First pass,
[
  [0, -1, +3, 2, 0],
  [1, +2, -5, +1, -3],
  [3, 0, 0, -4, -1]
]

Second pass,
[
  [0, +1, +3, 2, 0],
  [1, +2, +5, +1, +3],
  [3, 0, 0, +4, -1]
]

Third pass
[
  [0, +1, +3, 2, 0],
  [1, +2, +5, +1, +3],
  [3, 0, 0, +4, +1]
]
*/
function minimumPassesOfMatrix(matrix) {
  let [queue, total] = buildQueue(matrix);

  // if total is 0, it means all value in the matrix is positive
  // there's no need to convert, so just return 0
  if (total === 0) {
    return 0;
  }

  let res = 0;

  while (queue.length > 0) {
    let len = queue.length;

    // get the length of current pass
    while (len > 0) {
      const [row, col] = queue.shift();

      /*
      we can only convert if the value is less than 0
      then we 
      1. convert it to positive (by multiply -1)
      2. add it to the queue(to further convert it's adjacent negative element)
      3. minus total (means we convert a negative value to positive value)
      */
      if (row - 1 >= 0 && matrix[row - 1][col] < 0) {
        matrix[row - 1][col] *= -1;
        queue.push([row - 1, col]);
        total--;
      }
      if (row + 1 < matrix.length && matrix[row + 1][col] < 0) {
        matrix[row + 1][col] *= -1;
        queue.push([row + 1, col]);
        total--;
      }
      if (col - 1 >= 0 && matrix[row][col - 1] < 0) {
        matrix[row][col - 1] *= -1;
        queue.push([row, col - 1]);
        total--;
      }
      if (col + 1 < matrix[0].length && matrix[row][col + 1] < 0) {
        matrix[row][col + 1] *= -1;
        queue.push([row, col + 1]);
        total--;
      }

      len--;
    }

    res++;
  }

  // if we can't convert all negative to positive, return -1 (total !== 0 )
  /*
  It's important to res - 1 here
  because our last pass isn't convert -1 to 1, we only clear out the queue
  For example,
  At the end of the queue, it's [2,3] (just example)
  And all the value is positive
  So we only dequeue, remove element from the queue
  then res++
  and this plus one is not true, we're not converting
  so we need -1 here
  */
  return total === 0 ? res - 1 : -1;
}

function buildQueue(matrix) {
  const queue = [];
  let total = 0;

  for (let r = 0; r < matrix.length; r++) {
    for (let c = 0; c < matrix[0].length; c++) {
      // how many positive in the matrix is gonna to convert other negative?
      if (matrix[r][c] > 0) {
        queue.push([r, c]);
      }
      // how many negative in the matrix is gonna to convert?
      if (matrix[r][c] < 0) {
        total++;
      }
    }
  }

  return [queue, total];
}
