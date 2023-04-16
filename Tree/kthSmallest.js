//////////////////////////////////////////////////////
// *** Kth Smallest Element in a BST ***
//////////////////////////////////////////////////////
/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
 
Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104
 
Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
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
 * @param {number} k
 * @return {number}
 */
/*
This is acutally very easy
just use inorder traversal, and get the sorted array

************************************************************
Time: O(n)
Space: O(h)
*/
var kthSmallest = function (root, k) {
  const traverseArr = [];
  const stack = [];
  let node = root;

  while (stack.length > 0 || node !== null) {
    while (node !== null) {
      stack.push(node);
      node = node.left;
    }

    node = stack.pop();

    traverseArr.push(node.val);

    node = node.right;
  }

  return traverseArr[k - 1];
};
