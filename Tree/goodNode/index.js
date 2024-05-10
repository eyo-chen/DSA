//////////////////////////////////////////////////////
// ***  Count Good Nodes in Binary Tree ***
//////////////////////////////////////////////////////
/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.

Example 1:
Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.

Example 2:
Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

Example 3:
Input: root = [1]
Output: 1
Explanation: Root is considered as good.
 
Constraints:
The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].
*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
/*
This is the first solution I wrote
which is very intuitive

Just using path array to record all the node so far
And if the max value of all the current node value is less or equal to current node
then it's the good node
if (Math.max(...path) <= root.val) result++;
The pattern is very similar to backtracking problem

************************************************************
Time: O(n)
Space: O(h)
=> path array and call stack are both O(h)
*/
var goodNodes = function (root) {
  const path = [];

  return recursiveHelper(root, path);
};

function recursiveHelper(root, path) {
  if (root === null) return 0;

  let result = 0;

  path.push(root.val);

  if (Math.max(...path) <= root.val) result++;

  result += recursiveHelper(root.left, path);
  result += recursiveHelper(root.right, path);
  path.pop();

  return result;
}

/*
Another solution without path array

Just passing the max value so far

************************************************************
Time: O(n)
Space: O(h)
*/
var goodNodes = function (root) {
  return recursiveHelper(root, -Infinity);
};

function recursiveHelper(root, maxNum) {
  if (root === null) return 0;

  let result = 0;

  if (maxNum <= root.val) result++;

  result += recursiveHelper(root.left, Math.max(maxNum, root.val));
  result += recursiveHelper(root.right, Math.max(maxNum, root.val));

  return result;
}
