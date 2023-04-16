//////////////////////////////////////////////////////
// ***  Climbing Stairs ***
//////////////////////////////////////////////////////
/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 
Constraints:
1 <= n <= 45
*/
/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  const table = new Array(n + 1).fill(0);
  table[0] = 1;
  table[1] = 1;

  for (let i = 2; i <= n; i++) {
    for (let j = 1; j <= 2; j++) {
      if (i >= j) table[i] += table[i - j];
    }
  }

  return table[n];
};

var climbStairs1 = function (n, memo = []) {
  if (n === 0) return 1;
  if (n < 0) return 0;
  if (memo[n] !== undefined) return memo[n];

  let steps = 0;

  for (let i = 1; i <= 2; i++) {
    steps += climbStairs1(n - i, memo);
  }

  memo[n] = steps;
  return steps;
};

//////////////////////////////////////////////////////////////
// Variant from algoexpert
// given height and maxSteps

function staircaseTraversal(height, maxSteps) {
  if (height === 0) return 1;
  if (height < 0) return 0;

  let steps = 0;

  for (let i = 1; i <= maxSteps; i++) {
    steps += staircaseTraversal(height - i, maxSteps);
  }

  return steps;
}

function staircaseTraversal1(height, maxSteps, memo = []) {
  if (height === 0) return 1;
  if (height < 0) return 0;
  if (memo[height] !== undefined) return memo[height];

  let steps = 0;

  for (let i = 1; i <= maxSteps; i++) {
    steps += staircaseTraversal1(height - i, maxSteps, memo);
  }

  memo[height] = steps;
  return steps;
}

function staircaseTraversal(height, maxSteps) {
  const table = new Array(height + 1).fill(0);
  table[0] = 1;
  table[1] = 1;

  for (let i = 2; i <= height; i++) {
    for (let j = 1; j <= maxSteps; j++) {
      if (i >= j) table[i] += table[i - j];
    }
  }

  return table[height];
}
