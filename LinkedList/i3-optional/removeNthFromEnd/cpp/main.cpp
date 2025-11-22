struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

// Solution 1
class Solution {
 public:
  ListNode* removeNthFromEnd(ListNode* head, int n) {
    ListNode* dummyHead = new ListNode(0, head);
    ListNode* right = dummyHead;
    ListNode* left = dummyHead;

    for (int i = 0; i < n; i++) {
      right = right->next;
    }

    while (right != nullptr && right->next != nullptr) {
      right = right->next;
      left = left->next;
    }

    left->next = left->next->next;

    return dummyHead->next;
  }
};

// Solution 2
class Solution {
 public:
  ListNode* removeNthFromEnd(ListNode* head, int n) {
    ListNode* dummyHead = new ListNode(0, head);
    ListNode* right = head;
    ListNode* left = dummyHead;

    for (int i = 0; i < n; i++) {
      right = right->next;
    }

    while (right != nullptr) {
      right = right->next;
      left = left->next;
    }

    left->next = left->next->next;

    return dummyHead->next;
  }
};

// Solution 3
class Solution {
 public:
  ListNode* removeNthFromEnd(ListNode* head, int n) {
    int length = 0;
    ListNode* ptr = head;

    while (ptr != nullptr) {
      ptr = ptr->next;
      length++;
    }

    ListNode* dummyHead = new ListNode(0, head);
    ptr = dummyHead;
    int distanceToPrevRemovedNode = length - n;
    for (int i = 0; i < distanceToPrevRemovedNode; i++) {
      ptr = ptr->next;
    }

    ptr->next = ptr->next->next;

    return dummyHead->next;
  }
};