//////////////////////////////////////////////////////
// *** Remove Duplicates from Sorted List II ***
//////////////////////////////////////////////////////
/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Example 1:
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:
Input: head = [1,1,1,2,3]
Output: [2,3]

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
I done this problem by myself!
The problem is a little bit tircky
But it's not the difficult

I think the edge case is sth like this
1 -> 1 -> 2 -> 2 -> x => x
1 -> 1 -> 1 -> 1 -> x => x
How to handle these situations?

My solution is using three different pointers(ptrs)
1. prev
   => this is important
   => See the second situation above
   => when all the list is duplicate, we have to remove all the list
   => How are we gonna do this?
   => We have to reference the node before the head of linked list
   => we have to reference the node before 1
   => which means we have to create dummyHead
   => In short, prev is the node BEFORE duplicate linked list, which is the sublist we're gonna remove
   => For example,
   => Case 1: 1 -> 2 -> 2 -> 2 -> 3 -> x
             prev
   => Case 2: 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> x
                       prev
   => We have to always make sure prev is at the position where is BEFORE the head of duplicate linked list
   => so that we can later remove the duplicate sub list

2. cur
   => just reference the current node
   => and for checking if we're done iterating the linked list

3. deletePtr
   => we need this ptr when we have duplicate element
   => it helps us to find the duplicate sub list
   => in other words, it helps us to find the end of duplicate sub list
   => see the code implemention

For example, 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x

1. create dummyHead (DH)
x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
DH

2. Set the ptr
  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
prev  cur

3. Iterate the linked list until finding the duplicate
  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
prev  cur
=> cur.val !== cur.next.val
=> keep updating both prev and cur

  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
     prev  cur
=> cur.val !== cur.next.val
=> keep updating both prev and cur

  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
         prev  cur
=> cur.val !== cur.next.val
=> keep updating both prev and cur

  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
                prev  cur
=> cur.val === cur.next.val
=> now be ready to find the duplicate sub list
=> deletePtr = cur (dp)

  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
                prev  cur
                      dp
=> keep updating dp until not having duplicate element

  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
                prev  cur
                                dp
=>  deletePtr.val !== deletePtr.next.val
=> stop

  x -> 1 -> 2 -> 3 -> 4 -> 4 -> 4 -> 5 -> 5 -> 5 -> 6 -> x
                prev                cur
                                dp
=> cur = deletePtr.next;

  x -> 1 -> 2 -> 3 -> 5 -> 5 -> 5 -> 6 -> x
                prev cur
=> prev.next = cur;
=> cut off the sub list
=> it's okay we temporarily lose the reference of deletePtr
=> because we'll reference again when finding duplicate element
=> Now find the duplicate element again

  x -> 1 -> 2 -> 3 -> 5 -> 5 -> 5 -> 6 -> x
                prev cur
                     dp
=> keep updating dp until not having duplicate element

  x -> 1 -> 2 -> 3 -> 5 -> 5 -> 5 -> 6 -> x
                prev cur
                               dp
=> deletePtr.val !== deletePtr.next.val
=> stop

  x -> 1 -> 2 -> 3 -> 5 -> 5 -> 5 -> 6 -> x
                prev                cur
                               dp
=> cur = deletePtr.next;

  x -> 1 -> 2 -> 3 -> 6 -> x
                prev cur
=> prev.next = cur;
=> cut off the sub list
=> now we're at the end of linked list
=> break the while-loop 

Hope the process is clear

Now handle the edge case
1 -> 1 -> 1 -> 1 -> x
when will have error?

  x -> 1 -> 1 -> 1 -> 1 -> x
prev  cur
=> finding duplicate

  x -> 1 -> 1 -> 1 -> 1 -> x
prev  cur
      dp
=> keep updating dp until not having duplicate element

  x -> 1 -> 1 -> 1 -> 1 -> x
prev  cur
                      dp
=> note that we have to stop here
=> if we don't stop at here, deletePtr.next.val will have error
=> deletePtr.next is null, null.val give us error
=>  while (deletePtr.next && deletePtr.val === deletePtr.next.val)
=> this while-loop condition helps us to stop the updating before having error

  x -> x
prev  cur
=> cut off all the dupliate sub list
=>  while (cur !== null && cur.next !== null)
=> again, in the next while-loop condition 
=> cur !== null stop the condition earily
=> so that cur.next !== null won't give us error
=> return dummyHead.next -> great

************************************************************
Time: O(n)
Space: O(1)
*/
var deleteDuplicates = function (head) {
  // Base case
  if (head === null || head.next === null) return head;

  const dummyHead = new ListNode(null, head);
  // set three ptrs
  let prev = dummyHead,
    cur = head,
    deletePtr = null;

  while (cur !== null && cur.next !== null) {
    // find the duplicate, be ready to find the duplicate sub list
    if (cur.val === cur.next.val) {
      // set the deletePtr to find the duplicate sub list
      deletePtr = cur;

      //  keep updating deletePtr until not having duplicate element or at the end of linked list
      while (deletePtr.next && deletePtr.val === deletePtr.next.val) {
        deletePtr = deletePtr.next;
      }

      // update cur
      cur = deletePtr.next;
      // cur off the sublist
      prev.next = cur;
    }
    // if not finding duplicate element, just normally update two ptrs
    else {
      prev = prev.next;
      cur = cur.next;
    }
  }

  return dummyHead.next;
};
