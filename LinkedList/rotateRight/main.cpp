struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* rotateRight(ListNode* head, int k) {
    if (head == nullptr || head->next == nullptr || k == 0) return head;

    // find the length of the list, and the tail
    int len = 1;
    ListNode* tail = head;
    while (tail->next != nullptr) {
      tail = tail->next;
      len++;
    }

    // rotateCount is the number of nodes to rotate
    int rotateCount = k % len;

    // if rotateCount == 0, no need to rotate
    if (rotateCount == 0) return head;

    // find the new tail
    ListNode* newTail = head;
    int findNewTailIndex = len - rotateCount - 1;
    for (int i = 0; i < findNewTailIndex; i++) {
      newTail = newTail->next;
    }

    // find the new head
    ListNode* newHead = newTail->next;

    // wire the tail to the head
    tail->next = head;

    // cut off the new tail
    newTail->next = nullptr;

    return newHead;
  }
};

// This is the solution I come up with, but it is not as good as the above one
// because it has to handle the edge case where the new start node is the tail
// node.
// And the code is not clean
class Solution {
 public:
  ListNode* rotateRight(ListNode* head, int k) {
    if (head == nullptr || head->next == nullptr || k == 0) return head;

    // find the length of the list
    int len = 0;
    ListNode* ptr = head;
    while (ptr != nullptr) {
      ptr = ptr->next;
      len++;
    }

    // find the new start node
    ptr = new ListNode(0, head);
    for (int i = 0; i < k % len; i++) {
      ptr = ptr->next;
    }
    ListNode* newStartNode = head;
    // prt->next is edge case
    while (ptr->next != nullptr && ptr->next != newStartNode) {
      newStartNode = newStartNode->next;
      ptr = ptr->next;
    }

    // wire the tail to the head
    ptr = newStartNode;
    for (int i = 0; i < len - 1; i++) {
      if (ptr->next == nullptr) {
        ptr->next = head;
      }

      ptr = ptr->next;
    }

    // cut off the new tail
    ptr->next = nullptr;

    return newStartNode;
  }
};
