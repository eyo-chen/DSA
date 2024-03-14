//////////////////////////////////////////////////////
// *** Validate Binary Search Tree ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
 
Constraints:
The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1 
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
 * @return {boolean}
 */
/*
This problem becomes way more easier after watching the solution video
https://www.youtube.com/watch?v=MILxfAbIhrE

The main point is that we need to have our valid range for each state(call stack)
And have to properly update the valid range along with the process

For example,
                              5
                2                     9
           1         4           8         10
Use [] as valid range

                              5[-f, f]
                2                     9
           1         4           8         10
=> f means Infinity
=> at the root, root can be any value
=> go left first
=> Now we need to update the valid range
=> Because of going left, we know the node of left subtree has to be less than the current node
=> What does that mean?
=> It means the node of left subtree can be as small as possible
=> BUT it can not be greater than 5
=> Which means we need to update the upperbound, but lowerbound reamin the same
=> [-f, 5] -> pass to node of left subtree 

                              5[-f, f]
                2[-f, 5]                     9
           1         4           8         10
=> Is 2 in the valid range [5,f]?
=> YES
=> Keep going left
=> Same logic, update the upperbound, but lowerbound reamin the same
=> [2, f] -> pass to node of left subtree

                              5[-f, f]
                  2[-f, 5]                     9
           1[-f,2]         4           8             10
=> Is 1 in the valid range [-f,2] ?
=> YES
=> Keep going left
=> Hitting the base case (node === null)
=> return true
=> Now we're back at the state of 2[-f, 5]
=> Need to go right
=> Because of going right, we know the node of right subtree has to be greater than the current node
=> What does that mean?
=> It means the node of right subtree can be as large as possible
=> BUT it can not be less 2 
=> Which means we need to update the lowerbound, but upperbound reamin the same
=> [2, 5] -> pass to node of right subtree


                              5[-f, f]
                  2[-f, 5]                        9
           1[-f,2]         4[2,5]           8             10
=> Is 4 in the valid range [2,5] ?
=> YES
=> Keep going right
=> Hitting the base case (node === null)
=> return true
=> Now we're back at the state of 5[-f, f]
=> Need to go right
=> All the logic are the same

This is the result
                              5[-f, f]
                  2[-f, 5]                          9[5, f]
           1[-f,2]         4[2,5]           8[5,9]             10[10,f]
=> All node are in the valid range
=> And properly update the valid range
=> The main logic is 
=> If going left, the value can be as small as possible, but cannot be greater than the current node
=> Have to update the upperbound
=> If going right, the value can be as great as possible, but cannot be less than the current node
=> Have to update the lowerbound


Final thing to note
Initially, I write the solution is like
    if(root === null)
        return true;
    
    if(root.left === null && root.right === null)
        return true;
    
    if(root.left?.val >= root.val || root.right?.val <= root.val)
        return false;
    
    return isValidBST(root.left) && isValidBST(root.right);
=> This won't work
=> Because it only check the current node and it's left and right node
For example,
                              5
                2                     9
                                 3         10
=> In this case, the answer should be false
=> Tho 3 is less than 9, but it's also less than 5
=> Which is not valid
=> However, this case will return true in the code above
=> Because this code only check the current left and right node
=> 3 is less than 9, and 10 is greater than 9
=> then return true
=> But it's wrong

************************************************************
Time: O(n)
Space: O(h)
*/
var isValidBST = function (root) {
  return dfs(root, Infinity, -Infinity);
};

function dfs(node, maxVal, minVal) {
  if (node === null) return true;

  if (node.val >= maxVal || node.val <= minVal) return false;

  return dfs(node.left, node.val, minVal) && dfs(node.right, maxVal, node.val);
}
