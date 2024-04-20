//////////////////////////////////////////////////////
// *** Path Sum II ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: []

Example 3:
Input: root = [1,2], targetSum = 0
Output: []
 

Constraints:
The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
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
 * @param {number} targetSum
 * @return {number[][]}
 */
/*
BFS
The solution is very intutive
Very like backtracking problem

************************************************************
Time: O(n)
Space: O(h)
*/
var pathSum = function (root, targetSum) {
  const res = [];
  recursiveHelper(root, targetSum, res, []);
  return res;
};

function recursiveHelper(node, target, res, tmp) {
  // Base case
  if (node === null) {
    return;
  }

  // Hitting the leaf, check if can match the targetSum
  if (node.right === null && node.left === null) {
    // only add to the res array if targetSum is 0
    if (target - node.val === 0) {
      res.push([...tmp, node.val]);
    }
    return;
  }

  // choose
  tmp.push(node.val);

  // go left
  recursiveHelper(node.left, target - node.val, res, tmp);

  // go right
  recursiveHelper(node.right, target - node.val, res, tmp);

  // unchoose
  tmp.pop();
}

/*
BFS
Using iterative solution
The key point is that for each element in the stack
We have to store 3 element
1. the node itself
2. the current tmp array
3. the targetSum

The main pattern is very like pre-order BFS
*/
var pathSum = function (root, targetSum) {
  const stack = [[root, [], targetSum]];
  const res = [];

  while (stack.length > 0) {
    const [node, tmp, targetSum] = stack.pop();

    if (node === null) {
      continue;
    }

    // finding the leaf
    if (
      node.right === null &&
      node.left === null &&
      targetSum - node.val === 0
    ) {
      res.push([...tmp, node.val]);
    }

    // go right first(because of stack)
    if (node.right) {
      stack.push([node.right, [...tmp, node.val], targetSum - node.val]);
    }

    // go left
    if (node.left) {
      stack.push([node.left, [...tmp, node.val], targetSum - node.val]);
    }
  }

  return res;
};

/*
DFS
Using queue
Very like previous one, just change to using queue
*/
var pathSum = function (root, targetSum) {
  const queue = [[root, [], targetSum]];
  const res = [];

  while (queue.length > 0) {
    const [node, tmp, targetSum] = queue.shift();

    if (node === null) {
      continue;
    }

    if (
      node.right === null &&
      node.left === null &&
      targetSum - node.val === 0
    ) {
      res.push([...tmp, node.val]);
    }

    if (node.left) {
      queue.push([node.left, [...tmp, node.val], targetSum - node.val]);
    }

    if (node.right) {
      queue.push([node.right, [...tmp, node.val], targetSum - node.val]);
    }
  }

  return res;
};
