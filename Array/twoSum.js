//////////////////////////////////////////////////////
// *** Two Sum ***
//////////////////////////////////////////////////////
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
 
Constraints:
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
/*
The idea is intuitive, just using hashTable
hashTable
key -> value
num -> index

Note that if(hashTable[target - num]) may won't work in some cases
because index may be 0, so it will skip
Using (target - num in hashTable) or hashTable.hasOwnProperty(target - num)

************************************************************
Time compelxity: O(n)
Space comelxity: O(n)
*/
var twoSum = function (nums, target) {
  const hashTable = {};

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];

    if (hashTable.hasOwnProperty(target - num)) {
      return [[hashTable[target - num]], i];
    }

    hashTable[num] = i;
  }
};
