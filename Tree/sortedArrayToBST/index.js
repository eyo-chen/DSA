//////////////////////////////////////////////////////
// *** Convert Sorted Array to Binary Search Tree ***
//////////////////////////////////////////////////////
/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

Example 1:
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
 

Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.
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
 * @param {number[]} nums
 * @return {TreeNode}
 */
/*
I didn't solve it by myself at first
But this problem is actually very easy if i'm famialr with binary search
The core idea of this problem is divide and conquer

[-10,-3,0,5,9]

               0
         -3       5
    -10               9

As we can see, the root is the middle element
And it's left part is the left subtree, right part is right subtree

The process is
1. find the mid point
2. create the node
3. go build left subtree(cut the left array)(update rightIndex)
4. go build right subtree(cur the right array)(update leftIndex)

************************************************************
Time: O(n)
Space: O(h)
*/
var sortedArrayToBST = function (nums) {
  return recursiveHelper(nums, 0, nums.length);
};

function recursiveHelper(nums, leftIndex, rightIndex) {
  if (leftIndex >= rightIndex) return null;

  // const middle = left + Math.floor((right - left) / 2);
  const middle = Math.floor((leftIndex + rightIndex) / 2);
  const node = new TreeNode(nums[middle]);
  node.left = recursiveHelper(nums, leftIndex, middle);
  node.right = recursiveHelper(nums, middle + 1, rightIndex);

  return node;
}
