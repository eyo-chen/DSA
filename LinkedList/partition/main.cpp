struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* partition(ListNode* head, int x) {
    ListNode* ptr = head;
    ListNode* lessPtr = new ListNode();
    ListNode* greaterPtr = new ListNode();
    ListNode* lessHead = lessPtr;
    ListNode* greaterHead = greaterPtr;

    while (ptr != nullptr) {
      if (ptr->val < x) {
        lessPtr->next = ptr;
        lessPtr = ptr;
      } else {
        greaterPtr->next = ptr;
        greaterPtr = ptr;
      }

      ptr = ptr->next;
    }

    greaterPtr->next = nullptr;
    lessPtr->next = greaterHead->next;

    return lessHead->next;
  }
};