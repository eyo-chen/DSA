//////////////////////////////////////////////////////
// *** Lowest Common Ancestor of a Binary Search Tree ***
//////////////////////////////////////////////////////
/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [2,1], p = 2, q = 1
Output: 2
 
Constraints:
The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the BST.
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
/*
At first, I over think this problem
Which means I try to solve this problem in a more harder way
It turns out this problem should be very easy because it's BST(binary search tree)
And we know that two node p and q definitely in the tree


                              15
                10                     25
        5            13           23         30
    1     8                               29     33

For each node, we always have four situations
1. the value of node is greater than the value of p and q
=> For example, node is 15, p = 5, q = 5
=> The answer is on the left sub tree, go to left subtree

2. the value of node is smaller than the value of p and q
=> For example, node is 15, p = 30, q = 29
=> The answer is on the right sub tree, go to right subtree

3. the value of node is equal to the value of q or p
=> For example, node is 10, p = 10, q = 1
=> The answer is on the node, return it

4. the value of node is in the middle of value of p and q
=> For example, node is 25, p = 23, q = 33
=> The answer is on the node, return it 

************************************************************
Time: O(log(n))
=> Because we always cut the tree in half, aka only go to subtree

Space: O(h)
*/
var lowestCommonAncestor = function (root, p, q) {
  // the value of node is greater than the value of p and q
  if (root.val > p.val && root.val > q.val)
    return lowestCommonAncestor(root.left, p, q);

  // the value of node is smaller than the value of p and q
  if (root.val < p.val && root.val < q.val)
    return lowestCommonAncestor(root.right, p, q);

  // the value of node is equal to the value of q or p
  // the value of node is in the middle of value of p and q
  return root;
};

/*
This is iterative approach using queue
Basically same idea as before
*/
var lowestCommonAncestor = function (root, p, q) {
  const queue = [root];
  while (queue.length > 0) {
    const node = queue.pop();

    // the value of node is greater than the value of p and q
    if (q.val < node.val && p.val < node.val) {
      queue.push(node.left);
    }

    // the value of node is smaller than the value of p and q
    else if (q.val > node.val && p.val > node.val) {
      queue.push(node.right);
    }

    // the value of node is equal to the value of q or p
    // the value of node is in the middle of value of p and q
    else {
      return node;
    }
  }
};

/*
Simpler
*/

var lowestCommonAncestor = function (root, p, q) {
  while (root) {
    if (q.val < root.val && p.val < root.val) {
      root = root.left;
    } else if (q.val > root.val && p.val > root.val) {
      root = root.right;
    } else {
      return root;
    }
  }
};
