//////////////////////////////////////////////////////
// *** Longest Consecutive Sequence ***
//////////////////////////////////////////////////////
/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
 
Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
/**
 * @param {number[]} nums
 * @return {number}
 */
/*
Bruth force solution
Try every value in the nums
And use inner while-loop to find the longest consecutive sequence

************************************************************
Time complexity : O(n^3)
The outer loop runs exactly nn times, and because currentNum increments by 1 during each iteration of the while loop, it runs in O(n)O(n) time. Then, on each iteration of the while loop, an O(n)O(n) lookup in the array is performed. Therefore, this brute force algorithm is really three nested O(n)O(n) loops, which compound multiplicatively to a cubic runtime.

Space complexity : O(1)
The brute force algorithm only allocates a handful of integers, so it uses constant additional space.
*/
var longestConsecutive = function (nums) {
  let max = 0;

  for (const num of nums) {
    let acc = 0;
    let curNum = num;

    while (nums.includes(curNum)) {
      acc++;
      curNum++;
    }

    max = Math.max(max, acc);
  }

  return max;
};

/*
This is optimize solution
I try to find the optimize solution by myself
But it fail
Actually, the solution is a little tricky

For example, num = [100,4,200,1,3,2,101]
If we conceptually re-arrange this array
[1,2,3,4............, 100, 101, ....... 200]
How can we find the longest consecutive sequence
1. Find the starting point to start cout the sequence
=> In this case, there are three starting points
   1) 1
   2) 100
   3) 200
=> How to know it's a starting points?
=> Just check if it's leftside having a value
=> For example, 
   2 is not a starting point because it has 1 at it's left side
   4 is not a starting point because it has 3 at it's left side
   101 is not a starting point because it has 100 at it's left side
=> In short, a value is a starting point if there's no value at it's left side

2. From this starting points, count it's longest sequence
3. Done

The idea is quite simple
The hardest part is thinking of finding starting points (1.)

Note that it seems we have nested for-loops
But it's still O(n)
Just like our example,
We only do the while-loop on 1, 100 and 200
And we only touch each element once

And we first create an hash set or hashTable to have the O(1) lookup

************************************************************
Time complexity: O(n)
Space complexity: O(1)
*/
var longestConsecutive = function (nums) {
  // hashSet
  const set = new Set(nums);
  let max = 0;

  for (const num of nums) {
    // check if num is a starting point
    // it's a starting point if num - 1 is not in the hashSet
    if (!set.has(num - 1)) {
      let len = 0;
      let curNum = num;

      // keep counting sequence if curNum still in the hashSet
      while (set.has(curNum)) {
        len++;
        curNum++;
      }

      max = Math.max(len, max);
    }
  }

  return max;
};
