#include <set>

using namespace std;

struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

// First Solution: Using Hash Set
class Solution {
 public:
  bool hasCycle(ListNode* head) {
    if (head == nullptr || head->next == nullptr) return false;

    set<ListNode*> hashSet;
    while (head != nullptr) {
      if (hashSet.find(head) != hashSet.end()) return true;

      hashSet.insert(head);
      head = head->next;
    }

    return false;
  }
};

// Second Solution: Using Two Pointers
class Solution {
 public:
  bool hasCycle(ListNode* head) {
    if (head == nullptr || head->next == nullptr) return false;

    ListNode* slow = head;
    ListNode* fast = head;

    while (fast != nullptr && fast->next != nullptr) {
      slow = slow->next;
      fast = fast->next->next;

      if (slow == fast) return true;
    }

    return false;
  }
};