//////////////////////////////////////////////////////
// *** Subarray Product Less Than K ***
//////////////////////////////////////////////////////
/*
Given an array of integers nums and an integer k, return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.

Example 1:
Input: nums = [10,5,2,6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

Example 2:
Input: nums = [1,2,3], k = 0
Output: 0
 
Constraints:
1 <= nums.length <= 3 * 104
1 <= nums[i] <= 1000
0 <= k <= 106
*/
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
/*
This is bruth force solution
Very easy to understand

************************************************************
Time compelxity: O(n ^ 2)
Space comelxity: O(1)
*/
var numSubarrayProductLessThanK = function (nums, k) {
  let count = 0;

  for (let i = 0; i < nums.length; i++) {
    let accVal = nums[i];

    // nums[i] is less than k, increment count
    if (accVal < k) {
      count++;
    }

    for (let j = i + 1; j < nums.length; j++) {
      // keep accmulated accVal
      accVal *= nums[j];

      if (accVal < k) {
        count++;
      }
      // if accVal is greater than or eqaul to k, break
      else {
        break;
      }
    }
  }

  return count;
};
/*
This is optimize soltuion
I only came up with the idea
But can't implement the code

I know the overall idea of this solution
Using two pointers, and accVal variable
We just keep updating fast pointer
And accumulated accVal accVal *= nums[fast];
When accVal is eqaul to or greater than k
It's time to update slow pointer
This is just basic idea of sliding window
BUT
How do we increment the count variable
For example,
[10,5,2,6]
fast = 0
count = 1, add [10]
fast = 1
count = 2, add [10,5]
fast = 2
accVal = 100, move slow pointer
fast = 3, slow = 1
count = 3, add [5,2,6]
fast done
As we can see, we didn't count the case like [5] and [2]

So the hardest part is how do we count every case
when we update fast and slow pointer

the answer is this line of code
count += fast - slow + 1

For example,
[10,3,5,2,6], k = 100

fast = 0, slow = 0
accVal = 10
count += fast - slow + 1
count += 1
=> add [10]
=> count = 1
=> count = [10]
=> update fast pointer

fast = 1, slow = 0
accVal = 30
count += fast - slow + 1
count += 2
=> sliding window = [10,3]
=> add [10,3], [3]
=> count = 3
=> count = [10], [10,3], [3]
=> update fast pointer

fast = 2, slow = 0
accVal = 150
=> move slow pointer
=> and decrease accVal

fast = 2, slow = 1
accVal = 15
count += fast - slow + 1
count += 2
=> sliding window = [3,5]
=> add [3,5], [5]
=> count = 5
=> count = [10], [10,3], [3], [3,5], [5]
=> update fast pointer

fast = 3, slow = 1
accVal = 30
count += fast - slow + 1
count += 3
=> sliding window = [3,5,2]
=> add [3,5,2], [5,2], [2]
=> count = 8
=> count = [10], [10,3], [3], [3,5], [5], [3,5,2], [5,2], [2]
=> update fast pointer

fast = 4, slow = 1
accVal = 180
=> move slow pointer
=> and decrease accVal

fast = 4, slow = 2
accVal = 60
count += fast - slow + 1
count += 3
=> sliding window = [5,2,6]
=> add [5,2,6], [2,6], [6]
=> count = 11
=> count = [10], [10,3], [3], [3,5], [5], [3,5,2], [5,2], [2], [5,2,6], [2,6], [6]
=> update fast pointer

DONE

I hope the process is clear

Main takeaway is 
Try to observe the sliding window
And find the pattern to update count

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
var numSubarrayProductLessThanK = function (nums, k) {
  let fast = 0;
  let slow = 0;
  let accVal = 1;
  let count = 0;

  if (k <= 1) {
    return 0;
  }

  while (fast < nums.length) {
    accVal *= nums[fast];

    // if k is not greater than accVal, update slow pointer, and decrease the accVal
    while (accVal >= k) {
      accVal /= nums[slow];
      slow++;
    }

    count += fast - slow + 1;

    fast++;
  }

  return count;
};
