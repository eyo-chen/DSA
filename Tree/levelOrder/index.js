//////////////////////////////////////////////////////
// *** Binary Tree Level Order Traversal ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
 
Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
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
 * @return {number[][]}
 */
/*
BFS using queue
Very straightforward
Just see the code

************************************************************
Time: O(n)
Space: O(h)
*/
var levelOrder = function (root) {
  if (root === null) return [];

  const res = [];
  const queue = [root];

  while (queue.length > 0) {
    const len = queue.length;
    let tmp = [];

    for (let i = 0; i < len; i++) {
      const node = queue.shift();

      tmp.push(node.val);

      if (node.left !== null) queue.push(node.left);

      if (node.right !== null) queue.push(node.right);
    }

    res.push(tmp);
  }

  return res;
};

/*
DFS using recursion

************************************************************
Time: O(n)
Space: O(h)
*/
var levelOrder = function (root) {
  const res = [];
  recursiveHelper(root, res, 0);

  return res;
};

function recursiveHelper(node, res, level) {
  if (node === null) return;

  if (!res[level]) res[level] = [node.val];
  else res[level].push(node.val);

  recursiveHelper(node.left, res, level + 1);
  recursiveHelper(node.right, res, level + 1);
  return;
}
