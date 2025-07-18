//////////////////////////////////////////////////////
// ***   Palindrome Linked List ***
//////////////////////////////////////////////////////
/*
Given the head of a singly linked list, return true if it is a palindrome.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false
 

Constraints:
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

Follow up: Could you do it in O(n) time and O(1) space?
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
 * @return {boolean}
 */
/*
This problem is not that hard if using O(n) space

This solution is pretty intutive
Use stack to store the value
Stack is first in, last out
After creating the stack, and we again loop through the linked list to compare the value

For example, 1 -> 2 -> 3 -> 2 -> -> 1
stack: [1,2,3,2,1]
Then compare the value in stack with the value of linked list

1st iteration:
stack.pop() = 1
cur.val = 1
=> stack: [1,2,3,2]
=> linked list: 1 -> 2 -> 3 -> 2 -> -> 1
                     c

2nd iteration:
stack.pop() = 2
cur.val = 2
=> stack: [1,2,3]
=> linked list: 1 -> 2 -> 3 -> 2 -> -> 1
                          c
Hope this process is clead

So the main purpose of stack is storing the value
Later,
we can read the value BACKWARD, and compare the value with linked list
The reason we can read the value backward is because the stack is first in, last out

Again, if it's still not clear, graph the linked list and stack on the paper
************************************************************
Time: O(n)
Space: O(n)
*/
var isPalindrome = function (head) {
  const stack = [];
  let cur = head;

  // store value in the stack
  while (cur !== null) {
    stack.push(cur.val);
    cur = cur.next;
  }

  // compare the value
  while (head !== null) {
    // immediately return false if the value is not matching
    if (head.val !== stack.pop()) return false;
    head = head.next;
  }

  return true;
};

/*
Second solution
Tho this solution only use O(1) space, but it's very tricky

Cut the linked list into two half
Reverse the second half
Then compare first half and second half

For example, 1 -> 2 -> 2 -> 1
             1st half | second half
1st half: 1 -> 2
2nd half: 2 -> 1

Reverse 2nd half: 1 -> 2

Compare 1st half and 2nd half

BUT, there's one edge case, when the length of linked list is odd
1 -> 2 -> 3 -> 2 -> 1
1st half: 1 -> 2
2nd half: 3 -> 2 -> 1
There are two ways to handle this

This solution is use first way
Just need to handle the case when 1st half is null when compare the value

After reversing,
1st half: 1 -> 2
2nd half: 1 -> 2 -> 3
Now, we can compare
But there's the case when the cur node of 2nd half is 3, but the cur node of 1st half is null
So we have to check cur node of 1st half is null
if it's null, we don't need to do any comparison, just update the node of 2nd half
Again, the reason is that we can make sure the length of 2nd half is always eqaul or greater than 1st half

However, if we use this approach, there's yet another edge case
After reversing,
1st half: 1 -> 2
2nd half: 1 -> 2 -> 3
1st half is still point at 3
note that original linked list is 1 -> 2 -> 3 -> 2 -> 1
We NEVER cut off the connection between 2 and 3 (2 -> 3)
So even tho after reversing, 2 still point at 3
Which means the cur node of 1st half will not be null in correct position
We have to cur off that connection
How?
Have the referecne the node before 3, and cur off
************************************************************
Time: O(n)
Space: O(1)
*/
var isPalindrome1 = function (head) {
  // base case: single node is always palindrome
  if (head.next === null) return true;

  // duumyHead helps us to have the reference of the node before slow(aka, the head of 2nd linked list)
  const dummyHead = new ListNode(null, head);
  let fast = head,
    slow = head,
    nodeBeforeSlow = dummyHead,
    prev,
    cur,
    next;

  // because nodeBeforeSlow starts at dummyHead, so we can make sure nodeBeforeSlow is always before one position of slow
  while (fast !== null && fast.next !== null) {
    fast = fast.next.next;
    slow = slow.next;
    nodeBeforeSlow = nodeBeforeSlow.next;
  }
  // cut off the connection of 1st half and 2nd half
  nodeBeforeSlow.next = null;

  // reverse the 2nd half
  prev = null;
  cur = slow; // the head of 2nd half is slow ptr
  while (cur !== null) {
    next = cur.next;
    cur.next = prev;
    prev = cur;
    cur = next;
  }

  // compare two half of linked list
  while (prev !== null) {
    // always have to care about the cur node of 1st half (head represent the 1st half)
    // because 1st half is shorter than 2nd half
    if (head && head.val !== prev.val) return false;

    // again, care about the cur node of 1st half
    if (head) head = head.next;

    // the node of 2nd half is longer, don't need to handle the edge case
    prev = prev.next;
  }

  return true;
};

/*
This is second approach, cleaner

Even case: 1 -> 2 -> 2 -> 1
After using fast and slow ptr to find find the 1st half and 2nd half
1 -> 2 -> 2 -> 1 -> x
          s         f
=> fast at the null, and slow at the 2
=> Look, slow is exactly the place of the head of 2nd half
=> just reverse the 2nd half

Even case: 1 -> 2 -> 3 -> 2 -> 1
After using fast and slow ptr to find find the 1st half and 2nd half
1 -> 2 -> 3 -> 2 -> 1 -> x
          s         f
=> fast at the 1, and slow at the 3
=> we don't wanna start at 3 to reverse
=> so we update slow one position
1 -> 2 -> 3 -> 2 -> 1 -> x
               s    f
=> now 2 is the head of 2nd half
=> go reversing the 2nd half

So our approach to handle both odd case and even case is
use the fast position to determince if we're gonna update the slow
If fast is at the null, don't need to update slow because slow is exactly the place of the head of 2nd half
If fast is not null, have to update slow, so that slow is exactly the place of the head of 2nd half
*/
var isPalindrome2 = function (head) {
  if (head.next === null) return true;

  let fast = head,
    slow = head,
    prev,
    cur,
    next;

  // as same as before
  while (fast !== null && fast.next !== null) {
    fast = fast.next.next;
    slow = slow.next;
  }

  // If fast is not null, have to update slow, so that slow is exactly the place of the head of 2nd half
  if (fast !== null) slow = slow.next;

  // reverse the 2nd half
  prev = null;
  cur = slow;
  while (cur !== null) {
    next = cur.next;
    cur.next = prev;
    prev = cur;
    cur = next;
  }

  // Do not need to care about the edge case because now the length of 1st half and 2nd half is the same
  while (prev !== null) {
    if (head.val !== prev.val) return false;

    head = head.next;
    prev = prev.next;
  }

  return true;
};
