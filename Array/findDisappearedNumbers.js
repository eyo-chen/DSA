//////////////////////////////////////////////////////
// *** Find All Numbers Disappeared in an Array ***
//////////////////////////////////////////////////////
/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:
Input: nums = [1,1]
Output: [2]
 

Constraints:
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n

Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
*/
/**
 * @param {number[]} nums
 * @return {number[]}
 */
/*
This is non-optimize
************************************************************
Time compelxity: O(n)
Space comelxity: O(n)
*/
var findDisappearedNumbers = function (nums) {
  const output = [];
  const hashTable = {};

  // set 1 ~ nums.length to 1
  for (let i = 0; i < nums.length; i++) {
    hashTable[i + 1] = 1;
  }

  // loop through nums to minue the value
  for (let i = 0; i < nums.length; i++) {
    hashTable[nums[i]]--;
  }

  // if any value still 1, it means it didn't appear in the nums
  for (const [key, val] of Object.entries(hashTable)) {
    if (val === 1) {
      output.push(+key);
    }
  }

  return output;
};

/*
Using swap to sort the array

The main idea is that
we try to put every element into the correct position
The idea is simple, if nums[i] != i + 1 and nums[i] != nums[nums[i] - 1], then we swap nums[i] with nums[nums[i] - 1]
for example, nums[0] = 4 and nums[3] = 7
then we swap nums[0] with nums[3]
So In the end the array will be sorted and if nums[i] != i + 1, then i + 1 is missing.

For example, nums = [4,3,2,7,8,2,3,1]

i = 0, val = 4
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
4 !== 1 && 4 !== 7
swap(0, 3) -> 1th
[7,3,2,4,8,2,3,1]
=> 4 is in the correct position

Again, keep while-loop
i = 0, val = 7
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
7 !== 1 && 7 !== 3
swap(0, 6) -> 2th
[3,3,2,4,8,2,7,1]
=> 7 is in the correct position

Again, keep while-loop
i = 0, val = 3
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
3 !== 1 && 3 !== 2
swap(0, 2) -> 3th
[2,3,3,4,8,2,7,1]
=> 3 is in the correct position

Again, keep while-loop
i = 0, val = 2
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
2 !== 1 && 2 !== 3
swap(0, 1) -> 4th
[3,2,3,4,8,2,7,1]
=> 2 is in the correct position

Again, keep while-loop
i = 0, val = 3
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
3 !== 1 && 3 !== 3
-> false, don't swap, don't execute while-loop

i = 1, val = 2
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
2 !== 2 
-> false, don't swap, don't execute while-loop

i = 2, val = 3
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
3 !== 3 
-> false, don't swap, don't execute while-loop

i = 3, val = 4
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
4 !== 4
-> false, don't swap, don't execute while-loop

i = 4, val = 8
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
8 !== 5 && 8 !== 1
swap(4, 7) -> 5th
[3,2,3,4,1,2,7,8]
=> 8 is in the correct position

i = 4, val = 1
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
1 !== 5 && 1 !== 3
swap(4, 0) -> 6th
[1,2,3,4,3,2,7,8]
=> 1 is in the correct position

i = 4, val = 3
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
3 !== 5 && 3 !== 3
3 === 3
=> false
=> What does this mean?
=> It means nums[4], 3 is not in the correct position (nums[i] !== i + 1)
=> But, there's already a 3 in the correct position (nums[i] === nums[nums[i] - 1])
=> So that we don't need to swap

i = 5, val = 2
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
2 !== 6 && 2 !== 2
2 === 2
=> false
=> What does this mean?
=> It means nums[5], 2 is not in the correct position (nums[i] !== i + 1)
=> But, there's already a 2 in the correct position (nums[i] === nums[nums[i] - 1])
=> So that we don't need to swap

i = 6, val = 7
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
-> false, don't swap, don't execute while-loop

i = 7, val = 8
nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]
-> false, don't swap, don't execute while-loop



I hope the logic is clear
The for-loop
=> I just iterate through the whole nums

The while-loop
=> I just check if each element is in the correct position

Condition1
nums[i] !== i + 1
=> Is the element nums[i] at the correct position?
=> Note that index is is 0 based
=> So that we said nums[i] !== i + 1
=> For example, i = 2, val = 3, so that nums[2] should be eqault to 2 + 1
=> 3 === 3
=> Is at the correct position

Condition2
nums[i] !== nums[nums[i] - 1]
=> Is the correct position still have the correct value?
=> For example, i = 2, val = 3
=> So that if 2 === nums[2 - 1] 
=> It means the position(index = 1) already had the value 2
=> so that we don't need to swap

Final thing to note that this solution is O(n) time
time complexity cannot simply be concluded with the number of nested loops. You have to count the number of operations performed in the loops. The for loop will definitely iterate through all elements of the list once but the conditions of the while loop ensure that the loop will also only iterate through all elements at most once for the entirety of the for loop. Hence O(2n) -> O(n).
If we loop at the process above, indeed,
The total count of swap only happens 6th times

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
var findDisappearedNumbers = function (nums) {
  const output = [];

  for (let i = 0; i < nums.length; i++) {
    // Keep swaping until position i is has correct value
    while (nums[i] !== i + 1 && nums[i] !== nums[nums[i] - 1]) {
      const tmp = nums[i];
      nums[i] = nums[tmp - 1];
      nums[tmp - 1] = tmp;
    }
  }

  // checking if position is has correct value
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== i + 1) {
      output.push(i + 1);
    }
  }

  return output;
};

/*
This is tricky solution

The main idea is 
We want to notify that nums[i] has appeared in the array nums
How?
We set nums[nums[i] - 1] to negative
For example, nums = [4,3,2,7,8,2,3,1]
i = 0, val = 4
Hey, go to nums[4 - 1], and set it to negative number
It means value 4 appears in the array

i = 1, val = 3
Hey, go to nums[3 - 1], and set it to negative number
It means value 3 appears in the array

so on and so forth...

After that, we just loop through whole array
if nums[i] is still positive, it means i + 1 is not in the array


One thing to note that because nums[i] could be negative
So we have to use Math.abs() to make sure we get the correct index
Also we don't want to flip negative to positive again
so that we do Math.abs(nums[index]) * -1;
to maek sure value is always negative

https://www.youtube.com/watch?v=8i-f24YFWC4
************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
var findDisappearedNumbers = function (nums) {
  const output = [];

  for (let i = 0; i < nums.length; i++) {
    const index = Math.abs(nums[i]) - 1;
    nums[index] = Math.abs(nums[index]) * -1;
  }

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] > 0) {
      output.push(i + 1);
    }
  }

  return output;
};
