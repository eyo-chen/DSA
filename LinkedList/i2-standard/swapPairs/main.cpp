struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* swapPairs(ListNode* head) {
    ListNode* dummyHead = new ListNode(0, head);
    ListNode* cur = head;
    ListNode* prev = dummyHead;

    while (cur != nullptr && cur->next != nullptr) {
      ListNode* next = cur->next;
      ListNode* nextNext = next->next;

      cur->next = nextNext;
      next->next = cur;
      prev->next = next;

      prev = cur;
      cur = nextNext;
    }

    return dummyHead->next;
  }
};