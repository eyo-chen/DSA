//////////////////////////////////////////////////////
// *** Same Tree ***
//////////////////////////////////////////////////////
/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [1,2], q = [1,null,2]
Output: false

Example 3:
Input: p = [1,2,1], q = [1,1,2]
Output: false
 
Constraints:
The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104
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
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
/*
This problem is pretty easy to understand

************************************************************
Time: O(n)
Space: O(h)
*/
var isSameTree = function (p, q) {
  // both of them are null, kind of like base case, just return true
  if (p === null && q === null) {
    return true;
  }

  // if one of them is null, return false
  // Note that this condition only works because we have another condition above
  // so that we don't need to wory about both of them are null in this condition
  if (p === null || q === null) {
    return false;
  }

  // value is different
  if (p.val !== q.val) {
    return false;
  }

  // keep recursion
  return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
};
