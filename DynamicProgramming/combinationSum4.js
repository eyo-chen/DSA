//////////////////////////////////////////////////////
// *** Combination Sum IV ***
//////////////////////////////////////////////////////
/*
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.

Example 1:
Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.

Example 2:
Input: nums = [9], target = 3
Output: 0
 
Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 1000
All the elements of nums are unique.
1 <= target <= 1000
*/
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
/*
Unlike previous combinationSum problem, 
we're allowed to have duplicate array as our final answer
[1,1,2], [1,2,1], [2,1,1] count as three different answer

************************************************************
n = the length of nums array, t = target

Time compelxity: O(c * t)
=> Because of memoization, we'll have t subproblem at most
=> For each subproblems, we iterate n loop

Space comelxity: O(t)
=> the deepest height of recursive tree
*/
var combinationSum4 = function (nums, target, memo = {}) {
  if (target === 0) return 1;
  if (target < 0) return 0;
  if (memo[target] !== undefined) return memo[target];

  let result = 0;

  for (let i = 0; i < nums.length; i++) {
    result += combinationSum4(nums, target - nums[i], memo);
  }

  memo[target] = result;
  return result;
};

/*
DP table, bottom up
targert = 4, nums = [1,2,3]
0   1   2   3   4 -> table

target = 0
what's the total combination when target is 0?
0 is our base case, if target is 0, the total combination is always 0, which is []
0   1   2   3   4 -> table
1

target = 1
what's the total combination when target is 1?
If target is 1, we can only use 1 in nums array
After using 1, the target becomes 0, asks subproblem
what's the total combination when target is 0? -> 1
0   1   2   3   4 -> table
1   1

target = 2
what's the total combination when target is 2?
If target is 2, we can use 1 or 2 in nums array?
After using 1, the target becomes 1, asks subproblem
what's the total combination when target is 1? -> 1
After using 2, the target becomes 0, asks subproblem
what's the total combination when target is 0? -> 1
1 + 1 = 2

target = 3
what's the total combination when target is 3?
If target is 3, we can use 1, 2 and 3 in nums array?
After using 1, the target becomes 2, asks subproblem
what's the total combination when target is 2? -> 2
After using 2, the target becomes 1, asks subproblem
what's the total combination when target is 1? -> 1
After using 3, the target becomes 0, asks subproblem
what's the total combination when target is 0? -> 1
2 + 1 + 1 = 4

so on and so forth

Each cell of table represents
What's the total combination when target is i ?

DP[i] = Sum(DP[i - j])
where 0 <= j <= i && i includes nums[]

************************************************************
n = the length of nums array, t = target

Time compelxity: O(c * t)

Space comelxity: O(t)
*/

function combinationSum41(nums, target) {
  const table = new Array(target + 1).fill(0);
  table[0] = 1;

  for (let i = 1; i < table.length; i++) {
    for (let n = 0; n < nums.length; n++) {
      const num = nums[n];

      if (i >= num) {
        table[i] += table[i - num];
      }
    }
  }

  return table[target];
}
// console.log(combinationSum4([1, 2, 3, 5], 10));
// console.log(combinationSum41([1, 2, 3, 5], 10));
