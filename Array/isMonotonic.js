//////////////////////////////////////////////////////
// *** Monotonic Array ***
//////////////////////////////////////////////////////
/*
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is monotone decreasing if for all i <= j, nums[i] >= nums[j].

Given an integer array nums, return true if the given array is monotonic, or false otherwise.


Example 1:
Input: nums = [1,2,2,3]
Output: true

Example 2:
Input: nums = [6,5,4,4]
Output: true

Example 3:
Input: nums = [1,3,2]
Output: false
 
Constraints:
1 <= nums.length <= 105
-105 <= nums[i] <= 105
*/
/**
 * @param {number[]} nums
 * @return {boolean}
 */
/*
This problem is easy, just look at the code and understand

************************************************************
Time complexity: O(n)
Space complexity: O(1)
*/
var isMonotonic = function (nums) {
  let index = null;

  for (let i = 1; i < nums.length; i++) {
    // if index is null, it means we haven't found the pattern(increasing or descending)
    if (!index) {
      // increasing pattern
      if (nums[i] > nums[i - 1]) {
        index = 1;
      }
      // descending pattern
      else if (nums[i] < nums[i - 1]) {
        index = 2;
      }
    } else {
      // if it's increasing pattern, but nums[i] < nums[i - 1], it's false
      if (index === 1 && nums[i] < nums[i - 1]) {
        return false;
      }

      // if it's descending pattern, but nums[i] > nums[i - 1], it's false
      if (index === 2 && nums[i] > nums[i - 1]) {
        return false;
      }
    }
  }

  return true;
};

/*
This problem is easy, just look at the code and understand

************************************************************
Time complexity: O(n)
Space complexity: O(1)
*/
var isMonotonic = function (nums) {
  let inc = true;
  let dec = true;

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] > nums[i - 1]) {
      dec = false;
    }
    if (nums[i] < nums[i - 1]) {
      inc = false;
    }
  }

  // if one of them is true, then it's true
  // only if both of them are false, return fasle
  // it means both of them have ever been increasing and descending pattern
  return inc || dec;
};
