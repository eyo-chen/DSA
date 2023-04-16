//////////////////////////////////////////////////////
// *** Find All Duplicates in an Array ***
//////////////////////////////////////////////////////
/*
Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.

You must write an algorithm that runs in O(n) time and uses only constant extra space.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Example 2:
Input: nums = [1,1,2]
Output: [1]

Example 3:
Input: nums = [1]
Output: []

Constraints:
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
Each element in nums appears once or twice.
*/
/**
 * @param {number[]} nums
 * @return {number[]}
 */
/*
This problem is very similar to findDisappearedNumbers
And this solution is basically the same as the second solution of findDisappearedNumbers
Just go check there to see the detail explanation

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
var findDuplicates = function (nums) {
  const output = [];

  for (let i = 0; i < nums.length; i++) {
    while (nums[i] !== i + 1 && nums[nums[i] - 1] !== nums[i]) {
      const tmp = nums[i];
      nums[i] = nums[tmp - 1];
      nums[tmp - 1] = tmp;
    }
  }

  // after sorting, just check if current position has correct value
  // if not, then we know the value sitting on this position is duplicate
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== i + 1) {
      output.push(nums[i]);
    }
  }

  return output;
};

/*
This solution is basically the same as the third solution of findDisappearedNumbers
Just go check there to see the detail explanation

The main idea is loop through the array
when find a number i, flip the number at position i-1 to negative. 
if the number at position i-1 is already negative, i is the number that occurs twice.
*/
var findDuplicates = function (nums) {
  const output = [];

  for (let i = 0; i < nums.length; i++) {
    // note that nums[i] may filp to negative
    // so it's important to use Math.abs to make sure index is positive
    const index = Math.abs(nums[i]) - 1;

    if (nums[index] < 0) {
      output.push(index + 1);
    } else {
      nums[index] *= -1;
    }
  }

  return output;
};
