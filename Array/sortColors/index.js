//////////////////////////////////////////////////////
// *** Sort Colors ***
//////////////////////////////////////////////////////
/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]
 
Constraints:
n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.
 
Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
/*
I came up with this solution by myself
I think it's kind of hard to come up with the most optimize solution at the first time
But I did it
Although the code is actually kind of verbose, and not clean

This solution using the same idea as the most optimize solution(see below)
But just the code is not clean

Let's see the main idea
The main idea is not that hard
We use three pointers
1. zeroPtr
=> Indicates the position need to place 0 value
=> In other words, when we encounter 0, we know it's the position we're gonna swap

2. twoPtr
=> Indicates the position need to place 2 value
=> In other words, when we encounter 2, we know it's the position we're gonna swap

3. index
=> Just regular index to traverse the whole array

We use index pointer to traverse the whole array
When encounter 0, we just swap with the zeroPtr
Then move zeroPtr
When encounter 2, we just swap with the twoPtr
Them move twoPtr

But there' are couple edge cases need to be handled
I can't figure out the edge case right now
Just see the most-optimize solution above

************************************************************
Time complexity: O(n)
Space complexity: O(1)
*/
var sortColors = function (nums) {
  let zeroPtr = 0;
  let twoPtr = nums.length - 1;
  let index = 0;

  // we just move the zeroPtr to the correct position
  // to the position where it's not 0
  while (nums[zeroPtr] === 0) {
    zeroPtr++;
  }

  // index starts at zeroPtr
  index = zeroPtr;

  // we just move the twoPtr to the correct position
  // to the position where it's not 2
  while (nums[twoPtr] === 2) {
    twoPtr--;
  }

  // if index is over twoPtr, we can finish
  while (index <= twoPtr) {
    // encounter 0
    if (nums[index] === 0) {
      // do the swap
      swap(nums, index, zeroPtr);

      // update zeroPtr
      zeroPtr++;
    }
    if (nums[index] === 2) {
      swap(nums, index, twoPtr);
      twoPtr--;

      // handle the edge case
      if (nums[index] === 0 && index === zeroPtr) {
        while (nums[index] === 0) {
          zeroPtr++;
          index++;
        }
      }
    }

    if (nums[index] === 1) index++;
  }

  return nums;
};

function swap(arr, i, j) {
  const tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}

/*
This is non-optimize solution
The idea is just using hashTable to keep tracking the frequency of each value
And use second for-loop to mutate the original array

************************************************************
Time complexity: O(n)
Space complexity: O(n)
*/
var sortColors = function (nums) {
  const hashTable = {};

  for (const num of nums) {
    hashTable[num] = hashTable[num] ? hashTable[num] + 1 : 1;
  }

  for (let i = 0; i < nums.length; i++) {
    if (hashTable[0]) {
      nums[i] = 0;
      hashTable[0]--;
    } else if (hashTable[1]) {
      nums[i] = 1;
      hashTable[1]--;
    } else {
      nums[i] = 2;
    }
  }
};

function swap(arr, i, j) {
  const tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}

/*
This is the most-optimize solution
The main idea is as same as the first solution
But the code is a lot cleaner
Just see the video 
https://www.youtube.com/watch?v=4xbWSRZHqac

************************************************************
Time complexity: O(n)
Space complexity: O(1)
*/
var sortColors = function (nums) {
  let zeroPtr = 0;
  let twoPtr = nums.length - 1;
  let index = 0;

  while (index <= twoPtr) {
    if (nums[index] === 0) {
      swap(nums, index, zeroPtr);
      zeroPtr++;
    }
    if (nums[index] === 2) {
      swap(nums, index, twoPtr);
      twoPtr--;
      index--;
    }

    index++;
  }
};

function swap(arr, i, j) {
  const tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}
