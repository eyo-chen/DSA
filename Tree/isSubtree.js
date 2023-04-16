//////////////////////////////////////////////////////
// *** Subtree of Another Tree ***
//////////////////////////////////////////////////////
/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Example 1:
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Example 2:
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
 
Constraints:
The number of nodes in the root tree is in the range [1, 2000].
The number of nodes in the subRoot tree is in the range [1, 1000].
-104 <= root.val <= 104
-104 <= subRoot.val <= 104
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
 * @param {TreeNode} subRoot
 * @return {boolean}
 */
/*
This problem is very similar to isSameTree problem

We just traverse all the tree(root), and test if current subtree is as same as input subRoot
We do this by isSameTree problem

For example,
root:
             3
        4         5
    1      2

subRoot: 
           4
        1     2

First
Check if the subTree of root(3) is as same as subRoot
=> It's not
=> go to left and right subtree

Second
Check if the subTree of root(4) is as same as subRoot
=> Yes
=> return true

So basically, we'll traverse all the subtree, and touch each node multiple times
Note that the original tree is also one of the subtree

************************************************************
n = node of root, s = node of subRoot, h = height of subRoot
Time: O(n * s)
Space: O(h)
*/
var isSubtree = function (root, subRoot) {
  if (root === null) {
    return false;
  }

  if (isSameTree(root, subRoot)) {
    return true;
  }

  return isSubtree(root.right, subRoot) || isSubtree(root.left, subRoot);
};

function isSameTree(root, subRoot) {
  if (root === null && subRoot === null) {
    return true;
  }

  if (root === null || subRoot === null) {
    return false;
  }

  if (root.val !== subRoot.val) {
    return false;
  }

  return (
    isSameTree(root.right, subRoot.right) && isSameTree(root.left, subRoot.left)
  );
}
