//////////////////////////////////////////////////////
// *** Maximum Depth of Binary Tree ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2
 
Constraints:
The number of nodes in the tree is in the range [0, 104].
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
 * @return {number}
 */
/*
The code is very easy to understand the logic

************************************************************
Time: O(n)
Space: O(h)
*/
var maxDepth = function (root) {
  return recursiveHelper(root, 0);
};

function recursiveHelper(root) {
  if (root === null) {
    return 0;
  }

  // return the max height + 1
  return Math.max(recursiveHelper(root.left), recursiveHelper(root.right)) + 1;
}

/*
This is iterative approach to implement DFS

The main logic is same
But now we use array as stack
Also store 1) node 2) height in the stack
The code basically is the DFS

************************************************************
Time: O(n)
Space: O(h)
*/
var maxDepth = function (root) {
  // Base case
  if (root === null) return 0;

  const stack = [[root, 1]];
  let maxHeight = 0;

  while (stack.length > 0) {
    // store 1) node 2) height in the stack
    const [node, height] = stack.pop();

    // find the max depth
    maxHeight = Math.max(maxHeight, height);

    // only store the node if it's not null
    if (node.right !== null) stack.push([node.right, height + 1]);

    // only store the node if it's not null
    if (node.left !== null) stack.push([node.left, height + 1]);
  }

  return maxHeight;
};

/*
This is using queue to implement BFS
We go level by level
One thing to note
We'll use for-loop to run off the node in current level inside while-loop
For example, 
                              5
                3                     6
        1            2         7            8
    4    

First level
queue = [5]
queue.length = 1
=> for-loop one time
=> queue = [3, 6]

Second level
queue = [3, 6]
queue.length = 2
=> for-loop two times
=> queue = [1, 2, 7, 8]

Third level
queue = [1, 2, 7, 8]
queue.length = 4
=> for-loop four times
=> queue = [4]

Four level
=> done

So, for each level
we'll run off the all node in current level

************************************************************
Time: O(n)
Space: O(h)
*/
var maxDepth = function (root) {
  // Base case
  if (root === null) return 0;

  const queue = [root];
  let height = 0;

  while (queue.length > 0) {
    // after going one level depp, increment the height
    // it means we go to the next level, so the height is updated
    height++;

    // Have to store the current length of queue
    // Because in the for-loop below, we'll keep adding children node
    const curLength = queue.length;

    for (let i = 0; i < curLength; i++) {
      const node = queue.shift();

      if (node.left !== null) queue.push(node.left);

      if (node.right !== null) queue.push(node.right);
    }
  }

  return height;
};
