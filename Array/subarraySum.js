//////////////////////////////////////////////////////
// *** Rotate Image ***
//////////////////////////////////////////////////////
/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2

Constraints:
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
/*
Bruth force solution
It's very intuitive

Note that the problem is asking contiguous subarray
For example, [1, -1, 1, 1, 1, 1], k = 3
Start at index 0
[1, -1] = 0 
[1, -1, 1] = 1
[1, -1, 1, 1] = 2
[1, -1, 1, 1, 1] = 3, find it!!!
[1, -1, 1, 1, 1, 1] = 4

Start at index 1
[-1, 1] = 0
[-1, 1, 1] = 1
[-1, 1, 1, 1] = 2
[-1, 1, 1, 1, 1] = 3, find it!!!

So on and so forth
************************************************************
n = the legnth of array
Time compelxity: O(n ^ 2)
Space comelxity: O(1)
*/
var subarraySum = function (nums, k) {
  let res = 0;

  for (let i = 0; i < nums.length; i++) {
    let sum = nums[i];
    if (sum === k) {
      res++;
    }

    for (let j = i + 1; j < nums.length; j++) {
      sum += nums[j];

      if (sum === k) {
        res++;
      }
    }
  }

  return res;
};

/*
Optimize Solution
It's very non-intuitive

It's use two main ideas
1. prefix sum
2. hash table

Can first go watch this video https://www.youtube.com/watch?v=fFVZt-6sgyo&t=773s

If we look at the solution above, we do a lot of repeated works
Main idea is
1. iterate through the array once(0 ~ i)
2. using sum variable to keep tracking the current sum at position i
3. Add current sum(prefix sum) to hashTable
   (If exists, just increment the frequency)
4. Check if (sum - k) exists in the hashTable, then add the frequncy to count

I can't explain the logic very well right now, just go watching the video
*/
var subarraySum1 = function (nums, k) {
  const prefixSumTable = { 0: 1 };
  let sum = 0;
  let res = 0;

  for (let i = 0; i < nums.length; i++) {
    sum += nums[i];

    prefixSumTable[sum] = prefixSumTable[sum] ? prefixSumTable[sum] + 1 : 1;

    if (prefixSumTable[sum - k]) {
      res += prefixSumTable[sum - k];
    }
  }

  return res;
};
