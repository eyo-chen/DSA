//////////////////////////////////////////////////////
// *** Add Two Numbers ***
//////////////////////////////////////////////////////
/*
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [7,2,4,3], l2 = [5,6,4]
Output: [7,8,0,7]

Example 2:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [8,0,7]

Example 3:
Input: l1 = [0], l2 = [0]
Output: [0]

Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

Follow up: Could you solve it without reversing the input lists?
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
This is solution using stack
This idea is actually conceptually reverse the linked list

For example,
l1 = [7,2,4,3]
the stack would be [7,2,4,3]
later, we would read the element from the end (First in Last out)
Same thing for l2
So using stack helps us caculate the number from first digit, instead last digit
Watch this video https://www.youtube.com/watch?v=t0ljMKo2Hjs&t=1356s

One thing to note
Look how we construct the final returned linked list
Normally, we construct the linked list from the beginning
Like this
x -> 1
x -> 1 -> 2
x -> 1 -> 2 -> 3
But this won't work in this problem

stack1 = [7,2,4,3], stack2 = [5,6,4]
We have to read the element from the end becasue it's stack
4 + 3 = 7
x -> 7
6 + 4 = 10 (carry = 1)
x -> 7 -> 0
1 + 5 + 2
x -> 7 -> 0 -> 8
x -> 7 -> 0 -> 8 -> 7
This is not the question ask us ([7,8,0,7])

So we have to construct the linked list from the end, like this
4 + 3 = 7
7  <-  x
6 + 4 = 10 (carry = 1)
0 <- 7 <-  x
1 + 5 + 2
8 <- 0 <- 7 <-  x
7 <- 8 <- 0 <- 7 <- x
This is important and the main key takeaway of this problem

************************************************************
Time: O(n)
Space: O(n)
*/
var addTwoNumbers = function (l1, l2) {
  // create two stacks
  const stack1 = createStack(l1);
  const stack2 = createStack(l2);
  let node = null,
    newNode = null,
    carry = 0;

  // while-loop through both stack, read the value from the end
  while (stack1.length > 0 || stack2.length > 0) {
    let val = carry;
    if (stack1.length > 0) {
      val += stack1.pop();
    }

    if (stack2.length > 0) {
      val += stack2.pop();
    }

    // set the carry
    if (val >= 10) {
      val -= 10;
      carry = 1;
    } else carry = 0;

    // create new node, the next pointer point at old node
    newNode = new ListNode(val, node);

    // update node
    node = newNode;
  }

  // if there's carry, add it to the final linked list
  if (carry > 0) {
    newNode = new ListNode(carry, node);
    node = newNode;
  }

  return node;
};

function createStack(head) {
  const res = [];
  while (head !== null) {
    res.push(head.val);
    head = head.next;
  }

  return res;
}

/*
This solution without using stack

Input: l1 = [7,2,4,3], l2 = [5,6,4]
The process is like this

1. find the longList, shortList, and their length
Why?
Because later we need to add both linked list, and we need to know these variables
So that we can add tow digits in the correct position
For example, hundred digit, thousand digit, so on and so forth

2. building a sum linked list from the end
Now we know the lengt, and which is longer and shorter
sum linked list: 7 <- 7 <- 10 <- 7 <- x
Note that we do not take care of the carry here
All we do is adding the digit, and building the linked list from the end

3. Based on the computed sum linked list, construct the final returned linked list
Look the computed sum linked list above
The order is from x -> 7 -> 10 -> 7 -> 7
Now we can easily follow this order to build the final returned linked list again
Again, we have to build the final returned linked list from the start
7 <- x
0 <- 7 <- x
8 <- 0 <- 7 <- x
7 <- 8 <- 0 <- 7 <- x
This is our final answer

The main takeaway of this problem is very similar to the previous one
We build the linked list from the end twice, which is important to solve the problem
Watch this video https://www.youtube.com/watch?v=t0ljMKo2Hjs&t=1356s

************************************************************
Time: O(n)
Space: O(n)
*/
var addTwoNumbers = function (l1, l2) {
  // count the length
  const len1 = countLength(l1),
    len2 = countLength(l2);

  // set the variables
  let longList = len1 > len2 ? l1 : l2,
    shortList = len1 > len2 ? l2 : l1,
    longLen = len1 > len2 ? len1 : len2,
    shortLen = len1 > len2 ? len2 : len1,
    newNode = null,
    node = null,
    cur = node,
    carry = 0;

  // construct the linked list first time, do not need to cara about the carry
  while (longLen > 0) {
    let sum = 0;
    // add two digit when the length is equal, which means they're at same position
    if (longLen === shortLen) {
      sum += longList.val + shortList.val;
      longLen--;
      shortLen--;
      longList = longList.next;
      shortList = shortList.next;
    } else {
      sum += longList.val;
      longLen--;
      longList = longList.next;
    }

    // build the linked list from the end
    newNode = new ListNode(sum, node);
    node = newNode;
  }

  // build the linked list again, but now care about the carry
  cur = node;
  node = null;
  while (cur !== null) {
    let val = carry + cur.val;

    if (val >= 10) {
      carry = 1;
      val -= 10;
    } else carry = 0;

    // build the linked list from the end
    newNode = new ListNode(val, node);
    node = newNode;
    cur = cur.next;
  }

  // add the node if there's carry leftover
  if (carry > 0) {
    newNode = new ListNode(carry, node);
    node = newNode;
  }
  return node;
};

function countLength(head) {
  let len = 0;
  while (head !== null) {
    len++;
    head = head.next;
  }
  return len;
}
