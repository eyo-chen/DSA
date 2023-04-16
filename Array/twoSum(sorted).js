//////////////////////////////////////////////////////
// *** Two Sum II - Input Array Is Sorted ***
//////////////////////////////////////////////////////
/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
 
Constraints:
2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers is sorted in non-decreasing order.
-1000 <= target <= 1000
The tests are generated such that there is exactly one solution.
*/
/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
/*
The main difference of this problem is that the input array is sorted
So that we can use two-pointers, don't need to use hashTable(O(n))

For example, numbers = [2,7,11,15], target = 9
l -> left ptr, r -> right ptr

l           r
[2, 7, 11, 15]
sum = target - arr[l] - arr[r] = -8
=> It means the sum of l and r is too large
=> How can we decrease sum?
=> Move the right pointer inward
=> r--

l       r
[2, 7, 11, 15]
sum = target - arr[l] - arr[r] = -4
=> It means the sum of l and r is too large
=> How can we decrease sum?
=> Move the right pointer inward
=> r--

l   r
[2, 7, 11, 15]
sum = target - arr[l] - arr[r] = 0


For example, [-5, -3, -1, 0, 1, 3, 5], targer = 2

l                     r
[-5, -3, -1, 0, 1, 3, 5]
sum = target - arr[l] - arr[r] = 2
=> It means the sum of l and r is too small
=> How can we increase the sum?
=> Move the left pointer inward
=> l++

     l                 r
[-5, -3, -1, 0, 1, 3, 5]
sum = target - arr[l] - arr[r] = 0

Note that this problem can also be solved by using hashTable
But it make sense to use two pointers since the input array is sorted
************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(1)
*/
var twoSum = function (numbers, target) {
  let right = numbers.length - 1;
  let left = 0;

  while (right > left) {
    const sum = target - numbers[right] - numbers[left];

    if (sum === 0) {
      return [left + 1, right + 1];
    } else if (sum < 0) {
      right--;
    } else {
      left++;
    }
  }
};
