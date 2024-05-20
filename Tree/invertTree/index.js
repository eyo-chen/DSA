//////////////////////////////////////////////////////
// *** Invert Binary Tree ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []
 
Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
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
 * @return {TreeNode}
 */
/*
Recursive DFS
Solution is very easy to understand

************************************************************
Time: O(n)
Space: O(n)
*/
var invertTree = function (root) {
  if (root === null) {
    return null;
  }

  const left = invertTree(root.left);
  const right = invertTree(root.right);

  root.left = right;
  root.right = left;

  return root;
};

/*
Iterative DFS
*/
var invertTree = function (root) {
  const stack = [root];

  while (stack.length > 0) {
    const node = stack.pop();

    if (node === null) continue;

    const left = node.left;
    const right = node.right;

    if (right !== null) {
      stack.push(right);
    }

    if (left !== null) {
      stack.push(left);
    }

    node.left = right;
    node.right = left;
  }

  return root;
};

/*
Iterative BFS
*/
var invertTree = function (root) {
  const queue = [root];

  while (queue.length > 0) {
    const node = queue.shift();

    if (node === null) continue;

    const left = node.left;
    const right = node.right;

    if (left !== null) {
      queue.push(left);
    }

    if (right !== null) {
      queue.push(right);
    }

    node.left = right;
    node.right = left;
  }

  return root;
};
