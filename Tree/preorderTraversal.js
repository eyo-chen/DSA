//////////////////////////////////////////////////////
// *** Binary Tree Preorder Traversal ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,2,3]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]
 
Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
 
Follow up: Recursive solution is trivial, could you do it iteratively?
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
 * @return {number[]}
 */
/*
Recursive Solution
Very intuitive

************************************************************
Time: O(n)
Space: O(n)
*/
var preorderTraversal = function (root) {
  const res = [];

  recursiveHelper(root, res);
  return res;
};

function recursiveHelper(node, res) {
  if (node === null) return;

  // node, left, right
  res.push(node.val);
  recursiveHelper(node.left, res);
  recursiveHelper(node.right, res);

  return;
}

/*
Iterative Solution
This is also intuitive

However, one thing to note
res.push(node.val);
stack.push(node.right);
stack.push(node.left);
It seems we're doing node, right, left
But it's not
Because stack is First In Last Out
We store right first, and then store left
It makes sure next time we can add the value of left node first

For example,
                              5
                3                     6
        1            2         7            8
    4                       
res = []
stack = [] (First In Last Out)

Before starting while-loop, add root into stack
stack = [5]

(1)
node = 5
res = [5]
stack = [6, 3]

(2)
node = 3
res = [5, 3]
stack = [6, 2, 1]

(3)
node = 1
res = [5, 3, 1]
stack = [6, 2, 4]

(4)
node = 4
res = [5, 3, 1, 4]
stack = [6, 2]

(5)
node = 2
res = [5, 3, 1, 4, 2]
stack = [6]

(6)
node = 6
res = [5, 3, 1, 4, 2, 6]
stack = [8, 7]

(7)
node = 7
res = [5, 3, 1, 4, 2, 6, 7]
stack = [8]

(8)
node = 8
res = [5, 3, 1, 4, 2, 6, 7, 8]
stack = []

Stack is empty, done
Hope this process is clear why we put right into stack before left
https://www.youtube.com/watch?v=vMHaqhiTn7Y

************************************************************
Time: O(n)
Space: O(n)
*/
var preorderTraversal1 = function (root) {
  const res = [];
  const stack = [];
  stack.push(root);

  while (stack.length !== 0) {
    // get current node
    const node = stack.pop();

    // only do the operation if node is not null
    if (node) {
      // node, left, right
      res.push(node.val);
      stack.push(node.right);
      stack.push(node.left);
    }
  }
  return res;
};

/*
This is another iterative approach
But it's not intutive

For example,
                              5
                3                     6
        1            2         7            8
    4                       
res = []
stack = [] (First In Last Out)
node = 5

(1)
node = 5
res = [5]
stack = [6]
node = node.left -> node = 3

(2)
node = 3
res = [5, 3]
stack = [6, 2]
node = node.left -> node = 1

(3)
node = 1
res = [5, 3, 1]
stack = [6, 2]
node = node.left -> node = 4

(4)
node = 4
res = [5, 3, 1, 4]
stack = [6, 2]
node = node.left -> node = null

(5)
node = null
node = stack.pop() -> node = 2
res = [5, 3, 1, 4, 2]
stack = [6]
node = node.left -> node = null

(6)
node = null
node = stack.pop() -> node = 6
res = [5, 3, 1, 4, 2, 6]
stack = [8]
node = node.left -> node = 7

(7)
node = 7
res = [5, 3, 1, 4, 2, 6, 7]
stack = [8]
node = node.left -> node = null

(8)
node = null
node = stack.pop() -> node = 8
res = [5, 3, 1, 4, 2, 6, 7, 8]
stack = []

(9)
node = null
node = stack.pop() -> node = null
res = [5, 3, 1, 4, 2, 6, 7, 8]
stack = []
=> now stack is empty and node is null, stop the while-loop

This solution is not as intutive as previous one
But it's good to practice different solution
There are other ways to implement itertively on discussion of leetcode

************************************************************
Time: O(n)
Space: O(n)
*/
var preorderTraversal2 = function (root) {
  const res = [];
  const stack = [];
  let node = root;

  while (node !== null || stack.length !== 0) {
    if (node !== null) {
      res.push(node.val);
      stack.push(node.right);
      node = node.left;
    } else {
      node = stack.pop();
    }
  }

  return res;
};
