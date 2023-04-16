//////////////////////////////////////////////////////
// *** Move Zeroes ***
//////////////////////////////////////////////////////
/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]
 

Constraints:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Could you minimize the total number of operations done?
*/
/*
This is the most naive way
Just see the code

************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(n)
*/
var moveZeroes = function (nums) {
  let zeroCounts = 0;
  let tmp = [];

  /*
  Iterate through array
  if it's zero, increment the zeroCounts
  if it's not, add it to the tmp array
  */
  for (const num of nums) {
    if (num === 0) {
      zeroCounts++;
    } else {
      tmp.push(num);
    }
  }

  // create the array with length is equal to zeroCounts, and contains all 0
  const zeroArr = new Array(zeroCounts).fill(0);

  // combine both of them
  tmp = tmp.concat(zeroArr);

  // now we just need mutate the input array
  for (let i = 0; i < nums.length; i++) {
    nums[i] = tmp[i];
  }
};

/*
This solution uses the idea of keep tracking where is the last place of non-zero occurs
For example, input array = [0,1,0,3,12]
lastNonZero = 0

i = 0, it's 0, skip it

i = 1, it's 1, non-zero
=> nums[0] = nums[1]
=> array = [1,1,0,3,12]
=> lastNonZero = 1

i = 2, it's 0, skip it

i = 3, it's 3, non-zero
=> nums[1] = nums[3]
=> array = [1,3,0,3,12]
=> lastNonZero = 2

i = 4, it's 12, non-zero
=> nums[2] = nums[4]
=> array = [1,3,12,3,12]
=> lastNonZero = 3

And
We start at 3(lastNonZero), mutate all the value to 0
DONE

The main idea of this solution is that
We first move all the non-zero element in front
We accomplish this by using lastNonZero
then mutate the array to 0 in the end

************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(1)
*/
var moveZeroes = function (nums) {
  let lastNonZero = 0;

  // If the current element is not 0, then we need to
  // append it just in front of last non 0 element we found.
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== 0) {
      nums[lastNonZero] = nums[i];

      // we've found one non-zero element, update the variable
      // so that next swap can have the correct order
      lastNonZero++;
    }
  }

  // After we have finished processing new elements,
  // all the non-zero elements are already at beginning of array.
  // We just need to fill remaining array with 0's.
  for (let i = lastNonZero; i < nums.length; i++) {
    nums[i] = 0;
  }
};

/*
This is reference from https://www.youtube.com/watch?v=aayNRwUN3Do

We use two ptrs
zeroPtr -> indate the position of zero 
workingPtr -> working ptr to keep updating

The idea is that we keep moving the workingPtr
If we hitting non-zero element, then we swap it with the place where zeroPtr locate at
After that, update zeroPtr

The idea is pretty simple
We just keep finding non-zero element
Once finding it, we just swap this value with the last found zero ptr

For example, input array = [0,1,0,3,12]

zeroPtr = 0
workingPtr = 0
=> value is 0, skip it

zeroPtr = 0
workingPtr = 1
=> value is 1, do swap
=> swap(arr, 0, 1)
=> [1,0,0,3,12]
=> update zeroPtr

zeroPtr = 1
workingPtr = 2
=> value is 0, skip it

zeroPtr = 1
workingPtr = 3
=> value is 3, do swap
=> swap(arr, 1, 3)
=> [1,3,0,0,12]
=> update zeroPtr

zeroPtr = 2
workingPtr = 4
=> value is 12, do swap
=> swap(arr, 2, 4)
=> [1,3,12,0,0]
=> update zeroPtr

Done
************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(1)
*/
var moveZeroes = function (nums) {
  let zeroPtr = 0;
  let workingPtr = 0;

  while (workingPtr < nums.length) {
    if (nums[workingPtr] !== 0) {
      swap(nums, zeroPtr, workingPtr);
      zeroPtr++;
    }

    workingPtr++;
  }
};

/*
This is final solution
The idea is similar to previous one

The main idea is to think We're non-zero element
And our goal is to move in front of all the zero
How can we do that?
If we know how many zero in front of me, then we can easily swap
For example,
[0,0,3]
I'm 3, and I'm at index 2
I know there're two 0s in front of me, so now I know I can do swap(arr, 2, 2 - 2)
Kind of like moving forwad two positions

[1,0,0,3]
I'm 3, and I'm at index 3
I know there're two 0s in front of me, so now I know I can do swap(arr, 2, 3 - 2)
Kind of like moving forwad two positions

[1,2,0,0,3]
I'm 3, and I'm at index 4
I know there're two 0s in front of me, so now I know I can do swap(arr, 2, 4 - 2)
Kind of like moving forwad two positions

Hope can see the pattern
We just think ourselves like non-zero element
Once we know how many zero in front of me, I can easily do the swap

Note that it won't be a case like this [0,0,1,2,0,3]
When we're at 3, there's no way [0,0,1,2]
there's no way there's two 0s in front of 1 and 2
Because when we first met 1, it will swap(arr, 2, 2 -2)
So does 2

Follow this pattern, we can solve this problem
*/
var moveZeroes = function (nums) {
  let zeroCounts = 0;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== 0) {
      swap(nums, i, i - zeroCounts);
    } else {
      zeroCounts++;
    }
  }
};

function swap(arr, index1, index2) {
  const tmp = arr[index1];
  arr[index1] = arr[index2];
  arr[index2] = tmp;
}
