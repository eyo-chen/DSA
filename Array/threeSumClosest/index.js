//////////////////////////////////////////////////////
// *** 3Sum Closest ***
//////////////////////////////////////////////////////
/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

Example 1:
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

Example 2:
Input: nums = [0,0,0], target = 1
Output: 0
 
Constraints:
3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
/*
This problem is very similar to 3sum problem
If really understand that problem, this problem is very easy
I solve this problem by myself
But this problem is also relatively easy
Because we don't need to deal with duplicate case
However, we still do in this problem
Make the solution more optimize

************************************************************
Time compelxity: O(n ^ 2)
Space comelxity: O(1) or O(n)
=> It depeneds on how do we sort the array
*/
var threeSumClosest = function (nums, target) {
  /*
  Just give res large number
  -1000 <= nums[i] <= 1000
  So that we give 10 ^ 5
  */
  let res = 10 ** 5;
  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length; i++) {
    let left = i + 1;
    let right = nums.length - 1;

    while (left < right) {
      const sum = nums[i] + nums[left] + nums[right];
      // if sum is relative close to target, set res to sum
      if (Math.abs(target - sum) < Math.abs(target - res)) {
        res = sum;
      }

      if (sum > target) {
        const val = nums[right];

        // handle duplicate case
        // Note that we use different way to handle duplicate case (Go to 3sum)
        // both work
        while (left < right && val === nums[right]) {
          right--;
        }
      } else if (sum < target) {
        const val = nums[left];

        // same as above
        while (left < right && val === nums[left]) {
          left++;
        }
      }
      // if sum === target, just return target
      else {
        return target;
      }
    }

    // handle duplicate case
    while (i < nums.length - 1 && nums[i] === nums[i + 1]) {
      i++;
    }
  }

  return res;
};
