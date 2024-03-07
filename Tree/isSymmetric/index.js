//////////////////////////////////////////////////////
// *** Symmetric Tree ***
//////////////////////////////////////////////////////
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false

Constraints:
The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100
 
Follow up: Could you solve it both recursively and iteratively?
*/
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
/*
Recursive Solution
                                       1
                    2                                   2
            3              4               4                       3

The main point is after checking the first left and right node (2 & 2)
then we're gonna check left.right & right.left and left.left & right.right
so that we can check if it's symmetric

************************************************************
Time: O(n)
Space: O(h)
*/
var isSymmetric = function (root) {
  if (root === null) return true;
  return checkSymmetric(root.left, root.right);
};

function checkSymmetric(left, right) {
  // it's a leaf
  if (right === null && left === null) return true;

  if (right !== null && left !== null)
    return (
      right.val === left.val &&
      checkSymmetric(left.right, right.left) &&
      checkSymmetric(left.left, right.right)
    );

  // this is the case when one of leaf is null, and the other has value
  /*
  Like this
                                       1
                    2                                   null
  In this case, right === null, but left !== null, so it won't pass through first condition
  Also, it won't pass through second condition
  If one of leaf is null, just return false, it's not symmetric
  */
  return false;
}

/*
Iterative Approach
Using the idea of bfs

                                       1
                    2a                                   2b
            3a              4a               4b                       3b

q = [2a, 2b]
left = 2a
right = 2b
queue.push(left.left) -> [3a]
queue.push(right.right) -> [3a, 3b]
queue.push(left.right) -> [3a, 3b, 4a]
queue.push(right.left) -> [3a, 3b, 4a, 4b]

q = [3a, 3b, 4a, 4b]
left = 3a
right = 3b
queue.push(left.left) -> [4a, 4b, null]
queue.push(right.right) -> [4a, 4b, null, null]
queue.push(left.right) -> [4a, 4b, null, null, null]
queue.push(right.left) -> [4a, 4b, null, null, null, null]

q = [4a, 4b, null, null, null, null]
left = 4a
right = 4b
queue.push(left.left) -> [4a, 4b, null]
queue.push(right.right) -> [4a, 4b, null, null]
queue.push(left.right) -> [4a, 4b, null, null, null]
queue.push(right.left) -> [4a, 4b, null, null, null, null]
=> and the queue is all null, we'll keep shifting out of the value, and the queue will be empty

************************************************************
Time: O(n)
Space: O(h)
*/
var isSymmetric = function (root) {
  if (root === null) return true;

  const queue = [root.left, root.right];

  while (queue.length > 0) {
    // get out the value(the order does matter)
    const left = queue.shift();
    const right = queue.shift();

    // if it's null, just skip it, don't need to keep adding further value
    if (left === null && right === null) continue;

    // if one of node is null or the value is not equal, it's not symmetric
    if (left === null || right === null || left.val !== right.val) return false;

    // add the value into queue in order
    queue.push(left.left);
    queue.push(right.right);
    queue.push(left.right);
    queue.push(right.left);
  }

  return true;
};
