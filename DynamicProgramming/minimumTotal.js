//////////////////////////////////////////////////////
// *** Triangle ***(Minimum Weight Path In A Triangle)
//////////////////////////////////////////////////////
/*
Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
*/
/*
The idea is very non-intutive

Instead of going from top to bottom, we go from bottom to top
The idea is like this 
1. create an cache array for storing the element of last layer(row)(4th)([4, 1, 8, 3])
2. go upward to above layer(3th) [6, 5, 7]
3. loop through each element to find the each min total, we can either go left or right
   For 6, the total can be 6 + 4 and 6 + 1
   the min total would be 7, so we store 7 at 0 index of cache arr
   the cache arr would look like this
   [7, 1, 8, 3]

   For 5, the total can be 5 + 1 and 5 + 8
   the min total would be 6, so we store 6 at 1 index of cache arr
   the cache arr would look like this
   [7, 6, 8, 3]

   For 7, the total can be 7 + 8 and 7 + 3
   the min total would be 10, so we store 10 at 2 index of cache arr
   the cache arr would look like this
   [7, 6, 10, 3]

   the last element(3) doesn't matter anymore because next layer only has two elements
   and we will only touch the first three element in cache arr

   keep doing like this.....

   layer 2 => [3, 4]

   For 3, the total can be 3 + 7 and 3 + 6
   the min total would be 9, so we store 10 at 0 index of cache arr
   the cache arr would look like this
   [9, 6, 10, 3]

   For 4, the total can be 4 + 6 and 4 + 10
   the min total would be 10, so we store 10 at 1 index of cache arr
   the cache arr would look like this
   [9, 10, 10, 3]

   layer 1 => [2]

   For 2, the total can be 2 + 9 and 2 + 10
   the min total would be 11, so we store 10 at 1 index of cache arr
   the cache arr would look like this
   [11, 10, 10, 3]

4. As we could see, the first element of final cache array will always be our final answer
5. And because we store the current index at the cache array, we can keep going left and right for next element 
   For example, [4, 1, 8, 3], current element is 6, and min total is 7, then we store 7 at 0 index
   [7, 1, 8, 3], and next element can still go left(1) and right(8)
(this part is really hard to come up)


***********************************************************************
Time: O(n)
where n is the total element of triangle

Space: O(n)
store the largest row of triangle


///////////////////////////////////////////////////////////////////////
Takeaway
At the very beginning of looking at this problem, it's normal to see this problem from top to bottom
and try to find a way to solve the problem from top to bottom
It may work

BUT
it would be great to see the problem from bottom to top
it's simpler simetimes
*/
/**
 * @param {number[][]} triangle
 * @return {number}
 */
function minimumTotal(triangle) {
  const cacheArr = [];

  // store the element of last row in cache array
  for (let i = 0; i < triangle[triangle.length - 1].length; i++) {
    cacheArr.push(triangle[triangle.length - 1][i]);
  }

  // begin at the last two row of triangle
  for (let i = triangle.length - 2; i >= 0; i--) {
    const curRow = triangle[i];

    // loop through each element of row
    for (let j = 0; j < curRow.length; j++) {
      const curVal = curRow[j];

      const goLeft = curVal + cacheArr[j];
      const goRight = curVal + cacheArr[j + 1];

      cacheArr[j] = Math.min(goLeft, goRight);
    }
  }

  return cacheArr[0];
}

// bruth force solution
/*
  ***********************************************************************
  Time: O(h ^ 2)
  where h is the height of tirangle
  and for each recursive calls, we only have two decisions, go curIndex or go curIndex + 1
  
  Space: O(h)
  where h is the height of tirangle
  
  Note that this recusive does NOT have overlaping sub problem
  
     2
    3 4
   6 5 7
  4 1 8 3
  
  At first glance, it may seems we have
  for example, 
  for 3, we can choose 6 or 5
  for 4, we can choose 5 or 7
  see that we choose 5 twice
  and 5 will choose 1 and 8 twice
  
  but the tmpTotal is always different
  because for each unique path, we have different tmpTotal
  */
function minimumTotal1(triangle) {
  return minimumTotalRecursiveHelper(triangle, 1, 0, triangle[0][0], Infinity);
}

function minimumTotalRecursiveHelper(
  triangle,
  row,
  curIndex,
  tmpTotal,
  tmpRes
) {
  if (row >= triangle.length) {
    tmpRes = Math.min(tmpRes, tmpTotal);
    return tmpRes;
  }

  const curRowArr = triangle[row];

  const goLeft = tmpTotal + curRowArr[curIndex];
  const goRight = tmpTotal + curRowArr[curIndex + 1];

  const leftRes = minimumTotalRecursiveHelper(
    triangle,
    row + 1,
    curIndex,
    goLeft,
    tmpRes
  );

  const rightRes = minimumTotalRecursiveHelper(
    triangle,
    row + 1,
    curIndex + 1,
    goRight,
    tmpRes
  );

  return Math.min(rightRes, leftRes);
}

// console.log(
//   minimumTotal([[2], [3, 4], [6, 5, 7], [4, 1, 8, 3], [4, 12, 32, 43, 11]])
// );
// console.log(minimumTotal1([[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]));

// console.log(minimumTotal([[1], [2, 3]]));
// console.log(minimumTotal1([[1], [2, 3]]));
