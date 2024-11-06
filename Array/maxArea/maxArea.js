//////////////////////////////////////////////////////
// *** Container With Most Water ***
//////////////////////////////////////////////////////
/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1
 

Constraints:
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/
/**
 * @param {number[]} height
 * @return {number}
 */
/*
Using two pointers technique
First start on both side
Then we move the one which is smaller inward
For example, [1,8,6,2,5,4,8,3,7]
First ptr1 = 1, ptr2 = 7
in next while-loop, update 1 to 8 because 7 > 1

************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(1)
*/
var maxArea = function (height) {
  let res = 0;
  let ptr1 = 0;
  let ptr2 = height.length - 1;

  while (ptr2 > ptr1) {
    res = Math.max(res, (ptr2 - ptr1) * Math.min(height[ptr1], height[ptr2]));

    if (height[ptr1] > height[ptr2]) {
      ptr2--;
    } else {
      ptr1++;
    }
  }

  return res;
};
