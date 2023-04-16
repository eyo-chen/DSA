//////////////////////////////////////////////////////
// *** Construct Binary Tree from Inorder and Postorder Traversal ***
//////////////////////////////////////////////////////
/*
Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: inorder = [-1], postorder = [-1]
Output: [-1]
 
Constraints:
1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder and postorder consist of unique values.
Each value of postorder also appears in inorder.
inorder is guaranteed to be the inorder traversal of the tree.
postorder is guaranteed to be the postorder traversal of the tree.
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
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */
/*
This problem is very simlilar to previous one
But it gives us postorder
The logic is the same
But now we need to build right tree first

************************************************************
Time: O(n)
Space: O(n)
*/
var buildTree = function (inorder, postorder) {
  const cacheObj = inorder.reduce((acc, cur, index) => {
    acc[cur] = index;
    return acc;
  }, {});

  let postIndex = postorder.length - 1;

  function recursiveHelper(inorder, postorder, cacheObj, startIndex, endIndex) {
    if (startIndex > endIndex) {
      return null;
    }

    const rootVal = postorder[postIndex];
    const root = new TreeNode(rootVal);
    const rootIndex = cacheObj[rootVal];
    postIndex--;

    root.right = recursiveHelper(
      inorder,
      postorder,
      cacheObj,
      rootIndex + 1,
      endIndex
    );
    root.left = recursiveHelper(
      inorder,
      postorder,
      cacheObj,
      startIndex,
      rootIndex - 1
    );

    return root;
  }

  return recursiveHelper(inorder, postorder, cacheObj, 0, inorder.length - 1);
};
