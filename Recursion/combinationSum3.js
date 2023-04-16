//////////////////////////////////////////////////////
// *** Combination Sum III ***
//////////////////////////////////////////////////////
/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:

Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.

Example 1:
Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.

Example 2:
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.

Example 3:
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
 
Constraints:
2 <= k <= 9
1 <= n <= 60
*/
/*
This solution is pretty straight forward, just follow the backtracking pattern

************************************************************
n = the input n, k = the input k
Time complexity: O((9 ^ k) * k)
=> The largest branching factor is 9
=> The deepest recursive tree is k
=> res.push([...tmp]); -> O(k) work
=> this is very upper bound

Space complexity: O(k)
=> The deepest recursive tree is k
*/
/**
 * @param {number} k
 * @param {number} n
 * @return {number[][]}
 */
var combinationSum3 = function (k, n) {
  const res = [];

  recursiveHelper(k, n, 1, [], res);

  return res;
};

function recursiveHelper(remainingCounts, target, index, tmp, res) {
  // base case
  if (remainingCounts === 0 && target === 0) {
    res.push([...tmp]);
    return;
  }

  // base case: out of the bound
  if (remainingCounts < 0 || target < 0) return;

  // if i is greater that the target, there's no need to keep for-loop
  for (let i = index; i <= 9 && i <= target; i++) {
    // choose
    tmp.push(i);

    // explore
    recursiveHelper(remainingCounts - 1, target - i, i + 1, tmp, res);

    // unchoose
    tmp.pop();
  }

  return;
}
