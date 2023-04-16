//////////////////////////////////////////////////////
// ***  Rearrange Linked List | AlgoExpert ***
//////////////////////////////////////////////////////
/*
Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example 1:
Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]

Example 2:
Input: head = [2,1], x = 2
Output: [1,2]
 
Constraints:
The number of nodes in the list is in the range [0, 200].
-100 <= Node.val <= 100
-200 <= x <= 200
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
 * @param {number} x
 * @return {ListNode}
 */
/*
This solution should be very intuitive
Use two pointers
First pointer construct the linked list is less than x
Second pointer construct the linked list is greater or equal than x

Loop through the linked list, 
resprectively connect to two pointers according to their value

Finally, connect the tail of shorterLinkList to the head of greaterLinkList
Also, need to make sure the tail of greaterLinkList point to null
*/
var partition = function (head, x) {
  // set two pointers
  const largeHead = new ListNode(null),
    shortHead = new ListNode(null);
  let largeNode = largeHead,
    shortNode = shortHead,
    cur = head;

  while (cur !== null) {
    // construct greater linked list
    if (cur.val >= x) {
      largeNode.next = cur;
      largeNode = largeNode.next;
    }
    // construct shorter linked list
    else {
      shortNode.next = cur;
      shortNode = shortNode.next;
    }

    cur = cur.next;
  }

  // connect two sub linked list
  shortNode.next = largeHead.next;

  // make sure the tail of linked list point to null
  largeNode.next = null;

  return shortHead.next;
};

//////////////////////////////////////////////////////
// ***  Rearrange Linked List | AlgoExpert ***
//////////////////////////////////////////////////////
/*
Given an head, and k
rearrange the linked list, so that the value smaller than k is at first
then is the node as same as k
final is the node greater than k

For example, if head: 10 -> 1 -> 4 -> 3 -> 3 -> 3 -> 2 -> 0 -> 5, k = 3
Output: 1 -> 2 -> 0 -> 3 -> 3 -> 3 -> 10 -> 4 -> 5
The order in original list cannot be changed
And the order is smaller -> same -> larger
*/
// This is the class of the input linked list.
class LinkedList {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

function rearrangeLinkedList(head, k) {
  const greaterHead = new LinkedList(null);
  const smallerHead = new LinkedList(null);
  const sameHead = new LinkedList(null);
  let greaterCur = greaterHead,
    smallerCur = smallerHead,
    sameCur = sameHead;

  while (head !== null) {
    if (head.value > k) {
      greaterCur.next = head;
      greaterCur = greaterCur.next;
    } else if (head.value < k) {
      smallerCur.next = head;
      smallerCur = smallerCur.next;
    } else {
      sameCur.next = head;
      sameCur = sameCur.next;
    }

    head = head.next;
  }

  if (sameHead.next) {
    smallerCur.next = sameHead.next;
    sameCur.next = greaterHead.next;
    greaterCur.next = null; // have to cut off the connection, make sure the tail of new linked list point to null
  }

  // check if there's same value as k
  if (!sameHead.next) {
    smallerCur.next = greaterHead.next;
    greaterCur.next = null; // have to cut off the connection, make sure the tail of new linked list point to null
  }

  // handle the case when all the node is greater than k
  // simply return smallerHead.next is incorrect
  return smallerHead.next ? smallerHead.next : sameHead.next;
}
