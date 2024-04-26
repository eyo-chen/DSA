//////////////////////////////////////////////////////
// *** Balanced Binary Tree ***
//////////////////////////////////////////////////////
/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Example 3:
Input: root = []
Output: true
 
Constraints:
The number of nodes in the tree is in the range [0, 5000].
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
 * @param {TreeNode} root
 * @return {boolean}
 */
/*
The code should be very easy to understand the main logic

************************************************************
Time: O(n)
Space: O(h)
*/
var isBalanced = function (root) {
  // if returned value is false, just return false
  if (recursiveHelper(root) === false) return false;

  // if returned value is not false(it's integer), return true
  return true;
};

function recursiveHelper(root) {
  // base case
  if (root === null) return 0;

  const left = recursiveHelper(root.left); // go left
  const right = recursiveHelper(root.right); // go right

  /*
  if one of them is false,
  it means one of node is not balanced
  All we need to do is all the way to bubble up to return false
  */
  if (left === false || right === false) return false;

  // if height is not balanced, return false
  if (Math.abs(left - right) > 1) return false;

  // returns max height
  return Math.max(left, right) + 1;
}

/*
This is very similar to the solution above
The main difference is that
we immediately return false after finding the left sub tree is not balanced
Why do we do this?
Image the left and right sub tree is very height
And we've search to the left sub tree
then found out it's not balanced, we'll bubble up to return false
However, we don't need to search for the right subtree anymore
Because the answer is false

So after finding left subtree is false, we just immediately return false
*/
var isBalanced = function (root) {
  if (recursiveHelper(root) === false) return false;

  return true;
};

function recursiveHelper(root) {
  if (root === null) return 0;

  const left = recursiveHelper(root.left);

  // immediately return false
  if (left === false) return false;

  const right = recursiveHelper(root.right);

  if (right === false) return false;

  if (Math.abs(left - right) > 1) return false;

  return Math.max(left, right) + 1;
}
