//////////////////////////////////////////////////////
// *** Trapping Rain Water ***
//////////////////////////////////////////////////////
/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
*/
/*
It's easy to solve this problem by seeing the graph

In order to count the total traped water amount
We need to individually count each traper water amount for each index

How do we know the each traper water amount for each index?
We need two things
1. The left maximum height 
2. The right maximum height 

We just need to know the minimum between both of these
And subtract itself to get the traper water amount for each index
Why?
For example, [4,1,7]
How do we know the traper water amount for index 1
The maximum height for that index is determined by 4, not 7
And we just 4 - 1 to get the answer

***********************************************************************
Time: O(n)
Space: O(n)
*/
/**
 * @param {number[]} height
 * @return {number}
 */
function trap(heights) {
  const leftMax = new Array(heights.length).fill(0);
  const rightMax = new Array(heights.length).fill(0);
  leftMax[0] = heights[0];
  rightMax[heights.length - 1] = heights[heights.length - 1];
  let traperdWaterTotal = 0;
  let leftMaxNum = heights[0];
  let rightMaxNum = heights[heights.length - 1];

  // build the array represents the maximum left height for each index
  for (let i = 1; i < heights.length; i++) {
    leftMax[i] = Math.max(leftMaxNum, heights[i - 1]);
    leftMaxNum = Math.max(leftMaxNum, leftMax[i]);
  }

  // build the array represents the maximum right height for each index
  for (let i = heights.length - 2; i >= 0; i--) {
    rightMax[i] = Math.max(rightMaxNum, heights[i + 1]);
    rightMaxNum = Math.max(rightMaxNum, rightMax[i]);
  }

  /*
  count for each minimum and count the answer
  Note that if answer is negative, we just set it to 0
  because there's no way traped water is negative
  */
  for (let i = 0; i < heights.length - 1; i++) {
    const minHeightBetween = Math.min(rightMax[i], leftMax[i]);
    traperdWaterTotal +=
      minHeightBetween - heights[i] > 0 ? minHeightBetween - heights[i] : 0;
  }

  return traperdWaterTotal;
}

/*
Only build two arrays, and two iterations
but still O(n)
*/
function trap1(heights) {
  const table = new Array(heights.length).fill(0);
  let leftMaxNum = heights[0];
  let rightMaxNum = heights[heights.length - 1];
  let traperdWaterTotal = 0;

  for (let i = 1; i < heights.length; i++) {
    table[i] = Math.max(leftMaxNum, heights[i - 1]);
    leftMaxNum = Math.max(leftMaxNum, table[i]);
  }

  for (let i = heights.length - 2; i >= 0; i--) {
    rightMaxNum = Math.max(rightMaxNum, heights[i + 1]);

    const minHeightBetween = Math.min(rightMaxNum, table[i]);

    traperdWaterTotal +=
      minHeightBetween - heights[i] > 0 ? minHeightBetween - heights[i] : 0;
  }

  return traperdWaterTotal;
}

// console.log(trap1([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));
