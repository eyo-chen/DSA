//////////////////////////////////////////////////////
// *** Binary Tree Postorder Traversal ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

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
Very intuitve

************************************************************
Time: O(n)
Space: O(n)
*/
var inorderTraversal = function (root) {
  const res = [];
  recursiveHelper(root, res);
  return res;
};

function recursiveHelper(node, res) {
  if (!node) return;

  recursiveHelper(node.left, res);
  res.push(node.val);
  recursiveHelper(node.right, res);
}

/*
Iterative Solution

For example,
                              5
                3                     6
        1            2         7            8
    4      

res = []
stack = []
node = 5

res = []
stack = [5]
node = 3

res = []
stack = [5, 3]
node = 1

res = []
stack = [5, 3, 1]
node = 4

res = []
stack = [5, 3, 1, 4]
node = null
=> node = stack.pop() = 4
=> res = [4]
=> stack = [5, 3, 1]
=> node = node.right = 4.right = null

res = [4]
stack = [5, 3, 1]
node = null
=> node = stack.pop() = 1
=> res = [4, 1]
=> stack = [5, 3]
=> node = node.right = 1.right = null

res = [4, 1]
stack = [5, 3]
node = null
=> node = stack.pop() = 3
=> res = [4, 1, 3]
=> stack = [5]
=> node = node.right = 3.right = 2

res = [4, 1, 3]
stack = [5]
node = 2

res = [4, 1, 3]
stack = [5, 2]
node = null
=> node = stack.pop() = 2
=> res = [4, 1, 3, 2]
=> stack = [5]
=> node = node.right = 2.right = null

res =[4, 1, 3, 2]
stack = [5]
node = null
=> node = stack.pop() = 5
=> res = [4, 1, 3, 2, 5]
=> stack = []
=> node = node.right = 5.right = 6

res = [4, 1, 3, 2, 5]
stack = [6]
node = 7

res = [4, 1, 3, 2, 5]
stack = [6, 7]
node = null
=> node = stack.pop() = 7
=> res = [4, 1, 3, 2, 5, 7]
=> stack = [6]
=> node = node.right = 7.right = null

res = [4, 1, 3, 2, 5, 7]
stack = [6]
node = null
=> node = stack.pop() = 6
=> res = [4, 1, 3, 2, 5, 7, 6]
=> stack = []
=> node = node.right = 6.right = null
=> stack is empty, node is null
=> done

Go to leetcode to see the process of graph
https://leetcode.com/problems/binary-tree-inorder-traversal/
************************************************************
Time: O(n)
Space: O(n)
*/
var inorderTraversal = function (root) {
  const res = [];
  const stack = [];
  let node = root;

  while (stack.length > 0 || node !== null) {
    // keep going left
    while (node !== null) {
      stack.push(node);
      node = node.left;
    }
    // read the value
    node = stack.pop();
    res.push(node.val);

    // go right
    node = node.right;
  }

  return res;
};
