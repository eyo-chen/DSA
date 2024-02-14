struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* deleteDuplicates(ListNode* head) {
    if (head == nullptr) return head;

    ListNode* dummyHead = new ListNode(0, head);
    ListNode* p1 = dummyHead;
    ListNode* p2 = head;

    while (p2 != nullptr && p2->next != nullptr) {
      if (p2->val == p2->next->val) {
        while (p2->next != nullptr && p2->val == p2->next->val) {
          p2 = p2->next;
        }

        /*
        It's important to update p2 to next node after the while loop
        Because the condition of while-loop is p2->val == p2->next->val
        [1, 2a, 2b, 2c, 4]
        After the while loop, p2 will be at 2c
        Because 2c != 4
        So, we need to update p2 to 4
        To make sure p2 is updating to the node after the duplicate node list
        */
        p2 = p2->next;
        p1->next = p2;
        continue;
      }

      p1 = p1->next;
      p2 = p2->next;
    }

    return dummyHead->next;
  }
};