//////////////////////////////////////////////////////
// *** Combinations ***
//////////////////////////////////////////////////////
/*
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].

You may return the answer in any order.

Example 1:
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

Example 2:
Input: n = 1, k = 1
Output: [[1]]
 
Constraints:
1 <= n <= 20
1 <= k <= n
*/
/*
This question is quite easy after finishing previous several backtracking problem

Think the recursive tree like this when n = 4, k = 2
                                                             1,2,3,4
                                     1              2                3             4
                          2      3       4       3    4             4  
=> As we could see, each path is our answer
=> The largest branching factor is n
=> The deepest recursive tree is k
=> Once choosing i, we have to start iteration from i + 1 in next call stack
=> Look the left hand side, after choosing 1, we have to start the iteration from 2 in the next call stack

************************************************************
n = the input n, k = the input k
Time complexity: O((n ^ k) * k)
=> The largest branching factor is n
=> The deepest recursive tree is k
=> res.push([...tmp]); -> O(k) work

Space complexity: O(k)
*/
/**
 * @param {number} n
 * @param {number} k
 * @return {number[][]}
 */
var combine = function (n, k) {
  const res = [];
  recursiveHelper(n, k, 1, [], res);
  return res;
};

function recursiveHelper(n, k, index, tmp, res) {
  // Goal
  if (tmp.length === k) {
    res.push([...tmp]);
    return;
  }

  for (let i = index; i <= n; i++) {
    // Choose
    tmp.push(i);

    // Explore
    recursiveHelper(n, k, i + 1, tmp, res);

    // Unchoose
    tmp.pop();
  }

  return;
}

console.log(combine(3, 2));
