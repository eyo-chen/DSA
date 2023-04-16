//////////////////////////////////////////////////////
// ***  Linked List Cycle ***
//////////////////////////////////////////////////////
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

Constraints:
The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.
 
Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/
/**
 * @param {ListNode} head
 * @return {boolean}
 */
/*
Approach 1

This is pretty straightforward
just use set to find if it has cycle

************************************************************
Time: O(n)
Space: O(n)
*/
var hasCycle = function (head) {
  const set = new Set();

  while (head !== null) {
    if (set.has(head) === true) return true;
    set.add(head);
    head = head.next;
  }

  return false;
};

/*
Approach 2
This is not that intutive
I did not come up this solution by myself

The idea is just using two pointers, fast and slow
If the link list has cycle, it will eventually meet somwhere in the cycle

Why this is true?
Because slow always goes forward 1 position, and fast always goes forward 2 position
Why does this mean?
It means when slow and fast ptr moves forward one time, the gap between these two will decreas by 1
For example,
If the length of cycle is 10, it means now the distance between slow and fast is 10
It needs to take 10 iteration to meet each other because each iteration will decrease the gap of the length by 1

see the detial explanation here https://www.youtube.com/watch?v=gBTe7lFR3vc
************************************************************
Time: O(n)
Space: O(1)
*/
var hasCycle = function (head) {
  let fast = head,
    slow = head;

  while (fast !== null && fast.next !== null) {
    // have to move the ptr first
    // because fast and slow is intitally at the same position(head)
    fast = fast.next.next;
    slow = slow.next;

    // then check
    if (fast === slow) return true;
  }

  return false;
};
