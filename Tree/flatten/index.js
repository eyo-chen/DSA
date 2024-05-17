//////////////////////////////////////////////////////
// *** Flatten Binary Tree to Linked List ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.
 

Example 1:
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [0]
Output: [0]
 
Constraints:
The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
 
Follow up: Can you flatten the tree in-place (with O(1) extra space)?
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
 * @return {void} Do not return anything, modify root in-place instead.
 */
/*
This is the solution I wrote by myself
Using DFS, recursion, pre-order traversal

For example,
                       1
            2                       5
    3            4         null         6

Firsr using recursion to store the node in pre-order into stack
So the stack will be
[1, 2, 3, 4, 5, 6] after recursion

Then we pop the element off the stack,
set it's left to null
set it's right to previous node

Also, need to update the previous node

1. 
preNode = null
6.left = null
6.right = null

2. 
preNode = 6
5.left = 6
5.right = null

3. 
preNode = 5
4.left = 5
3.right = null

... so on and so forth

************************************************************
Time: O(n)

Space: O(n)
=> Using stack to store all the node
*/
var flatten = function (root) {
  const stack = [];
  recursiveHelper(root, stack);
  let preNode = null;
  let curNode = null;

  while (stack.length > 0) {
    curNode = stack.pop();

    curNode.left = null;
    curNode.right = preNode;

    preNode = curNode;
  }
};

function recursiveHelper(root, stack) {
  if (root === null) return;

  stack.push(root);
  recursiveHelper(root.left, stack);
  recursiveHelper(root.right, stack);
}

/*
This is the solution from discuss
Using the reverse pre-order

                       1
            2                       5
    3            4         null         6

post-order -> left, right, node
reverse post-order -> right, left, node
If we're using reverse post-order, the array would be like this
[6,5,4,3,2,1]
which is very similar to the final result we want

The following pattern is very similar to previous solution
Note that the first node to handle is 6

1. 
preNode = null
6.left = null
6.right = null

2. 
preNode = 6
5.left = null
5.right = 6

3. 
preNode = 5
4.left = 5
3.right = null

... so on and so forth

************************************************************
Time: O(n)
Space: O(h)
*/
var flatten = function (root) {
  let preNode = null;

  function recursiveHelper(root) {
    if (root === null) return null;

    recursiveHelper(root.right);
    recursiveHelper(root.left);

    root.left = null;
    root.right = preNode;

    preNode = root;
  }

  recursiveHelper(root);
};

/*
Yet another great solution
The idea is also using stack

Just see the process
                       1
            2                       5
    3            4         null         6

1.
curNode = 1
stack = [5, 2]
curNode.right = stack[stack.length - 1];
=> 1.right = 2
=> 1.left = null

2. 
curNode = 2
stack = [5, 4, 3]
curNode.right = stack[stack.length - 1];
=> 2.right = 3
=> 2.left = null

3. 
curNode = 3
stack = [5, 4]
curNode.right = stack[stack.length - 1];
=> 3.right = 4
=> 3.left = null

4. 
curNode = 4
stack = [5]
curNode.right = stack[stack.length - 1];
=> 4.right = 5
=> 4.left = null

so on and so forth ...

Go to see video https://www.youtube.com/watch?v=vssbwPkarPQ

This is pretty hard to come up at first time
************************************************************
Time: O(n)
Space: O(h)
*/
var flatten = function (root) {
  if (root === null) return root;

  const stack = [root];

  while (stack.length > 0) {
    const curNode = stack.pop();

    if (curNode.right !== null) {
      stack.push(curNode.right);
    }

    if (curNode.left !== null) {
      stack.push(curNode.left);
    }

    if (stack.length > 0) {
      curNode.right = stack[stack.length - 1];
    }

    curNode.left = null;
  }
};
