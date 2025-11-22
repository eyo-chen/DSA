/*
Iterative Approach

There are several approaches to do iteratively
Hard to explain the logic here
Just grab a paper, and graph the actual linked list

There are two thing to note
1. Instead of using condition guard in while-statement, i use if-statement to break out the while-loop
=> Why do this?
=> The first part is rewiring, and the second part is updating
=> Imagine we put the condition in the while-loop
=> First example, x -> 1 -> 2 -> 3 -> 4
=> x -> 1 -> 2 -> 3 -> 4
   p    a    b
=> After first rewiring,
=> x -> 2 -> 1 -> 3 -> 4
   p    b    a
=> Go updating
=> x -> 2 -> 1 -> 3 -> 4
             p    a    b
=> And now the while-loop will stop
=> Why?
=> because a.next.next === null
=> so we won't even swap the last part, which is not we want
=> Rewiring again,
=> x -> 2 -> 1 -> 4 -> 3
             p    b    a
=> Again, it's kind of hard to explain here, just grab the paper and graph the linked list

2. if (a.next === null || a.next.next === null) break;
=> a.next === null is when the case like this
=> x -> 2 -> 1 -> 4 -> 3
             p    b    a
=> a.next.next === null is when the case like this
=> x -> 2 -> 1 -> 4 -> 3 -> 5
             p    b    a
=> In both case, we have to stop the while-loop

************************************************************
Time: O(n)
Space: O(1)
*/
var swapPairs = function (head) {
  if (head === null || head.next === null) return head;

  const dummyHead = new ListNode(null);
  dummyHead.next = head;
  let a = head,
    b = head.next,
    prev = dummyHead;

  while (true) {
    // rewiring
    a.next = b.next;
    b.next = a;
    prev.next = b;

    // stop the while-loop if hitting the end
    if (a.next === null || a.next.next === null) break;

    // updating
    prev = a;
    a = a.next;
    b = a.next;
  }

  return dummyHead.next;
};

/*
Recursive Approach

!! Still need to review, can't understand the logic
*/
var swapPairs1 = function (head) {
  if (head === null || head.next === null) return head;

  const next = head.next;
  head.next = swapPairs1(head.next.next);
  next.next = head;

  return next;
};

////////////////////////////////////////////////////////////////////////////////
// Helper Function
////////////////////////////////////////////////////////////////////////////////
const a = [1, 2, 3, 4];
function createList(arr) {
  let index = 0;
  const dummyHead = new Node(null);
  let curNode = dummyHead;

  while (index < arr.length) {
    const node = new Node(arr[index]);
    curNode.next = node;
    curNode = curNode.next;
    index++;
  }

  return dummyHead.next;
}
// const res = createList(a);
