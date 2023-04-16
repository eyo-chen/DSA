//////////////////////////////////////////////////////
// *** House Robber ***(Max Subset Sum No Adjacent)
//////////////////////////////////////////////////////
/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
*/
/*
The main approach of solving this problem is not remembering the formula or pattern
That's not helpful
Instead, always try to start with the small sub problem and try to firgure out the pattern

EX1
input: [1]
=> the output is just one, we have NO choices
=> It seems this may be our base case

EX2
input: [1,2]
=> in this input, we could either choose 1 or 2, which means now we have two choices
=> we choose the maximun one which is 2

EX3
input: [1,2,3]
=> If we closely look this input array, this just the EX2 array concating with [3]
=> which means we've solved two sub problems
=> for [1], the answer is 1, it means up to the index 0, the maximum is 1
=> for [1,2], the answer is 2, it means up to this point(index 1), the maximum is 2
=> for [1,2,3], for the index 2, we now have two choices
   1) include itself, so we can add nums[i - 2] -> 3 + nums[i - 2]
   2) exclude itself, so only choose 2, nums[i - 1]
=> the value is 3 + 1
=> What does this value represent at this index?
   It means up to this point where index is 2, the maximum value is 4, it may be include itself or not, but it doesn't matter

EX4
input: [1,2,3,4]
=> for [1,2,3,4], for the index 3, we now have two choices
   1) include itself, so we can add nums[i - 2], 4 + nums[i - 2]
   2) exclude itself, soe only choose nums[i - 1]
      => nums[i - 1] = nums[2], which is our solved sub problem, and the value is 4


We may start seeing the pattern of this problem
For the input [1,2,3,4,5]
The final problem is up to the last index(index 4), what's the maximum sum value of non-adjacent?
The answer could be
1) the answer of "up to the last index(index 3), what's the maximum sum value of non-adjacent?"
Which means the subproblm [1,2,3,4]
Because once choosing 4, we're not allowed to choose 5

2) the answer of "up to the last index(index 2), what's the maximum sum value of non-adjacent?" + "5"
Which means the sub problem [1,2,3] + 5

so on and so forth

The pattern would be
ans[0 : i] = Math.max(ans[0: i - 1], ans[0: i - 2] + ans[i])

the solution is going from start to the end
so at the end of value is our final answer

************************************************************
n = the length of array

Time compelxity: O(n)

Space comelxity: O(n)
*/
/**
 * @param {number[]} nums
 * @return {number}
 */
function rob(nums) {
  // small edge cases
  if (nums.length === 0) return 0;
  if (nums.length === 1) return nums[0];

  const table = new Array(nums.length).fill(0);

  // base case
  table[0] = nums[0];
  table[1] = Math.max(nums[0], nums[1]);

  // main logic
  for (let i = 2; i < nums.length; i++) {
    table[i] = Math.max(nums[i] + table[i - 2], table[i - 1]);
  }

  return table[nums.length - 1];
}
/*
  How to optimize?
  Time complexity?
  => It seems no way
  So try to optimize space complexity
  
  
  The pattern
  ans[0 : i] = Math.max(ans[0: i - 1], ans[0: i - 2] + ans[i])
  
  For each index i, we only care i, i - 1, and i - 2
  So we can use pointer, and not use array
  */
function rob1(nums) {
  if (nums.length === 0) return 0;
  if (nums.length === 1) return nums[0];

  let p1 = nums[0];
  let p2 = Math.max(nums[0], nums[1]);

  for (let i = 2; i < nums.length; i++) {
    const current = Math.max(nums[i] + p1, p2);

    // update pointer
    p1 = p2;
    p2 = current;
  }

  return p2;
}
// console.log(rob([]));
// console.log(rob1([]));
