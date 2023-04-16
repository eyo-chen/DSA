//////////////////////////////////////////////////////
// *** Delete Node in a BST ***
//////////////////////////////////////////////////////
/*
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.

Example 1:
Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:
Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.

Example 3:
Input: root = [], key = 0
Output: []
 
Constraints:
The number of nodes in the tree is in the range [0, 104].
-105 <= Node.val <= 105
Each node has a unique value.
root is a valid binary search tree.
-105 <= key <= 105 

Follow up: Could you solve it with time complexity O(height of tree)?
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
 * @param {number} key
 * @return {TreeNode}
 */
/*
This problem is kind of hard
The process is like
1. keep finding the target node in a BST
2. After finding, there are four different situations
   1) both left and right are null
          6
    4         10
            x    x
    => target node = 10
    => If we wanna delete 10, just return null to 6
    => 6.right = null
    
    2) left is not null, but right is null
          6
    4         10
            7     x
    => target node = 10
    => If we wanna delete 10, just return 7 to 6
    => 6.right = 7

    3) right is not null, but left is null
          6
    4         10
            x     12
    => target node = 10
    => If we wanna delete 10, just return 12 to 6
    => 6.right = 12

    4) both right and left are not null
                               50
              4                                   70
                                      65                     80
                                                    75              100
                                              73        77
    => target node = 70
    => After deletion
                               50
              4                                   73
                                      65                     80
                                                    75              100
                                                        77
    => 1) we have to first go to right
    => node = 70.right = 80
    => 2) all the way to find the min number in the left subtree
    => node = 73
    => 3) replace the value of 70 to 73
    => 4) use 80 as root, invoke the function to delete 73

This is the whole process
Once understanding the logic, it's not that hard
The fourth case is very tricky

************************************************************
Time: O(h)
Space: O(h)
*/
var deleteNode = function (root, key) {
  if (root === null) return null;

  if (root.val > key) root.left = deleteNode(root.left, key);
  else if (root.val < key) root.right = deleteNode(root.right, key);
  else {
    // First situation
    if (root.right === null && root.left === null) return null;

    // Second situation
    if (root.left !== null && root.right === null) return root.left;

    // Third situation
    if (root.right !== null && root.left === null) return root.right;

    // Fourth situation
    let node = root.right;

    while (node.left !== null) {
      node = node.left;
    }

    root.val = node.val;
    root.right = deleteNode(root.right, node.val);

    return root;
  }

  return root;
};
