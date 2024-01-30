var rotateRight = function (head, k) {
  /*
    Base case
    1. the length of head is 0
    2. the length of head is 1
    3. k = 0, which means don't need to do any rotation
  */
  if (head === null || head.next === null || k === 0) return head;

  let len = 1,
    tail = head,
    newTail = head,
    newHead = null,
    rightShiftIndex = null;

  // find the tail and the length of linked list
  while (tail.next !== null) {
    len++;
    tail = tail.next;
  }

  // k % len === 0 means do need to do any rotation
  // for example, the length of linked list is 5, and k = 5, the rotation after that is just original linked list
  if (k % len === 0) return head;

  // the index helps us to find the new tail
  rightShiftIndex = len - (k % len) - 1;

  // find new tail
  while (rightShiftIndex !== 0) {
    newTail = newTail.next;
    rightShiftIndex--;
  }

  // pointer rewiring
  tail.next = head;
  newHead = newTail.next;
  newTail.next = null;

  return newHead;
};

// basically same idea, the code implementation is different
var rotateRight1 = function (head, k) {
  if (head === null || head.next === null || k === 0) return head;

  let len = 1,
    tail = head;

  while (tail.next !== null) {
    len++;
    tail = tail.next;
  }

  let rightShiftIndex = k % len,
    newTail = head;
  tail = head;

  if (rightShiftIndex === 0) return head;

  while (tail.next !== null) {
    if (rightShiftIndex === 0) {
      tail = tail.next;
      newTail = newTail.next;
    } else {
      rightShiftIndex--;
      tail = tail.next;
    }
  }

  const newHead = newTail.next;
  tail.next = head;
  newTail.next = null;

  return newHead;
};
