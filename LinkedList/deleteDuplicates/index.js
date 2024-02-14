//////////////////////////////////////////////////////
// *** Remove Duplicates from Sorted List ***
//////////////////////////////////////////////////////
/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]
 
Constraints:
The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
/*
Iterative Approach
Very intutive and straightforward

There's recursvie approach, go to discuss of leetcode

************************************************************
Time: O(n)
Space: O(1)
*/
var deleteDuplicates = function (head) {
  if (head === null || head.next === null) return head;

  let indexVal = head.val;
  let curNode = head;

  while (curNode.next !== null) {
    // find the duplicate value, remove it
    if (curNode.next.val === indexVal) curNode.next = curNode.next.next;
    // No duplicate value, update ptr and change the indexVal
    else {
      curNode = curNode.next;
      indexVal = curNode.val;
    }
  }

  return head;
};
