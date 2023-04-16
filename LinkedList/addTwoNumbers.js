//////////////////////////////////////////////////////
// *** Add Two Numbers ***
//////////////////////////////////////////////////////
/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]
.
Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
 
Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
/*
Again, this solution is pretty straightforward
Because digits are stored in reverse order, so we can just simply add each digit number from the head to tail

The small edge case of the problem is caring about the carry when while-loop is breaking
For example,
l1: 1 -> 3 -> 6
l2: 2 -> 5 -> 5
result: 6 -> 8 -> 1 -> 1
As we can see, after while-loop, we only have three digits in the result because we won't have a chance to add finial digit
So we have to check if curCarry is equal to 1, which means the last digits create curCarry
so we have to add it

************************************************************
n = the length of l1, m = the length of l2
Time: O(max(n + m))
Space: O(1)
=> If we don't count the output
=> If we do count, the space complexity is also O(max(n + m))
*/
var addTwoNumbers = function (l1, l2) {
  const dummyHead = new ListNode(null);
  let curVal,
    node,
    curNode = dummyHead,
    curCarry = 0;

  while (l1 !== null || l2 !== null) {
    curVal = 0;

    if (l1 !== null) {
      curVal += l1.val;
      l1 = l1.next;
    }

    if (l2 !== null) {
      curVal += l2.val;
      l2 = l2.next;
    }

    curVal += curCarry;

    if (curVal >= 10) curCarry = 1;
    else curCarry = 0;

    node = new ListNode(curVal % 10);
    curNode.next = node;
    curNode = curNode.next;
  }

  if (curCarry === 1) {
    node = new ListNode(1);
    curNode.next = node;
  }

  return dummyHead.next;
};

/*
The is recursive approach which I referecne from leetcode
Sitll not familar with it
*/
var addTwoNumbers1 = function (l1, l2, carry = 0) {
  let curNode = null;

  if (l1 || l2) {
    let curVal = 0;

    if (l1 !== null) {
      curVal += l1.val;
      l1 = l1.next;
    }

    if (l2 !== null) {
      curVal += l2.val;
      l2 = l2.next;
    }

    curVal += carry;

    let nextCarry = 0;
    if (curVal >= 10) nextCarry = 1;

    curNode = new ListNode(curVal % 10);
    curNode.next = addTwoNumbers1(l1, l2, nextCarry);
  } else if (carry > 0) {
    curNode = new ListNode(1);
    curNode.next = null;
  }

  return curNode;
};
