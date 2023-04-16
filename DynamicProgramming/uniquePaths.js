//////////////////////////////////////////////////////
// *** Unique Paths ***(Number of Ways To Traverse A Matrix)
//////////////////////////////////////////////////////
/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
/*
Given an integer value m (rows of a matrix), and an interger value n (columns of a matrix), return the total possible unique, simple, paths from the top-left of the matrix to the bottom-right with restricted moves.

You may only make one of these moves at each position:
Down 1 cell
Right 1 cell

Example
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
*/
/*
Bottom-Up approach

One key to note that if it's 2x2 grid
like this
--------------
| 0,0 | 0,1 |
--------------
| 1,0 | 1,1 |
--------------
for the unique path for (1,1), is the sum of (0,1) and (1,0)

The very top row and very left column are always be 1 

***********************************************************************
Time: O(m * n)
We have to traverse every each spot, which is what the nested for-loop doing

Space: O(m * n)
*/
// m => rows, n => cols
function uniquePaths(m, n) {
  const matrix = [];

  for (let i = 0; i < m; i++) {
    const rows = new Array(n);
    rows.fill(0);

    matrix.push(rows);
  }

  // very top row
  for (let i = 0; i < m; i++) {
    matrix[i][0] = 1;
  }

  // the very right column
  for (let i = 0; i < n; i++) {
    matrix[0][i] = 1;
  }

  // for the unique path for (1,1), is the sum of (0,1) and (1,0)
  for (let row = 1; row < m; row++) {
    for (let col = 1; col < n; col++) {
      matrix[row][col] = matrix[row - 1][col] + matrix[row][col - 1];
    }
  }

  return matrix[m - 1][n - 1];
}

/*
  Top-Down approach with memoization
  
  If we're given [2,3], then we have to figure out that there are only two ways to go to this spot
  Which is from [1,3] and [2,2] because we can only either go right or go down
  So if we're asked what's the unique paths for [2,3], the answer is unique paths of [1,3] + unique paths of [2,2]
  
  First, have to think about the base case
  1. whenever we have the 0, like [0,1], [10,0], it means there's no way for this path, so just return 0
  2. if we're given [1,1], then just return 1
  
  Second, have to think about what's our decision for each recursive calls
  According to the description of this problem, we know that we can only either go "down" or go "right"
  This is our decision
  
  The recursive tree would look like this
                                       [2,3]
                        [1,3]                       [2,2]
                [0,3]        [1,2]             [1,2]       [2,1]
                        [0,2]      [1,1]    [0,2]   [1,1] [1,1]   [2,0]
  Simple, for each recursive calls, we either go down or go right, which is row - 1 or col - 1
  
  ***********************************************************************
  Before optimization
  Time: O(2 ^ max(m,n) )
  Because braching factor is 2, and the deepest recursive tree would be max(m,n)
  
  Space: O(max(m,n))
  
  
  After optimization
  Time: O(m * n)
  Because we won't have any recursive call when we useing memo object
  For example, m = 3, n = 2
  For m, we could have {0,1,2,3}
  For n, we could have {0,1,2}
  m * n would be the total maximum recursive calls
  
  Space: O(max(m,n))
  
  Note that the reason with 2 keys in memo (path1 and path2)
  is because the answer for [1,2] and [2,1] are the same
  So our actual time complexity would be O(m*n / 2)
  */
function uniquePaths1(m, n, memo = {}) {
  const path1 = `${m}-${n}`;
  const path2 = `${n}-${m}`;

  if (path1 in memo) return memo[path1];
  if (path2 in memo) return memo[path2];

  // base case
  if (m === 0 || n === 0) return 0;
  if (m === 1 && n === 1) return 1;

  // either go right or go down
  memo[path1] = uniquePaths1(m - 1, n, memo) + uniquePaths1(m, n - 1, memo);
  memo[path2] = memo[path1];

  return memo[path1];
}

// console.log(uniquePaths(8, 4));
// console.log(uniquePaths1(8, 4));
