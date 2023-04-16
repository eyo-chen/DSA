//////////////////////////////////////////////////////
// *** Copy List with Random Pointer ***
//////////////////////////////////////////////////////
/*
Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Example 2:
Input: head = [5], left = 1, right = 1
Output: [5]
 
Constraints:
The number of nodes in the list is n.
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

Follow up: Could you do it in one pass?
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
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
/*
This problem is not very hard,
but it's hard to manage the edge case, and make the code looks clean
Normally, the edge case would be the length of 1 or 2
(this problem said the length of linked list is at least one)

This solution is wrote by myself,
it works, and also handle the edge case, but the code is not that clean

The idea is very intutive
In order to reverse the sublist, we just reference the head and tail of sublist
and reverse it, after reversing, just merge back to original list
But, we have to reference multiple reference, so that we won't loose the node along with the process

tail -> the node after the tail of sublist 
head -> the node before the head of sublist
reverseHead -> the head of sublist
reverseTail -> the tail of sublist
pre, next, cur -> is the ptr help us to reverse the linked list
dummyHead -> this is very important, it helps us to handle the edge case

For example, left = 3, right = 4
In this case,
head -> 2, tail -> 5, reverseHead -> 3, reverseTail -> 4

1 -> 2 -> 3 -> 4 -> 5 -> x
     h   rh   rt    t
Here, we can konw why we need head and tail
Because we're gonna reverse 3 -> 4
After reversing, the sublist is 4 -> 3
And we want to know merge back to original list
so 2 should connect to 4, and 3 should connect to 5

So the reason we reference head and tail(the node after and before the head and tail of sublist)
Is just helps us to merge back after reversing

Why dummyHead is important in this problem?
Imagine the list is 1 -> 2 -> 3 -> 4 -> 5 -> x, and left = 1, right = 5
How can we access the head, aka the node before the head of sublist
Now sublist is just 1 -> 2 -> 3 -> 4 -> 5 -> x
So we need dummyHead to help us to manage this situation
Not only in this situation, dummyHead works in every situation of this problem
Also
we have to move tail, head, reverseHead, reverseTail to the correct position
And dummyHead helps us to do that
For example, left = 3, right = 4
dummyHead -> 1 -> 2 -> 3 -> 4 -> 5 -> x
   c
=> we move cur ptr from dummyHead
=> after left - 1 moving, we get the head, aka the node before the head of sublist
=> after left moving, we get the reverseHead
=> after right moving, we get the reverseTail
=> so dummyHead is also helps us to move all the ptrs to correct position

************************************************************
Time: O(n)
Space: O(1)
*/
var reverseBetween = function (head, left, right) {
  const dummyHead = new ListNode(null, head);
  // all ptrs
  let tail,
    reverseHead,
    reverseTail,
    prev,
    next,
    cur = dummyHead; // note here, we start traversing from dummyHead

  // move head and reverseHead to the correct position
  while (left >= 0) {
    // head is the node BEFORE the head of sublist
    if (left === 1) head = cur;
    if (left === 0) reverseHead = cur;

    left--;
    cur = cur.next;
  }

  // move tail and reverseTail to the correct position
  cur = dummyHead;
  while (right >= 0) {
    if (right === 0) reverseTail = cur;

    right--;
    cur = cur.next;
  }
  // tail is the node AFTER the tail of sublist
  tail = reverseTail.next;

  // isolate the sublist, this is very important, we have to make sure isolate the sublist
  // so the later cur can stop properly
  reverseTail.next = null;

  // ready to reverse the sublist(all the logic is as same as reverse linked list)
  cur = reverseHead;
  prev = null;
  while (cur !== null) {
    next = cur.next;
    cur.next = prev;
    prev = cur;
    cur = next;
  }

  // merge sublist back to original list
  head.next = prev;
  reverseHead.next = tail;

  return dummyHead.next;
};
/*
This is second solution, and the code looks cleaner
we get rid of idea of isolation the sublist, then reverse, and merge back finally
All we do here is just move the node step by step, and get out final result

For example, 1 -> 2 -> 3 -> 4 -> 5 -> x, left = 2, right = 5
Again, we still need to referecne the node BEFORE the head of sublist
In this case, is 1
This is very important, we want to reverse the sublist,
so we have to reference the node 1 (nbr = nodeBeforeRemove)
nbr  cur
 1 -> 2 -> 3 -> 4 -> 5 -> x
The process is sth like this,
In each step, we move the next node of cur to the head of sublist
in other words, we move the next node of cur to in front of the sublist
in other words, we always move the next node of cur to the next node of nbr

nbr  cur
 1 -> 2 -> 3 -> 4 -> 5 -> x
=> next = cur.next -> 3
=> move 3 to the next node of nbr
=> so that we want to sth like this 1 -> 3

First 
nbr       cur
 1 -> 3 -> 2  -> 4 -> 5 -> x
=> move 3 to the next node of nbr
=> now the new next node is 4, because cur.next = 4

Second 
nbr            cur
 1 -> 4 -> 3 -> 2 -> 5 -> x
=> move 4 to the next node of nbr
=> now the new next node is 4, because cur.next = 5
=> note that we always want to move next node to the next node of nbr
=> in other words, 4 should be the next node of nbr

Third 
nbr                 cur
 1 -> 5 -> 4 -> 3 -> 2 -> x
=> move 5 to the next node of nbr
=> Finish

So the process is sth like this

How many steps we take?
right - left + 1 step
5     -  3   + 1 = 3 step

Just like previous solution
dummyHead helps us to move the ptr to correct position

************************************************************
Time: O(n)
Space: O(1)
*/
var reverseBetween1 = function (head, left, right) {
  const dummyHead = new ListNode(null, head);
  let cur = dummyHead,
    next = null,
    nodeBeforeRemove = dummyHead;

  // move nodeBeforeRemove and cur to correct position
  for (let i = 0; i < left; i++) {
    if (i === left - 1) nodeBeforeRemove = cur;
    cur = cur.next;
  }

  // step by step to reverse the sublist
  // left < right help us to have to correct number of moving position
  while (left < right) {
    // reference the next ptr
    next = cur.next;

    // rewiring part, make the next node to be the next node of nodeBeforeRemove
    cur.next = next.next;
    next.next = nodeBeforeRemove.next;
    nodeBeforeRemove.next = next;
    left++;
  }

  return dummyHead.next;
};
