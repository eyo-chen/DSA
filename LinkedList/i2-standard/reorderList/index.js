/*
This is the solution using O(n) space

The idea is pretty simple
1. create hash table, key is the order of linked list, value is the corresponding node
2. loop through linked list, re-connect the node

The idea is quite simple, but it's not that easy to implement in the code
There are two details have to care about
1. create the hashTable and count the length of linked list
   => len start at 0
   => we have to increment the len, and add it to the hashTable
   => Again, the order of increment and adding to the hashTable does matter here
   => try to change the order, and see the result

2. have to cut off all the connection when storing the node in hashTable
   => hold the reference of next node (next = cur.next)
   => build the hashTable (hashTable[len] = cur)
   => cur off the .next connection  (hashTable[len].next = null)
   => Why we need to do this?
   => If will cause the cycle later when we re-order the linked list without this
   => we only want to store the single node in each key-value pair, not entire linked list

************************************************************
Time: O(n)
Space: O(n)
*/
var reorderList = function (head) {
  const hashTable = {};
  let cur = head,
    len = 0,
    index = 2, // why 2? => later, we want to re-order the linked list from second node because first node don't need to change
    next = null;

  // build the hashTable
  while (cur !== null) {
    next = cur.next; // hold the reference of next node
    len++;
    hashTable[len] = cur;
    hashTable[len].next = null; // cut off the connection
    cur = next; // update to the next node
  }

  /*
  Two indexes are important
  They help us know what's the current order need to be connected

  indexOdd => In odd position, we want the node from start
  indexEven => In even position, we want the node from the end
  */
  let indexOdd = 2, // start from 2 because first node is fixed
    indexEven = len;
  cur = head;

  // re-order the linked list
  while (index <= len) {
    // even position, connect with the node of indexEven
    if (index % 2 === 0) {
      cur.next = hashTable[indexEven];
      indexEven--;
    }
    // odd position, connect with the node of indexOdd
    else {
      cur.next = hashTable[indexOdd];
      indexOdd++;
    }
    index++;
    cur = cur.next;
  }
};

/*
This is solution using O(1) space
The idea is not that intutive as previous one, but the code implementation is that hard

The idea is 
1. find the middle of linked list, cut off the linked list into first part and second part
2. reverse the linked list of second part
3. re-oreder the list

Why this works?
For example, input is 1 -> 2 -> 3 -> 4 -> 5
1. find the middle of linked list, cut off the linked list into first part and second part
   => First part: 1 -> 2
   => Second part: 3 -> 4 -> 5
2. reverse the linked list of second part
   => Second part: 5 -> 4 -> 3
3. re-oreder the list
   => First part: 1 -> 2
   => Second part: 5 -> 4 -> 3
   => Now we can connect
   => 1 -> 5
   => 5 -> 2
   => 2 -> 4
   => done

There are two edge cases need to be careful
1. we need nodeBeforeSlow
   => why?
   => because we need cut off the connection between first part and second part
   => wihtout this, first part would be like this  1 -> 2 -> 3 -> 4
   => See, first part still point to the end of second part
   => this is important because it will cause the erro later when we re-oreder the lined list without this
2. cur.next !== null and cur.next = prev;
   => why we only iterate throughcur.next !== null and need cur.next = prev; after the while-loop?
   => First part: 1 -> 2
   => Second part: 5 -> 4 -> 3
   => Now we can connect
   => 1 -> 5
   => 5 -> 2
   => at this momenet, cur.next is null (2 -> x)
   => we have to stop here
   => because after while-loop, we just need simply connect to the prev, and we're done
   => 2 -> 4
   => done

************************************************************
Time: O(n)
Space: O(1)
*/
var reorderList1 = function (head) {
  if (!head.next) return head;

  let fast = head,
    slow = head,
    nodeBeforeSlow = new ListNode(null, head),
    cur = null,
    prev = null,
    next = null,
    tmp1 = null,
    tmp2 = null;

  // find the middle of linked list
  while (fast !== null && fast.next !== null) {
    fast = fast.next.next;
    slow = slow.next;
    nodeBeforeSlow = nodeBeforeSlow.next;
  }
  // cut off the linked list to be two parts
  nodeBeforeSlow.next = null;

  // reverse the second parts
  cur = slow;
  while (cur !== null) {
    next = cur.next;
    cur.next = prev;
    prev = cur;
    cur = next;
  }

  // re-order the linked list
  cur = head;
  while (cur.next !== null) {
    // hold the reference of next node
    tmp1 = cur.next;
    tmp2 = prev.next;

    // re-order
    cur.next = prev;
    prev.next = tmp1;

    // update to next node
    cur = tmp1;
    prev = tmp2;
  }

  // final step
  cur.next = prev;
};
