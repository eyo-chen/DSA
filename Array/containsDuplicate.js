//////////////////////////////////////////////////////
// *** Contains Duplicate ***
//////////////////////////////////////////////////////
/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct. 

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4] 
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
 
Constraints:
1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
/**
 * @param {number[]} nums
 * @return {boolean}
 */
/*
************************************************************
Time compelxity: O(n ^ 2)
Space comelxity: O(1)
*/
var containsDuplicate = function (nums) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[j] === nums[i]) {
        return true;
      }
    }
  }

  return false;
};

/*
************************************************************
Time compelxity: O(n * log(n))
Space comelxity: O(1)
*/
var containsDuplicate = function (nums) {
  nums.sort((a, b) => a - b);

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] === nums[i - 1]) {
      return true;
    }
  }

  return false;
};

/*
************************************************************
Time compelxity: O(n)
Space comelxity: O(n)
*/
var containsDuplicate = function (nums) {
  const hashTable = {};

  for (const num of nums) {
    if (hashTable[num]) {
      return true;
    } else {
      hashTable[num] = true;
    }
  }

  return false;
};
