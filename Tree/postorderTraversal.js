//////////////////////////////////////////////////////
// *** Binary Tree Postorder Traversal ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [3,2,1]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]
 
Constraints:
The number of the nodes in the tree is in the range [0, 100].
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
Recursive solution is very intutive

************************************************************
Time: O(n)
Space: O(n)
*/
var postorderTraversal = function (root) {
  const res = [];
  recursiveHelper(root, res);
  return res;
};

function recursiveHelper(node, res) {
  if (!node) return;
  // left, right, node
  recursiveHelper(node.left, res);
  recursiveHelper(node.right, res);
  res.push(node.val);
  return;
}

/*
Iterative solution is very non-intutive


For example,
                              5
                3                     6
        1            2         
                         5
                     10     18

res = []
stack = []
node = 5

Start

res = []
stack = [5]
node = 3

res = []
stack = [5, 3]
node = 1

res = []
stack = [5, 3, 1]
node = null

res = []
stack = [5, 3, 1]
node = null
tmp = 1.right = null
=> tmp = stack.pop() = 1
=> res = [1]
=> stack = [5, 3]
=> tmp === stack[stack.length - 1].right -> false

res = [1]
stack = [5, 3]
node = null
=> tmp = 3.right = 2
=> node = 2

res = [1]
stack = [5, 3, 2]
node = null
=> tmp = 2.right = 5
=> node = 5

res = [1]
stack = [5, 3, 2, 5]
node = 10

res = [1]
stack = [5, 3, 2, 5, 10]
node = null
tmp = 10.right = null
=> tmp = stack.pop() = 10
=> res = [1, 10]
=> stack = [5, 3, 2, 5]
=> tmp === stack[stack.length - 1].right
=> 10 === 5.right -> false

res = [1, 10]
stack = [5, 3, 2, 5]
node = null
tmp = 5.right = 18
=> node = 18

res = [1, 10]
stack = [5, 3, 2, 5, 18]
node = null
tmp = 18.right = null
=> tmp = stack.pop() = 18
=> res = [1, 10, 18]
=> stack = [5, 3, 2, 5]
=> tmp === stack[stack.length - 1].right
=> 18 = 5.right -> true
=> tmp = stock.pop() = 5
=> res = [1, 10, 18, 5]
=> stack = [5, 3, 2]
=> tmp === stack[stack.length - 1].right
=> 5 = 2.right -> true
=> tmp = stack.pop() = 2
=> res = [1, 10, 18, 5, 2]
=> stack = [5, 3]
=> tmp === stack[stack.length - 1].right
=> 2 = 3.right -> true
=> tmp = stack.pop() = 3
=> res = [1, 10, 18, 5, 2, 3]
=> stack = [5]
=> tmp === stack[stack.length - 1].right
=> 3 = 5.right -> false

res = [1, 10, 18, 5, 2, 3]
stack = [5]
node = null
=> tmp = 5.right = 6
=> node = 6

res = [1, 10, 18, 5, 2, 3]
stack = [5, 6]
node = null
=> tmp = 6.right = null
=> tmp = stack.pop() = 6
=> res = [1, 10, 18, 5, 2, 3, 6]
=> stack = [5]
=> tmp === stack[stack.length - 1].right
=> 6 = 5.right -> true
=> tmp = stack.pop() = 5
=> res = [1, 10, 18, 5, 2, 3, 6, 5]
=> stack is empty
=> node = null
=> done

This solution is very hard
go to this video to see more clear
https://www.youtube.com/watch?v=xLQKdq0Ffjg

************************************************************
Time: O(n)
Space: O(n)
*/
var postorderTraversal = function (root) {
  const res = [];
  const stack = [];
  let node = root;

  while (stack.length !== 0 || node !== null) {
    // go left
    if (node !== null) {
      stack.push(node);
      node = node.left;
    }
    // can't go left, go right
    else {
      // right node
      let tmp = stack[stack.length - 1].right;

      // right node is null -> have already go both left and right
      // get the value of node
      if (tmp === null) {
        tmp = stack.pop();
        res.push(tmp.val);

        // important part -> see if the next node has already go both left and right
        // also get the value of node
        while (stack.length !== 0 && tmp === stack[stack.length - 1].right) {
          tmp = stack.pop();
          res.push(tmp.val);
        }
      }
      // right node is not null, keep going left first
      else {
        node = tmp;
      }
    }
  }

  return res;
};
