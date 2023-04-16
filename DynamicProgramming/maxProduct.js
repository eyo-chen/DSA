//////////////////////////////////////////////////////
// *** Maximum Product Subarray ***
//////////////////////////////////////////////////////
/*
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 
Constraints:
1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/
/**
 *
 * @param {number[]} nums
 * @return {number}
 */
/*
Bruth force soultion

Try every sub array, and find the maximum

************************************************************
n = the legnth of array
Time compelxity: O(n ^ 2)

Space comelxity: O(1)
*/
function maxProduct(nums) {
  let res = nums[0];
  for (let i = 0; i < nums.length; i++) {
    res = Math.max(res, nums[i]);

    let tmp = nums[i];
    for (let j = i + 1; j < nums.length; j++) {
      tmp *= nums[j];
      res = Math.max(res, tmp);
    }
  }

  return res;
}

/*
Use dynamic programming

If all values in array is positive, it's quite easy
but it's tricky once having negative value

For example,[-1,-2,-3], think of subproblem
Kinda easy?
Iterate through the array, store the max product value in each index

Imagine we can have two choices, extend the pre max, or start the new sub array
[-1] -> What's the max product in [-1] -> -1
[-1,-2] -> What's the max product in [-1, -2] -> Max(-2, -1 * -2) -> 2
[-1,-2,-3] -> What's the max product in [-1, -2, -3] -> Max(-3, 2 * -3) -> -3 ???

This is not gonna work obviously

Instead of keeping track the max product value for each element of array
We have to both remember max and min product value
Because once the value of element is negative, 
our answer is gonna be (value of element * min product value)

For example,[-1,-2,-3]
[-1] -> What's the max and min product in [-1]
=> min: -1
=> max: -1

[-1,-2] -> What's the max and min product in [-1, -2]
=> min: -2
=> max: 2

[-1,-2,-3] -> What's the max and min product in [-1, -2, -3]
=> min: -6
=> max: 6

So now the problem may lead to how to store the min and max along with the iteration
Here is first solution, it's kind intuitive
For each iteration, we know we always have three values
1. value of element itself
2. value of element * min (which has been set before)
3. value of element * max (which has been set before)
=> Just find which one is min, which one is max

For example, [-1,-2,-3], when index = 2, value = -3
1. value of element itself
=> -3
2. value of element * min (which has been set before)
=> -3 * -2
3. value of element * max (which has been set before)
=> -3 * 2

************************************************************
n = the legnth of array
Time compelxity: O(n)

Space comelxity: O(1)
*/
function maxProduct1(nums) {
  let min = nums[0],
    max = nums[0],
    res = nums[0];

  for (let i = 1; i < nums.length; i++) {
    const firstVal = min * nums[i]; // the value multiply previous min
    const seconVal = max * nums[i]; // the value multiply previous max

    // find the min, and max among three different values
    min = Math.min(firstVal, seconVal, nums[i]);
    max = Math.max(firstVal, seconVal, nums[i]);

    // update global max value
    res = Math.max(res, max);
  }

  return res;
}

/*
The core idea is same
But the main difference is how to find min and max value

There are two different scenario
1. the value of element is positive
=> just multiply previous min and max to get the new ones

2. the value of element is negative
=> need to swap first

For example, now min = 3, min = 10
If the value of element is 2
then new min and new max is just
old min * 2 = 6
old max * 2 = 20

If the value of element is -2
then first swap min and max, then multiply
10 * -2 = -20 -> new min
3 * -2 = -6  -> new max

************************************************************
n = the legnth of array
Time compelxity: O(n)

Space comelxity: O(1)
*/
function maxProduct2(nums) {
  let min = nums[0],
    max = nums[0],
    res = nums[0];

  for (let i = 1; i < nums.length; i++) {
    // swap
    if (nums[i] < 0) [max, min] = [min, max];

    // still need to compare with current value of element
    min = Math.min(nums[i] * min, nums[i]);
    max = Math.max(nums[i] * max, nums[i]);

    res = Math.max(res, max);
  }

  return res;
}

console.log(maxProduct([-1]));
console.log(maxProduct1([-1]));
console.log(maxProduct2([-1]));
