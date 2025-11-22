struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* swapNodes(ListNode* head, int k) {
    ListNode* p1 = head;

    // Find the first node to swap (kth node from the beginning)
    for (int i = 0; i < k - 1; i++) {
      p1 = p1->next;
    }

    // Find the second node to swap (kth node from the end)
    ListNode* p2 = head;
    ListNode* fast = p1;
    while (fast->next != nullptr) {
      fast = fast->next;
      p2 = p2->next;
    }

    // Swap the values of the two nodes
    int tempVal = p1->val;
    p1->val = p2->val;
    p2->val = tempVal;

    return head;
  }
};