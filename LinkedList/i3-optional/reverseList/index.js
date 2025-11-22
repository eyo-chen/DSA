//////////////////////////////////////////////////////
// *** Reverse Linked List ***
//////////////////////////////////////////////////////
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
The solution should be straightforward

************************************************************
Time: O(n)
Space: O(1)
*/
var reverseList = function (head) {
  let prev = null;
  let cur = head;
  let next;

  while (cur !== null) {
    next = cur.next;

    cur.next = prev;

    prev = cur;
    cur = next;
  }

  return prev;
};
