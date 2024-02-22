#include <algorithm>
#include <cstdlib>
#include <stack>

using namespace std;

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
    // Create two stacks to store the digits of the two lists
    stack<int> stack1 = toStack(l1);
    stack<int> stack2 = toStack(l2);

    // Build the new list by popping the digits from the stacks(in reverse
    // order)
    int nextDigit = 0;
    ListNode* cur = nullptr;
    while (stack1.size() > 0 || stack2.size() > 0) {
      int val = nextDigit;

      // only pop the stack if it is not empty
      if (stack1.size() > 0) {
        val += stack1.top();
        stack1.pop();
      }

      // only pop the stack if it is not empty
      if (stack2.size() > 0) {
        val += stack2.top();
        stack2.pop();
      }

      // handle the carry
      if (val >= 10) {
        nextDigit = 1;
        val -= 10;
      } else {
        nextDigit = 0;
      }

      ListNode* node = new ListNode(val, cur);
      cur = node;
    }

    // Edge case: if there is a carry after the last digit, add a new node
    if (nextDigit == 1) {
      ListNode* node = new ListNode(1, cur);
      cur = node;
    }

    return cur;
  }

 private:
  stack<int> toStack(ListNode* l) {
    stack<int> stack;
    ListNode* cur = l;

    while (cur != nullptr) {
      stack.push(cur->val);
      cur = cur->next;
    }

    return stack;
  }
};

class Solution {
 public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    // Count the length of each list
    int len1 = countLength(l1);
    int len2 = countLength(l2);

    // Find the longer list and the shorter list
    int maxLen = max(len1, len2);
    ListNode* maxList;
    ListNode* minList;
    if (len1 > len2) {
      maxList = l1;
      minList = l2;
    } else {
      maxList = l2;
      minList = l1;
    }

    // Build the new list without considering the carry
    // The new list is built in reverse order on purpose
    int diff = abs(len1 - len2);
    ListNode* ptr = nullptr;
    for (int i = 0; i < maxLen; i++) {
      int val = 0;
      if (i >= diff) {
        val += minList->val;
        minList = minList->next;
      }

      val += maxList->val;
      maxList = maxList->next;

      ListNode* node = new ListNode(val, ptr);
      ptr = node;
    }

    // Reverse the new list, and deal with the carry
    ListNode* prev = nullptr;
    ListNode* cur = ptr;
    int nextDigit = 0;
    while (cur != nullptr) {
      cur->val += nextDigit;
      if (cur->val >= 10) {
        cur->val -= 10;
        nextDigit = 1;
      } else {
        nextDigit = 0;
      }

      ListNode* next = cur->next;
      cur->next = prev;
      prev = cur;
      cur = next;
    }

    // Edge case:
    // If there is a carry after the last digit, add a new node
    if (nextDigit == 1) {
      ListNode* node = new ListNode(1, prev);
      prev = node;
    }

    return prev;
  }

 private:
  int countLength(ListNode* l) {
    int len = 0;
    ListNode* cur = l;
    while (cur != nullptr) {
      len++;
      cur = cur->next;
    }
    return len;
  }
};
