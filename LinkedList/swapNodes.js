//////////////////////////////////////////////////////
// *** Swapping Nodes in a Linked List ***
//////////////////////////////////////////////////////
/*
You are given the head of a linked list, and an integer k.

Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [1,4,3,2,5]

Example 2:
Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]
 
Constraints:
The number of nodes in the list is n.
1 <= k <= n <= 105
0 <= Node.val <= 100
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
 * @param {number} k
 * @return {ListNode}
 */
/*
This is my initial answer
Kind of verbose, but it's easy to understand

1. find the length of linked list
2. let toSecondPtr = len - k;
=> find the distance to second ptr
3. use two while-loop to the ptr is gonna swap the value

************************************************************
Time: O(n)
Space: O(1)
*/
var swapNodes = function (head, k) {
  let len = 0;
  let cur = head;
  let ptr = head;
  let ptr2 = head;

  // find the length
  while (cur !== null) {
    cur = cur.next;
    len++;
  }

  let toSecondPtr = len - k;

  // find first ptr
  while (k > 1) {
    ptr = ptr.next;
    k--;
  }

  // find second ptr
  while (toSecondPtr) {
    ptr2 = ptr2.next;
    toSecondPtr--;
  }

  // swap them
  swap(ptr, ptr2);

  return head;
};

/*
I reference this from the discuss
The idea is smart
When we reach k-th node, we set p1 to the current node, and p2 - to the head.
We continue traversing the list, but now we also move p2. When we reach the end, p2 will points to k-th node from end.

For example, 
linked list: 1 -> 2 -> 3 -> 4 -> 5 -> x

Go into the for-loop,
ptr = head,
--k,
k = 1
 ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x

ptr = ptr.next,
--k
k = 0
      ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x
=> p1 = ptr
=> p2 = head
  p2  ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x

ptr = ptr.next
       p2  ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x

ptr = ptr.next
            p2  ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x

ptr = ptr.next
                  p2  ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x

                  p2      ptr
  1 -> 2 -> 3 -> 4 -> 5 -> x
ptr === null, break the for-loop
Note that here, p2 won't move
Because the order of for-loop is
1) update ptr (ptr = ptr.next)
2) check the condition (ptr !== null)
3) then break

This solution is really smart, it uses for-loop to update the ptr
Again, really think about the order of for-loop
So that we can move two ptrs to the right place
1. At the very beginning, we first minue the k, and do nothing, the order is like this
1. ptr = head
2. go into the for-loop
3. --k, nothing happen because k is not eqaul to 0
4. go into second for-loop, ptr = ptr.next
5. --k = 0, and find the first ptr

************************************************************
Time: O(n) (only one pass)
Space: O(1)
*/
var swapNodes = function (head, k) {
  let p1 = null;
  let p2 = null;

  for (let ptr = head; ptr !== null; ptr = ptr.next) {
    // we only want to update the second ptr when we find the place for first ptr
    // if p2 still null, then remain it to null
    // else, update the p2 (p2 = p2.next)
    p2 = p2 === null ? null : p2.next;

    // if k is 0, then set two ptrs
    if (--k === 0) {
      p1 = ptr;
      p2 = head;
    }
  }

  swap(p1, p2);

  return head;
};

function swap(p1, p2) {
  let tmp = p1.val;

  p1.val = p2.val;
  p2.val = tmp;
}
