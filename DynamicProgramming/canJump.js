//////////////////////////////////////////////////////
// *** Jump Game ***
//////////////////////////////////////////////////////
/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
 
Constraints:
1 <= nums.length <= 104
0 <= nums[i] <= 105
*/
/**
 * @param {number[]} nums
 * @return {boolean}
 */
/*
Recursive with memoization

************************************************************
n =  the length of nums
Time: O(n * n)
=> We would at least do n time works, or at least have n nodes
=> For each works, we do another n iteration, the reason we assume it's n iteration
=> Is because the maximum value in nums is n, if maximum value is greater than n, we still do n iteration at most, because once hitting the last index, we return true

Space: O(n)
The highest length of recusive tree
*/
function canJump(nums) {
  return recursiveHelper(nums, 0, {});

  function recursiveHelper(nums, index, memo) {
    // arrive last index
    if (index === nums.length - 1) return true;

    // when hitting 0 value, it means we can't go anywhere, or can't move
    if (nums[index] === 0) return false;

    // return memo answer
    if (memo[index] !== undefined) return memo[index];

    // what's steps can i take in this index?
    const availableSteps = nums[index];

    // try each index
    for (let i = 1; i <= availableSteps; i++) {
      const subAnswer = recursiveHelper(nums, index + i, memo);

      if (subAnswer) {
        memo[index] = subAnswer;
        return true;
      }
    }

    // if none of step can return true, we know there's no way to get the end
    memo[index] = false;
    return false;
  }
}

/*
DP table

************************************************************
n =  the length of nums
Time: O(n * n)
=> Similar to above, we assume the inner loop will do n works at most

Space: O(n)
*/
function canJump1(nums) {
  const table = new Array(nums.length).fill(false);
  // base case
  table[0] = true;

  for (let i = 0; i < nums.length; i++) {
    const availableSteps = nums[i];

    // if there's no steps we can take, the cell of table remain false
    // if there's no way to get this index before, which means from previous index, there's no way to get this point, remain false
    if (availableSteps >= 1 && table[i]) {
      // k + i < nums.length gurantee this loops only does n works at most
      for (let k = 1; k <= availableSteps && k + i < nums.length; k++) {
        table[i + k] = true;
      }
    }
  }

  return table[nums.length - 1];
}
/*
This use clever mindset to solve this question
Instead of cracking the problem from start, cracking the problem from the end

The orginal problem is kinda asking us can i get the final goal?
And the final goal is just last index

So we start at last two index
And keep asking,
Can you get to the current goal by using your max available steps from your index?
If the answer is yes, we change our goal to current index
If not, don't change the current gaol

For example, [3, 2, 1, 0, 4], current goal is 4 (last index of array)
Start at index 3, value 0, and ask 
Can you get to the current goal by using your max available steps from your index?
What's max available steps? -> 0
What's current index? -> 3
There's no way to get the goal 4 by using 0 max available steps from index 3
=> Don't change the goal

index 2, value 1, and ask
Can you get to the current goal by using your max available steps from your index?
What's max available steps? -> 1
What's current index? -> 2
There's no way to get the goal 4 by using 1 max available steps from index 2
=> Don't change the goal

index 1, value 2, and ask
Can you get to the current goal by using your max available steps from your index?
What's max available steps? -> 2
What's current index? -> 1
There's no way to get the goal 4 by using 2 max available steps from index 1
=> Don't change the goal

so on and so forth

************************************************************
n =  the length of nums
Time: O(n * n)

Space: O(1)
*/
function canJump2(nums) {
  // the current goal (initialize with last index of nums)
  let goal = nums.length - 1;

  // start from last two index
  for (let i = nums.length - 2; i >= 0; i--) {
    // Can you get to the current goal by using your max available steps from your index?
    if (i + nums[i] >= goal) goal = i;
  }

  return goal === 0;
}

// console.log(canJump([3, 2, 1, 0, 4]));
// console.log(canJump1([3, 2, 1, 0, 4]));
// console.log(canJump2([3, 2, 1, 0, 4]));
// [2,3,1,1,4]
