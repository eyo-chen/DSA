/*
First solution
It's intuitive and straightforward

Four things to note
1. Look at the base case
=> head === null means the length of head is 0, like this []
=> head.next === null means the length of head is 1, like this [1]
=> In both cases, we could just simple return head
=> Do not overlook the base case, the length of linked list is 0 or 1

2. Have to hold the reference of evenHead beforeHand
=> Why?
=> After iteration, we'll have the tail of even list and odd list
=> And we have to connect tail of odd list with the head of even list
=> odd list: 1 -> 3 -> 5
=> even list: 2 -> 4 
=> Have to connect 5 with 2

3. Initialize index with 1
=> Why it's not 0?
=> Because index helps us to keep tracking the current position is odd or even
=> It's abviois the very first position of head should be 1
=> So that the next one is 2, so on and so forth

4. even.next = null; is necessary
=> Why?
=> This is like the edge case when the length of head is odd
=> input head: 1 -> 2 -> 3 -> 4 -> 5
=> After iteration,
=> odd list: 1 -> 3 -> 5
=> even list: 2 -> 4 
=> After while-loop, 4 still point to 5
=> We have to disconnect
=> try to paint on the papar

************************************************************
Time: O(n)
Space: O(1)
*/
var oddEvenList = function (head) {
  if (head === null || head.next === null) return head;

  let odd = new ListNode(null),
    even = new ListNode(null),
    cur = head,
    evenHead = head.next,
    index = 1;

  while (cur !== null) {
    if (index % 2 === 0) {
      even.next = cur;
      even = cur;
    } else {
      odd.next = cur;
      odd = cur;
    }

    cur = cur.next;
    index++;
  }

  even.next = null;
  odd.next = evenHead;

  return head;
};

/*
This is second solution without using index

Note two things
1. Try to paint on the paper, and see the logic inside while-loop

2. Why using even inside validation?
=> Look at the logic inside while-loop
=> Every time after while-loop, even is always in front of the odd
=> For example,
=> 1 -> 2 -> 3 -> 4
             o    e     
=> 1 -> 2 -> 3 -> 4 -> 5
                       o    e  
=> Again, try to write on the paper to see clearly
=> The main point is that we can make sure even pointer is always in front of the odd pointer
=> because of this two operation
=>  even.next = odd.next;
    even = even.next;

************************************************************
Time: O(n)
Space: O(1)
*/
var oddEvenList1 = function (head) {
  if (head === null || head.next === null) return head;
  let odd = head,
    even = head.next,
    evenHead = head.next;

  while (even !== null && even.next !== null) {
    odd.next = even.next;
    odd = odd.next;
    even.next = odd.next;
    even = even.next;
  }

  odd.next = evenHead;

  return head;
};
