/*
Even though this question ask us to solve by only one pass
this problem is still pretty intutive

The hard part of this function is handle the edge case
Like lenght two and length one linked list

This solution is wrote by myself

For example, if linked list is 1 -> 2 -> 3 -> 4 -> 5 -> x, and n = 2
the node we're gonna remove is 4, so new linked list is 1 -> 2 -> 3 -> 5 -> x
The idea is quite easy
We just use two pointers (c: right ptr(cur), r: left ptr(removePointer))
n = 2
1 -> 2 -> 3 -> 4 -> 5 -> x
c
r

n = 1
1 -> 2 -> 3 -> 4 -> 5 -> x
     c
r

n = 0
1 -> 2 -> 3 -> 4 -> 5 -> x
          c
r

and move c all way to the end of linked list
1 -> 2 -> 3 -> 4 -> 5 -> x
                   c
          r
Now we have access to 3, so we can easily remove 4

BUT what if the length of linked list is only 2, like this
1 -> 2, n = 2
c
r

1 -> 2, n = 1  
     c
r

Now we have two problems
1. we want to remove 1, but we can't access any node before 1
2. n is not 0, and we're already at the end of linked list

Another simialr case
If 
1 -> 2, n = 1
c
r

1 -> 2, n = 0  
     c
r
Now we can easily remove 2

So we have to use some if-statement to handle these edge cases
************************************************************
Time: O(n)
Space: O(1)
*/
var removeNthFromEnd = function (head, n) {
  /*
  Specail base case
  Why we can just return null if the length of linked list is 1
  Becasue the description of problem say n is equal or greater than 1
  So if the length of linked list is 1, just remove this node
  simply return node
  */
  if (head.next === null) return null;

  // two pointers
  let cur = head,
    removePointer = head;

  // move two pointers until cur is at the end of linked list
  while (cur.next !== null) {
    if (n === 0) {
      cur = cur.next;
      removePointer = removePointer.next;
    } else {
      n--;
      cur = cur.next;
    }
  }

  // Edge cases: the length of linked list is 2, and n = 2, we want to remove the first node
  // so just return head.next
  if (removePointer === head && n === 1) return head.next;

  // other simple case, remove the node
  removePointer.next = removePointer.next.next;

  return head;
};

/*
This is second solution, this is cleaner

We don't need special base case or if-statement to handle the edge case
Just use the dummyHead to help us

First put right at the head, and move n position
x -> 1 -> 2, n = 2
     r

x -> 1 -> 2, n = 1
          r

x -> 1 -> 2 -> x, n = 0
               r

Put left at the dummyHead
x -> 1 -> 2 -> x, n = 0
               r
l

Because of right is at the end of linked list, so we don't need to move left ptr
And now left ptr at the dummyHead, so we can easily remove the first node

So dummyHead help us to handle the edge case

Just graph to understand clearly

************************************************************
Time: O(n)
Space: O(1)
*/
var removeNthFromEnd1 = function (head, n) {
  const dummyHead = new ListNode(null, head);

  // two ptrs
  let right = head,
    left = dummyHead;

  // move right n position
  for (let i = 0; i < n; i++) {
    right = right.next;
  }

  // move the right all the way to the end of linked list(null)
  while (right !== null) {
    right = right.next;
    left = left.next;
  }

  // remove the node
  left.next = left.next.next;

  return dummyHead.next;
};
