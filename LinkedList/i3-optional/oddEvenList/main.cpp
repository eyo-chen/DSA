struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

// First Solution: use index to determine odd or even
class Solution {
 public:
  ListNode* oddEvenList(ListNode* head) {
    ListNode* ptr = head;
    ListNode* oddList = new ListNode();
    ListNode* evenList = new ListNode();
    ListNode* evenHead = evenList;
    int index = 1;

    while (ptr != nullptr) {
      if (index % 2 == 0) {
        evenList->next = ptr;
        evenList = ptr;
      } else {
        oddList->next = ptr;
        oddList = ptr;
      }

      index++;
      ptr = ptr->next;
    }

    evenList->next = nullptr;
    oddList->next = evenHead->next;

    return head;
  }
};

// Second Solution: use two pointers to wire up the odd and even list
class Solution {
 public:
  ListNode* oddEvenList(ListNode* head) {
    if (head == nullptr || head->next == nullptr) return head;

    ListNode* oddPtr = head;
    ListNode* evenPtr = head->next;
    ListNode* evenHead = evenPtr;

    // use evenPtr as validation because evenPtr is always ahead of oddPtr
    while (evenPtr != nullptr && evenPtr->next != nullptr) {
      oddPtr->next = evenPtr->next;
      oddPtr = oddPtr->next;

      evenPtr->next = oddPtr->next;
      evenPtr = evenPtr->next;
    }

    oddPtr->next = evenHead;

    return head;
  }
};