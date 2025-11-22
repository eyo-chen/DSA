struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

// First Solution
class Solution {
 public:
  ListNode* reverseBetween(ListNode* head, int left, int right) {
    ListNode* dummyHead = new ListNode(0, head);
    ListNode* nodeBeforeStrPtr = dummyHead;
    ListNode* strPtr = head;

    for (int i = 0; i < left - 1; i++) {
      nodeBeforeStrPtr = nodeBeforeStrPtr->next;
      strPtr = strPtr->next;
    }

    ListNode* prev = nullptr;
    ListNode* cur = strPtr;
    int reverseCount = right - left + 1;
    for (int i = 0; i < reverseCount && cur != nullptr; i++) {
      ListNode* next = cur->next;
      cur->next = prev;

      prev = cur;
      cur = next;
    }

    strPtr->next = cur;
    nodeBeforeStrPtr->next = prev;

    return dummyHead->next;
  }
};

// Second Solution
class Solution {
 public:
  ListNode* reverseBetween(ListNode* head, int left, int right) {
    ListNode* dummyHead = new ListNode(0, head);
    ListNode* nodeBeforeStrPtr = dummyHead;

    for (int i = 0; i < left - 1; i++) {
      nodeBeforeStrPtr = nodeBeforeStrPtr->next;
    }

    ListNode* prev = nodeBeforeStrPtr;
    ListNode* cur = prev->next;
    int reverseCount = right - left;
    for (int i = 0; i < reverseCount; i++) {
      ListNode* next = cur->next;

      cur->next = next->next;
      next->next = prev->next;
      prev->next = next;
    }

    return dummyHead->next;
  }
};