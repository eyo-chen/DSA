//////////////////////////////////////////////////////
// *** Product of Array Except Self ***
//////////////////////////////////////////////////////
/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/
/**
 * @param {number[]} nums
 * @return {number[]}
 */
/*
The solution of this problem is very non-intuitive

Use prefix array to represent the total amount from the left side
prefix[k] = prefix[k - 1] * array[k] 
Use suffix array to represent the total amount from the right side
suffix[k] = suffix[k + 1] * array[k]


For example, we're given 
         [1, 2, 3, 4]

prefix = [1, 2, 6, 24]
each element in the prefix array represent
the total accumulated amount caculating from left side
for example, for index 2, the value is 6, it means the total accumulated amount from index 0 ~ 2 is 1 * 2 * 3 = 6

suffix = [24, 24, 12, 4]
each element in the suffix array represent
the total accumulated amount caculating from right side
for example, for index 1, the value is 24, it means the total accumulated amount from index 3 ~ 1 is 4 * 3 * 2 = 24

After building those two arrays, we can build the result array
Except for the very first and last element
If index is 2, then we know 
the index 1 of prefix represent the total total accumulated amount caculating from index 0 ~ 1
the index 3 of suffix represent the total total accumulated amount caculating from index 3 ~ 3
We get the answer once we multiply those two elements

However, the edge case is at the very first and last element
so we're supposed the value before the first element and after the last elemet is 1
If index is 0, then index 1 of suffix represent the total accumulated amount caculating from 3 ~ 1
If index is 3, then index 2 of prefix represent the total accumulated amount caculating from 0 ~ 2
*/
function productExceptSelf(nums) {
  const prefix = [];
  const suffix = [];
  const res = new Array(nums.length).fill(1);
  const len = nums.length;

  prefix[0] = nums[0];
  for (let i = 1; i < len; i++) {
    prefix[i] = prefix[i - 1] * nums[i];
  }

  suffix[len - 1] = nums[len - 1];
  for (let i = len - 2; i >= 0; i--) {
    suffix[i] = suffix[i + 1] * nums[i];
  }

  // for (let i = 0; i < len; i++) {
  //   if (i === 0) {
  //     res[i] = 1 * suffix[i + 1];
  //   } else if (i === len - 1) {
  //     res[i] = 1 * prefix[i - 1];
  //   } else res[i] = prefix[i - 1] * suffix[i + 1];
  // }

  // same for-loop as above one, but it's cleaner
  // have to initialize each element with value 1
  for (let i = 0; i < len; i++) {
    if (i !== 0) {
      res[i] *= prefix[i - 1];
    }

    if (i !== len - 1) {
      res[i] *= suffix[i + 1];
    }
  }

  return res;
}
/*
  This solution is very similar to the one above, but build the prefix and suffix differently
  
  Instead of initializing the first and last element with original value, 
  it initializes to 1
  
  For example, we're given 
           [1, 2, 3, 4]
  
  prefix = [1, 1, 2, 6]
  for example, for index 2, the value is 2, it means the total accumulated amount from index 0 ~ 1 is 1 * 2 = 2
  
  suffix = [24, 12, 4, 1]
  for example, for index 2, the value is 12, it means the total accumulated amount from index 3 ~ 2 is 4 * 3 = 12
  
  The main difference is caculaing the value exclusively
  each element in the prefix represent the total amount before the index
  each element in the suffix represent the total amount after the index
  
  So at the end of for-loop we can just multiply the element in the suffix and prefix
  
  ************************************************************
  both soultion are the same
  
  Time compelxity: O(n)
  
  Space comelxity: O(n)
  */
function productExceptSelf1(nums) {
  const prefix = [];
  const suffix = [];
  const res = [];

  prefix[0] = 1;
  for (let i = 1; i < nums.length; i++) {
    prefix[i] = nums[i - 1] * prefix[i - 1];
  }

  suffix[nums.length - 1] = 1;
  for (let i = nums.length - 2; i >= 0; i--) {
    suffix[i] = nums[i + 1] * suffix[i + 1];
  }

  for (let i = 0; i < nums.length; i++) {
    res.push(suffix[i] * prefix[i]);
  }

  return res;
}

/*
  This solution only use O(1) space (The output array does not count as extra space for space complexity analysis.)
  
  Instead of creating two arrays(prefix and suffix)
  
  As same as before, build the prefix part in the result array
  
  But we use suffixPointer instead, mutiply each other and also accumulated the suffixPointer
  */
function productExceptSelf2(nums) {
  const res = [];
  let suffixPointer = 1;

  res[0] = 1;
  for (let i = 1; i < nums.length; i++) {
    res[i] = nums[i - 1] * res[i - 1];
  }

  for (let i = nums.length - 1; i >= 0; i--) {
    res[i] = suffixPointer * res[i];
    suffixPointer *= nums[i];
  }

  return res;
}

// console.log(productExceptSelf([1, 2, 3, 4]));
// console.log(productExceptSelf1([1, 2, 3, 4]));
// console.log(productExceptSelf2([1, 2, 3, 4]));
