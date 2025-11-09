//////////////////////////////////////////////////////
// *** Convert Sorted List to Binary Search Tree ***
//////////////////////////////////////////////////////
/*
Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:
Input: head = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.

Example 2:
Input: head = []
Output: []
 
Constraints:
The number of nodes in head is in the range [0, 2 * 104].
-105 <= Node.val <= 105
*/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
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
 * @param {ListNode} head
 * @return {TreeNode}
 */
/*

************************************************************
Time: O(n * log(n))
Space: O(h)
*/
var sortedListToBST = function (head) {
  return recursiveHelper(head, null);
};

function recursiveHelper(head, tail) {
  if (head === tail) {
    return null;
  }

  let fast = head,
    slow = head;

  while (fast !== tail && fast.next !== tail) {
    fast = fast.next.next;
    slow = slow.next;
  }

  const root = new TreeNode(slow.val);

  root.left = recursiveHelper(head, slow);

  /*
    slow.next is very important
    Should write the process on the paper again
  */
  root.right = recursiveHelper(slow.next, tail);

  return root;
}
