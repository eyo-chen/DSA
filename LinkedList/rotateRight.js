//////////////////////////////////////////////////////
// ***  Rotate List ***
//////////////////////////////////////////////////////
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/*
Given the head of a linked list, rotate the list to the right by k places.

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]

Example 2:
Input: head = [0,1,2], k = 4
Output: [2,0,1]
 

Constraints:
The number of nodes in the list is in the range [0, 500].
-100 <= Node.val <= 100
0 <= k <= 2 * 109
*/
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
/*
I didn't come up the solution, it's little tricky

We don't need to care about the process of rotation
We only need to care about the result after rotation
What does this mean?
Look at the example, if head = [1,2,3,4,5], k = 2
What's the result after doing rotation 2 times?
it's [4,5,1,2,3]

Can you see the pattern
We only need to reference following things
1. the head
2. the tail
3. [3] => the tail of new list
4. [4] => the head of new list

That's all we need to have the result

Now the main question is how we find the [3] and [4]
It's quite easy

Move forward (length - k - 1) times to find the [3] (the tail of new list)
and move forward one more time to find the [4] (the head of new list)

Note the we need the exact (length - k - 1)
This help us to find the tail of new list

The edge case
What if k > the length of list?

It's easy, all we need to do is use % 
if head = [1,2,3,4,5],
k = 7 is equal to k = 2
Just graph the linked list to see clearly

************************************************************
Time: O(n)
Space: O(1)
*/
var rotateRight = function (head, k) {
  /*
    Base case
    1. the length of head is 0
    2. the length of head is 1
    3. k = 0, which means don't need to do any rotation
  */
  if (head === null || head.next === null || k === 0) return head;

  let len = 1,
    tail = head,
    newTail = head,
    newHead = null,
    rightShiftIndex = null;

  // find the tail and the length of linked list
  while (tail.next !== null) {
    len++;
    tail = tail.next;
  }

  // k % len === 0 means do need to do any rotation
  // for example, the length of linked list is 5, and k = 5, the rotation after that is just original linked list
  if (k % len === 0) return head;

  // the index helps us to find the new tail
  rightShiftIndex = len - (k % len) - 1;

  // find new tail
  while (rightShiftIndex !== 0) {
    newTail = newTail.next;
    rightShiftIndex--;
  }

  // pointer rewiring
  tail.next = head;
  newHead = newTail.next;
  newTail.next = null;

  return newHead;
};

// basically same idea, the code implementation is different
var rotateRight1 = function (head, k) {
  if (head === null || head.next === null || k === 0) return head;

  let len = 1,
    tail = head;

  while (tail.next !== null) {
    len++;
    tail = tail.next;
  }

  let rightShiftIndex = k % len,
    newTail = head;
  tail = head;

  if (rightShiftIndex === 0) return head;

  while (tail.next !== null) {
    if (rightShiftIndex === 0) {
      tail = tail.next;
      newTail = newTail.next;
    } else {
      rightShiftIndex--;
      tail = tail.next;
    }
  }

  const newHead = newTail.next;
  tail.next = head;
  newTail.next = null;

  return newHead;
};
