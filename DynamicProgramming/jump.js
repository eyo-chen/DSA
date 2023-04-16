//////////////////////////////////////////////////////
// *** Jump Game II ***
//////////////////////////////////////////////////////
/*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.

Example 1:
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [2,3,0,1,4]
Output: 2
 
Constraints:
1 <= nums.length <= 104
0 <= nums[i] <= 1000
*/
/**
 * @param {number[]} nums
 * @return {number}
 */
/*
Recursive Approach

For example, the input array is [2,3,1,1,4]

For 0 index, has two choies, jump 1 and 2
subproblem: Min([3,1,1,4], [1,1,4])
Again, the original problem is what's the minimum number of jumps of [2,3,1,1,4]?
The answer is Min([3,1,1,4], [1,1,4]) + 1
Once I knows the answer of [3,1,1,4] and [1,1,4], then I just choose the minimum and plus 1

Subproblem: [3,1,1,4]
For 0 index, has three choies, jump 1, 2 and 3
subproblem: Min([1,1,4], [1,4], [4])
Same logic as above

This the the core logic of this approach

Also, as we could see, we'll solve the [1,1,4] twice, which means we have duplicate work
So we use memo to cache the answer

************************************************************
n = the legnth of array
Time compelxity: O(n * 2)
=> We would have at least n node(call) in the tree because we will at least touch every single element in the input array with memoazation
=> For each node, we'll do n iteration at worst
   For example, [3,3,3,3]
   Index 0 -> 3 times work (n - 1)
   Index 1 -> 2 times work (n - 2)
   .....
   (n - 1) + (n - 2) + (n - 3) + ....... 1 

Space comelxity: O(n)
*/
function jump(nums) {
  return recuersiveHelper(nums, 0, {});

  function recuersiveHelper(nums, index, memo) {
    // find the answer
    if (index >= nums.length - 1) return 0;

    // can't reach anywhere
    if (nums[index] === 0) return null;
    if (memo[index] !== undefined) return memo[index];

    /*
      Initialize the invalid number
      We could just initialize with Infinity
      But the minimum invalid number just the length of input
      Because the longest jump won't greater than the length of input
      In other words, the final answer is gurantee less than the length of input
      */
    let res = nums.length;

    const availableSteps = nums[index];

    // i start at one -> the minimum step is alaways one
    // index + i < nums.length guarantee the iteration does O(n) work
    for (let i = 1; i <= availableSteps && index + i < nums.length; i++) {
      const subAnswer = recuersiveHelper(nums, index + i, memo);

      // if subAnswer is null, we skip it
      // plus one means jump once
      if (subAnswer !== null) res = Math.min(res, subAnswer + 1);
    }

    memo[index] = res;
    return res;
  }
}

/*
  DP Table approach
  
  For example, we're given [2,3,1,1,4]
  We start from the end
  Intialzie with the length of array because that's the minmum of invalid number
  
  2   3   1   1   4
  5   5   5   5   5 
  
  [4] -> what's the minimum number of jumps of [4]?
  => 0 (base case)
  2   3   1   1   4
  5   5   5   5   0
  
  [1] ->  what's the minimum number of jumps of [1, 4]?
  => From this index, I can only jump 1
  => So what's the minimum number of jumps of [4]? -> 0
  => 1 + 0 = 0
  2   3   1   1   4
  5   5   5   1   0
  
  [1] ->  what's the minimum number of jumps of [1, 1, 4]?
  => From this index, I can only jump 1
  => So what's the minimum number of jumps of [1, 4]? -> 1
  => 1 + 1 = 2
  2   3   1   1   4
  5   5   2   1   0
  
  [3] ->  what's the minimum number of jumps of [3, 1, 1, 4]?
  => From this index, I can jump 1, 2 and 3
  => So i'm gonna find the minimum of jumping 1,2 and 3
  2   3   1   1   4
  5   5   2   1   0
  => Min(2, 1, 0)
  => the minimum is 0 (jump 3)
  => 0 + 1 = 1
  2   3   1   1   4
  5   1   2   1   0
  
  [2] ->  what's the minimum number of jumps of [2, 3, 1, 1, 4]?
  => From this index, I can jump 1 and 2
  => So i'm gonna find the minimum of jumping 1 and 2
  2   3   1   1   4
  5   1   2   1   0
  => Min(1,2)
  => the minimum is 1 (jump 1)
  => 1 + 1 = 2
  2   3   1   1   4
  2   1   2   1   0
  
  ************************************************************
  n = the legnth of array
  Time compelxity: O(n * 2)
  => The outer loop abviously is O(n) works 
  => The logic of inner loop is similar to the recursive approach
  => At worst, we'll do another O(n) works to find the minimum number at worst
  
  Space comelxity: O(n)
  */
function jump1(nums) {
  const table = new Array(nums.length).fill(nums.length);
  table[nums.length - 1] = 0;

  // i start from last two index
  for (let i = nums.length - 2; i >= 0; i--) {
    const availableSteps = nums[i];

    // no steps
    if (availableSteps === 0) {
      nums[i] = null;
      continue;
    }

    let min = nums.length;

    // k + i < nums.length is same logic as above
    for (let k = 1; k <= availableSteps && k + i < nums.length; k++) {
      if (table[k + i] !== null) {
        min = Math.min(min, table[k + i]);
      }
    }

    table[i] = min + 1;
  }

  return table[0];
}

// greedy solution
function jump2(nums) {
  if (nums.length === 1) return 0;
  let maxReach = nums[0];
  let steps = nums[0];
  let res = 0;

  for (let i = 1; i < nums.length - 1; i++) {
    maxReach = Math.max(maxReach, nums[i] + i);
    steps--;

    if (steps === 0) {
      res++;
      steps = maxReach - i;
    }
  }

  return res + 1;
}

// console.log(jump([1, 1]));
// console.log(jump1([1, 1]));
// console.log(jump2([1, 1]));
