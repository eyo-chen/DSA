struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    int nextDigit = 0;
    ListNode* head = new ListNode(0, nullptr);
    ListNode* ptr = head;

    while (l1 != nullptr || l2 != nullptr) {
      int amount = nextDigit;

      if (l1 != nullptr) {
        amount += l1->val;
        l1 = l1->next;
      }

      if (l2 != nullptr) {
        amount += l2->val;
        l2 = l2->next;
      }

      if (amount >= 10) {
        nextDigit = amount / 10;
        amount -= nextDigit * 10;
      } else {
        nextDigit = 0;
      }

      ListNode* node = new ListNode(amount, nullptr);
      ptr->next = node;
      ptr = ptr->next;
    }

    if (nextDigit > 0) {
      ListNode* node = new ListNode(nextDigit, nullptr);
      ptr->next = node;
    }

    return head->next;
  }
};