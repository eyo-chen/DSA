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
  void reorderList(ListNode* head) {
    if (head->next == nullptr) return;
    ListNode* ptr = head;

    while (ptr != nullptr && ptr->next != nullptr) {
      // find the last node and the node before the last node
      ListNode* tail = ptr;
      ListNode* nodeBeforeTail = new ListNode(0, ptr);
      while (tail->next != nullptr) {
        tail = tail->next;
        nodeBeforeTail = nodeBeforeTail->next;
      }

      // reference the next node of ptr
      ListNode* curNextNode = ptr->next;
      if (curNextNode == tail) break;

      // re-wire the node
      ptr->next = tail;
      tail->next = curNextNode;

      // cut off the node before the last node
      // so that it becomes the last node
      nodeBeforeTail->next = nullptr;

      // update ptr
      ptr = curNextNode;
    }
  }
};

// Second Solution
class Solution {
 public:
  void reorderList(ListNode* head) {
    if (head->next == nullptr) return;

    // find the middle of the list
    ListNode* fast = head;
    ListNode* slow = head;
    while (fast != nullptr && fast->next != nullptr) {
      slow = slow->next;
      fast = fast->next->next;
    }

    // reverse the second half of the list
    ListNode* cur = slow->next;
    ListNode* prev = nullptr;
    while (cur != nullptr) {
      ListNode* next = cur->next;
      cur->next = prev;
      prev = cur;
      cur = next;
    }
    slow->next = nullptr;

    // merge the two halves
    ListNode* firstList = head;
    ListNode* secondList = prev;
    while (secondList != nullptr) {
      ListNode* fNext = firstList->next;
      ListNode* sNext = secondList->next;

      firstList->next = secondList;
      secondList->next = fNext;

      firstList = fNext;
      secondList = sNext;
    }
  }
};