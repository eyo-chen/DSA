//////////////////////////////////////////////////////
// *** Copy List with Random Pointer ***
//////////////////////////////////////////////////////
/*
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.

Example 1:
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

Example 2:
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

Example 3:
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
 

Constraints:
0 <= n <= 1000
-104 <= Node.val <= 104
Node.random is null or is pointing to some node in the linked list.
*/
/**
 * Definition for a Node.
 * function Node(val, next, random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */
/**
 * @param {Node} head
 * @return {Node}
 */
/*
Solution 1: use Map
Even tho this solution is intutive after explanation, i didn't figure it out by myself

The idea is 
1. create Map to store key-value pair
   key: the node in original list
   value: the clone node
   For example,
   input: 1 -> 2 -> 3 -> 4 -> x
   (key, value)
   => (1, 1c), (2, 2c), (3, 3c), (4, 4c)
   => c represenet for clone
   => note that we did NOT create the connection, we just store key-value pair

2. Iterate the link list 
=> now we have the original node and clone node
=> the conversation of clone node is like
=> Hey, original node, where does your next and random point at?
=> I just simple connect with the clone node of node you point at
=> For example,
=> original list: 1 -> 2 -> 3 -> 4 -> x
=> Assume 1.next = 2, 1.random = 4
=> Hey 1, where does your next and random point at?
=> It points ar 2 and 4
=> okay, go to 2 and 4
=> And because of map, now 2 and 4 both have clone node
=> so i just make the connection

There's one edge case to node in this solution
=> Initially, we store original node <-> clone node in the Map
=> Later, map.get(curNode.next) and map.get(curNode.random) may give us undefined
=> Why?
=> Go looking at the first while-loop
=> We never store the null in Map
=> However, there are chance that curNode.next and curNode.random may give us null
=> map.get(null) returns undefined
=> But as problem said, we have to point to null, instead of undefiend
=> so have to make sure we point at null

Hope this clear
If it's still confused, go to https://youtu.be/OvpKeraoxW0, or graph the list on papaer

************************************************************
Time: O(n)
Space: O(n)
*/
var copyRandomList = function (head) {
  if (head === null) return null;

  const map = new Map();
  let curNode = head,
    newCurNode = null;

  // First Iteration: create the key <-> value pair (original node <-> clone node)
  while (curNode !== null) {
    newCurNode = new Node(curNode.val);
    map.set(curNode, newCurNode);
    curNode = curNode.next;
  }

  // Second Iteration: wire up, connect the next and random pointer
  curNode = head;
  while (curNode !== null) {
    // use Map to get the new clone node
    newCurNode = map.get(curNode);

    // connect next and random ptr
    // note two edge cases
    newCurNode.next = map.get(curNode.next) ? map.get(curNode.next) : null;
    newCurNode.random = map.get(curNode.random)
      ? map.get(curNode.random)
      : null;

    // update
    curNode = curNode.next;
  }

  return map.get(head);
};

/*
Solution 2
This is more tricky

This is really hard to describe by words
Go to https://youtu.be/OvpKeraoxW0, or graph the list on papaer

Again, this solution is tricky
Have to graph the linked list to understand deeply
Also, there are the edge cases in this soultion
Have to graph the linked list to understand deeply

************************************************************
Time: O(n)
Space: O(1)
*/
var copyRandomList1 = function (head) {
  if (head === null) return null;

  let curNode = head,
    newCurNode = null;

  // First Iteration: connect orinigal list and clone list as one linked list
  while (curNode !== null) {
    newCurNode = new Node(curNode.val);
    newCurNode.next = curNode.next;
    curNode.next = newCurNode;

    // update ptr
    curNode = curNode.next.next;
  }

  // Second Iteration: connect random ptr connection
  curNode = head;
  while (curNode !== null) {
    // Note that curNode.random may be null, null.next will have error
    curNode.next.random = curNode.random ? curNode.random.next : null;

    // update ptr
    curNode = curNode.next.next;
  }

  // Hode the reference of the head of new clone list
  const dummyHead = new Node(null, head.next);

  // Third Iteration: re-construct the connection(both original list and clone list)
  // Just simply re-connect the next ptr
  curNode = head;
  while (curNode !== null) {
    newCurNode = curNode.next;
    curNode.next = newCurNode.next;

    // edge case: the next ptr of last node point to null, null.next will have error
    newCurNode.next = newCurNode.next ? newCurNode.next.next : null;

    // update ptr
    curNode = curNode.next;
  }

  return dummyHead.next;
};
