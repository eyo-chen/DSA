//////////////////////////////////////////////////////
// *** Binary Tree Level Order Traversal II ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).
 
Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]

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

It's very similar to previous one
But now we use stack to temporarily store the result
After that, we pop out to store into the res
So that the order is reverse

************************************************************
Time: O(n)
Space: O(h)
*/
var levelOrderBottom = function (root) {
  if (root === null) return [];

  const queue = [root];
  const stack = [];
  const res = [];

  while (queue.length > 0) {
    const len = queue.length;
    const tmp = [];

    for (let i = 0; i < len; i++) {
      const node = queue.shift();
      tmp.push(node.val);

      if (node.left !== null) queue.push(node.left);

      if (node.right !== null) queue.push(node.right);
    }

    stack.push(tmp);
  }

  while (stack.length > 0) {
    res.push(stack.pop());
  }

  return res;
};

/*
Using hashTable to store the level and the res
After that, just store into result in a reverse order

************************************************************
Time: O(n)
Space: O(h)
*/
var levelOrderBottom = function (root) {
  const hashTable = {};
  const res = [];

  recursiveHelper(root, hashTable, 0);

  const levelKey = Object.keys(hashTable)
    .map(key => Number(key))
    .sort((a, b) => b - a);

  for (let i = 0; i < levelKey.length; i++) {
    res.push(hashTable[String(levelKey[i])]);
  }

  return res;
};

function recursiveHelper(root, hashTable, level) {
  if (root === null) return;

  if (hashTable[level]) hashTable[level].push(root.val);
  else hashTable[level] = [root.val];

  recursiveHelper(root.left, hashTable, level + 1);
  recursiveHelper(root.right, hashTable, level + 1);

  return;
}
