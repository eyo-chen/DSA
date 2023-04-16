//////////////////////////////////////////////////////
// *** Minimum Depth of Binary Tree ***
//////////////////////////////////////////////////////
/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 2

Example 2:
Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5
 
Constraints:
The number of nodes in the tree is in the range [0, 105].
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
 * @return {number}
 */
/*
This is wrote by myself
The code does work, but it's not clean

Because the problem is finding the min depth
We need to handle different case
       5  
         6
           7
             8
=> If the tree looks like this, we can not just return 1

The idea is 
1. Inititally set left and right to false
=> It means "haven't traverse"

2. If the subtree is not null, then we can go down to traverse the tree

3. If left and right both are false, it means it's leaf node
=> Just return 1

4. If left subtree is false, the tree looks like this
       5
null      4
       3    1
=> just return the height of right subtree

5. Same thing for right subtree
6. If both right and left subtree are not null, just return the min depth of tow of them

************************************************************
Time: O(n)
Space: O(h)
*/
var minDepth = function (root) {
  if (root === null) return 0;

  let left = false,
    right = false;

  if (root.left !== null) left = minDepth(root.left) + 1;

  if (root.right !== null) right = minDepth(root.right) + 1;

  if (left === false && right === false) return 1;
  if (left === false) return right;
  if (right === false) return left;

  return Math.min(left, right);
};

/*
This code is kind of similar to previous one
But the code is cleaner
The logic is slightl different

return left && right ? Math.min(left, right) + 1 : left + right + 1;
=> Look at the base case, if hitting the null, we'll return 0
=> If one of left and right is null, the tree is like this
      5
null    6
=> We can't just return Math.min(left, right) + 1
=> Because the returned value is 0 + 1
=> We have to return right + 1
=> Which is identical to left + right + 1
=> Because left is 0

So the main logic is 
If one of left and right is 0, we'll return the other one + 1
=> Because one of them is 0, so we can just say return left + right + 1;
=> I don't care which one is 0,

If both of them is not 0
=> We can just return Math.min(left, right) + 1

************************************************************
Time: O(n)
Space: O(h)
*/
var minDepth = function (root) {
  if (root === null) return 0;

  const left = minDepth(root.left);
  const right = minDepth(root.right);

  return left && right ? Math.min(left, right) + 1 : left + right + 1;
};

/*
This is BFS solution which is more efficient than DFS
The worst case is still the same, but it's more efficent in some edge case
Image if there is a tree, it's left subtree has heigh 20, and right subtree is 2
In DFS, we have to all the way go as deep as the height of left subtree
And get the height is 20,
then go to right, and find out the min depth is 2
This is not efficient

In BFS, we go level by level
Once finding the leaf(node.left === null && node.right === null))
We know we find the min depth
I don't care it's in the right or left subtree
It's definitely the min depth

One thing to note
In the process of while-loop, we have to store the 1) node 2) height

************************************************************
Time: O(n)
Space: O(h)
*/
var minDepth = function (root) {
  if (root === null) return 0;

  // store the 1) node 2) height
  const queue = [[root, 1]];

  while (queue.length > 0) {
    const [node, height] = queue.shift();

    // if node is null, just skip it
    /*
         5
     4       2
3         3    1
                 3
    => the right subtree of 4 is null, just skip it, it's not leaf
    => Which means there's deeper height at the left subtree
    */
    if (node === null) continue;

    // find the leaf, return the height
    if (node.left === null && node.right === null) return height;

    queue.push([node.left, height + 1]);
    queue.push([node.right, height + 1]);
  }
};
