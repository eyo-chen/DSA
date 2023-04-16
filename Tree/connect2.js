//////////////////////////////////////////////////////
// *** Populating Next Right Pointers in Each Node II ***
//////////////////////////////////////////////////////
/*
Given a binary tree

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Example 1:
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Example 2:
Input: root = []
Output: []
 
Constraints:
The number of nodes in the tree is in the range [0, 6000].
-100 <= Node.val <= 100
 
Follow-up:
You may only use constant extra space.
The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.
*/
/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */
/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */
/**
 * @param {Node} root
 * @return {Node}
 */
/*
BFS, queue, iterative
This is solution I wrote by myself

The idea is very similar to previous variant
We need to handle left and right subtree
But now it's not perfect binary tree
So we need wireNextNode function
This function only does one thing
=> Find the next pointer to wire up

*/
var connect = function (root) {
  const queue = [root];

  while (queue.length > 0) {
    const curNode = queue.shift();

    // skip the process when curNode is null
    if (curNode === null) continue;

    // handle the left subtree
    if (curNode.left !== null) {
      let node = curNode;

      /*
      right does exist, just easily wire up
           4
        3   1
      */
      if (node.right !== null) node.left.next = node.right;
      /*
      right does not exist, find the next value
             5         8
        1      null  9    10
      => wire 1 -> 9
      */ else {
        wireNextNode(node.left, node);
      }
    }

    // handle right subtree, same thing to check if it has next value
    if (curNode.right !== null && curNode.next !== null) {
      let node = curNode;

      /*
      We don't need to check if next.left exist
      Because it's part of the logic of wireNextNode
      Just find the next value to wire up
      */
      wireNextNode(node.right, node);
    }

    queue.push(curNode.left);
    queue.push(curNode.right);
  }

  return root;
};
/*
There are three situations
First 
                5           8
           1      null   9    10
=> In this case, we have to find the next value to wire 1.next
=> We just have to find the first value to wire up
=> In this case, the first next value is 9
=> So 1 -> 9

Second
                5             8
           1      null   null    10
=> In this case, the first next value is 10
=> So 1 -> 10

Third
                5             8
           1      null   null    null
=> In this case, there's no next value
=> Don't need to do anything
*/
function wireNextNode(targetNode, node) {
  while (true) {
    // keep going next(right) to find the value(ptr)
    node = node.next;

    // if one of them is true, just break out, we find it
    if (node === null || node.left !== null || node.right !== null) break;
  }

  // if it's null, third situation, just don't do anything
  if (node !== null) {
    /*
    we have to first check left
    In this case
                    5           8
                1      null   9    10
    => even tho 10 does exist, we have to wire the first next value
    => In this case, is 9
    */
    if (node.left !== null) targetNode.next = node.left;
    // Second case
    else if (node.right !== null) targetNode.next = node.right;
  }
}

/*
DFS, recursive

All the logic is same

Note that we have to wire right subtree first
Because there's one edge case like this
                    5           8
                1      null   9    10

Imagine
5 is at the right subtree of root, 
and 8 is at the left subtree of root
When we want to wire next ptr of 1
We need to do sth like this 
5.next.left 
However, 5 haven't wire to 8
So  5.next is null

So we have to wire right subtree first

[2,1,3,0,7,9,1,2,null,1,0,null,null,8,8,null,null,null,null,7]
try this testing case with left is in front of the right
*/
var connect = function (root) {
  if (root === null) return null;

  if (root.left !== null) {
    let curNode = root;

    if (curNode.right !== null) root.left.next = curNode.right;
    else {
      wireNextNode(root.left, curNode);
    }
  }

  if (root.right !== null && root.next !== null) {
    let curNode = root;
    wireNextNode(root.right, curNode);
  }

  root.right = connect(root.right);
  root.left = connect(root.left);

  return root;
};
