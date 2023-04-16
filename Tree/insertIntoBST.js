//////////////////////////////////////////////////////
// *** Insert into a Binary Search Tree ***
//////////////////////////////////////////////////////
/*
You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

Example 1:
Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:

Example 2:
Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]

Example 3:
Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]
 
Constraints:
The number of nodes in the tree will be in the range [0, 104].
-108 <= Node.val <= 108
All the values Node.val are unique.
-108 <= val <= 108
It's guaranteed that val does not exist in the original BST.
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
 * @param {number} val
 * @return {TreeNode}
 */
/*
This is the solution I wrote by myself before seeing the solution

The code does work, but the code is not clean, and it's verbose
The main logic is simple
=> We don't check if current node is greater or less than target val
=> Just keeping going right and left
=> If the value of current node is out of the bound, we know we have to go back up
  if (val <= start || val >= end) {
    return root;
  }
=> If now the current node is leaf, we're gonna check if we can insert the node below this leaf
  if (root.left === null && root.right === null) {
    if (val > start && val < root.val) {
      root.left = new TreeNode(val);
    } else if (val > root.val && val < end) {
      root.right = new TreeNode(val);
    }
    return root;
  }
=> If hitting the null, we're gonna check if we can insert the node in this position

The main logic is very simple,
We don't check the value beforehand, we just do the whole search
When hitting the base case(leaf or null)
Then we check if it's okay to insert the node
Along with the whole process, we need to have a boundary[start, end]

************************************************************
Time: O(n)
Space: O(h)
*/
var insertIntoBST = function (root, val) {
  // Base case
  if (root === null) {
    return new TreeNode(val);
  }
  return recursiveHelper(root, val, -Infinity, Infinity);
};

function recursiveHelper(root, val, start, end) {
  // the value of current node is out of the bound, stop, and go back up
  if (val <= start || val >= end) {
    return root;
  }

  // it's a leaf
  if (root.left === null && root.right === null) {
    // insert the node at the left subtree
    if (val > start && val < root.val) {
      root.left = new TreeNode(val);
    }
    // insert the node at the right subtree
    else if (val > root.val && val < end) {
      root.right = new TreeNode(val);
    }

    // can't insert, just return the node
    return root;
  }

  // it's null, check if it's okay to insert the node here
  if (root.left === null && val > start && val < root.val) {
    root.left = new TreeNode(val);
    return root;
  }

  // it's null, check if it's okay to insert the node here
  if (root.right === null && val > root.val && val < end) {
    root.right = new TreeNode(val);
    return root;
  }

  root.left = recursiveHelper(root.left, val, start, root.val);
  root.right = recursiveHelper(root.right, val, root.val, end);

  return root;
}

/*
This code is cleaner and shorter

The main difference is that now we're gonna choose one path
Check the value beforehand
Then decide we're gonna go down to left or right
When hitting the null, we know we're gonna insert the node here

https://leetcode.com/problems/insert-into-a-binary-search-tree/discuss/1683942/Well-Detailed-Explaination-Java-C%2B%2B-Python-oror-Easy-for-mind-to-Accept-it

************************************************************
Time: O(n)
=> Note that the best case is O(log(n))
=> Because we're gonna choose one path
=> Thw worst case is when BST is like linked list
                 7
                   9
                     12
                        16
=> Like this, then we would traverse n node

Space: O(h)
*/
var insertIntoBST = function (root, val) {
  if (root === null) {
    return new TreeNode(val);
  }

  if (val < root.val) {
    root.left = insertIntoBST(root.left, val);
  } else {
    root.right = insertIntoBST(root.right, val);
  }

  return root;
};

/*
This is the same idea as before
But using iterative approach
*/
var insertIntoBST = function (root, val) {
  if (root === null) return new TreeNode(val);

  let curNode = root;

  while (true) {
    // go left subtree
    if (val < curNode.val) {
      if (curNode.left === null) {
        curNode.left = new TreeNode(val);
        break;
      }
      curNode = curNode.left;
    }
    // go right subtree
    else {
      if (curNode.right === null) {
        curNode.right = new TreeNode(val);
        break;
      }
      curNode = curNode.right;
    }
  }

  return root;
};

/*
Another approach to use iterative approach
But using parent node
https://www.youtube.com/watch?v=bmaeYtlO2OE
*/
var insertIntoBST = function (root, val) {
  if (root === null) return new TreeNode(val);

  let curNode = root,
    parentNode = null;

  while (curNode !== null) {
    parentNode = curNode;

    // go left
    if (val < curNode.val) curNode = curNode.left;
    // go right
    else curNode = curNode.right;
  }

  /*
  Now curNode is null
  but parentNode is the node above curNode, which is the parentNode
  By using parentNode to compare with the val, we then know we're gonna insert at left or right subtree
  */
  if (val < parentNode.val) parentNode.left = new TreeNode(val);
  else parentNode.right = new TreeNode(val);

  return root;
};
