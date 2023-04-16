//////////////////////////////////////////////////////
// *** Maximum Subarray ***
//////////////////////////////////////////////////////
/**
 * @param {number[]} nums
 * @return {number}
 */
/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/
/*
Bruth force solution
Try to find all the contiguous subarray, and get the best answer

************************************************************
n = the legnth of array
Time compelxity: O(n ^ 2)
=> Nested for loop

Space comelxity: O(1)
*/
function maxSubArray(nums) {
  // simple base case
  if (nums.length === 1) return nums[0];
  let res = nums[0];

  for (let i = 0; i < nums.length; i++) {
    // tmpMax helps us to keep track the total value of every subarray
    let tmpMax = 0;

    for (let k = i; k < nums.length; k++) {
      tmpMax += nums[k];

      // find the global best answer
      res = Math.max(res, tmpMax);
    }
  }

  return res;
}

/*
  Using the concept of dynamic programming
  
  For each cell, the subproblem is what's the Maximum Subarray at this index?
  
  For example, [-2,1,-3,4,-1,2,1,-5,4]
  
  -2   1   -3   4   -1   2   1   -5   4
                .
  what's the Maximum Subarray at the index 4?
  
  -2   1   -3   4   -1   2   1   -5   4
                     .
  what's the Maximum Subarray at the index 5?
  
  How do we find the Maximum Subarray at the any given index?
  
  For each subproblem, we can
  1. Start the new subarray from this given index
  2. Extend the previous subarray
  
  How can we choose either one?
  Because the question is asking Maximum, so we compare the value
  If (my value) < (my value) + (previous Maximum Subarray value),
  then we want to extend the maximum subarray
  But
  If (my value) > (my value) + (previous Maximum Subarray value)
  then we want to start the new subarray from this given index
  
  i = 0 is base case
  At this point, the maximum subarray value is -2
  
  i = 1
  1 > 1 + -2 (prev maximum subarray value)
  => start new subarray
  -2   1   -3   4   -1   2   1   -5   4
       .
  
  i = 2
  -3 < -3 + 1 (prev maximum subarray value)
  => extend the prev subarray
  -2   1   -2   4   -1   2   1   -5   4
            .
  
  i = 3
  4 > 4 + -2 (prev maximum subarray value)
  => start new subarray
  -2   1   -2   4   -1   2   1   -5   4
                .
  
  i = 4
  -1 < -1 + 4 (prev maximum subarray value)
  => extend the prev subarray
  -2   1   -2   4    3   2   1   -5   4
                     .
  
  i = 5
  2 < 2 + 3 (prev maximum subarray value)
  => extend the prev subarray
  -2   1   -2   4    3   5   1   -5   4
                         .
  
  i = 6
  1 < 1 + 5 (prev maximum subarray value)
  => extend the prev subarray
  -2   1   -2   4    3   5   6   -5   4
                             .
  
  i = 7
  -5 < -5 + 6 (prev maximum subarray value)
  => extend the prev subarray
  -2   1   -2   4    3   5   6   1   4
                                 .
  
  i = 8
  4 < 4 + 1 (prev maximum subarray value)
  => extend the prev subarray
  -2   1   -2   4    3   5   6   1   5
                                     .
  
  As we could see, we mutated the input array, and store the maximum subarray value at each index
  Again, now the value at any given index represent the maximum subarray value
  And we find the best answer is 6
  
  The problem is mutating the input array, we can just create another DP table
  
  ************************************************************
  n = the legnth of array
  Time compelxity: O(n)
  
  Space comelxity: O(1)
  => It would be O(n) if building DP table
  */
function maxSubArray1(nums) {
  if (nums.length === 1) return nums[0];

  let res = nums[0];

  for (let i = 1; i < nums.length; i++) {
    // extend the prev subarray
    if (nums[i] < nums[i] + nums[i - 1]) {
      nums[i] = nums[i] + nums[i - 1];
    }

    res = Math.max(res, nums[i]);
  }

  return res;
}
/*
  This does NOT mutate the input array

  ************************************************************
  n = the legnth of array
  Time compelxity: O(n)
  
  Space comelxity: O(1)
*/
function maxSubArray2(nums) {
  if (nums.length === 1) return nums[0];

  let res = nums[0];
  let prevMaxSum = nums[0];

  for (let i = 1; i < nums.length; i++) {
    // preMaxSum helps us to keep track do we want to extend the prev subarray or start the new subarray
    // nums[i] => start the new subarray
    // prevMaxSum + nums[i] => extend the prev subarray
    prevMaxSum = Math.max(nums[i], prevMaxSum + nums[i]);

    res = Math.max(res, prevMaxSum);
  }

  return res;
}

console.log(maxSubArray([2, 3, -2, 4]));
// console.log(maxSubArray1([5, 4, -1, 7, 8]));
// console.log(maxSubArray2([5, 4, -1, 7, 8]));
