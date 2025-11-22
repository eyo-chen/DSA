var addTwoNumbers = function (l1, l2) {
  const dummyHead = new ListNode(null);
  let curVal,
    node,
    curNode = dummyHead,
    curCarry = 0;

  while (l1 !== null || l2 !== null) {
    curVal = 0;

    if (l1 !== null) {
      curVal += l1.val;
      l1 = l1.next;
    }

    if (l2 !== null) {
      curVal += l2.val;
      l2 = l2.next;
    }

    curVal += curCarry;

    if (curVal >= 10) curCarry = 1;
    else curCarry = 0;

    node = new ListNode(curVal % 10);
    curNode.next = node;
    curNode = curNode.next;
  }

  if (curCarry === 1) {
    node = new ListNode(1);
    curNode.next = node;
  }

  return dummyHead.next;
};

/*
The is recursive approach which I referecne from leetcode
Sitll not familar with it
*/
var addTwoNumbers1 = function (l1, l2, carry = 0) {
  let curNode = null;

  if (l1 || l2) {
    let curVal = 0;

    if (l1 !== null) {
      curVal += l1.val;
      l1 = l1.next;
    }

    if (l2 !== null) {
      curVal += l2.val;
      l2 = l2.next;
    }

    curVal += carry;

    let nextCarry = 0;
    if (curVal >= 10) nextCarry = 1;

    curNode = new ListNode(curVal % 10);
    curNode.next = addTwoNumbers1(l1, l2, nextCarry);
  } else if (carry > 0) {
    curNode = new ListNode(1);
    curNode.next = null;
  }

  return curNode;
};
