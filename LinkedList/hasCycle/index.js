var hasCycle = function (head) {
  const set = new Set();

  while (head !== null) {
    if (set.has(head) === true) return true;
    set.add(head);
    head = head.next;
  }

  return false;
};

var hasCycle = function (head) {
  let fast = head,
    slow = head;

  while (fast !== null && fast.next !== null) {
    // have to move the ptr first
    // because fast and slow is intitally at the same position(head)
    fast = fast.next.next;
    slow = slow.next;

    // then check
    if (fast === slow) return true;
  }

  return false;
};
