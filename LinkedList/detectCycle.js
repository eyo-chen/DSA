//////////////////////////////////////////////////////
// ***  Linked List Cycle II ***
//////////////////////////////////////////////////////
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/*
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
Input: head = [1], pos = -1
Output: no cycle
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
This problem is an variant of Linked List Cycle
Instead of just returning true or false, we want to return the node of cycle

This solution is intutive, just use set

************************************************************
Time: O(n)
Space: O(n)
*/
var detectCycle = function (head) {
  const set = new Set();

  while (head !== null) {
    if (set.has(head)) return head;
    set.add(head);
    head = head.next;
  }

  return null;
};

/*
Second approach

This is very hard to come up with by myself

The process of solution is like this
1. Use fast and slow ptr to find if it has cycle
2. If it does not have cycle, we can just simply return null
3. If it does have cycle, we then find the lengt of cycle
4. create an window with length of cycle
5. move the window foward
6. we then find the node of cycle when the head of window and tail of window meet

Again, this is kind of hard to come up at first
But we can keep in mind we now have two technique when solving linked list problems
1. use two pointers 
2. use window
*/
var detectCycle = function (head) {
  let fast = head,
    slow = head,
    meet = null,
    cycleLength = 1;

  // find if has cycle (use fast and slow ptr)
  while (fast !== null && fast.next !== null) {
    fast = fast.next.next;
    slow = slow.next;

    if (fast === slow) {
      meet = fast;
      break;
    }
  }

  // if no cycle, just return null
  if (meet === null) return null;

  // count the length of cycle
  slow = slow.next;
  while (slow !== fast) {
    slow = slow.next;
    cycleLength++;
  }

  // create the window
  // the length of window is the length of cycle
  fast = head;
  slow = head;
  while (cycleLength !== 0) {
    fast = fast.next;
    cycleLength--;
  }

  // move window forward to find the node of cycle
  while (slow !== fast) {
    slow = slow.next;
    fast = fast.next;
  }

  return slow;
};
