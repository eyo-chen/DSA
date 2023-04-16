//////////////////////////////////////////////////////
// *** Construct Binary Tree from Preorder and Inorder Traversal ***
//////////////////////////////////////////////////////
/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]
 

Constraints:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
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
 * @param {number[]} preorder
 * @param {number[]} inorder
 * @return {TreeNode}
 */
/*
I actually try to find the solution before watching the video
I got the 60% the main logic of this problem, but still fail
The main logic is not that hard to understand after finding the logic

preorder
=> Give the "root" information

inorder
=> Give the "left and right subtree" information

left and right give us the boundary
Help us to hitting the base case

preIndex helps us to get the current root of subtree

For example,
preorder = [3, 9, 8, 12, 13, 20, 15, 7];
inorder = [12, 8, 13, 9, 3, 15, 20, 7];


Stack 1
boundary = [0, 7]
The first element in the prorder is the root of tree
node = 3
preIndex = 1 (after preIndex ++)
mid = 4
Go left
                       3

Stack 2
boundary = [0, 3]
node = preorder[preIndex] = 9
preIndex = 2 (after preIndex ++)
mid = 3
Go left
                       3
                9

Stack 3
boundary = [0, 2] [0, preMid - 1]
node = preorder[preIndex] = 8
preIndex = 3 (after preIndex ++)
mid = 1
Go left
                       3
                9
        8

Stack 4
boundary = [0, 0] [0, preMid - 1]
node = preorder[preIndex] = 12
preIndex = 4 (after preIndex ++)
mid = 0
Go left
                       3
                9
        8
    12

Stack 5
boundary = [0, -1] [0, preMid - 1]
Base case
                       3
                9
        8
    12

Stack 4
boundary = [0, 0] [0, preMid - 1]
node = preorder[preIndex] = 13
preIndex = 5 (after preIndex ++)
mid = 2
Go left
Go right
                       3
                9
        8
    12   13

Stack 5
boundary = [3, 0] [preMid + 1, 0]
Base case
                       3
                9
        8
    12   13

Stack 2
boundary = [0, 3]
node = preorder[preIndex] = 9
preIndex = 5 
mid = 3
Go left
Go right
                       3
                9
        8
    12   13

Stack 3
boundary = [4, 3] [preMid + 1, 3]
Base case
                       3
                9
        8
    12   13

Stack 1
boundary = [0, 7]
The first element in the prorder is the root of tree
node = 3
preIndex = 5
mid = 4
Go left
Go right
                       3
                9
        8
    12   13

Stack 2
boundary = [4, 7] [preMid + 1, 7]
node = 20
preIndex = 6
mid = 6
Go left
                       3
                9          20
        8
    12   13

Stack 3
boundary = [4, 6] [4, preMid - 1]
node = 15
preIndex = 7
mid = 5
Go left
                       3
                9          20
        8              15
    12   13
....


Hope can see the pattern
There are two things to note along with the process
1. We have to put preIndex outside the recursive function
=> preIndex gives us the root of sub tree
=> When we all the way go left, we'll keep updating
=> So it can make sure we have the correct preIndex

3. The reason to find min in each call stack is because later we'll use that to cut off the left and right subtree

4. When going left, we need to update right index(mid - 1)
5. When going right, we need to update left index(mid + 1)
=> For example,
preorder = [3, 9, 8, 12, 13, 20, 15, 7];
inorder = [12, 8, 13, 9, 3, 15, 20, 7];
=> First finding 3 as root
=> the left sub tree is [12, 8, 13, 9]
=> the right sub tree is [15, 20, 7]
=> note that inorder give us the left and right subtree information

************************************************************
Time: O(n ^ 2)
=> Traverse all the node
=> const mid = inorder.findIndex(val => val === preorder[preIndex]);
=> O(n) works

Space: O(h)
*/
var buildTree = function (preorder, inorder) {
  let preIndex = 0;
  function recursiveHelper(preorder, inorder, left, right) {
    if (left > right) {
      return null;
    }

    // current node
    const node = new TreeNode(preorder[preIndex]);

    // the mid point of inorder(later need to cut off the left and right subtree)
    const mid = inorder.findIndex(val => val === preorder[preIndex]);

    // keep updating preIndex
    preIndex++;

    node.left = recursiveHelper(preorder, inorder, left, mid - 1);
    node.right = recursiveHelper(preorder, inorder, mid + 1, right);

    return node;
  }

  return recursiveHelper(preorder, inorder, 0, preorder.length - 1);
};

/*
Same logic
But use Map to do O(1) work when finding the mid
Build the key-value pair of inorder array

************************************************************
Time: O(n)
Space: O(n)
=> Build map -> O(n) work
=> Call stack -> O(h)
*/
var buildTree = function (preorder, inorder) {
  const map = new Map();
  let preIndex = 0;

  // Build map
  for (let i = 0; i < inorder.length; i++) {
    map.set(inorder[i], i);
  }

  function recursiveHelper(preorder, inorder, left, right) {
    if (left > right) {
      return null;
    }

    const node = new TreeNode(preorder[preIndex]);

    // O(1) work
    const mid = map.get(preorder[preIndex]);
    preIndex++;

    node.left = recursiveHelper(preorder, inorder, left, mid - 1);
    node.right = recursiveHelper(preorder, inorder, mid + 1, right);

    return node;
  }

  return recursiveHelper(preorder, inorder, 0, preorder.length - 1);
};

var buildTree = function (preorder, inorder) {
  let preIndex = 0;
  // We can use reduce to create hashTable !!!NICE!!!
  const hashTable = inorder.reduce((acc, cur, index) => {
    acc[cur] = index;
    return acc;
  }, {});

  function recursiveHelper(preorder, inorder, left, right) {
    if (left > right) {
      return null;
    }

    const node = new TreeNode(preorder[preIndex]);
    const mid = hashTable[preorder[preIndex]];
    preIndex++;

    node.left = recursiveHelper(preorder, inorder, left, mid - 1);
    node.right = recursiveHelper(preorder, inorder, mid + 1, right);

    return node;
  }

  return recursiveHelper(preorder, inorder, 0, preorder.length - 1);
};
