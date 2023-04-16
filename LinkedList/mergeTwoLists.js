//////////////////////////////////////////////////////
// *** Merge Two Sorted Lists ***
//////////////////////////////////////////////////////
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
 
Constraints:
The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
/*
This problem is pretty intutive, but again have to handle the edge case
This is the first solution, which is more verbose to handle the edge case

The idea is 
1. iterate through both list1 and list2
2. use newCur to traverse and as cur node
3. and always compare current node of list1 and list2
4. find the smaller one,
   connect newCur to smaller one, and update the node


This is how I handle the edge case because the input may be [] and [1] or [] and []
  if (!list1) return list2;
  if (!list2) return list1;

  if (list1.val < list2.val) {
    newHead = list1;
    newCur = list1;
    list1 = list1.next;
  } else {
    newHead = list2;
    newCur = list2;
    list2 = list2.next;
  }
It's more verbose, but it's worked

************************************************************
Time: O(n)
Space: O(1)
*/
var mergeTwoLists = function (list1, list2) {
  if (!list1 && !list2) return null;
  if (!list1) return list2; // if list1 is null, we can simply return list2
  if (!list2) return list1;

  let newHead = null, // reference the head of returned list
    newCur = null; // reference the cur node of returned list

  // handle the first case
  // find the head of returned list
  if (list1.val < list2.val) {
    newHead = list1;
    newCur = list1;
    list1 = list1.next;
  } else {
    newHead = list2;
    newCur = list2;
    list2 = list2.next;
  }

  // iteration
  while (list1 !== null || list2 !== null) {
    /*
    What does this condition mean?
    Note that the operation inside condition is connect the node of list2
    And the while-loop only stop if both list1 and list2 are null

    list2 -> if list2 is null, we don't wanna do any operation of list2
    (list1 === null || list2.val < list1.val)
    -> if list1 === null, it means we're hitted the end of list1, so just traverse the rest of list2
    -> list2.val < list1.val, it means the value of list2 is smaller, connect the smaller value of node
    */
    if (list2 && (list1 === null || list2.val < list1.val)) {
      newCur.next = list2; // make the connection with new returned list
      list2 = list2.next; // update the node
    } else if (list1 && (list2 === null || list1.val <= list2.val)) {
      newCur.next = list1;
      list1 = list1.next;
    }

    // update the node
    newCur = newCur.next;
  }

  return newHead;
};

/*
The way to handle the edge case is verbose
We can just use dummyHead to avoid that

NOTE
Compare this solution with previous one
Really understand why dummyHead is helping us so much

  newCur  ->  x
dummyHead
Now we can go into while-loop, and connect the smaller value of node
*/
var mergeTwoLists1 = function (list1, list2) {
  if (!list1 && !list2) return null;

  const dummyHead = new ListNode(null);
  let newCur = dummyHead;

  while (list1 !== null || list2 !== null) {
    if (list2 && (list1 === null || list2.val < list1.val)) {
      newCur.next = list2;
      list2 = list2.next;
    } else if (list1 && (list2 === null || list1.val <= list2.val)) {
      newCur.next = list1;
      list1 = list1.next;
    }

    newCur = newCur.next;
  }

  return dummyHead.next;
};

/* 
This solution is cleaner

The main difference is the condition inside while-loop
Why?
Because once either list1 or list2 hitting the end of linked list, aka null
We can make sure all the value of the rest linked list is greater
So we can stop the while-loop, and connect afterward

See the video https://www.youtube.com/watch?v=GfRQvf7MB3k

For example,
list 1: 1 -> 3 -> 5 -> x
list 2: 2 -> 6 -> 10 -> x


1st
list 1: 1 -> 3 -> 5 -> x
            l1
list 2: 2 -> 6 -> 10 -> x
        l2
new list: dummyHead -> 1
                      cur
=> compare l1 and l2

2nd
list 1: 1 -> 3 -> 5 -> x
            l1
list 2: 2 -> 6 -> 10 -> x
            l2
new list: dummyHead -> 1 -> 2
                           cur
=> compare l1 and l2

3rd
list 1: 1 -> 3 -> 5 -> x
                  l1
list 2: 2 -> 6 -> 10 -> x
            l2
new list: dummyHead -> 1 -> 2 -> 5
                                cur
=> compare l1 and l2

4th
list 1: 1 -> 3 -> 5 -> x
                      l1
list 2: 2 -> 6 -> 10 -> x
            l2
new list: dummyHead -> 1 -> 2 -> 5
                                cur
=> l1 hit the end of linked list
=> what does this mean?
=> the rest of value in list2 are all greater than 5
=> break out the while-loop
=> let whatever the list is not null connect to the new returned list
=> new list: dummyHead -> 1 -> 2 -> 5 -> 6 -> 10 -> x

*/
var mergeTwoLists2 = function (list1, list2) {
  if (!list1 && !list2) return null;

  const dummyHead = new ListNode(null);

  let newCur = dummyHead;

  // break out the while-loop when one of list is null (aka, hitting the end of linked list)
  while (list1 && list2) {
    if (list2.val < list1.val) {
      newCur.next = list2;
      list2 = list2.next;
    } else {
      newCur.next = list1;
      list1 = list1.next;
    }

    newCur = newCur.next;
  }

  // let whatever the list is not null connect to the new returned list
  newCur.next = list1 || list2;

  return dummyHead.next;
};
